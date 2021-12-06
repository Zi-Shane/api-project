package services

import "api/database"

func ReadLanguage(id string) []database.Languages {
	var languages []database.Languages
	database.DB.Find(&languages, id)

	return languages
}
