package repository

import (
	proto "rpm/microservices/core/proto"
)

// GetShoppingCartByUserId ..
func (ar *abstractRepository) GetShoppingCartByUserId(userId string) (*proto.ShoppingCartResponse, error) {
	var (
		result = &proto.ShoppingCartResponse{
			Cart: &proto.Cart{
				UserInfo:    &proto.UserInfoResponse{},
				ProductCart: []*proto.ProductCart{},
			},
		}
	)

	records, err := ar.GetShoppingCartByUserIdQuery(userId)
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

			quantity, errQuantity := ar.utils.StringToInt(records[i][4])
			if errQuantity != nil {
				return nil, errQuantity
			}

			productCart := &proto.ProductCart{
				Product: &proto.Product{
					ProductName:        records[i][0],
					ProductDescription: records[i][1],
					Price:              int32(price),
					Stock:              int32(stock),
				},
				Quantity: int32(quantity),
			}
			result.Cart.ProductCart = append(result.Cart.ProductCart, productCart)

			result.Cart.UserInfo = &proto.UserInfoResponse{
				UserID:       records[i][5],
				UserName:     records[i][6],
				UserLocation: records[i][7],
			}
		}
	}
	return result, nil
}
