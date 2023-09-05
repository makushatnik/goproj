package user

import (
	"goproj/internal/entity"
	"goproj/internal/generated/models"
	"goproj/internal/generated/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	HandleGetUserByEmail(params operations.GetUserEmailParams) middleware.Responder
	HandleGetUserList(_ operations.GetUserParams) middleware.Responder
	HandlePostUser(params operations.PostUserParams) middleware.Responder
}

type handlers struct {
	storage map[string]entity.User
}

func New() Handlers {
	return &handlers{
		storage: make(map[string]entity.User),
	}
}

func (h *handlers) HandleGetUserByEmail(params operations.GetUserEmailParams) middleware.Responder {
	user, ok := h.storage[params.Email]
	if !ok {
		errorMsg := "user doesn't exist"
		return operations.NewGetUserEmailInternalServerError().WithPayload(
			&models.FailResponse{Error: &errorMsg},
		)
	}

	return operations.NewGetUserEmailOK().WithPayload(&models.User{
		Email:    &user.Email,
		FullName: &user.FullName,
		IsAdmin:  &user.IsAdmin,
	})
}

func (h *handlers) HandleGetUserList(_ operations.GetUserParams) middleware.Responder {
	result := make([]*models.User, 0, len(h.storage))
	for _, user := range h.storage {
		var userModel models.User
		userModel.Email = &user.Email
		userModel.FullName = &user.FullName
		userModel.IsAdmin = &user.IsAdmin

		result = append(result, &userModel)
	}

	return operations.NewGetUserOK().WithPayload(result)
}

func (h *handlers) HandlePostUser(params operations.PostUserParams) middleware.Responder {
	user := *params.User
	if *user.Payload.Email == "" && *user.Payload.FullName == "" {
		errorMsg := "empty params"
		return operations.NewPostUserBadRequest().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	if _, ok := h.storage[*user.Payload.Email]; ok {
		errorMsg := "user already exist"
		return operations.NewPostUserInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMsg,
		})
	}

	h.storage[*user.Payload.Email] = entity.User{
		Email:    *user.Payload.Email,
		FullName: *user.Payload.FullName,
		IsAdmin:  *user.Payload.IsAdmin,
	}

	return operations.NewPostUserCreated()
}
