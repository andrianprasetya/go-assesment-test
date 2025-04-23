package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUniqueNumber() string {
	// Seed random
	rand.Seed(time.Now().UnixNano())

	// Get current time as MMSS (Minutes and Seconds)
	t := time.Now()
	timestampPart := t.Format("0405") // HHMM or MMSS (4 digits)

	// Generate random 4-digit number
	randomPart := rand.Intn(10000) // 0 - 9999
	randomStr := fmt.Sprintf("%04d", randomPart)

	// Concatenate: "25" + timestamp(4) + random(4) = 10 digits
	return "25" + timestampPart + randomStr
}
