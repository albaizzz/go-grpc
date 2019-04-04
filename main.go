package main

import (
	"bytes"
	"fmt"
	"go-grpc1/models"
	"os"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var user = &models.User{
		Id:       "asd123",
		Name:     "Albai",
		Password: "123123",
		Gender:   models.UserGender_MALE,
	}

	var userList = &models.UserList{
		List: []*models.User{
			user,
		},
	}

	var garage = &models.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &models.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &models.GarageList{
		List: []*models.Garage{
			garage,
		},
	}

	fmt.Println(userList)
	fmt.Println(garageList)

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)
}
