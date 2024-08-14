package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	pb "auth/genprotos"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

const emailRegex = `^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func isValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
func (u *UserStorage) RegisterUser(user *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	if !isValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4)`
	_, err = u.db.Exec(query, user.Email, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterUserResponse{}, nil
}

func (u *UserStorage) LoginUser(user *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	query := `
	SELECT id, email, password_hash, first_name, last_name FROM users WHERE email = $1 AND deleted_at = 0
	`
	row := u.db.QueryRow(query, user.Email)
	res := pb.LoginUserResponse{}
	err := row.Scan(
		&res.Id,
		&res.Email,
		&res.PasswordHash,
		&res.FirstName,
		&res.LastName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	return &res, nil
}

func (u *UserStorage) GetByIdUser(id *pb.GetByIdUserRequest) (*pb.GetByIdUserResponse, error) {
	query := `
		SELECT id, email, first_name, last_name FROM users 
		WHERE id = $1 AND deleted_at = 0
	`
	row := u.db.QueryRow(query, id.Id)

	user := pb.GetByIdUserResponse{}
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserStorage) UpdateUser(req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    query := `UPDATE users SET `
    var condition []string
    var args []interface{}

    if req.Email != "" && req.Email != "string" {
        condition = append(condition, fmt.Sprintf("email = $%d", len(args)+1))
        args = append(args, req.Email)
    }
    if req.FirstName != "" && req.FirstName != "string" {
        condition = append(condition, fmt.Sprintf("first_name = $%d", len(args)+1))
        args = append(args, req.FirstName)
    }

	if req.LastName != "" && req.LastName != "string" {
		condition = append(condition, fmt.Sprintf("last_name = $%d", len(args)+1))
        args = append(args, req.LastName)
	}

    if len(condition) == 0 {
        return nil, errors.New("nothing to update")
    }

    query += strings.Join(condition, ", ")
    query += fmt.Sprintf(" WHERE id = $%d RETURNING id, email, first_name, last_name", len(args)+1)
    args = append(args, req.Id)

    res := pb.UpdateUserResponse{}
    row := u.db.QueryRow(query, args...)

    err := row.Scan(&res.Id, &res.Email, &res.FirstName, &res.LastName)
    if err != nil {
        return nil, err
    }
    return &res, nil
}

func (u *UserStorage) DeleteUser(id *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	query := `
		UPDATE users
		SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := u.db.Exec(query, id.Id, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{}, nil
}

func (u *UserStorage) ChangePassword(password *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	// Fetch the current hashed password from the database
	var currentHashedPassword string
	query := `
		SELECT password_hash
		FROM users
		WHERE id = $1 AND deleted_at = 0
	`
	err := u.db.QueryRow(query, password.Id).Scan(&currentHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	// Compare the current password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(currentHashedPassword), []byte(password.CurrentPassword))
	if err != nil {
		return nil, fmt.Errorf("invalid current password")
	}

	// Hash the new password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash new password")
	}

	// Update the password in the database
	updateQuery := `
		UPDATE users
		SET password_hash = $2, updated_at = $3
		WHERE id = $1 AND deleted_at = 0
	`
	_, err = u.db.Exec(updateQuery, password.Id, hashedNewPassword, time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{}, nil
}

func (p *UserStorage) ForgotPassword(forgotPass *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	return &pb.ForgotPasswordResponse{}, nil
}

func (s *UserStorage) GetUserByEmail(email string) (*pb.UpdateUserResponse, error) {
	var user pb.UpdateUserResponse
	query := "SELECT id, email, first_name, last_name FROM users WHERE email = $1"
	row := s.db.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}


func (p *UserStorage) ResetPassword(resetPass *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPass.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE users
		SET password_hash = $2, updated_at = $3
		WHERE id = $1 AND deleted_at = 0
	`
	_, err = p.db.Exec(query, resetPass.Id, string(hashedPassword), time.Now())
	if err != nil {
		return nil, err
	}

	return &pb.ResetPasswordResponse{}, nil
}