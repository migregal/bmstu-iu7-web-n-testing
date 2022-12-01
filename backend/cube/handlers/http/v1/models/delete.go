package models

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	ID string `uri:"model_id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

// Registration  godoc
// @Summary      Delete model info
// @Description  Deletes model info owned by authorized user
// @Tags         models
// @Param        model_id   path string true "Model ID to delete"
// @Success      200 "Model info deleted"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to delete model info from storage"
// @Router       /v1/models/{model_id} [delete]
// @security     ApiKeyAuth
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailDelete.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		statFailDelete.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req DeleteRequest
	if err := c.BindUri(&req); err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	defer func() {
		lg.Info("attempt to delete model from cache")
		_ = h.cache.Delete(modelStorage, req.ID)
	}()

	lg.WithFields(map[string]any{"req": req}).Info("attempt to delete model")
	err := h.resolver.Delete(c, usrID, req.ID)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete model: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, "failed to delete model info")
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to delete model info")
		}
		return
	}

	statOKDelete.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
