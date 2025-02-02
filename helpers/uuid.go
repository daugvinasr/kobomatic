package helpers

import "github.com/google/uuid"

func GetDeterministicUUID(input string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(input))
}
