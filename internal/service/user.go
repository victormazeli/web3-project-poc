package service

import (
	"context"
	"fmt"
	"goApiStartetProject/internal/storages/postgres/repository"
	"goApiStartetProject/internal/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	// "goApiStartetProject/internal/domain"
)

// UserService is an interface for interacting with user-related business logic
type UserServiceInterface interface {
	// Register registers a new user
	Register(ctx context.Context, user domain.NewUserRequestPayload) (uuid.UUID, error)
	// GetUser returns a user by id
	GetUser(ctx context.Context, id uint64) (*repository.User, error)
	// ListUsers returns a list of users with pagination
	ListUsers(ctx context.Context, skip, limit uint64) ([]repository.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user domain.NewUserRequestPayload) (*repository.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uint64) error
}

/**
 * UserService implements port.tUserService interface
 * and provides an access to the user repository
 * and cache service
 */
type UserService struct {
	repo *repository.Repository
}

func NewUserService(db *sqlx.DB) UserServiceInterface {
	return &UserService{
		repo: repository.NewRepository(db),
	}
}

// DeleteUser implements UserServiceInterface.
func (u *UserService) DeleteUser(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// GetUser implements UserServiceInterface.
func (u *UserService) GetUser(ctx context.Context, id uint64) (*repository.User, error) {
	panic("unimplemented")
}

// ListUsers implements UserServiceInterface.
func (u *UserService) ListUsers(ctx context.Context, skip uint64, limit uint64) ([]repository.User, error) {
	panic("unimplemented")
}

// Register implements UserServiceInterface.
func (u *UserService) Register(ctx context.Context, user domain.NewUserRequestPayload) (uuid.UUID, error) {
	userPayload := repository.User{
		ID:        uuid.New(),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		MiddleName: user.MiddleName,
		Email:     user.Email,
		Password:  user.Password,
		Verified:  false,
		Profile: repository.Profile{
			DOB:      user.DOB,
			Phone:    user.CountryCode + user.Phone,
			Username: user.Username,
		},
		Address: repository.Address{
			Street:  user.Street,
			ZipCode: user.ZipCode,
			State:   user.State,
			City:    user.City,
			Country: user.Country,
			Type:    user.AddressType,
		},
	}

	fmt.Println("Creating")
	respPayload, err := u.repo.UserRepo.CreateUser(ctx, &userPayload)
	if err != nil {
		return uuid.UUID{}, err
	}

	return respPayload, nil
}

// UpdateUser implements UserServiceInterface.
func (u *UserService) UpdateUser(ctx context.Context, user domain.NewUserRequestPayload) (*repository.User, error) {
	panic("unimplemented")
}
