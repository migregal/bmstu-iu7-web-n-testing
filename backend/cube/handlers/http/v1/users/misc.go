package users

import "neural_storage/cube/core/entities/user"

func fromBL(info user.Info) UserInfo {
	return UserInfo{
		Id:       info.ID(),
		Email:    info.Email(),
		Username: info.Username(),
		Fullname: info.Fullname(),
	}
}
