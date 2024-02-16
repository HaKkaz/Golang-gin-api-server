package controller

import "go-gin-api-server/utils"

type ad struct {
	Title  string `json:"title"`
	End_At int64  `json:"endAt"`
}

type adResult struct {
	Title  string `json:"title"`
	End_At string `json:"endAt"`
}

type adResponse struct {
	Ads []adResult `json:"items"`
}

func AdToadResult(ads []ad) []adResult {
	var results []adResult

	for _, ad := range ads {
		endAt := utils.TimestampToDate(ad.End_At)

		result := adResult{
			Title:  ad.Title,
			End_At: endAt,
		}

		results = append(results, result)
	}

	return results
}
