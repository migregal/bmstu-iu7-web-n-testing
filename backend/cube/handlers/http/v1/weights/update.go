package weights

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UpdateRequestURI struct {
	WeightsID string `uri:"weight_id" binding:"required"`
}

type UpdateRequest struct {
	ModelID      string                `form:"model_id" binding:"required"`
	WeightsTitle string                `form:"weights_title"`
	Weights      *multipart.FileHeader `form:"weights"`
}

// Registration  godoc
// @Summary      Update model info
// @Description  Update such model info as weights, weights titles
// @Tags         weights
// @Accept       multipart/form-data
// @Param        weight_id     path     string true  "Model Weights ID to update"
// @Param        weight_title  formData string false "Model Weights Title to set"
// @Param        weights       formData file   false "Model Weights to Update/Add"
// @Success      200 "Model info updated"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to update model info"
// @Router       /v1/weights/{weight_id} [patch]
// @security     ApiKeyAuth
func (h *Handler) Update(c *gin.Context) {
	statCallUpdate.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailUpdate.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok || usrID == "" {
		statFailUpdate.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var reqURI UpdateRequestURI
	if err := c.BindUri(&reqURI); err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	content, err := req.Weights.Open()
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	r, err := gzip.NewReader(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read gzipped weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer r.Close()

	plan, err := io.ReadAll(r)
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to read weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var w weights.Info
	err = json.Unmarshal(plan, &w)
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to parse weights info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid weights format")
		return
	}
	w.ID = reqURI.WeightsID

	if req.WeightsTitle != "" {
		w.Name = req.WeightsTitle
	}

	lg.WithFields(map[string]any{"user": usrID, "id": req.ModelID}).Info("attempt to update model")
	err = h.resolver.UpdateStructureWeights(c, usrID, req.ModelID, weightToBL(w))
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to update model info: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "model weights info update failed")
		} else {
			c.JSON(http.StatusInternalServerError, "model weights info update failed")
		}
		return
	}

	lg.Info("attempt to delete model from cache")
	_ = h.cache.Delete(modelStorage, req.ModelID)
	lg.Info("attempt to delete weight from cache")
	_ = h.cache.Delete(weightStorage, reqURI.WeightsID)

	statOKUpdate.Inc()
	lg.Info("success")

	c.JSON(http.StatusOK, nil)
}
