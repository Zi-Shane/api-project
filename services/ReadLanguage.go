package services

import "api/database"

func ReadLanguage(id string) []database.Languages {
	var languages []database.Languages
	// SELECT * FROM language WHERE id=1
	database.DB.Find(&languages, id)

	return languages
}
