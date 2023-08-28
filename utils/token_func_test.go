package utils

import (
	"testing"

	mock_utils "goproj/utils/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBearerExtractSuccess(t *testing.T) {
	req := require.New(t)
	bearer := "Bearer fgbdgsefsd"
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract(bearer)

	req.NoError(err)
	req.Contains(bearer, token)
}

func TestBearerExtractLowerCase(t *testing.T) {
	req := require.New(t)
	bearer := "bearer fgbdgsefsd"
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract(bearer)

	req.NoError(err)
	req.Contains(bearer, token)
}

func TestBearerExtractDifferentCases(t *testing.T) {
	req := require.New(t)
	bearer := "BeARer    fgbdgsefsd"
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract(bearer)

	req.NoError(err)
	req.Contains(bearer, token)
}

func TestBearerExtractEmpty(t *testing.T) {
	req := require.New(t)
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract("")

	req.Error(err)
	req.Contains(err.Error(), "bearer is empty")
	req.Equal("", token)
}

func TestBearerExtractWithSpaces(t *testing.T) {
	req := require.New(t)
	bearer := "Bearer            fgbdgsefsd "
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract(bearer)

	req.NoError(err)
	req.Contains(bearer, token)
}

func TestBearerExtractInvalidBearer(t *testing.T) {
	req := require.New(t)
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract("   fgbdgsefsd")

	req.Error(err)
	req.Contains(err.Error(), "invalid bearer")
	req.Equal("", token)
}

func TestBearerExtractEmptyToken(t *testing.T) {
	req := require.New(t)
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract("Bearer ")

	req.NoError(err)
	req.Contains(err.Error(), "token is empty")
	req.Equal("", token)
}

func TestBearerExtractInvalidToken(t *testing.T) {
	req := require.New(t)
	bearerExtractor := &bearerExtractor{schema: "bearer"}
	token, err := bearerExtractor.Extract("Bearer 	   ")

	req.NoError(err)
	req.Contains(err.Error(), "token is empty")
	req.Equal("", token)
}

func TestValidate(t *testing.T) {
	req := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockExtractor := mock_utils.NewMockExtractor(mockController)
	mockExtractor.EXPECT().Extract(gomock.Any()).Return("token", nil).Times(1)

	err := validate("bearer ssss", mockExtractor)
	req.NoError(err)
}
