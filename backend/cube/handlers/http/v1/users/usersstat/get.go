package usersstat

import (
	"net/http"
	"neural_storage/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	From         time.Time `form:"from"`
	To           time.Time `form:"to"`
	Registration bool      `form:"registration"`
	Update       bool      `form:"update"`
}

type StatInfo struct {
	ID   string    `json:"id"   example:"3f8bf2a3-01cc-4c4a-9759-86cec9cf8da9"`
	Time time.Time `json:"time" example:"2006-01-02T15:04:05Z07:00"`
}

type UserStatInfo struct {
	Registrations []StatInfo `json:"registration,omitempty"`
	Edits         []StatInfo `json:"edit,omitempty"`
} // @name UserStatInfoResponse

// Registration  godoc
// @Summary      Get users stat info
// @Description  Get such user stat info as registration and edit stat per period
// @Tags         users
// @Produces     json
// @Param        from     query string   false "Time to start from, RFC3339" format("2006-01-02T15:04:05Z07:00")
// @Param        to       query string   false "Time to stop at, RFC3339" format("2006-01-02T15:04:05Z07:00")
// @Param        load     query boolean  false "Search for load stat"
// @Param        update   query boolean  false "Search for update stats"
// @Success      200 {object} []UserStatInfo "Users stat info found"
// @Failure      400 "Invalid request"
// @Failure      401 "Unauthorized"
// @Failure      500 "Failed to get user stat info"
// @Router       /v1/users/stats [get]
// @security     ApiKeyAuth
func (h *Handler) Get(c *gin.Context) {
	statCallGet.Inc()
	lg := h.lg.WithFields(map[string]any{logger.ReqIDKey: c.Value(logger.ReqIDKey)})

	var req request
	if err := c.ShouldBind(&req); err != nil {
		statFailGet.Inc()
		lg.Errorf("failed to bind request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	resp := UserStatInfo{}

	if req.Registration {
		lg.WithFields(map[string]any{"req": req}).Info("attempt to get load stat")
		data, err := h.resolver.GetUserRegistrationStat(c, req.From, req.To)
		if err != nil {
			statFailGet.Inc()
			lg.Errorf("failed to get load stat data: %v", err)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		for _, v := range data {
			resp.Registrations = append(resp.Registrations, fromBL(v))
		}
	}

	if req.Update {
		lg.WithFields(map[string]any{"req": req}).Info("attempt to get update stat")
		data, err := h.resolver.GetUserEditStat(c, req.From, req.To)
		if err != nil {
			statFailGet.Inc()
			lg.Errorf("failed to get update stat data: %v", err)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		for _, v := range data {
			resp.Edits = append(resp.Edits, fromBL(v))
		}
	}

	statOKGet.Inc()
	lg.WithFields(map[string]any{"res": resp}).Info("success")
	c.JSON(http.StatusOK, resp)
}
