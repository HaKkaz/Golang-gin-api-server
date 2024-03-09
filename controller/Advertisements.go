package controller

import (
	"github.com/lib/pq"
)

type Advertisements struct {
	Title    string         `json:"title"`
	StartAt  int64          `json:"start_at" gorm:"index:idx_ad_start_end_age, priority:2"`
	EndAt    int64          `json:"end_at" gorm:"index:idx_ad_start_end_age, priority:1"`
	AgeStart int32          `json:"age_start,omitempty" gorm:"index"`
	AgeEnd   int32          `json:"age_end,omitempty" gorm:"index"`
	Gender   pq.StringArray `json:"gender,omitempty" gorm:"type:text[]"`
	Country  pq.StringArray `json:"country,omitempty" gorm:"type:text[]"`
	Platform pq.StringArray `json:"platform,omitempty" gorm:"type:text[]"`
}
