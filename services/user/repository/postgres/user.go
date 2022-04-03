package postgres

import (
	"fmt"
)

func (userRepository *Repository) GetUserInfoByIDQuery(userID string) ([][]string, error) {
	qapiapp := `
	SELECT
		user_id,
		user_name,
		user_location
	FROM
		users
	WHERE 
		user_id= '%s'
	`
	qapiapp = fmt.Sprintf(qapiapp, userID)

	return userRepository.Connection.ExecuteDBQuery(qapiapp)
}

func (userRepository *Repository) ListOfUserQuery() ([][]string, error) {
	qapiapp := `
	SELECT
		user_id,
		user_name,
		user_location
	FROM
		users
	`
	qapiapp = fmt.Sprintf(qapiapp)

	return userRepository.Connection.ExecuteDBQuery(qapiapp)
}
