package userblock

import "neural_storage/cube/core/entities/user"

func fromBL(info user.Info) BlockInfo {
	return BlockInfo{ID: info.ID(), Until: info.BlockedUntil()}
}
