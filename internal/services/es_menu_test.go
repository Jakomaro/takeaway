package services_test

import (
	"fmt"
	"testing"

	"github.com/jakomaro/takeaway/internal/repository"
	"github.com/jakomaro/takeaway/internal/services"
	"github.com/stretchr/testify/require"
)

func TestESMenuService(t *testing.T) {

	es, err := repository.NewElasticDB()
	require.NoError(t, err)

	esMenuService := services.NewESMenuService(es)
	menu, err := esMenuService.GetMenu(t.Context())
	require.NoError(t, err)

	fmt.Println(menu)
}
