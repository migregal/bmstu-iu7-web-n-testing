package userblock

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type updateRequestURI struct {
	UserId string    `uri:"user_id" binding:"required"`
}

type updateRequestJSON struct {
	Until  time.Time `json:"until" binding:"required"`
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

	var reqURI updateRequestURI
	if err := c.BindUri(&reqURI); err != nil || reqURI.UserId == "" {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request uri: %v", err)
		return
	}
	var reqJSON updateRequestJSON
	if err := c.Bind(&reqJSON); err != nil || reqJSON.Until.IsZero() {
		statFailUpdate.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	lg.WithFields(map[string]any{"req": reqURI}).Info("attempt to block user")
	err := h.resolver.Block(c, reqURI.UserId, reqJSON.Until)
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
