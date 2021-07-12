package controllers

import (
	argonpass "github.com/dwin/goArgonPass"
)

// Hash a password with argon
func Hash(plain string) (string, error) {
	hash, err := argonpass.Hash(plain, &argonpass.ArgonParams{
		Time:       4,
		Memory:     4096,
		OutputSize: 32,
		SaltSize:   24,
	})
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Verify a password witha argon
func Verify(plain string, hashed string) (bool, error) {
	err := argonpass.Verify(plain, hashed)
	if err != nil {
		return false, err
	}
	return true, nil
}
