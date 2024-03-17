package test

import (
	"net/http"
)

// Integration test cases
func initIntegrationTestCases() []testcase {
	postWrongCountryCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "No Conditions",
			"startAt": "2023-12-10T03:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"country": ["SDKNFJ"]
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Conditions.Country[0]' Error:Field validation for 'Country[0]' failed on the 'iso3166_1_alpha2' tag"}`,
	}

	postWrongPlatformCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "AD 55",
			"startAt": "2023-12-10T03:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 30,
				"country": ["TW", "JP"],
				"platform": ["android", "ios", "haha"]
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Conditions.Platform[2]' Error:Field validation for 'Platform[2]' failed on the 'oneof' tag"}`,
	}

	postAgeOutOfRangeCase1 := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "age out of range",
			"startAt": "2021-01-08T13:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 10
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Conditions.AgeStart' Error:Field validation for 'AgeStart' failed on the 'ltefield' tag"}`,
	}

	postAgeOutOfRangeCase2 := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "age out of range",
			"startAt": "2021-01-08T13:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": -1,
				"ageEnd": 10
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Conditions.AgeStart' Error:Field validation for 'AgeStart' failed on the 'gte' tag"}`,
	}

	postAgeOutOfRangeCase3 := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "age out of range",
			"startAt": "2021-01-08T13:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 101
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Conditions.AgeEnd' Error:Field validation for 'AgeEnd' failed on the 'lte' tag"}`,
	}

	postEmptyTitleCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "",
			"startAt": "2021-01-08T13:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 100
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
	}

	postEmptyStartAtCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "Empty StartAt",
			"endAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 50
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.StartAt' Error:Field validation for 'StartAt' failed on the 'required' tag"}`,
	}

	postEmptyEndAtCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "Empty StartAt",
			"StartAt": "2025-12-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 59
			}
		}`,
		400,
		`{"error":"Key: 'AdCreationRequesetBody.EndAt' Error:Field validation for 'EndAt' failed on the 'required' tag"}`,
	}

	postGeneralCase1 := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "General Case 1",
			"startAt": "2018-05-10T03:00:00.000Z",
			"endAt": "2025-01-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 18,
				"ageEnd": 30,
				"country": ["TW", "JP", "US"],
				"platform": ["android", "ios", "web"]
			}
		}`,
		200,
		`{"message":"Advertisement created successfully"}`,
	}

	postGeneralCase2 := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "General Case 2",
			"startAt": "2022-01-10T03:00:00.000Z",
			"endAt": "2025-01-31T16:00:00.000Z",
			"conditions": {
				"ageStart": 20,
				"ageEnd": 80,
				"country": ["TW"],
				"platform": ["android", "ios", "web"]
			}
		}`,
		200,
		`{"message":"Advertisement created successfully"}`,
	}

	postNoConditionCase := testcase{
		"/api/v1/ad",
		http.MethodPost,
		`{
			"title": "No Conditions",
			"startAt": "2023-12-10T03:00:00.000Z",
			"endAt": "2025-12-31T16:00:00.000Z"
		}`,
		200,
		`{"message":"Advertisement created successfully"}`,
	}

	// Should show all ads
	getWithoutConditionCase := testcase{
		"/api/v1/ad",
		http.MethodGet,
		"",
		200,
		`{"items":[{"title":"General Case 1","endAt":"2025-01-31T16:00:00Z"},{"title":"General Case 2","endAt":"2025-01-31T16:00:00Z"},{"title":"No Conditions","endAt":"2025-12-31T16:00:00Z"}]}`,
	}

	// Should show ads with country JP or no country condition
	getWithCountryJPCase := testcase{
		"/api/v1/ad?country=JP",
		http.MethodGet,
		"",
		200,
		`{"items":[{"title":"General Case 1","endAt":"2025-01-31T16:00:00Z"},{"title":"No Conditions","endAt":"2025-12-31T16:00:00Z"}]}`,
	}

	return []testcase{
		postAgeOutOfRangeCase1,
		postAgeOutOfRangeCase2,
		postAgeOutOfRangeCase3,
		postWrongCountryCase,
		postWrongPlatformCase,
		postEmptyTitleCase,
		postEmptyStartAtCase,
		postEmptyEndAtCase,

		postGeneralCase1,
		postGeneralCase2,
		postNoConditionCase,

		getWithoutConditionCase,
		getWithCountryJPCase,
	}
}
