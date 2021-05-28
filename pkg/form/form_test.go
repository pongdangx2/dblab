package form_test

import (
	"testing"

	"github.com/danvergara/dblab/pkg/command"
	"github.com/danvergara/dblab/pkg/form"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var cases = []struct {
		name   string
		given  command.Options
		wanted bool
	}{
		{
			name:   "total empty",
			given:  command.Options{},
			wanted: true,
		},
		{
			name: "ignoring ssl",
			given: command.Options{
				SSL: "disable",
			},
			wanted: true,
		},
		{
			name: "not empty",
			given: command.Options{
				Driver: "postgres",
				Host:   "localhost",
				Port:   "5432",
				User:   "user",
				Pass:   "password",
				DBName: "users",
			},
			wanted: false,
		},
		{
			name: "not empty with url",
			given: command.Options{
				URL: "postgres://user:password@host:port/db?sslmode=mode",
			},
			wanted: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := form.IsEmpty(tc.given)

			assert.Equal(t, tc.wanted, actual)
		})
	}
}