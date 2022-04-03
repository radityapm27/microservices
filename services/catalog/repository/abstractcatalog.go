package repository

import (
	proto "rpm/microservices/core/proto"
)

// GetProductCatalog ..
func (ar *abstractRepository) GetProductCatalog() (*proto.ProductCatalogResponse, error) {
	var (
		result = &proto.ProductCatalogResponse{
			Catalogs: &proto.Catalog{
				Products: []*proto.Product{},
			},
		}
	)

	records, err := ar.GetProductCatalogQuery()
	if err != nil {
		return nil, err
	}

	for i := range records {
		if i > 0 {
			price, errPrice := ar.utils.StringToInt(records[i][2])
			if errPrice != nil {
				return nil, errPrice
			}

			stock, errStock := ar.utils.StringToInt(records[i][3])
			if errStock != nil {
				return nil, errStock
			}
			product := &proto.Product{
				ProductName:        records[i][0],
				ProductDescription: records[i][1],
				Price:              int32(price),
				Stock:              int32(stock),
			}

			result.Catalogs.Products = append(result.Catalogs.Products, product)
		}
	}
	return result, nil
}
