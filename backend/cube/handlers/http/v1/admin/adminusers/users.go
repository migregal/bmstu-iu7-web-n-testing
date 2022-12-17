package adminusers

import (
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/pkg/logger"
	"neural_storage/pkg/stat"
)

var (
	statCallDelete stat.Counter
	statFailDelete stat.Counter
	statOKDelete   stat.Counter
)

func init() {
	statCallDelete = stat.NewCounter("v1", "cube_admin_users_call_delete", "The total number of deleting admin users attempts")
	statFailDelete = stat.NewCounter("v1", "cube_admin_users_fail_delete", "The total number of deleting admin users fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_admin_users_ok_delete", "The total number of deleted admin users")
}

type Handler struct {
	resolver interactors.UserInfoInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.UserInfoInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}
