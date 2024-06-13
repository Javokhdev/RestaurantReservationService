package postgres

import (
	"database/sql"
	"fmt"

	"github.com/dilshodforever/reservation-service/config"
	st "github.com/dilshodforever/reservation-service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db                *sql.DB
	Restorans         st.Restoran
	Reservations      st.Reservation
	ReservationOrders st.ReservationOrder
	Menus             st.Menu
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

	return &Storage{Db: db, Restorans: &Restorantorage{db}, Reservations: &Reservationstorage{db}, ReservationOrders: &ReservationOrderstorage{db}, Menus: &Menustorage{db}}, nil

}

func (s *Storage) Restoran() st.Restoran {
	if s.Restorans == nil {
		s.Restorans = &Restorantorage{s.Db}
	}
	return s.Restorans
}

func (s *Storage) Reservation() st.Reservation {
	if s.Reservations == nil {
		s.Reservations = &Reservationstorage{s.Db}
	}
	return s.Reservations
}

func (s *Storage) ReservationOrder() st.ReservationOrder {
	if s.ReservationOrders == nil {
		s.ReservationOrders = &ReservationOrderstorage{s.Db}
	}
	return s.ReservationOrders
}

func (s *Storage) Menu() st.Menu {
	if s.Menus == nil {
		s.Menus = &Menustorage{s.Db}
	}
	return s.Menus
}
