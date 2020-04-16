package auth

import "github.com/markbates/goth"

// TODO: This matcher is temporal, in the future, this program needs a dynamic matcher based on
// complex admin configuration

func staticMatcher(user goth.User) (string, error) {
	// user.
	return "", nil
}
