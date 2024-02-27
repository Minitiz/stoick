package db

import (
	"fmt"
)

func ReadDataAndIncrementAccess(short string) (*Data, error) {

	query := `SELECT url, short, access, expiry FROM urlshortened WHERE short = $1`
	tx, err := DBClient.Begin()
	if err != nil {
		return nil, fmt.Errorf("échec de la transaction: %v", err)
	}

	rows, err := tx.Query(query, short)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("échec de l'exécution de la requête SELECT: %v", err)
	}
	defer rows.Close()

	var result Data

	for rows.Next() {
		if err := rows.Scan(&result.URL, &result.Short, &result.Access, &result.Expiry); err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("fail line cursor: %v", err)
		}
		continue
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fail cursor: %v", err)
	}
	updateQuery := `UPDATE urlshortened SET access = access + 1 WHERE short = $1`
	if _, err := tx.Exec(updateQuery, short); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("fail UPDATE: %v", err)
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("fail transaction: %v", err)
	}

	return &result, nil
}
