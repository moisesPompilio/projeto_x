package userhandles

import "net/http"

type IUserHandles interface {
	LoginHandle(w http.ResponseWriter, r *http.Request)
	DeleteUserHandles(w http.ResponseWriter, r *http.Request)
	GetByIDHandles(w http.ResponseWriter, r *http.Request)
	CreateUserHandle(w http.ResponseWriter, r *http.Request)
}
