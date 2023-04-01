package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCustomerNumber() *string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomInt := rand.Intn(1000)
	customerNumber := fmt.Sprintf("%d-%d", timestamp, randomInt)

	return &customerNumber
}
