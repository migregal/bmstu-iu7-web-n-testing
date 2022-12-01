package models

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"

	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
)

type UpdateRequest struct {
	ModelID      string                `form:"id" binding:"required"`
	WeightsTitle string                `form:"weights_title"`
	Weights      *multipart.FileHeader `form:"weights"`
}

// Registration  godoc
// @Summary      Update model info
// @Description  Update such model info as weights, weights titles
// @Tags         models
// @Accept       multipart/form-data
// @Param        model_id      path     string true  "Model ID to update"
// @Param        weights_title formData string false "Model Weights Title to set"
// @Param        weights       formData file   false "Model Weights to Update/Add"
// @Success      200 "Model info updated"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to update model info"
// @Router       /v1/models/{model_id} [patch]
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
	plan, err := io.ReadAll(content)
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

	lg.WithFields(map[string]any{"user": usrID, "id": req.ModelID}).Info("attempt to update model")
	err = h.resolver.UpdateStructureWeights(c, usrID, req.ModelID, *weightToBL(w))
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to update model info: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "model update failed")
		} else {
			c.JSON(http.StatusInternalServerError, "model update failed")
		}
		return
	}

	lg.Info("attempt to delete model from cache")
	_ = h.cache.Delete(modelStorage, req.ModelID)

	statOKUpdate.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
