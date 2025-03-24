package services

import "github.com/jakomaro/takeaway/model"

type MenuServicer interface {
	GetMenu() []model.Item
}

type MenuService struct {
	Menu []model.Item
}

func (ms *MenuService) GetMenu() []model.Item {
	return ms.Menu
}
