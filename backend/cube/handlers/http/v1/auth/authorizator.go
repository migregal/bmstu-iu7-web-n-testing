package auth

import (
	"neural_storage/cube/handlers/http/jwt"

	"github.com/gin-gonic/gin"
)

func (m *Handler) Authorizator(roles []uint64) func(c *gin.Context, data any) bool {
	return func(c *gin.Context, data any) bool {
		usr, ok := data.(*User)

		c.Set(jwt.IdentityKey, usr.ID)
		c.Set(jwt.IdentityFlags, usr.Flags)

		allowed := false
		for _, role := range roles {
			if usr.Flags&role == role {
				allowed = true
			}
		}

		return ok && allowed
	}
}
