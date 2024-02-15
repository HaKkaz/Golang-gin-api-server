package utils

import (
	"fmt"
	"time"
)

func DateToTimestamp(dateStr string) int32 {
	t, err := time.Parse(time.RFC3339Nano, dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return -1
	}

	// Convert the time value to a Unix timestamp (seconds)
	timestamp := t.Unix()

	fmt.Println("Timestamp (seconds):", timestamp)
	return int32(timestamp)
}
