package validator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_check(t *testing.T) {

	type DummyReq struct {
		Name        string `validate:"required" json:"name"`
		Description string `validate:"required" json:"description"`
	}

	type TestCase struct {
		Name         string
		DataNotValid bool
		ReqBody      string
	}

	cases := []TestCase{
		{
			Name:         "When name not present",
			DataNotValid: true,
			ReqBody:      `{"description": "foobar"}`,
		},
		{
			Name:         "When description not present",
			DataNotValid: true,
			ReqBody:      `{"name": "foobar"}`,
		},
		{
			Name:         "When name and description are present",
			DataNotValid: false,
			ReqBody:      `{"name": "foo", "description": "bar"}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var req DummyReq
			_ = json.Unmarshal([]byte(tc.ReqBody), &req)
			isErr := Check(&req)
			assert.Equal(t, tc.DataNotValid, isErr)
		})
	}

}
