package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
	{
		in: User{
			ID:     "123456789012345678901234567890123456",
			Name:   "Jane Doe",
			Age:    17,
			Email:  "janedoe@example.com",
			Role:   "admin",
			Phones: []string{"12345678901", "09876543210"},
		},
		expectedErr: errors.New("Age: must be at least 18"),
	},
	{
		in: User{
			ID:     "123456789012345678901234567890123456",
			Name:   "John Doe",
			Age:    30,
			Email:  "johndoe@example.com",
			Role:   "invalid",
			Phones: []string{"12345678901", "09876543210"},
		},
		expectedErr: errors.New("Role: must be one of [admin stuff]"),
	},
	{
		in: User{
			ID:     "123456789012345678901234567890123456",
			Name:   "John Doe",
			Age:    30,
			Email:  "johndoe@example.com",
			Role:   "admin",
			Phones: []string{"1234567890121", "09876543210"},
		},
		expectedErr: errors.New("Phones: length must be 11"),
	},
	{
		in: App{
			Version: "1.2.3.4",
		},
		expectedErr: errors.New("Version: length must be 5"),
	},
	{
		in: Response{
			Code: 201,
			Body: "Success",
		},
		expectedErr: errors.New("Code: must be one of [200 404 500]"),
	},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
		if tt.expectedErr != nil {
			assert.EqualError(t, err, tt.expectedErr.Error())
		} else {
			assert.NoError(t, err)
		}
		})
	}
}

