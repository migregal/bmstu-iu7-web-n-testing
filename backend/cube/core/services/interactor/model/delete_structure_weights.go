package model

import (
	"context"
	"neural_storage/cube/core/entities/structure/weights"
	"neural_storage/pkg/logger"
)

func (i *Interactor) DeleteStructureWeights(ctx context.Context, ownerID, weightsId string) error {
	lg := i.lg.WithFields(map[string]any{logger.ReqIDKey: ctx.Value(logger.ReqIDKey)})

	lg.WithFields(map[string]any{"owner": ownerID, "weights": weightsId}).Info("model delete called")
	info := *weights.NewInfo(weightsId, "", nil, nil)

	lg.Info("attempt to delete weights info")
	return i.weightsInfo.Delete([]weights.Info{info})
}
