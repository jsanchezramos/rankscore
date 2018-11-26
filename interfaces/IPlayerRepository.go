package interfaces

import "socialpoint/model"

type IPlayerRepository interface {
	GetAllPlayers() []model.User
	AddPlayer(user model.User)
	UpdatePlayer(user model.User)
}
