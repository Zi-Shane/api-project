package services

import "api/database"

func ReadLanguage(id string) (database.Languages, error) {
	var languages database.Languages
	// SELECT * FROM `language` WHERE id=1
	result := database.DB.Where("Language_id = ?", id).Find(&languages)

	return languages, result.Error
}

func ReadLanguages(startId string, endId string) ([]database.Languages, error) {
	var languages []database.Languages
	// SELECT * FROM `language` WHERE `Language_id` BETWEEN 10 AND 20
	result := database.DB.Where("Language_id BETWEEN ? AND ?", startId, endId).Find(&languages)

	return languages, result.Error
}

func InsertLanguage(items []database.Languages) (int64, error) {
	// INSERT INTO `Language` (`Language_id`, `Language`) VALUES (500, "Test1"), (500, "Test1")
	result := database.DB.Create(items)

	return result.RowsAffected, result.Error
}

func DeleteLanguage(languageName string) (int64, error) {
	// DELETE FROM `Language` WHERE `Language` = "Test1"
	result := database.DB.Where("Language = ?", languageName).Delete(&database.Languages{})

	return result.RowsAffected, result.Error
}

func UpdateLanguage(item database.Languages) (int64, error) {
	// UPDATE `Language` SET `Language` = "UpdatedLanguage" WHERE `Language_id` = 500
	result := database.DB.Model(&database.Languages{}).Where("Language_id = ?", item.Language_id).Update("Language", item.Language)

	return result.RowsAffected, result.Error
}
