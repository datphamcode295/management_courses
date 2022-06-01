package errors

import "net/http"

var (
	ErrAdventureNotfound             = NewStringError("Adventure not found", http.StatusNotFound)
	ErrMonsterNotfound               = NewStringError("Monster not found", http.StatusNotFound)
	ErrAdventureLogNotfound          = NewStringError("Adventure log not found", http.StatusNotFound)
	ErrAdventureBattleNotfound       = NewStringError("Battle not found", http.StatusNotFound)
	ErrCharacterIsAlreadyInAdventure = NewStringError("Character is already in adventure", http.StatusBadRequest)
	ErrRunOutOfAdventureChance       = NewStringError("Run out of adventure chance", http.StatusBadRequest)
	ErrRunOutOfHealth                = NewStringError("Run out of health", http.StatusBadRequest)
	ErrInvalidAdventureLevel         = NewStringError("Invalid adventure level", http.StatusBadRequest)
	ErrStillInCoolingTime            = NewStringError("Character still in cooling time", http.StatusBadRequest)
)
