package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/reservation-service/genprotos"

	"github.com/google/uuid"
)

type Menustorage struct {
	db *sql.DB
}

func NewMenustorage(db *sql.DB) *Menustorage {
	return &Menustorage{db: db}
}

func (p *Menustorage) CreateMenu(menu *pb.Menu) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO menus (id, restaurant_id, name, description, price)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, menu.RestaurantId, menu.Name, menu.Description, menu.Price)
	return nil, err
}

func (p *Menustorage) GetByIdMenu(id *pb.ById) (*pb.Menu, error) {
	query := `
			SELECT restaurant_id, name, description, price from menus 
			where id =$1 and delated_at=0
		`
	row := p.db.QueryRow(query, id.Id)

	var Menu pb.Menu

	err := row.Scan(&Menu.RestaurantId, &Menu.Name, &Menu.Description, &Menu.Price)
	if err != nil {
		return nil, err
	}

	return &Menu, nil
}

func (p *Menustorage) GetAllMenu(m *pb.Menu) (*pb.GetAllMenus, error) {
	var arr []interface{}
	count := 1
	query := `
			SELECT restaurant_id, name, description, price from menus 
			where delated_at=0 `

	if len(m.Name) > 0 {
		query += fmt.Sprintf(" and name=$%d", count)
		count++
		arr = append(arr, m.Name)
	}
	if m.Price != 0 {
		query += fmt.Sprintf(" and user_name=$%d", count)
		count++
		arr = append(arr, m.Price)
	}
	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var Menus pb.GetAllMenus
	for rows.Next() {
		var Menu pb.Menu
		err := rows.Scan(&Menu.RestaurantId, &Menu.Name, &Menu.Description, &Menu.Price)
		if err != nil {
			return nil, err
		}
		Menus.Menus = append(Menus.Menus, &Menu)
	}
	return &Menus, nil
}

func (p *Menustorage) UpdateMenu(menu *pb.Menu) (*pb.Void, error) {
	query := `
		UPDATE menus
		SET restaurant_id = $1, name = $2, description = $3, price = $4
		WHERE id = $5
	`
	_, err := p.db.Exec(query, menu.RestaurantId, menu.Name, menu.Description, menu.Price, menu.Id)
	return nil, err
}

func (p *Menustorage) DeleteMenu(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE menus
		SET delated_at = $1
		WHERE id = $2
	`
	_, err := p.db.Exec(query, time.Now().Unix(), id.Id)
	return nil, err
}
