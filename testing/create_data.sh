#!/bin/bash

# General Case
curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "General Case 1",
	"startAt": "2023-12-10T03:00:00.000Z",
	"endAt": "2025-12-31T16:00:00.000Z",
	"conditions": {
		"ageStart": 20,
		"ageEnd": 30,
		"country": ["TW", "JP"],
		"platform": ["android", "ios"]
	}
}'

echo ""

curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "General Case 2",
	"startAt": "2022-01-10T03:00:00.000Z",
	"endAt": "2025-01-31T16:00:00.000Z",
	"conditions": {
		"ageStart": 20,
		"ageEnd": 80,
		"country": ["TW"],
		"platform": ["android", "ios", "web"]
	}
}'

echo ""

# No conditions
curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "No Conditions",
	"startAt": "2023-12-10T03:00:00.000Z",
	"endAt": "2025-12-31T16:00:00.000Z"
}'

echo ""

# US
curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "US",
	"startAt": "1999-12-10T03:00:00.000Z",
	"endAt": "2030-12-31T16:00:00.000Z",
	"conditions": {
		"ageStart": 20,
		"ageEnd": 30,
		"country": ["TW", "JP", "US"],
		"platform": ["android", "ios"]
	}
}'
echo ""

curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "Very Old",
	"startAt": "1999-12-10T03:00:00.000Z",
	"endAt": "2000-12-31T16:00:00.000Z",
	"conditions": {
		"ageStart": 20,
		"ageEnd": 30,
		"gender": ["F"],
		"country": ["TW", "JWWP", "US"],
		"platform": ["android", "ios"]
	}
}'
echo ""

# No Country
curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "No Country",
	"startAt": "1999-12-10T03:00:00.000Z",
	"endAt": "2030-12-31T16:00:00.000Z",
	"conditions": {
		"ageStart": 20,
		"ageEnd": 30,
		"platform": ["android", "ios"]
	}
}'
echo ""

# Wrong Country 
curl -X POST -H "Content-Type: application/json" \
"http://localhost:8080/api/v1/ad" \
--data '{
	"title": "No Conditions",
	"startAt": "2023-12-10T03:00:00.000Z",
	"endAt": "2025-12-31T16:00:00.000Z",
	"conditions": {
		"country": ["SDKNFJ"]
	}
}'
echo ""
