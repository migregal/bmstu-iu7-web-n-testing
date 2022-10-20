package models

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"neural_storage/cube/core/entities/model"
	httpmodel "neural_storage/cube/handlers/http/v1/entities/model"
	"neural_storage/cube/handlers/http/v1/entities/structure"
	"neural_storage/cube/handlers/http/v1/entities/structure/weights"

	"neural_storage/cube/handlers/http/jwt"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	ModelTitle string                `form:"title" binding:"required"`
	Structure  *multipart.FileHeader `form:"structure" binding:"required"`
	Weights    *multipart.FileHeader `form:"weights" binding:"required"`
}

// Registration  godoc
// @Summary      Create new model
// @Description  Adds such model info as title, structure, weights
// @Tags         models
// @Accept       multipart/form-data
// @Param        title           formData string true "Model Title to create"
// @Param        structure_title formData string true "Model Structure Title to add"
// @Param        structure       formData file   true "Model Structure to add"
// @Param        weights_title   formData string true "Model Weights Title to add"
// @Param        weights         formData file   true  "Model Weights to add"
// @Success      200 {object} httpmodel.Info "Model created"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      403 "Invalid token, user id not specified"
// @Failure      500 "Failed to create model info"
// @Router       /v1/models [post]
// @security     ApiKeyAuth
func (h *Handler) Add(c *gin.Context) {
	statCallAdd.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	claimID, ok := c.Get(jwt.IdentityKey)
	if !ok {
		statFailAdd.Inc()
		lg.Error("access token missing")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}
	usrID, ok := claimID.(string)
	if !ok {
		statFailAdd.Inc()
		lg.Error("invalid access token")
		c.JSON(http.StatusForbidden, "invalid access token")
		return
	}

	var req AddRequest

	if err := c.ShouldBind(&req); err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	content, err := req.Structure.Open()
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to find structure info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	r, err := gzip.NewReader(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read gzipped structure info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer r.Close()

	plan, err := io.ReadAll(r)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read structure info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var structure structure.Info
	err = json.Unmarshal(plan, &structure)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to parse structure info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid structure format")
		return
	}

	content, err = req.Weights.Open()
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to find weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rw, err := gzip.NewReader(content)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read gzipped weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer rw.Close()

	plan, err = io.ReadAll(rw)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to read weights info: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var w weights.Info
	err = json.Unmarshal(plan, &w)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to parse weights info: %v", err)
		c.JSON(http.StatusBadRequest, "invalid weights format")
		return
	}
	structure.Weights = []weights.Info{w}

	lg.WithFields(map[string]any{"user": usrID, "title": req.ModelTitle}).Info("attempt to add new model")
	model := model.NewInfo("", usrID, req.ModelTitle, structToBL(structure))
	modelID, err := h.resolver.Add(c, *model)
	if err != nil {
		statFailAdd.Inc()
		lg.Errorf("failed to add new model: %v", err)
		c.JSON(http.StatusInternalServerError, "model creation failed")
		return
	}

	var res httpmodel.Info = modelFromBL(model)
	res.ID = modelID
	if stringified, err := jsonGzip(res); err == nil {
		_ = h.cache.Update(modelStorage, modelID, stringified)
	}

	statOKAdd.Inc()
	lg.Info("success")
	c.JSON(http.StatusOK, res)
}
