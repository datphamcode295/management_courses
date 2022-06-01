package errors

import "net/http"

var (
	ErrUserNotfound        = NewStringError("User not found", http.StatusNotFound)
	ErrReferrerNotfound    = NewStringError("Referrer not found", http.StatusNotFound)
	ErrUserAlreadyBonded   = NewStringError("User already bonded", http.StatusBadRequest)
	ErrInvalidReferralCode = NewStringError("Invalid referral code", http.StatusBadRequest)
)
