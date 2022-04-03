package postgres

import (
	"fmt"
)

func (cartRepository *Repository) GetShoppingCartByUserIdQuery(userID string) ([][]string, error) {
	qapiapp := `
	SELECT
		ctlg.product_name,
		ctlg.product_desc,
		ctlg.price,
		ctlg.stock,
		crt.quantity,
		usr.user_id,
		usr.user_name,
		usr.user_location
	FROM
		cart crt
		LEFT JOIN "catalog" ctlg ON crt.product_id = ctlg.product_id
		LEFT JOIN users usr ON crt.user_id = usr.user_id
	WHERE 
		usr.user_id= '%s'
	`
	qapiapp = fmt.Sprintf(qapiapp, userID)

	return cartRepository.Connection.ExecuteDBQuery(qapiapp)
}
