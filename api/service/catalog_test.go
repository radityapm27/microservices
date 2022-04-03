package service

import (
	"context"
	"testing"

	"rpm/microservices/api/graph/model"
	proto "rpm/microservices/core/proto"

	"github.com/stretchr/testify/assert"
)

func TestService_GetProductCatalog(t *testing.T) {
	t.Parallel()

	var (
		test = assert.New(t)
		s    = new(service)
	)

	t.Run("success", func(t *testing.T) {
		var (
			response = &proto.ProductCatalogResponse{
				Catalogs: &proto.Catalog{
					Products: []*proto.Product{
						&proto.Product{
							ProductName:        "Product Name 1",
							ProductDescription: "Product Description 1",
							Price:              10000,
							Stock:              100,
						},
						&proto.Product{
							ProductName:        "Product Name 2",
							ProductDescription: "Product Description 2",
							Price:              20000,
							Stock:              200,
						},
					},
				},
			}
			err error
			ctx = context.Background()
			obj = &model.AbstractModel{}
		)

		svc := resetCatalog(s)
		svc.On("GetProductCatalog", ctx, &proto.CoreRequest{}).Return(response, err)
		resp, respErr := s.GetProductCatalog(ctx, obj)

		svc.AssertExpectations(t)

		test.NotNil(resp)
		test.Nil(respErr)
	})
}
