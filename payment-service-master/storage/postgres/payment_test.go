package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/payment-service/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	payment := &pb.Payment{
		ReservationId: "UUID",
		Amount:        15,
		PaymentMethod: "Online",
		PaymentStatus: "Payed",
	}

	result, err := stg.Payment().CreatePayment(payment)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetByIdUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.Payment().GetByIdPayment(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	users, err := stg.Payment().GetAllPayments(&pb.Payment{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUpdateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	payment := &pb.Payment{
		ReservationId: "UUID",
		Amount:        15,
		PaymentMethod: "Online",
		PaymentStatus: "Payed",
	}

	result, err := stg.Payment().UpdatePayment(payment)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Payment().DeletePayment(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
