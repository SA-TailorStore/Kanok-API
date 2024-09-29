package mysql

import (
	"context"

	"github.com/SA-TailorStore/Kanok-API/entities"
	"github.com/SA-TailorStore/Kanok-API/exceptions"
	"github.com/SA-TailorStore/Kanok-API/reposititories"
	"github.com/SA-TailorStore/Kanok-API/requests"
	"github.com/SA-TailorStore/Kanok-API/responses"
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
func (u *UserMySQL) FindAllUser(ctx context.Context) ([](entities.User), error) {
	rows, err := u.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.ID, &user.Username, &user.Password,
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

	user := responses.UsernameResponse{}

	err := u.db.GetContext(ctx, &user, "SELECT username FROM users WHERE username = ?", req)

	if err != nil {
		err = exceptions.ErrDuplicatedUsername
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername implements reposititories.UserRepository.
func (u *UserMySQL) GetUserByUsername(ctx context.Context, req *requests.UsernameRequest) (*responses.UserResponse, error) {
	user := responses.UserResponse{}
	err := u.db.GetContext(ctx, &user, "SELECT * FROM users WHERE username = ?", req.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
