package services

import (
	"api/database"
)

func DeleteLanguage(languageName string) (error, int64) {
	result := database.DB.Where("Language = ?", languageName).Delete(&database.Languages{})

	return nil, result.RowsAffected
}
