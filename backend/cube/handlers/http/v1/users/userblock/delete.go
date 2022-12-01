package userblock

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	UserId string `uri:"user_id"`
}

// Registration  godoc
// @Summary      Delete user block info
// @Description  Deletes user block info by user id
// @Tags         blocks
// @Param        user_id path string true "User ID to unblock"
// @Success      200 "User unblocked"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to delete user block info from storage"
// @Router       /v1/blocks/users/{user_id} [delete]
// @security     ApiKeyAuth
func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req deleteRequest
	if err := c.BindUri(&req); err != nil || req.UserId == "" {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request %v", err)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get user block info")
	err := h.resolver.Block(c, req.UserId, time.Now().Add(5*time.Minute))
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to unblock user %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "failed to unblock user")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to unblock user")
		}
		return
	}

	statOKDelete.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
