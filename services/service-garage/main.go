package servicegarage

import (
	"context"
	"log"

	"go-grpc1/models"
	"go-grpc1/services"

	"github.com/golang/protobuf/ptypes/empty"
)

var localStorage *models.GarageListByUser

func init() {
	localStorage = new(models.GarageListByUser)
	localStorage.List = make(map[string]*models.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *services.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(models.GarageList)
		localStorage.List[userId].List = make([]*models.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(empty.Empty), nil
}

func (GaragesServer) List(ctx context.Context, param *models.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}
