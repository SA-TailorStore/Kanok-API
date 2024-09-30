package mysql

import (
	"context"
	"database/sql"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/entities"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserMySQL struct {
	db *sqlx.DB
}

func NewUserMySQL(db *sqlx.DB) reposititories.UserRepository {
	return &UserMySQL{
		db: db,
	}
}

// Create implements reposititories.UserRepository.
func (u *UserMySQL) Create(ctx context.Context, req *requests.UserRegisterRequest) error {
	user_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = u.db.QueryContext(ctx, "INSERT INTO users (user_id, username, password, phone_number, display_name, user_profile_url, role, address) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)",
		user_id, req.Username, req.Password, req.Phone_number, "-", "-", "-", "-")

	return err
}

// FindAllUser implements reposititories.UserRepository.
func (u *UserMySQL) GetAllUser(ctx context.Context) ([](entities.User), error) {
	rows, err := u.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.User_id, &user.Username, &user.Password,
			&user.Created_at,
			&user.Phone_number,
			&user.User_profile_url,
			&user.Role,
			&user.Display_name,
			&user.Address); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByUsername implements reposititories.UserRepository.
func (u *UserMySQL) FindByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UsernameResponse, error) {

	var user responses.UsernameResponse

	err := u.db.GetContext(ctx, &user, "SELECT username FROM users WHERE username = ?", req.Username)
	switch err {
	case sql.ErrNoRows:
		return &user, nil
	case nil:
		return nil, exceptions.ErrDuplicatedUsername
	default:
		return nil, err
	}

}

// GetUserByUsername implements reposititories.UserRepository.
func (u *UserMySQL) GetPasswordByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UserLoginResponse, error) {

	var user responses.UserLoginResponse

	err := u.db.GetContext(ctx, &user, "SELECT user_id,password FROM users WHERE username = ?", req.Username)
	if err != nil {
		return nil, exceptions.ErrUserNotFound
	}

	return &user, nil
}

func (u *UserMySQL) GetUserByUserID(ctx context.Context, req *requests.UserIDRequest) (*responses.UserResponse, error) {
	var user responses.UserResponse

	err := u.db.GetContext(ctx, &user, "SELECT user_id,username,display_name,user_profile_url,role,phone_number,address,created_at FROM users WHERE user_id = ?", req.User_id)

	if err != nil {
		return nil, exceptions.ErrUserNotFound
	}

	return &user, nil
}
