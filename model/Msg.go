package model

type Msg struct {
	User User
	Adduser bool
	Top bool
	Numtop int
	Relative bool
	Posrelative int
	Aroundrelative int
	Users []User
	AddScore bool
	Newscore string
}
