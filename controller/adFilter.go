package controller

import "fmt"

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
		Limit: 5,
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
