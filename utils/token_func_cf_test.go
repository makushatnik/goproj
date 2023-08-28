package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBearerExtractCF(t *testing.T) {
	req := require.New(t)

	t.Run("success", func(t *testing.T) {
		bearer := "Bearer fgbdgsefsd"
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract(bearer)

		req.NoError(err)
		req.Contains(bearer, token)
	})

	t.Run("lowercase", func(t *testing.T) {
		bearer := "bearer fgbdgsefsd"
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract(bearer)

		req.NoError(err)
		req.Contains(bearer, token)
	})

	t.Run("with spaces", func(t *testing.T) {
		bearer := "Bearer            fgbdgsefsd "
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract(bearer)

		req.NoError(err)
		req.Contains(bearer, token)
	})

	t.Run("different cases", func(t *testing.T) {
		bearer := "BeARer    fgbdgsefsd"
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract(bearer)

		req.NoError(err)
		req.Contains(bearer, token)
	})

	t.Run("empty bearer", func(t *testing.T) {
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract("")

		req.Error(err)
		req.Contains(err.Error(), "bearer is empty")
		req.Equal("", token)
	})

	t.Run("invalid bearer", func(t *testing.T) {
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract("   fgbdgsefsd")

		req.Error(err)
		req.Contains(err.Error(), "invalid bearer")
		req.Equal("", token)
	})

	t.Run("empty token", func(t *testing.T) {
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract("Bearer ")

		req.NoError(err)
		req.Contains(err.Error(), "token is empty")
		req.Equal("", token)
	})

	t.Run("invalid token", func(t *testing.T) {
		bearerExtractor := &bearerExtractor{schema: "bearer"}
		token, err := bearerExtractor.Extract("Bearer 	   ")

		req.NoError(err)
		req.Contains(err.Error(), "token is empty")
		req.Equal("", token)
	})
}
