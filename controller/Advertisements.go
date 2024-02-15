package controller

import "github.com/lib/pq"

type Advertisements struct {
	Title     string         `json:"title"`
	Start_At  string         `json:"start_at" gorm:"index:idx_ad_start_end_age, priority:2"`
	End_At    string         `json:"end_at" gorm:"index:idx_ad_start_end_age, priority:1"`
	Age_Start int            `json:"age_start,omitempty" gorm:"index"`
	Age_End   int            `json:"age_end,omitempty" gorm:"index"`
	Gender    pq.StringArray `json:"gender,omitempty" gorm:"type:text[]"`
	Country   pq.StringArray `json:"country,omitempty" gorm:"type:text[]"`
	Platform  pq.StringArray `json:"platform,omitempty" gorm:"type:text[]"`
}
