package output

import "github.com/google/uuid"

type UserOutpDTO struct {
	Id       uuid.UUID `db:"id" json:"id"`
	NickName string    `db:"nick_name" json:"nick_name"`
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email"`
}
