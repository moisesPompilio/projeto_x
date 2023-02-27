package userhandles

import "net/http"

type IUserHandles interface {
	LoginHandle(w http.ResponseWriter, r *http.Request)
	DeleteUserHandles(w http.ResponseWriter, r *http.Request)
	// GetByID(ctx context.Context, ID int64) (*domain.UserWithoutPassword, error)
	// GetAll(ctx context.Context) ([]domain.UserWithoutPassword, error)
	CreateUserHandle(w http.ResponseWriter, r *http.Request)
}
