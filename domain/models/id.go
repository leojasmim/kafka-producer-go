package models

import "github.com/lucsky/cuid"

const STRING_EMPTY = ""

func GenerateID(prefix string) string {
	id := cuid.New()

	if prefix == STRING_EMPTY {
		return id
	}

	return prefix + "_" + id
}
