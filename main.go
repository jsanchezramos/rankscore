package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"socialpoint/interfaces"
	"socialpoint/model"
	"socialpoint/repositori"
	srv "socialpoint/service"
)


func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	fmt.Fprintf(w, "%s", content)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)

	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	repo := new(repositori.PlayerRepository)
	srv := srv.PlayerService{repo}
	loadDummyUsers(repo)

	go echo(conn,repo, &srv)
}

func echo(conn *websocket.Conn,repo interfaces.IPlayerRepository,srv interfaces.IPlayerService) {
	for {
		m := model.Msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		if m.Adduser{
			repo.AddPlayer(m.User)
			fmt.Print(repo.GetAllPlayers())
		}else if m.Top {
			m.Users = srv.GetScoreTop(m.Numtop,repo.GetAllPlayers())
			fmt.Print(m.Users)
		}else if m.Relative {
			m.Users = srv.GetRelativeRank(m.Posrelative,m.Aroundrelative,repo.GetAllPlayers())
			fmt.Print(m.Users)
		}else if m.AddScore{
			srv.PutScoreUser(m.User.Name,m.Newscore)
		}

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}
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