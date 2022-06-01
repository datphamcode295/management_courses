package errors

import "net/http"

var (
	ErrInvalidNFTCharacter          = NewStringError("Character type not existed", http.StatusBadRequest)
	ErrCharacterNotfound            = NewStringError("Character not found", http.StatusNotFound)
	ErrNotOwnerCharacter            = NewStringError("Not an owner's character", http.StatusBadRequest)
	ErrInvalidPrice                 = NewStringError("Invalid price", http.StatusBadRequest)
	ErrInvalidAttributeName         = NewStringError("Invalid attribute name", http.StatusInternalServerError)
	ErrInvalidTokenId               = NewStringError("Invalid token id", http.StatusInternalServerError)
	ErrAttributeRequirementNotMatch = NewStringError("Character's attribute not matched with requirement", http.StatusInternalServerError)
	ErrInvalidCharacterLevel        = NewStringError("Character's level not enough", http.StatusBadRequest)
)
