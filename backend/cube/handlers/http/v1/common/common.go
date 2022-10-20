package common

import (
	"neural_storage/cube/core/roles"
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	adminHandler func(c *gin.Context)
	userHandler func(c *gin.Context)
}

func New(admin, user func(c *gin.Context)) Handler {
	return Handler{adminHandler: admin, userHandler: user}
}

func (h Handler) Handle(c *gin.Context) {
	flags := c.GetUint64(jwt.IdentityFlags)

	if flags&roles.RoleAdmin == roles.RoleAdmin {
		h.adminHandler(c)
	} else {
		h.userHandler(c)
	}
}
