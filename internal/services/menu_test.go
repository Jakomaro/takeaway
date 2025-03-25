package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jakomaro/takeaway/internal/repository"
	"github.com/jakomaro/takeaway/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupTMPDatabase(t *testing.T, ctx context.Context, db *sqlx.DB) {
	t.Helper()

	createSchema := `
CREATE SCHEMA IF NOT EXISTS testdata;`
	_, err := db.ExecContext(ctx, createSchema)
	require.NoError(t, err)

	createMenu := `
CREATE TABLE IF NOT EXISTS testdata.menu (
  item_id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  price NUMERIC(10,2) NOT NULL
);`
	_, err = db.ExecContext(ctx, createMenu)
	require.NoError(t, err)

	insertMenu := `
INSERT INTO testdata.menu (
  name, price
) 
VALUES 
  ('focaccia', 125.00),
  ('biancaneve', 125.50),
  ('margherita', 126.5);
`
	_, err = db.ExecContext(ctx, insertMenu)
	require.NoError(t, err)

}

func teardownTMPDatabase(t *testing.T, ctx context.Context, db *sqlx.DB) {
	t.Helper()

	dropMenu := `DROP TABLE IF EXISTS testdata.menu;`
	_, err := db.ExecContext(ctx, dropMenu)
	require.NoError(t, err)

	dropSchema := `DROP SCHEMA IF EXISTS testdata;`
	_, err = db.ExecContext(ctx, dropSchema)
	require.NoError(t, err)
}

func TestPGMenuService(t *testing.T) {

	connString := "postgresql://postgres:postgres@localhost:5432/takeaway?sslmode=disable"
	db, err := repository.NewPostgresDB(connString)
	require.NoError(t, err)

	ctx := t.Context()

	schemaValue := "testdata"
	ctx = context.WithValue(ctx, "schemaID", schemaValue)

	setupTMPDatabase(t, ctx, db)
	defer teardownTMPDatabase(t, ctx, db)

	pgMenuService := services.NewPGMenuService(db)
	menu, err := pgMenuService.GetMenu(ctx)
	require.NoError(t, err)

	fmt.Println(menu)
}
