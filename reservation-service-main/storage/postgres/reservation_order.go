package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/reservation-service/genprotos"

	"github.com/google/uuid"
)

type ReservationOrderstorage struct {
	db *sql.DB
}

func NewReservationOrderstorage(db *sql.DB) *ReservationOrderstorage {
	return &ReservationOrderstorage{db: db}
}

func (p *ReservationOrderstorage) CreateReservationOrder(reservationOrder *pb.ReservationOrder) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO reservation_orders (id, reservation_id, menu_item_id, quantity)
		VALUES ($1, $2, $3, $4)
	`
	_, err := p.db.Exec(query, id, reservationOrder.ReservationId, reservationOrder.MenuItemId, reservationOrder.Quantity)
	return nil, err
}

func (p *ReservationOrderstorage) GetByIdReservationOrder(id *pb.ById) (*pb.ReservationOrder, error) {
	query := `
			SELECT reservation_id, menu_item_id, quantity from reservation_orders 
			where id =$1 and delated_at=0
		`
	row := p.db.QueryRow(query, id.Id)

	var ReservationOrder pb.ReservationOrder

	err := row.Scan(&ReservationOrder.ReservationId, &ReservationOrder.MenuItemId, &ReservationOrder.Quantity)
	if err != nil {
		return nil, err
	}

	return &ReservationOrder, nil
}

func (p *ReservationOrderstorage) GetAllReservationOrder(reservationOrder *pb.ReservationOrder) (*pb.GetAllReservationOrders, error) {
	ReservationOrders := &pb.GetAllReservationOrders{}
	var query string
	query = ` SELECT reservation_id, menu_item_id, quantity from reservation_orders 
			where delated_at=0 `
	var arr []interface{}
	count := 1
	if len(reservationOrder.Quantity) > 0 {
		query += fmt.Sprintf(" and quantity=$%d", count)
		count++
		arr = append(arr, reservationOrder.Quantity)
	}
	if len(reservationOrder.MenuItemId) > 0 {
		query += fmt.Sprintf(" and menu_item_id=$%d", count)
		count++
		arr = append(arr, reservationOrder.MenuItemId)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err

	}
	for row.Next() {
		var ReservationOrder pb.ReservationOrder
		err := row.Scan(&ReservationOrder.ReservationId, &ReservationOrder.MenuItemId, &ReservationOrder.Quantity)
		if err != nil {
			return nil, err
		}
		ReservationOrders.ReservationOrders = append(ReservationOrders.ReservationOrders, &ReservationOrder)
	}
	return ReservationOrders, nil
}

func (p *ReservationOrderstorage) UpdateReservationOrder(reservationOrder *pb.ReservationOrder) (*pb.Void, error) {
	query := `
		UPDATE reservation_orders
		SET reservation_id = $1, menu_item_id = $2, quantity = $3
		WHERE id = $4
	`
	_, err := p.db.Exec(query, reservationOrder.ReservationId, reservationOrder.MenuItemId, reservationOrder.Quantity, reservationOrder.Id)
	return nil, err
}

func (p *ReservationOrderstorage) DeleteReservationOrder(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE reservation_orders
		SET delated_at = $1
		WHERE id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
