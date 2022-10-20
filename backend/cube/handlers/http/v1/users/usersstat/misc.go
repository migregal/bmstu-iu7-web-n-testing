package usersstat

import "neural_storage/cube/core/entities/user/userstat"

func fromBL(v *userstat.Info) StatInfo {
	return StatInfo{ID: v.ID(), Time: v.CreatedAt()}
}
