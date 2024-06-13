package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	"github.com/stretchr/testify/assert"
)

func TestCreateMenu(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	menu := &pb.Menu{
		Id:           "b409ff53-ff2b-4033-84b4-4ce555081647",
		RestaurantId: "b409ff53-ff2b-4033-84b4-4ce555081647",
		Name:         "Abu Saxiy",
		Description:  "Katta hola",
		Price:        777,
	}

	result, err := stg.Menu().CreateMenu(menu)

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestGetByIdMenu(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.Menu().GetByIdMenu(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllMenu(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	Menus, err := stg.Menu().GetAllMenu(&pb.Menu{})
	assert.NoError(t, err)
	assert.NotNil(t, Menus)
}

func TestUpdateMenu(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	menu := &pb.Menu{
		Id:           "b409ff53-ff2b-4033-84b4-4ce555081647",
		RestaurantId: "b409ff53-ff2b-4033-84b4-4ce555081647",
		Name:         "Abu Saxiy",
		Description:  "Katta hola",
		Price:        777,
	}

	result, err := stg.Menu().UpdateMenu(menu)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteMenu(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Menu().DeleteMenu(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

