package repository

import (
	"database/sql"
	"fmt"

	"github.com/abraaoduarte/ekklesia-backend-golang/internal/models"
)


type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (userRepo *UserRepository) GetUserByID(userID int) (*models.User, error) {
    sqlQuery := `SELECT id, name FROM users WHERE id=$1`

    row := userRepo.DB.QueryRow(sqlQuery, userID)

	var user models.User

	err := row.Scan(&user.ID, &user.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found with ID: %v", userID)
        }
        return nil, err
    }

	return &user, nil

}