package randname

import (
	"math/rand"
	"strconv"
)

func GenerateRandomName() string {
	return words[rand.Intn(len(words))]
}

func GenerateRandomNameNumber() string {
	return GenerateRandomNameNumberRang(100, 99999899)
}

func GenerateRandomNameNumberRang(min, max int) string {
	return GenerateRandomName() + strconv.Itoa(min+rand.Intn(max-min))
}
