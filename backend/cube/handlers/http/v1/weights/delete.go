package weights

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ID string `json:"id" form:"id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

// Registration  godoc
// @Summary      Delete model info
// @Description  Deletes model info from any user
// @Tags         weights
// @Param        weight_id  path string true "Model ID to delete"
// @Success      200 "Model info deleted"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to delete model info from storage"
// @Router       /v1/weights/{weight_id} [delete]
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

	var req deleteRequest
	if err := c.Bind(&req); err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	defer func() {
		lg.Info("attempt to delete weight from cache")
		_ = h.cache.Delete(weightStorage, req.ID)
	}()

	lg.WithFields(map[string]any{"user": usrID, "id": req.ID}).Info("attempt to delete weights")
	err := h.resolver.DeleteStructureWeights(c, usrID, req.ID)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete weights info: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "failed to delete model weights info")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to delete model weights info")
		}
		return
	}

	statOKDelete.Inc()
	c.AbortWithStatus(http.StatusOK)
}
