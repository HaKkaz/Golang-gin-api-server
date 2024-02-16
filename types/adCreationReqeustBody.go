package types

import "fmt"

type AdCondition struct {
	AgeStart int32    `json:"ageStart" binding:"omitempty,gte=1,ltefield=AgeEnd"`
	AgeEnd   int32    `json:"ageEnd" binding:"omitempty,gte=1,lte=100"`
	Gender   []string `json:"gender" binding:"omitempty,dive,oneof=M F"`
	Country  []string `json:"country" binding:"omitempty,dive,iso3166_1_alpha2"`
	Platform []string `json:"platform" binding:"omitempty,dive,oneof=android ios web"`
}

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
