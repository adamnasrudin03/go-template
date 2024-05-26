package models

const (
	ROOT  = `ROOT`
	ADMIN = `ADMIN`
	USER  = `USER`
)

var (
	IsUserRoleValid = map[string]bool{
		ADMIN: true,
		USER:  true,
	}
)
