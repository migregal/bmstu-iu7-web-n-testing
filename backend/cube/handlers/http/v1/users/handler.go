package users

import (
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)


var (
	statCall stat.Counter
	statFail stat.Counter
	statOK   stat.Counter
)

func init() {
	statCall = stat.NewCounter("v1", "cube_users_call_read", "The total number of getting user info attempts")
	statFail = stat.NewCounter("v1", "cube_users_fail_read", "The total number of getting user info fails")
	statOK = stat.NewCounter("v1", "cube_users_ok_read", "The total number of login attempts")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{lg: lg, resolver: resolver}
}
