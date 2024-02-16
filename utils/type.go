package utils

import (
	"fmt"
	"time"
)

func DateToTimestamp(dateStr string) int64 {
	t, err := time.Parse(time.RFC3339Nano, dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return -1
	}

	timestamp := t.Unix()

	fmt.Println("Timestamp (seconds):", timestamp)
	return timestamp
}

func TimestampToDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)

	dateStr := t.Format(time.RFC3339Nano)

	return dateStr
}
