package sql

import "github.com/lib/pq"

const uniqueViolation = "23505"

func IsUniqueViolation(err error) bool {
	if pqerr, ok := err.(*pq.Error); ok && pqerr.Code == uniqueViolation {
		return true
	}
	return false
}
