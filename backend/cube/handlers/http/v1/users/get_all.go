package users

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type getAllRequest struct {
	Page     int    `form:"page"`
	PerPage  int    `form:"per_page"`
}

// Registration  godoc
// @Summary      Find users info
// @Description  Find such users info as id, username, email and fullname
// @Tags         users
// @Accept       json
// @Param        username query string false "Username to search for"
// @Param        email    query string false "Email to search for"
// @Param        page     query int false "Page number for pagination"
// @Param        per_page query int false "Page size for pagination"
// @Success      200 {object} []UserInfo "Users info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      500 "Failed to get user info from storage"
// @Router       /v1/users [get]
// @security     ApiKeyAuth
func (h *Handler) GetAll(c *gin.Context) {
	statCall.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req getAllRequest
	if err := c.ShouldBind(&req); err != nil {
		statFail.Inc()
		lg.Error("failed to bind request")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	filter := interactors.UserInfoFilter{}

	filter.Offset = req.Page

	if req.PerPage == 0 {
		req.PerPage = 10
	}
	filter.Limit = req.PerPage

	lg.WithFields(map[string]any{"filter": filter}).Info("attempt to find user info")
	infos, err := h.resolver.Find(c, filter)
	if err != nil {
		statFail.Inc()
		if errors.Is(err, interactors.ErrNotFound) {
			lg.Info("no users found")
			c.JSON(http.StatusNotFound, "failed to fetch user info")
		} else {
			lg.Error("failed to fetch user info")
			c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		}
		return
	}

	var res []UserInfo
	for _, val := range infos {
		res = append(res, fromBL(val))
	}

	statOK.Inc()
	lg.WithFields(map[string]any{"res": res}).Info("success")
	c.JSON(http.StatusOK, res)
}
