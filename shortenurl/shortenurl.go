package shortenurl

import (
	"fmt"
	"math/rand"
)

func ShortenURL(url string) string {
	s := ""

	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97)
	}

	return fmt.Sprintf("http://localhost:8000/%s", s)
}
