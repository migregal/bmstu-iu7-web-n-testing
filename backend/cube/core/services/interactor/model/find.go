package model

import (
	"context"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/ports/interactors"
	"neural_storage/cube/core/ports/repositories"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Find(ctx context.Context, filter interactors.ModelInfoFilter) ([]*model.Info, int64, error) {
	lg := i.lg.WithFields(map[string]any{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]any{"filter": filter}).Info("model find called")

	lg.Info("attempt to find model info")
	return i.modelInfo.Find(
		repositories.ModelInfoFilter{
			Owners: filter.Owners,
			IDs:    filter.IDs,
			Offset: filter.Offset,
			Limit:  filter.Limit,
		},
	)
}
