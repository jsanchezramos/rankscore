package repositori

import "socialpoint/model"

type PlayerRepository struct {
	users []model.User
}


func (repo *PlayerRepository) AddPlayer(user model.User)  {
	repo.users = append(repo.users, user)
}
func (repo *PlayerRepository) GetAllPlayers() []model.User {
	return repo.users
}
func (repo *PlayerRepository)UpdatePlayer(user model.User){
	for i := 0; i < len(repo.users); i++ {
		var userObj = repo.users[i]
		if user.Name == userObj.Name {
			repo.users[i] = user
		}
	}
}

