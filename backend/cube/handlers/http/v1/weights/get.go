package weights

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getRequest struct {
	ID string `uri:"weight_id" binding:"required"`
}

// Registration  godoc
// @Summary      Find model weights info
// @Description  Find such model weights info as id, name, link weights and neuron offsets
// @Tags         weights
// @Param        weight_id    path string    true "Weight ID to search for"
// @Success      200 {object} []weights.Info "Model weights info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to get model weights info from storage"
// @Router       /v1/weights/{weight_id} [get]
// @security     ApiKeyAuth
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getRequest
	if err := c.ShouldBindUri(&req); err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	lg.WithFields(map[string]any{"req": req}).Info("attempt to get weight info into cache")
	if info, err := h.cache.Get(weightStorage, req.ID); err != nil || len(info) < 2 {
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
	if req.ID != "" {
		filter.IDs = append(filter.IDs, req.ID)
	}

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

	res := weightFromBL(*infos[0])
	if data, err := jsonGzip(res); err == nil {
		_ = h.cache.Update(weightStorage, req.ID, data)
	}

	statOKGet.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)
}
