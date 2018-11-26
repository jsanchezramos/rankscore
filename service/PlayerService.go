package service

import (
	"fmt"
	"socialpoint/interfaces"
	"socialpoint/model"
	"sort"
	"strconv"
	"strings"
)

type PlayerService struct {
	interfaces.IPlayerRepository
}

func (service *PlayerService) GetScoreTop(topNumber int,listUsers []model.User) []model.User {
    var userScoreList []model.User

    listUsers = service.orderArrayScore(listUsers)

	if topNumber <= len(listUsers) {
		for i := 0; i < topNumber; i++ {
			userScoreList = append(userScoreList, listUsers[i])
		}
	}

	return userScoreList
}

func (service *PlayerService) GetRelativeRank(posUser int , posAround int,listUsers []model.User) []model.User{
	var userScoreRedundant []model.User
	var total = len(listUsers)

	listUsers = service.orderArrayScore(listUsers)

	if total>0 {

		for i := posUser + 1 ; i < posUser + posAround ; i++{
			if i < total {
				userScoreRedundant = append(userScoreRedundant,listUsers[i])
			}
		}

		for i := (posUser-1) - posAround ; i <= posUser ; i++{
			if i >= 0 {
				userScoreRedundant = append(userScoreRedundant,listUsers[i])
			}
		}
	}

	return service.orderArrayScore(userScoreRedundant)
}

func (service *PlayerService) PutScoreUser(username string, score string) {
	i, err := strconv.Atoi(score)

	users := service.GetAllPlayers()
	user := model.User{}

	for i := 0; i < len(users); i++ {
		if users[i].Name == username{
			user = users[i]
		}
	}

	if user.Name == username {
		if strings.Contains(score, "+") || strings.Contains(score, "-"){
			user.Total = user.Total + i
		}else {
			user.Total = i
		}

		service.UpdatePlayer(user)
	}

	if err != nil {
		fmt.Print(i)
	}
}


func (service *PlayerService) orderArrayScore (arrayUser []model.User)([]model.User){

	sort.Slice(arrayUser[:], func(i, j int) bool {
		return arrayUser[i].Total > arrayUser[j].Total
	})

	return arrayUser
}







