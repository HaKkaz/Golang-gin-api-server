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

// Convert ad to adResult, adResult is the spcefic format for response
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
