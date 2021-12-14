package services

import (
	"api/database"
)

func InsertLanguage(items []database.Languages) (error, int64) {
	result := database.DB.Create(items)

	return result.Error, result.RowsAffected
}
