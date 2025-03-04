package tests_test

import (
	"encoding/json"
	"errors"
	"testing"

	testMaxLength "github.com/atombender/go-jsonschema/tests/data/validation/maxLength"
	testMinLength "github.com/atombender/go-jsonschema/tests/data/validation/minLength"
	"github.com/atombender/go-jsonschema/tests/helpers"
)

func TestMaxStringLength(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
	}{
		{
			desc:    "no violations",
			data:    `{"myString": "hi"}`,
			wantErr: nil,
		},
		{
			desc:    "myString too long",
			data:    `{"myString": "hello"}`,
			wantErr: errors.New("field myString length: must be <= 5"),
		},
		{
			desc:    "myString not present",
			data:    `{}`,
			wantErr: errors.New("field myString in MaxLength: required"),
		},
		{
			desc:    "myNullableString too long",
			data:    `{"myString": "hi","myNullableString": "hello world"}`,
			wantErr: errors.New("field myNullableString length: must be <= 10"),
		},
		{
			desc:    "myString and myNullableString too long",
			data:    `{"myString": "hello","myNullableString": "hello world"}`,
			wantErr: errors.New("field myNullableString length: must be <= 10"),
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testMaxLength.MaxLength{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}

func TestMinStringLength(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		data    string
		wantErr error
	}{
		{
			desc:    "no violations",
			data:    `{"myString": "hello"}`,
			wantErr: nil,
		},
		{
			desc:    "myString too short",
			data:    `{"myString": "hi"}`,
			wantErr: errors.New("field myString length: must be >= 5"),
		},
		{
			desc:    "myString not present",
			data:    `{}`,
			wantErr: errors.New("field myString in MinLength: required"),
		},
		{
			desc:    "myNullableString too short",
			data:    `{"myString": "hello","myNullableString": "hi"}`,
			wantErr: errors.New("field myNullableString length: must be >= 10"),
		},
		{
			desc:    "myString and myNullableString too short",
			data:    `{"myString": "hi","myNullableString": "hello"}`,
			wantErr: errors.New("field myNullableString length: must be >= 10"),
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			model := testMinLength.MinLength{}

			err := json.Unmarshal([]byte(tC.data), &model)

			helpers.CheckError(t, tC.wantErr, err)
		})
	}
}
