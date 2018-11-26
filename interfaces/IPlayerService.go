package interfaces

import "socialpoint/model"

type IPlayerService interface {
	GetScoreTop(topNumber int,listUsers []model.User) []model.User
	GetRelativeRank(posUser int , posAround int,listUsers []model.User)[]model.User
	PutScoreUser(userName string,score string )
}

