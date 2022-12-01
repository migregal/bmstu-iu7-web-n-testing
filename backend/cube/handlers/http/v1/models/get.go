package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"neural_storage/cube/core/ports/interactors"
	_ "neural_storage/cube/handlers/http/v1/entities/model"
	"neural_storage/pkg/logger"
)

type getRequest struct {
	ModelID   string `uri:"model_id" binding:"required"`
}

// Registration  godoc
// @Summary      Find model info
// @Description  Find such model info as id, username, email and fullname
// @Tags         models
// @Param        model_id  path string true "Model ID to search for"
// @Success      200 {object} []model.Info "Model info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to get model info from storage"
// @Router       /v1/models/{model_id} [get]
// @security     ApiKeyAuth
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.BindUri(&req); err != nil {
		lg.Errorf("failed to bind request: %v", err)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get model info into cache")
	if info, err := h.cache.Get(modelStorage, req.ModelID); err != nil && len(info) < 2 {
		lg.Errorf("failed to get model info from cache: %v", err)
	} else {
		lg.Info("success to get model info from cache")
		resp, err := unGzip(info[1].([]byte))
		if err == nil {
			statOKGet.Inc()
			c.Data(http.StatusOK, "application/json", resp)
			return
		}
		lg.Errorf("failed to ungzip model info from cache: %v", err)
	}

	filter := interactors.ModelInfoFilter{}
	if req.ModelID != "" {
		filter.IDs = []string{req.ModelID}
	}

	lg.WithFields(map[string]any{"filter": filter}).Info("attempt to find model info")
	infos, _, err := h.resolver.Find(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to find model info: %v", err)

		if errors.Is(err, interactors.ErrNotFound) {
			lg.Info("no models found")
			c.JSON(http.StatusNotFound, "failed to fetch user info")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		}
		return
	}

	res := modelFromBL(infos[0])
	if data, err := jsonGzip(res); err == nil {
		_ = h.cache.Update(modelStorage, req.ModelID, data)
	}

	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)
}
