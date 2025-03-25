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

func (ms *PGMenuService) GetMenu(ctxWithValue context.Context) ([]model.Item, error) {

	var query = `SELECT * FROM menu`

	schemaAny := ctxWithValue.Value("schemaID")
	if schemaAny != nil {
		schema := schemaAny.(string)
		query = fmt.Sprintf(`SELECT * FROM %v.menu`, schema)
	}

	var menu []model.Item
	err := ms.db.SelectContext(ctxWithValue, &menu, query)
	if err != nil {
		return nil, fmt.Errorf("error getting menu item. Error: %w", err)
	}
	return menu, nil
}
