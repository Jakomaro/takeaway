package services

import (
	"context"

	"github.com/jakomaro/takeaway/internal/model"
)

type MenuService struct {
	Menu []model.Item
}

func (ms *MenuService) GetMenu(ctx context.Context) ([]model.Item, error) {
	return ms.Menu, nil
}
