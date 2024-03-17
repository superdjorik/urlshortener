package randomizer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandom_RandSeq(t *testing.T) {
	//init()
	randomStrings := make(map[string]int)

	charLimit := 8
	stringLimit := 100
	// Проверка на несовпадение выборки из 100 сгенерированных строк.
	for i := 0; i < stringLimit; i++ {
		randomStrings[Randomaizer(charLimit)]++
	}
	for _, count := range randomStrings {
		require.Equal(t, count, 1)
	}

}
