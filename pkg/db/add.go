package db

import (
	"fmt"
)

// exists, err := URLExists(db, url)
// if err != nil {
//     return fmt.Errorf("échec de la vérification de l'existence de l'URL: %v", err)
// }

// if exists {
//     // L'URL existe déjà, donc mise à jour du champ expiry
//     query := `UPDATE url_shortened SET expiry = $1 WHERE url = $2`
//     _, err := db.Exec(query, newExpiry, url)
//     if err != nil {
//         return fmt.Errorf("échec de la mise à jour du champ expiry: %v", err)
//     }
// } else {
//     // L'URL n'existe pas, retourner une erreur
//     return fmt.Errorf("l'URL '%s' n'existe pas dans la base de données", url)
// }

// Fonction pour vérifier si une URL existe déjà dans la base de données

func (d Data) InsertURL() error {
	// check if already in DB
	query := `SELECT EXISTS(SELECT 1 FROM urlshortened WHERE url = $1)`

	// Exécuter la requête
	var exists bool
	err := DBClient.QueryRow(query, d.URL).Scan(&exists)
	if err != nil {
		return fmt.Errorf("échec de l'exécution de la requête SQL: %v", err)
	}
	if exists {
		// Update
		query := `UPDATE urlshortened SET expiry = $1 WHERE url = $2`
		_, err := DBClient.Exec(query, d.Expiry, d.URL)
		if err != nil {
			return fmt.Errorf("échec de la mise à jour du champ expiry: %v", err)
		}
	} else {
		// Add
		query := `INSERT INTO urlshortened (url, short, access, expiry) VALUES ($1, $2, $3, $4)`
		_, err := DBClient.Exec(query, d.URL, d.Short, d.Access, d.Expiry)
		if err != nil {
			return fmt.Errorf("échec de l'insertion des données: %v", err)
		}
	}

	return nil
}
