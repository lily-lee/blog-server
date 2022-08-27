package common

import (
	"math/rand"
	"time"
)

func Random(n int) string {
	seed := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	rand.Seed(time.Now().UnixNano())

	random := ""
	for i := 0; i < n; i++ {
		random += string(seed[rand.Intn(len(seed)-1)])
	}

	return random
}
