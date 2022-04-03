// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Carts struct {
	UserInfo *UserInfo      `json:"userInfo"`
	Cart     []*ProductCart `json:"cart"`
}

type Catalogs struct {
	Catalogs []*Product `json:"catalogs"`
}

type ListUser struct {
	Users []*UserInfo `json:"users"`
}

type Product struct {
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	Price              int    `json:"price"`
	Stock              int    `json:"stock"`
}

type ProductCart struct {
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}

type UserInfo struct {
	UserID       string `json:"userID"`
	UserName     string `json:"userName"`
	UserLocation string `json:"userLocation"`
}