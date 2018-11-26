package service

import (
	"socialpoint/interfaces"
	"socialpoint/model"
	"socialpoint/repositori"
	"testing"
)

func TestGetAllPlayers(t *testing.T) {
	var repo interfaces.IPlayerRepository  = new(repositori.PlayerRepository)

	loadDummyUsers(repo)

	repo.GetAllPlayers()

	t.Log(repo.GetAllPlayers())

}

func TestGetAbosoluteRank(t *testing.T) {
	var repo interfaces.IPlayerRepository  = new(repositori.PlayerRepository)
	var service interfaces.IPlayerService = new(PlayerService)

	loadDummyUsers(repo)

	t.Log(service.GetScoreTop(4,repo.GetAllPlayers()))

}

func TestGetRelativeRank(t *testing.T) {
	var repo interfaces.IPlayerRepository  = new(repositori.PlayerRepository)
	var service interfaces.IPlayerService = new(PlayerService)

	loadDummyUsers(repo)

	t.Log(service.GetRelativeRank(4,2,repo.GetAllPlayers()))
}

func TestGetPutScoreUser(t *testing.T) {

	var repo interfaces.IPlayerRepository  = new(repositori.PlayerRepository)
	var service interfaces.IPlayerService = &PlayerService{repo}

	loadDummyUsers(repo)

	service.PutScoreUser("pedro","+33")

	t.Log(repo.GetAllPlayers())
}

func loadDummyUsers(repo interfaces.IPlayerRepository){

	user := model.User{}
	user.Name =  "Manolo"
	user.Total = 10

	user2 := model.User{}
	user2.Name =  "Joselito"
	user2.Total = 23432

	user3 := model.User{}
	user3.Name =  "Jorgito"
	user3.Total = 20

	user4 := model.User{}
	user4.Name =  "Mateo"
	user4.Total = 20304

	user5 := model.User{}
	user5.Name =  "Mateo Junior"
	user5.Total = 12

	user6 := model.User{}
	user6.Name =  "David";
	user6.Total = 20;

	user7 := model.User{}
	user7.Name =  "Jose"
	user7.Total = 43

	user8 := model.User{}
	user8.Name =  "Manolito"
	user8.Total = 87

	user9 := model.User{}
	user9.Name =  "Loli"
	user9.Total = 10000

	repo.AddPlayer(user)
	repo.AddPlayer(user2)
	repo.AddPlayer(user3)
	repo.AddPlayer(user4)
	repo.AddPlayer(user5)
	repo.AddPlayer(user6)
	repo.AddPlayer(user7)
	repo.AddPlayer(user8)
	repo.AddPlayer(user9)
}
