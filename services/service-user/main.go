package serviceuser

import (
	"context"
	"go-grpc1/models"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
)

var localStorage *models.UserList

func init() {
	localStorage = new(models.UserList)

	localStorage.List = make([]*models.User, 0)
}

type UserServer struct {
}

func (s *UserServer) Register(ctx context.Context, param *models.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}

func (s *UserServer) List(ctx context.Context, void *empty.Empty) (*models.UserList, error) {
	return localStorage, nil
}
