package repositorie

import (
	"github.com/jmoiron/sqlx"

	userrepositorie "github.com/moisesPompilio/projeto_x/src/adapters/repositorie/UserRepositorie"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User userrepositorie.UserRepositorie
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		User: userrepositorie.NewUserRepository(opts.WriterSqlx, opts.ReaderSqlx),
	}
}
