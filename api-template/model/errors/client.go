package errors

import "net/http"

var (
	ErrAccountNotfound = NewStringError("Account does not exist", http.StatusNotFound)
	ErrUnAuthorized    = NewStringError("User not authorized", http.StatusUnauthorized)
	// ErrAtNotfound = NewStringError("CourseAtClass not found", http.StatusNotFound)
	// ErrReferrerNotfound    = NewStringError("Referrer not found", http.StatusNotFound)
	// ErrUserAlreadyBonded   = NewStringError("User already bonded", http.StatusBadRequest)
	// ErrInvalidReferralCode = NewStringError("Invalid referral code", http.StatusBadRequest)
)
