package main

import (
	"context"
	"go-grpc1/models"
)

var localStorage *models.UserList

func init() {
	localStorage = new(models.UserList)

	localStorage.List = make([]*models.User, 0)
}

type UserServer struct {
}

func (s *UserServer) Register(ctx context.Context, param *models.User) {

}
