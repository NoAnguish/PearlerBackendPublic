package formatters

import (
	"math/rand"
	"time"
)

// never change it
const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
const idLength = 18
const imageIdLength = 36

var randProducer *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateId() string {
	id := make([]byte, idLength)

	for i := range id {
		id[i] = charset[randProducer.Intn(len(charset))]
	}

	return string(id)
}

func GenerateImageId() string {
	id := make([]byte, imageIdLength)

	for i := range id {
		id[i] = charset[randProducer.Intn(len(charset))]
	}

	return string(id)
}
