package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/payment-service/genprotos"

	"github.com/google/uuid"
)

type PaymentStorage struct {
	db *sql.DB
}

func NewPaymentStorage(db *sql.DB) *PaymentStorage {
	return &PaymentStorage{db: db}
}

func (p *PaymentStorage) CreatePayment(payment *pb.Payment) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO payments (id, reservation_id, amount, payment_method, payment_status)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, payment.ReservationId, payment.Amount, payment.PaymentMethod, 
		payment.PaymentStatus)
	return nil, err
}

func (p *PaymentStorage) GetByIdPayment(id *pb.ById) (*pb.Payment, error) {
	query := `
			SELECT reservation_id, amount, payment_method, payment_status from payments 
			where id =$1 and delated_at=0
		`
	row := p.db.QueryRow(query, id.Id)

	var Payment pb.Payment
	err := row.Scan(&Payment.ReservationId, &Payment.Amount, &Payment.PaymentMethod, &Payment.PaymentStatus)
	if err != nil {
		return nil, err
	}

	return &Payment, nil
}


func (p *PaymentStorage) GetAllPayments(rest *pb.Payment) (*pb.GetAllPayments, error) {
	Payment := &pb.GetAllPayments{}
	var query string
	query = ` SELECT reservation_id, amount, payment_method, payment_status from payments 
			where delated_at=0 `
	var arr []interface{}
	count := 1
	if len(rest.ReservationId) > 0 {
		query += fmt.Sprintf(" and name=$%d", count)
		count++
		arr = append(arr, rest.Amount)
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var r pb.Payment
		err = row.Scan(&r.ReservationId, &r.Amount, &r.PaymentMethod, &r.PaymentStatus)
		if err != nil {
			return nil, err
		}
		Payment.Payments = append(Payment.Payments, &r)
	}
	return Payment, nil
}

func (p *PaymentStorage) UpdatePayment(Payment *pb.Payment) (*pb.Void, error) {
	query := `
		UPDATE payments
		SET reservation_id = $1, amount = $2, payment_method = $3, payment_status = $4
		WHERE id = $5
	`
	_, err := p.db.Exec(query, Payment.ReservationId, Payment.Amount, Payment.PaymentMethod, Payment.PaymentStatus)
	return nil, err
}

func (p *PaymentStorage) DeletePayment(id *pb.ById) (*pb.Void, error) {
	query := `
		update from payments set delated_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}





