package services

import (
	"context"
	"fmt"

	"github.com/jakomaro/takeaway/internal/model"
	"github.com/jmoiron/sqlx"
)

type MenuServicer interface {
	GetMenu(ctx context.Context) ([]model.Item, error)
}

type MenuService struct {
	Menu []model.Item
}

func (ms *MenuService) GetMenu(ctx context.Context) ([]model.Item, error) {
	return ms.Menu, nil
}

/***********************  DB IMPLEMENTATION	***********************/
type PGMenuService struct {
	db *sqlx.DB
}

func NewPGMenuService(db *sqlx.DB) *PGMenuService {
	return &PGMenuService{db: db}
}

func (ms *PGMenuService) GetMenu(ctx context.Context) ([]model.Item, error) {

	query := `SELECT * FROM menu`

	var menu []model.Item
	err := ms.db.SelectContext(ctx, &menu, query)
	if err != nil {
		return nil, fmt.Errorf("error getting menu item. Error: %w", err)
	}
	return menu, nil
}
