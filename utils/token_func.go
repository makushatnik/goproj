package utils

import (
	"errors"
	"fmt"
	"strings"
)

type Extractor interface {
	Extract(value string) (string, error)
}

type bearerExtractor struct {
	schema string
}

const INVALID_TOKEN = "invalid token"

var ErrExtractToken = errors.New("failed extract token")

func validate(bearer string, extractor Extractor) error {
	token, err := extractor.Extract(bearer)
	if err != nil {
		return err
	}

	if token == INVALID_TOKEN {
		return errors.New(INVALID_TOKEN)
	}

	return nil
}

func (ce bearerExtractor) Extract(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("%w: bearer is empty", ErrExtractToken)
	}

	bearerFields := strings.Fields(value)
	if len(bearerFields) == 0 || !strings.EqualFold(bearerFields[0], ce.schema) {
		return "", fmt.Errorf("%w: invalid bearer", ErrExtractToken)
	}

	tokenWithSpaces := value[len(ce.schema):]
	token := strings.TrimSpace(tokenWithSpaces)
	if len(token) == 0 {
		return "", fmt.Errorf("%w: token is empty", ErrExtractToken)
	}

	return token, nil
}
