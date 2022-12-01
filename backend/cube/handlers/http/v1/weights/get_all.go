package weights

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getAllRequest struct {
	StructureID string `form:"structure_id"`
	Page        int    `form:"page"`
	PerPage     int    `form:"per_page"`
}

type getAllResponse struct {
	Infos []weights.Info `json:"infos"`
}

// Registration  godoc
// @Summary      Find structure weights info
// @Description  Find such model info as id, username, email and fullname
// @Tags         weights
// @Param        structure_id query string true "Structure ID to search for"
// @Param        page         query int    false "Page number for pagination"
// @Param        per_page     query int    false "Page size for pagination"
// @Success      200 {object} []weights.Info "Model weights info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to get model weights info from storage"
// @Router       /v1/weights [get]
// @security     ApiKeyAuth
func (h *Handler) GetAll(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getAllRequest
	if err := c.Bind(&req); err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to bind request: %v", err)
		return
	}
	if structID := c.Param("structure_id"); structID != "" {
		req.StructureID = structID
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get weight info into cache")
	if info, err := h.cache.Get(weightStorage, req.StructureID); err != nil || len(info) < 2 {
		lg.Errorf("failed to get weight info from cache: %v", err)
	} else {
		lg.Info("success to get weight info from cache")
		resp, err := unGzip(info[1].([]byte))
		if err == nil {
			statOKGet.Inc()
			c.Data(http.StatusOK, "application/json", resp)
			return
		}
		lg.Errorf("failed to ungzip weight info from cache: %v", err)
	}

	filter := interactors.ModelWeightsInfoFilter{}
	if req.StructureID != "" {
		filter.Structures = []string{req.StructureID}
	}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	lg.WithFields(map[string]any{"filter": filter}).Info("attempt to get weights")
	infos, err := h.resolver.FindStructureWeights(c, filter)
	if err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		if errors.Is(err, interactors.ErrNotFound) {
			lg.Info("no weights found")
			c.JSON(http.StatusNotFound, "failed to fetch weights info")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to fetch weights info")
		}
		return
	}

	var res []weights.Info
	for _, val := range infos {
		res = append(res, weightFromBL(*val))
	}

	resp := getAllResponse{res}
	if req.StructureID != "" {
		if data, err := jsonGzip(resp); err == nil {
			_ = h.cache.Update(weightStorage, req.StructureID, data)
		}
	}
	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, resp)
}
