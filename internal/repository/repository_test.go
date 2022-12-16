package repository

import "testing"

func Test(t *testing.T) {
	NewRepository(Source{
		User:         "app",
		Password:     "password",
		Host:         "localhost",
		Port:         5432,
		DatabaseName: "hackowagonetka",
	})
}
