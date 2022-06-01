package model

type TryXactAdvisoryLock struct {
	PgTryAdvisoryXactLock bool `json:"pg_try_advisory_xact_lock"`
}
