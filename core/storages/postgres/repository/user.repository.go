package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	AddressType = "Home"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Verified  bool      `db:"verified"`
	Profile   Profile   `db:"profile"`
	Address   Address   `db:"address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Profile struct {
	Username string `db:"username"`
	DOB      string `db:"dob"`
	Phone    string `db:"phone"`
}

type Address struct {
	Type    string `db:"type"`
	ZipCode string `db:"zip_code"`
	Street  string `db:"street"`
	City    string `db:"city"`
	State   string `db:"state"`
	Country string `db:"country"`
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) (uuid.UUID, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *sqlx.DB
	// User User
}

func NewUserRepository(db *sqlx.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *User) (uuid.UUID, error) {
	query := `INSERT INTO users (id, firstname, lastname, address.state, address.type, address.zip_code, address.street, verified, address.city, address.country, profile.username, email, password, profile.dob, profile.phone, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING id`

	row := r.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Firstname,
		user.Lastname,
		user.Address.State,
		user.Address.Type,
		user.Address.ZipCode,
		user.Address.Street,
		user.Verified,
		user.Address.City,
		user.Address.Country,
		user.Profile.Username,
		user.Email, user.Password,
		user.Profile.DOB,
		user.Profile.Phone,
		time.Now(),
		time.Now())
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to save user details: %w", err)
	}
	return id, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `SELECT id, firstname, lastname, address.state, address.type, address.zip_code, address.street, verified, address.city, address.country, profile.username, email, password, profile.dob, profile.phone, created_at, updated_at FROM users WHERE id = $1`
	var user User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, firstname, lastname, address.state, address.type, address.zip_code, address.street, verified, address.city, address.country, profile.username, email, password, profile.dob, profile.phone, created_at, updated_at FROM users WHERE id = $1`
	var user User
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *User) error {
	query := `UPDATE users SET address.state=$1, address.type=$2, address.zip_code=$3, address.street=$4, verified=$5, address.city=$6, profile.username=$7, profile.dob=$8, profile.phone=$9, user.address.country=$10, updated_at=$11 WHERE id=$5`
	_, err := r.db.ExecContext(ctx, query, user.Address.State, user.Address.Type, user.Address.ZipCode, user.Address.Street, user.Verified, user.Address.City, user.Profile.Username, user.Profile.DOB, user.Profile.Phone, user.Address.Country, time.Now(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
