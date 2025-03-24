package services

import (
	"context"
	"fmt"

	"github.com/jakomaro/takeaway/model"
	"github.com/jmoiron/sqlx"
)

type MenuServicer interface {
	GetMenu() ([]model.Item, error)
}

type MenuService struct {
	Menu []model.Item
}

func (ms *MenuService) GetMenu() ([]model.Item, error) {
	return ms.Menu, nil
}

/***********************  DB IMPLEMENTATION	***********************/
type PGMenuService struct {
	db *sqlx.DB
}

func (ms *PGMenuService) GetMenu() ([]model.Item, error) {

	query := `SELECT * FROM menu`

	var menu []model.Item
	err := ms.db.SelectContext(context.TODO(), &menu, query)
	if err != nil {
		return nil, fmt.Errorf("error getting menu item. Error: %w", err)
	}
	return menu, nil
}
