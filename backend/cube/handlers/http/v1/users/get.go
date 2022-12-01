package users

import (
	"errors"
	"net/http"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"

	"github.com/gin-gonic/gin"
)

type request struct {
	UserId string `uri:"user_id" binding:"required"`
}

type UserInfo struct {
	Id       string `json:"id,omitempty" example:"f6457bdf-4e67-4f05-9108-1cbc0fec9405"`
	Username string `json:"username,omitempty" example:"awesome_username"`
	Email    string `json:"email,omitempty" example:"my_awesome@email.com"`
	Fullname string `json:"fullname,omitempty" example:"Ivanov Ivan Ivanovich"`
} // @name UserInfoResponse

// Registration  godoc
// @Summary      Find user info
// @Description  Find such users info as id, username, email and fullname
// @Tags         users
// @Accept       json
// @Param        user_id  query string true "UserId to search for"
// @Success      200 {object} []UserInfo "Users info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      404 "Not Found"
// @Failure      500 "Failed to get user info from storage"
// @Router       /v1/users/{user_id} [get]
// @security     ApiKeyAuth
func (h *Handler) Get(c *gin.Context) {
	statCall.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req request
	if err := c.BindUri(&req); err != nil {
		statFail.Inc()
		lg.Error("failed to bind request")
		return
	}

	filter := interactors.UserInfoFilter{}
	if req.UserId != "" {
		filter.Ids = []string{req.UserId}
	}

	lg.WithFields(map[string]any{"filter": filter}).Info("attempt to find user info")
	infos, _, err := h.resolver.Find(c, filter)
	if err != nil {
		statFail.Inc()
		lg.Error("failed to fetch user info")
		if errors.Is(err, interactors.ErrNotFound) {
			c.JSON(http.StatusNotFound, "failed to fetch user info")
		} else {
			c.JSON(http.StatusInternalServerError, "failed to fetch user info")
		}
		return
	}

	res := fromBL(infos[0])

	statOK.Inc()
	lg.WithFields(map[string]any{"res": res}).Info("success")
	c.JSON(http.StatusOK, res)
}
