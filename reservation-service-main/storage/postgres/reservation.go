package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/reservation-service/genprotos"

	"github.com/google/uuid"
)

type Reservationstorage struct {
	db *sql.DB
}

func NewReservationstorage(db *sql.DB) *Reservationstorage {
	return &Reservationstorage{db: db}
}

func (p *Reservationstorage) CreateReservation(reservation *pb.Reservation) (*pb.Void, error) {
	cheak, err:=p.CheakBron(reservation.ReservationTime)
	if err!=nil{
		return nil, err
	}
	if cheak{
		return &pb.Void{}, fmt.Errorf("bu joy band")
	}
	query := `
		update  reservations set user_id=$1, restaurant_id=$2, reservation_time=$3, status='busy'
	`
	_, err = p.db.Exec(query, reservation.UserId, reservation.RestaurantId, reservation.ReservationTime)
	if err!=nil{
		panic(err)
	}
	return nil, err
}

func (p *Reservationstorage) GetByIdReservation(id *pb.ById) (*pb.Reservation, error) {
	query := `
			SELECT user_id, restaurant_id, reservation_time, status from reservations 
			where id =$1 and delated_at=0 and status='empty'
		`
	row := p.db.QueryRow(query, id.Id)

	var Reservation pb.Reservation

	err := row.Scan(&Reservation.UserId, &Reservation.RestaurantId, &Reservation.ReservationTime, &Reservation.Status)
	if err != nil {
		return nil, err
	}

	return &Reservation, nil
}

func (p *Reservationstorage) GetAllReservation(res *pb.Reservation) (*pb.GetAllReservations, error) {
	Reservations := &pb.GetAllReservations{}
	var query string
	query = ` SELECT r.user_id,  r.reservation_time,  re.name,
				re.address, re.phone_number, re.description from reservations r
				join restaurants re 
				on r.restaurant_id=re.id
			    where delated_at=0 and status='empty'`
	var arr []interface{}
	count := 1
	if len(res.UserId) > 0 {
		query += fmt.Sprintf(" and user_id=$%d", count)
		count++
		arr = append(arr, res.UserId)
	}
	if len(res.ReservationTime) > 0 {
		query += fmt.Sprintf(" and reservation_time=$%d", count)
		count++
		arr = append(arr, res.ReservationTime)
	}


	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		res:=&pb.Restoran{}
		var Reservation pb.Reservation
		err := row.Scan(&Reservation.UserId,
			&Reservation.ReservationTime, 
			&res.Name, 
			&res.Address, 
			&res.PhoneNumber, 
			&res.Description)
		if err != nil {
			return nil, err
		}
		Reservation.RestaurantId=res
		Reservations.Reservations = append(Reservations.Reservations, &Reservation)
	}


	return Reservations, nil
}
		

func (p *Reservationstorage) UpdateReservation(reservation *pb.Reservation) (*pb.Void, error) {
	query := `
		UPDATE reservations
		SET user_id = $1, restaurant_id = $2, reservation_time = $3, status = $4
		WHERE id = $5
	`
	_, err := p.db.Exec(query, reservation.UserId, reservation.RestaurantId, reservation.ReservationTime, reservation.Status, reservation.Id)
	return nil, err
}
func (p *Reservationstorage) DeleteReservation(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE reservations
		SET delated_at = $1
		WHERE id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}


func (p *Reservationstorage) Reservations(reservation *pb.Reservation) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO reservations (id, reservation_time)
		VALUES ($1, $2)
	`
	_, err := p.db.Exec(query, id, reservation.ReservationTime)
	if err!=nil{
		panic(err)
	}
	return nil, err
}



func (p *Reservationstorage) CheakBron(time string) ( bool, error) {
	query := `
			SELECT  status from reservations 
			where reservation_time =$1 and deleted_at=0 
		`
	row := p.db.QueryRow(query, time)

	var Reservation pb.Reservation

	err := row.Scan(&Reservation.Status)
	if err != nil {
		return false, err
	}
	if Reservation.Status=="busy"{
		return true, nil
	}
	return false, nil
}