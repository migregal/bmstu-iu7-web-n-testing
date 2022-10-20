package modelsstat

import "neural_storage/cube/core/entities/model/modelstat"

func fromBL(v *modelstat.Info) StatInfo {
	return StatInfo{ID: v.ID(), Time: v.CreatedAt()}
}
