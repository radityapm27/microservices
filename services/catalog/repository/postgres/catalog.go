package postgres

import (
	"fmt"
)

func (catalogRepository *Repository) GetProductCatalogQuery() ([][]string, error) {
	qapiapp := `
	SELECT
		product_name,
		product_desc,
		price,
		stock
	FROM
		catalog
	`
	qapiapp = fmt.Sprintf(qapiapp)

	return catalogRepository.Connection.ExecuteDBQuery(qapiapp)
}
