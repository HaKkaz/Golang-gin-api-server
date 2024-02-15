package types

import "fmt"

type AdCreationRequesetBody struct {
	Title      string      `json:"title" binding:"required"`
	StartAt    string      `json:"startAt" binding:"required"`
	EndAt      string      `json:"endAt" binding:"required"`
	Conditions AdCondition `json:"conditions" binding:"omitempty"`
}

func (ad AdCreationRequesetBody) Print() {
	fmt.Println("Title", ad.Title)
	fmt.Println("StartAt", ad.StartAt)
	fmt.Println("EndAt", ad.EndAt)
	fmt.Printf("AgeStart %d\n", ad.Conditions.AgeStart)
	fmt.Printf("AgeEnd %d\n", ad.Conditions.AgeEnd)
	fmt.Printf("Platform %v\n", ad.Conditions.Country)
}
