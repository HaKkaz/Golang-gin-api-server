package controller

import (
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_NewAdFilter(t *testing.T) {
	filter := NewAdFilter()
	assert.Equal(t, int32(0), filter.Offset)
	assert.Equal(t, int32(5), filter.Limit)
}

func Test_IsInCountry(t *testing.T) {
	countries := pq.StringArray{"TW", "JP", "US"}
	assert.True(t, isInCountry("TW", countries))
	assert.False(t, isInCountry("CN", countries))

	countries = pq.StringArray{}
	assert.True(t, isInCountry("TW", countries))
}

func Test_IsInGender(t *testing.T) {
	genders := pq.StringArray{"F"}
	assert.True(t, isInGender("F", genders))
	assert.False(t, isInGender("M", genders))

	genders = pq.StringArray{}
	assert.True(t, isInGender("F", genders))
	assert.True(t, isInGender("M", genders))
}

func Test_IsInPlatform(t *testing.T) {
	platforms := pq.StringArray{"iOS", "Android"}
	assert.True(t, isInPlatform("iOS", platforms))
	assert.False(t, isInPlatform("Web", platforms))

	platforms = pq.StringArray{}
	assert.True(t, isInPlatform("iOS", platforms))
}
