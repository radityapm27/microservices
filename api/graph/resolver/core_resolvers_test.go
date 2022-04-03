package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"testing"

	"rpm/microservices/api/graph/model"

	"github.com/stretchr/testify/assert"
)

func TestQueryResolver_ShoppingCart(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		r    = new(Resolver)
	)

	t.Run("success", func(t *testing.T) {

		resp := new(model.AbstractModel)
		ctx := context.Background()

		res, err := r.Query().ShoppingCart(ctx)

		test.NotNil(res)
		test.Equal(res, resp)
		test.Nil(err)
	})
}

func TestQueryResolver_ProductCatalog(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		r    = new(Resolver)
	)

	t.Run("success", func(t *testing.T) {

		resp := new(model.AbstractModel)
		ctx := context.Background()

		res, err := r.Query().ProductCatalog(ctx)

		test.NotNil(res)
		test.Equal(res, resp)
		test.Nil(err)
	})
}

func TestQueryResolver_User(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		r    = new(Resolver)
	)

	t.Run("success", func(t *testing.T) {

		resp := new(model.AbstractModel)
		ctx := context.Background()

		res, err := r.Query().User(ctx)

		test.NotNil(res)
		test.Equal(res, resp)
		test.Nil(err)
	})
}
