package pg

import (
	"gorm.io/gorm"
)

func getByFieldSort(query *gorm.DB, field string, sortOrder string) *gorm.DB {
	if len(sortOrder) > 0 {

		if sortOrder == "asc" {
			query = query.Order(field + " asc")
		} else {
			query = query.Order(field + " desc")
		}
	}
	return query
}

// func getTotalScore(character *model.Character) int {
// 	return character.Strength + character.Stamina + character.Courage + character.Dexterity + character.Intelligence + character.Vitality
// }
