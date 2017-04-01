package random

import "math/rand"

func Word(w []string) string {
	return w[rand.Intn(len(w))]
}
