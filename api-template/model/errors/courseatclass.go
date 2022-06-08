package errors

import "net/http"

var (
	ErrAtNotfound  = NewStringError("CourseAtClass not found", http.StatusNotFound)
	ErrAtsNotfound = NewStringError("Record not found", http.StatusNotFound)
	// ErrUserAlreadyBonded   = NewStringError("User already bonded", http.StatusBadRequest)
	// ErrInvalidReferralCode = NewStringError("Invalid referral code", http.StatusBadRequest)
)
