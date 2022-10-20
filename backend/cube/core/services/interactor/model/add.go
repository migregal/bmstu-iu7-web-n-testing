package model

import (
	"context"
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/pkg/logger"
)

func (i *Interactor) Add(ctx context.Context, info model.Info) (string, error) {
	lg := i.lg.WithFields(map[string]any{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]any{"info": info}).Info("model add called")
	if err := i.validator.ValidateModelInfo(&info); err != nil {
		lg.Error("invlaid model info")
		return "", fmt.Errorf("invalid model info: %w", err)
	}

	lg.Info("attempt to add model info")
	return i.modelInfo.Add(info)
}
