package services

import (
	"context"

	"github.com/jakomaro/takeaway/internal/model"
)

type MenuServicer interface {
	GetMenu(ctx context.Context) ([]model.Item, error)
}
