package user

import (
	"context"
	"errors"
	"goproj/internal/entity"
	"goproj/pkg/userservice"
)

type Handlers interface {
	CreateUser(ctx context.Context, user *userservice.User) (*userservice.UserState, error)
	GetUserList(ctx context.Context, _ *userservice.Empty) (*userservice.UserList, error)
	GetUserByEmail(ctx context.Context, _ *userservice.EmailFilter) (*userservice.User, error)
}

type handlers struct {
	userservice.UnimplementedUserServiceServer
	storage map[string]entity.User
}

func New() *handlers {
	return &handlers{
		storage: make(map[string]entity.User),
	}
}

func (h *handlers) CreateUser(_ context.Context, user *userservice.User) (*userservice.UserState, error) {
	if user.Email == "" && user.FullName == "" {
		return &userservice.UserState{Success: false}, errors.New("empty params")
	}

	if _, ok := h.storage[user.Email]; ok {
		return &userservice.UserState{Success: false}, errors.New("user already exist")
	}

	h.storage[user.Email] = entity.User{
		Email:    user.Email,
		FullName: user.FullName,
		IsAdmin:  user.IsAdmin,
	}

	return &userservice.UserState{Success: true}, nil
}

func (h *handlers) GetUserList(_ context.Context, _ *userservice.Empty) (*userservice.UserList, error) {
	result := make([]*userservice.User, 0, len(h.storage))
	for _, user := range h.storage {
		result = append(result, &userservice.User{
			Email:    user.Email,
			FullName: user.FullName,
			IsAdmin:  user.IsAdmin,
		})
	}

	return &userservice.UserList{List: result}, nil
}

func (h *handlers) GetUserByEmail(_ context.Context, filter *userservice.EmailFilter) (*userservice.User, error) {
	user, ok := h.storage[filter.Email]
	if !ok {
		return nil, errors.New("User doesn't exist")
	}

	return &userservice.User{
		Email:    user.Email,
		FullName: user.FullName,
		IsAdmin:  user.IsAdmin,
	}, nil
}
