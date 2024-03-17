package randomizer

import (
	"math/rand"
	"time"
)

var randomChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var randRange *rand.Rand

// Инициализация рандомайзера
func init() {
	source := rand.NewSource(time.Now().UnixNano())
	randRange = rand.New(source)
}

// Генератор случайной строки
func Randomaizer(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randomChars[randRange.Intn(len(randomChars))]
	}
	return string(b)
}
