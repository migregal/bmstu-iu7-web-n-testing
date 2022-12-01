package userblock

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}

type BlockInfo struct {
	ID    string    `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Until time.Time `json:"blocked_until,omitempty" example:"2025-08-09T15:00:00.053Z"`
} // @name userblockUserResponse

// Registration  godoc
// @Summary      Find user block info
// @Description  Find such users info as id and block time
// @Tags         blocks
// @Produce      json
// @Param        user_id path string    true "UserId to search for"
// @Success      200 {object} BlockInfo "Users block info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to get user info from storage"
// @Router       /v1/blocks/users/{user_id} [get]
// @security     ApiKeyAuth
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.BindUri(&req); err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get user block info")
	info, err := h.resolver.Get(c, req.UserID)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to get user block info: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "failed to fetch user info")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		}
		return
	}

	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, fromBL(info))
}
