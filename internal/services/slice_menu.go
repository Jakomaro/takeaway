package services

import (
	"context"

	"github.com/jakomaro/takeaway/internal/model"
)

type SMenuService struct {
	Menu []model.Item
}

func NewSMenuService() *SMenuService {
	return &SMenuService{
		Menu: []model.Item{
			{ItemID: 1, Name: "focaccia", Price: 5},
			{ItemID: 2, Name: "biancaneve", Price: 5.5},
			{ItemID: 3, Name: "margherita", Price: 6.5},
		},
	}
}

func (ms *SMenuService) GetMenu(ctx context.Context) ([]model.Item, error) {
	return ms.Menu, nil
}
