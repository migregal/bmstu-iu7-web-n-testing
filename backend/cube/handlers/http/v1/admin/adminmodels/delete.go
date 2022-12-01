package adminmodels

import (
	"net/http"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	ID string `form:"id" json:"id" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
}

func (h *Handler) Delete(c *gin.Context) {
	statCallDelete.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req deleteRequest
	if err := c.Bind(&req); err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to delete model")
	err := h.resolver.Delete(c, "", req.ID)
	if err != nil {
		statFailDelete.Inc()
		lg.Errorf("failed to delete model: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to delete model info")
		return
	}

	statOKDelete.Inc()
	lg.Info("success")
	c.AbortWithStatus(http.StatusOK)
}
