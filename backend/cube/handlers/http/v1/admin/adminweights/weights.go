package adminweights

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
	statCallDelete = stat.NewCounter("v1", "cube_admin_weights_call_delete", "The total number of deleting admin weights attempts")
	statFailDelete = stat.NewCounter("v1", "cube_admin_weights_fail_delete", "The total number of deleting admin weights fails attempts")
	statOKDelete = stat.NewCounter("v1", "cube_admin_weights_ok_delete", "The total number of deleted admin weights")
}

type Handler struct {
	resolver interactors.NeuralNetworkInteractor

	lg *logger.Logger
}

func New(lg *logger.Logger, resolver interactors.NeuralNetworkInteractor) Handler {
	return Handler{resolver: resolver, lg: lg}
}
