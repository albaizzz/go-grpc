package main

import (
	"go-grpc/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// var user = &models.User{
	// 	Id:       "asd123",
	// 	Name:     "Albai",
	// 	Password: "123123",
	// 	Gender:   models.UserGender_MALE,
	// }

	// var userList = &models.UserList{
	// 	List: []*models.User{
	// 		user,
	// 	},
	// }

	// var garage = &models.Garage{
	// 	Id:   "g001",
	// 	Name: "Kalimdor",
	// 	Coordinate: &models.GarageCoordinate{
	// 		Latitude:  23.2212847,
	// 		Longitude: 53.22033123,
	// 	},
	// }

	// var garageList = &models.GarageList{
	// 	List: []*models.Garage{
	// 		garage,
	// 	},
	// }

	// fmt.Println(userList)
	// fmt.Println(garageList)

	// var buf bytes.Buffer
	// err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	// if err1 != nil {
	// 	fmt.Println(err1.Error())
	// 	os.Exit(0)
	// }
	// jsonString := buf.String()
	// fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)

	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
