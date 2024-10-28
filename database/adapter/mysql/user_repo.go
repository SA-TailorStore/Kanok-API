package mysql

import (
	"context"
	"database/sql"

	"github.com/SA-TailorStore/Kanok-API/database/requests"
	"github.com/SA-TailorStore/Kanok-API/database/responses"
	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/SA-TailorStore/Kanok-API/domain/reposititories"
	"github.com/SA-TailorStore/Kanok-API/utils"
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

func (u *UserMySQL) CreateUser(ctx context.Context, req *requests.UserRegister) error {

	if err := utils.CheckUsernameDup(u.db, ctx, req.Username); err != nil {
		return err
	}

	query := `
	INSERT INTO USERS
	(user_id, username, password, phone_number, address) 
	VALUES ( ?, ?, ?, ?, ?)
	`

	user_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = u.db.QueryContext(ctx, query,
		user_id,
		req.Username,
		req.Password,
		req.Phone_number,
		"-")

	return err
}

func (u *UserMySQL) CreateTailor(ctx context.Context, req *requests.UserRegister) error {

	if err := utils.CheckUsernameDup(u.db, ctx, req.Username); err != nil {
		return err
	}

	query := `
	INSERT INTO USERS
	(user_id, username, password, phone_number,role, address) 
	VALUES ( ?, ?, ?, ?, ?)
	`

	user_id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	_, err = u.db.QueryContext(ctx, query,
		user_id,
		req.Username,
		req.Password,
		req.Phone_number,
		"tailor",
		"-")

	return err
}

func (u *UserMySQL) GetAllUser(ctx context.Context, req *requests.UserRole) ([]*responses.User, error) {

	query := `
	SELECT 
		user_id, 
		username, 
		display_name, 
		user_profile_url, 
		role, 
		phone_number, 
		address, 
		timestamp 
	FROM USERS 
	`

	rows, err := u.db.QueryContext(ctx, query)

	switch req.Role {
	case "user", "tailor", "store":
		query = query + "WHERE role = ?"
		rows, err = u.db.QueryContext(ctx, query, req.Role)
	case "":
	default:
		return nil, exceptions.ErrRoleNotHave

	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*responses.User
	for rows.Next() {
		var user responses.User
		if err := rows.Scan(
			&user.User_id,
			&user.Username,
			&user.Display_name,
			&user.User_profile_url,
			&user.Role,
			&user.Phone_number,
			&user.Address,
			&user.Timestamp,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *UserMySQL) GetAllTailor(ctx context.Context, req *requests.UserRole) ([]*responses.UserTailor, error) {

	query := `
	SELECT 
		u.user_id, 
		u.username, 
		u.display_name, 
		u.user_profile_url, 
		u.role, 
		u.phone_number, 
		u.address, 
		u.timestamp,
        CASE 
        	WHEN o.tailor_id = u.user_id && o.status NOT IN ("cancel","success_tailor", "recevied_shop", "success_shop", "recevied_user", "success_user") THEN SUM(p.process_quantity)
        	ELSE 0 
    	END AS process,
    	CASE 
        	WHEN o.tailor_id = u.user_id && o.status NOT IN ("cancel","success_tailor", "recevied_shop", "success_shop", "recevied_user", "success_user") THEN SUM(p.total_quantity)
        	ELSE 0 
    	END AS total
	FROM 
		USERS u
	LEFT JOIN
		ORDERS o ON o.tailor_id = u.user_id
    LEFT JOIN
		PRODUCTS p ON p.created_by = o.order_id
    WHERE u.role = ?
    GROUP BY user_id;
	`

	rows, err := u.db.QueryContext(ctx, query, req.Role)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*responses.UserTailor
	for rows.Next() {
		var user responses.UserTailor
		if err := rows.Scan(
			&user.User_id,
			&user.Username,
			&user.Display_name,
			&user.User_profile_url,
			&user.Role,
			&user.Phone_number,
			&user.Address,
			&user.Timestamp,
			&user.Product_Process,
			&user.Product_Total,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *UserMySQL) GetByUsername(ctx context.Context, req *requests.Username) error {

	query := `SELECT username FROM USERS WHERE username = ?`
	var user responses.Username

	err := u.db.GetContext(ctx, &user, query, req.Username)
	switch err {
	case sql.ErrNoRows: // user found
		return nil
	case nil:
		return exceptions.ErrUsernameDuplicated
	default:
		return err
	}
}

func (u *UserMySQL) GetPasswordByUsername(ctx context.Context, req *requests.Username) (*responses.UserLogin, error) {

	query := `SELECT user_id,password FROM USERS WHERE username = ?`

	var user responses.UserLogin

	err := u.db.GetContext(ctx, &user, query, req.Username)
	if err != nil {
		return nil, exceptions.ErrWrongUsername
	}

	return &user, nil
}

func (u *UserMySQL) GetUserByUserID(ctx context.Context, req *requests.UserID) (*responses.User, error) {

	if err := utils.CheckUserByID(u.db, ctx, req.User_id); err != nil {
		return nil, err
	}

	query := `
	SELECT 
		user_id,
		username,
		display_name,
		user_profile_url,
		role,
		phone_number,
		address,
		timestamp 
	FROM USERS WHERE user_id = ?
	`

	var user responses.User

	err := u.db.GetContext(ctx, &user, query, req.User_id)

	if err != nil {
		return &user, nil
	}

	return &user, nil
}

func (u *UserMySQL) UpdateAddress(ctx context.Context, req *requests.UserUpdate) error {

	if err := utils.CheckUserByID(u.db, ctx, req.Token); err != nil {
		return err
	}

	query := `
	UPDATE USERS 
	SET 
		display_name = ?, 
		phone_number = ?, 
		address = ? 
	WHERE user_id = ?
	`

	_, err := u.db.ExecContext(ctx, query, req.Display_name, req.Phone_number, req.Address, req.Token)

	if err != nil {
		return err
	}

	return err
}

func (u *UserMySQL) UpdateImage(ctx context.Context, req *requests.UserUploadImage) error {

	if err := utils.CheckUserByID(u.db, ctx, req.Token); err != nil {
		return err
	}

	query := `
	UPDATE USERS 
	SET 
		user_profile_url = ? 
	WHERE user_id = ?
	`

	_, err := u.db.ExecContext(ctx, query, req.Image, req.Token)
	if err != nil {
		return err
	}

	return err
}
