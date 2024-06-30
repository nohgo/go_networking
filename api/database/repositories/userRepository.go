package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/api/models"
)

type UserRepository interface {
	Add(models.User) error
	GetAll() ([]models.User, error)
}

type postgresUserRepository struct {
	pool *sql.DB
}

func NewUserRepository() *postgresUserRepository {
	return &postgresUserRepository{db.Pool}
}

func (ur *postgresUserRepository) Add(user models.User) error {
	_, err := ur.pool.Exec(fmt.Sprintf("INSERT INTO users (username, password) VALUES ('%v', '%v')", user.Username, user.Password))
	return err
}

func (*postgresUserRepository) GetAll() ([]models.User, error) {
	return nil, nil
}
