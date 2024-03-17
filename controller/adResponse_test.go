package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdToadResult(t *testing.T) {
	ads := []ad{
		{
			Title:  "title1",
			End_At: 1617225600,
		},
		{
			Title:  "title2",
			End_At: 1717225600,
		},
	}

	results := AdToadResult(ads)
	assert.Equal(t, 2, len(results))
	assert.Equal(t, "title1", results[0].Title)
	assert.Equal(t, "2021-03-31T21:20:00Z", results[0].End_At)
	assert.Equal(t, "title2", results[1].Title)
	assert.Equal(t, "2024-06-01T07:06:40Z", results[1].End_At)
}
