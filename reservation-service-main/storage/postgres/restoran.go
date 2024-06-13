package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/reservation-service/genprotos"

	"github.com/google/uuid"
)

type Restorantorage struct {
	db *sql.DB
}

func NewRestorantorage(db *sql.DB) *Restorantorage {
	return &Restorantorage{db: db}
}

func (p *Restorantorage) CreateRestoran(restoran *pb.Restoran) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO restaurants (id, name, address, phone_number, description)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, restoran.Name, restoran.Address,
		restoran.PhoneNumber, restoran.Description)
	return nil, err
}

func (p *Restorantorage) GetByIdRestoran(id *pb.ById) (*pb.Restoran, error) {
	query := `
			SELECT name, address, phone_number, description from restaurants 
			where id =$1 and delated_at=0 
		`
	row := p.db.QueryRow(query, id.Id)

	var Restoran pb.Restoran

	err := row.Scan(&Restoran.Name, &Restoran.Address, &Restoran.PhoneNumber, &Restoran.Description)
	if err != nil {
		return nil, err
	}

	return &Restoran, nil
}

func (p *Restorantorage) GetAllRestoran(rest *pb.Restoran) (*pb.GetAllRestorans, error) {
	Restoran := &pb.GetAllRestorans{}
	var query string
	query = ` SELECT name, address, phone_number, description from restaurants 
			where delated_at=0`
	var arr []interface{}
	count := 1
	if len(rest.Name) > 0 {
		query += fmt.Sprintf(" and name=$%d", count)
		count++
		arr = append(arr, rest.Name)
	}
	if len(rest.Address) > 0 {
		query += fmt.Sprintf(" and address=$%d", count)
		count++
		arr = append(arr, rest.Address)
	}
	if len(rest.PhoneNumber) > 0 {
		query += fmt.Sprintf(" and phone_number=$%d", count)
		count++
		arr = append(arr, rest.PhoneNumber)
	}
	if len(rest.Description) > 0 {
		query += fmt.Sprintf(" and description=$%d", count)
		count++
		arr = append(arr, rest.Description)
	}
	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var r pb.Restoran
		err = row.Scan(&r.Name, &r.Address, &r.PhoneNumber, &r.Description)
		if err != nil {
			return nil, err
		}
		Restoran.Restorans = append(Restoran.Restorans, &r)
	}
	return Restoran, nil
}

func (p *Restorantorage) UpdateRestoran(Restoran *pb.Restoran) (*pb.Void, error) {
	query := `
		UPDATE restaurants
		SET name = $1, address = $2, description = $3
		WHERE id = $4
	`
	_, err := p.db.Exec(query, Restoran.Name, Restoran.Address, Restoran.Description, Restoran.Id)
	return nil, err
}

func (p *Restorantorage) DeleteRestoran(id *pb.ById) (*pb.Void, error) {
	query := `
		update from restaurants set delated_at=$1
		where id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
