package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{"simple data", args{uuid: "example"}, &User{UUID: "example"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewUser(tt.args.uuid), "NewUser(%v)", tt.args.uuid)
		})
	}
}
