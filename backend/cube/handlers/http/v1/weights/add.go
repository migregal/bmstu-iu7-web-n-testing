package weights

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"

	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

type AddRequest struct {
	ModelID      string                `form:"model_id" binding:"required"`
	Weights      *multipart.FileHeader `form:"weights" binding:"required"`
}

// Registration  godoc
// @Summary      Create new model weights info
// @Description  Adds model weights info to existing model
// @Tags         weights
// @Accept       multipart/form-data
// @Param        model_id        formData string true "Model ID to add weights to"
// @Param        weights         formData file   true "Model Weights to add"
// @Success      200 "Weights added"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to create model weights info"
// @Router       /v1/weights [post]
// @security     ApiKeyAuth
func (h *Handler) Add(c *gin.Context) {
	statCallAdd.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailAdd.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		statFailAdd.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req AddRequest

	if err := c.ShouldBind(&req); err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, err := req.Weights.Open()
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	rw, err := gzip.NewReader(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read gzipped weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer rw.Close()
	plan, err := io.ReadAll(rw)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var w weights.Info
	err = json.Unmarshal(plan, &w)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to parse weights info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid weights format")
		return
	}

	lg.WithFields(map[string]any{"user": usrID, "id": req.ModelID}).Info("attempt to add weights")
	id, err := h.resolver.AddStructureWeights(c, usrID, req.ModelID, weightToBL(w))
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to add weights info: %v", err)
		c.JSON(http.StatusInternalServerError, "model weights creation failed")
		return
	}

	lg.Info("attempt to delete model from cache")
	_ = h.cache.Delete(modelStorage, req.ModelID)

	lg.Info("attempt to delete add weight to cache")
	w.ID = id
	if stringified, err := jsonGzip(weightToBL(w)); err == nil {
		_ = h.cache.Update(modelStorage, id, stringified)
	}

	statOKAdd.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, w)
}
