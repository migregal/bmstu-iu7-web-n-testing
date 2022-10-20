package userblock

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	UserId string    `uri:"user_id" binding:"required"`
	Until  time.Time `json:"until" form:"until" binding:"required"`
}

// Registration  godoc
// @Summary      Block user
// @Description  Blocks user until specified moment
// @Tags         blocks
// @Param        user_id path string  true "User ID to block"
// @Param        until   query string true "Time to block until"
// @Success      200 "User blocked"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to block user "
// @Router       /v1/blocks/users/{user_id} [patch]
// @security     ApiKeyAuth
func (h *Handler) Update(c *gin.Context) {
	statCallUpdate.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req updateRequest
	if err := c.ShouldBindUri(&req); err != nil || req.UserId == "" {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := c.ShouldBind(&req); err != nil || req.UserId == "" {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to block user")
	err := h.resolver.Block(c, req.UserId, req.Until)
	if err != nil {
		statFailUpdate.Inc()
		lg.Errorf("failed to block user: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "failed to block user")
		}
		c.JSON(http.StatusInternalServerError, "failed to block user")
		return
	}

	statOKUpdate.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
