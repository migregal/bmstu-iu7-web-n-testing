package models

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/v1/entities/model"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getAllRequest struct {
	OwnerID string `form:"user_id"`
	Page    int    `form:"page"`
	PerPage int    `form:"per_page"`
}

type getAllResponse struct {
	Infos []model.Info `json:"infos"`
	Total int64        `json:"total"`
} // @name GetAllUsersResponse

// Registration  godoc
// @Summary      Find models info
// @Description  Find such model info as id, username, email and fullname
// @Tags         models
// @Param        page     query int    false "Page number for pagination"
// @Param        per_page query int    false "Page size for pagination"
// @Success      200 {object} []model.Info "Model info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      500 "Failed to get model info from storage"
// @Router       /v1/models [get]
// @security     ApiKeyAuth
func (h *Handler) GetAll(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getAllRequest
	if err := c.Bind(&req); err != nil {
		lg.Errorf("failed to bind request: %v", err)
		return
	}
	if owner_id := c.Param("user_id"); owner_id != "" {
		req.OwnerID = owner_id
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get models info into cache")
	if info, err := h.cache.Get(modelStorage, req.OwnerID); err != nil || len(info) < 2 {
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
	if req.OwnerID != "" {
		filter.Owners = append(filter.Owners, req.OwnerID)
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	lg.WithFields(map[string]any{"filter": filter}).Info("attempt to find model info")
	infos, total, err := h.resolver.Find(c, filter)
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

	var res getAllResponse
	for _, val := range infos {
		res.Infos = append(res.Infos, modelFromBL(val))
	}
	res.Total = total

	if req.OwnerID != "" {
		if data, err := jsonGzip(res); err == nil {
			_ = h.cache.Update(modelStorage, req.OwnerID, data)
		}
	}

	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)
}
