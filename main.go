package main

import (
	"fmt"
	"learn-protobuf/model"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Anugrah Wardhani",
		Password: "Password123",
		Gender:   model.UserGender_MALE,
	}

	_ = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	_ = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	jsonb, err1 := protojson.Marshal(garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	protoObject := new(model.GarageList)
	err2 := protojson.Unmarshal(jsonb, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# === Original \n         %#v  \n", user1)
	fmt.Printf("# === As String\n   %s  \n", user1.String())
	fmt.Printf("# === As JSON String \n       %s    \n", string(jsonb))
	fmt.Printf("# === As String \n     %s   \n", protoObject.String())
}
