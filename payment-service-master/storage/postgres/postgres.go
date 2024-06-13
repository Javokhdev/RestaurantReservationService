package postgres

import (
	"database/sql"
	"fmt"
	"github.com/dilshodforever/payment-service/config"
	st "github.com/dilshodforever/payment-service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	Payments 	st.Payment
}

func NewPostgresStorage() (st.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db, Payments: &PaymentStorage{db}}, nil

}

func (s *Storage) Payment() st.Payment{
	if s.Payments == nil {
		s.Payments = &PaymentStorage{s.Db}
	}
	return s.Payments
}






