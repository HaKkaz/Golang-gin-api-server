package controller

import (
	"fmt"

	"github.com/lib/pq"
)

type AdFilter struct {
	Offset   int32   `form:"offset"`
	Limit    int32   `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Age      *int32  `form:"age" binding:"omitempty,gte=1,lte=100"`
	Gender   *string `form:"gender"`
	Country  *string `form:"country"`
	Platform *string `form:"platform"`
}

func NewAdFilter() *AdFilter {
	return &AdFilter{
		Offset: 0,
		Limit:  5,
	}
}

func (filter AdFilter) check() {
	fmt.Println("-------------------")
	fmt.Printf("Offset = %d\n", filter.Offset)
	fmt.Printf("Limit = %d\n", filter.Limit)
	if filter.Age != nil {
		fmt.Printf("Age = %v\n", *filter.Age)
	}
	if filter.Gender != nil {
		fmt.Printf("Gender = %v\n", *filter.Gender)
	}
	if filter.Country != nil {
		fmt.Printf("Country = %v\n", *filter.Country)
	}
	if filter.Platform != nil {
		fmt.Printf("Platform = %v\n", filter.Platform)
	}
	fmt.Println("-------------------")
}

func isInCountry(s string, countries pq.StringArray) bool {
	if len(countries) == 0 {
		return true
	}
	for _, country := range countries {
		if s == country {
			return true
		}
	}
	return false
}

func isInGender(s string, genders pq.StringArray) bool {
	if len(genders) == 0 {
		return true
	}
	for _, gender := range genders {
		if s == gender {
			return true
		}
	}
	return false
}

func isInPlatform(s string, platforms pq.StringArray) bool {
	if len(platforms) == 0 {
		return true
	}
	for _, platform := range platforms {
		if s == platform {
			return true
		}
	}
	return false
}
