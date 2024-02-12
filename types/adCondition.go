package types

type AdCondition struct {
	AgeStart int      `json:"ageStart" binding:"omitempty,gte=1,ltefield=AgeEnd"`
	AgeEnd   int      `json:"ageEnd" binding:"omitempty,gte=1,lte=100"`
	Gender   []string `json:"gender" binding:"omitempty,dive,oneof=M F"`
	Country  []string `json:"country" binding:"omitempty,dive,iso3166_1_alpha2"`
	Platform []string `json:"platform" binding:"omitempty,dive,oneof=android ios web"`
}
