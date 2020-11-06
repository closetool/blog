package constants

import "time"

const (
	AuthHeader = "Authorization"
	TokenTTL   = time.Duration(time.Second * 7 * 24 * 60 * 60)

	RoleUser  = 1
	RoleAdmin = 2

	AccountUnlocked = 0
	AccountLocked   = 1

	SocialIsHome  = 1
	SocialEnabled = 1

	TreePath = "."
)

var (
	Roles = map[int64]string{RoleUser: "user", RoleAdmin: "admin"}
)
