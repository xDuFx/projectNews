package models

type News struct {
	ID    int
	Title string
	Body  string
	Image string
	Mark  string
	Reliz string
}

type PathFromServer struct{
	Path string
}

type UserDataLogin struct{
	Login string
	Hashpass string

}