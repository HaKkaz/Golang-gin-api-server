#!/bin/bash

for T in {0..1000}; do
    echo "$T"
    # Generate random title
    title=$(head /dev/urandom | tr -dc A-Za-z0-9 | head -c 10)

    # Generate random startAt and endAt dates within a range
    startAt=$(date -d "$((RANDOM%365)) days ago" +"%Y-%m-%dT%H:%M:%S.000Z")
    endAt=$(date -d "$((RANDOM%365)) days" +"%Y-%m-%dT%H:%M:%S.000Z")

    # Generate random ageStart and ageEnd within a range
    ageStart=$((RANDOM%50 + 10))
    ageEnd=$((ageStart + RANDOM%20))

    # Generate random country list
    countries=("TW" "JP" "AW" "AZ" "BH" "BY")
    random_countries=()
    for i in {1..2}; do
        random_countries+=("${countries[RANDOM%${#countries[@]}]}")
    done

    # Generate random platform list
    platforms=("android" "ios" "web")
    random_platforms=()
    for i in {1..2}; do
        random_platforms+=("${platforms[RANDOM%${#platforms[@]}]}")
    done

    # Construct JSON data
    json_data=$(cat <<EOF
{
    "title": "$title",
    "startAt": "$startAt",
    "endAt": "$endAt",
    "conditions": {
        "ageStart": $ageStart,
        "ageEnd": $ageEnd,
        "country": $(printf '%s\n' "${random_countries[@]}" | jq -R . | jq -s .),
        "platform": $(printf '%s\n' "${random_platforms[@]}" | jq -R . | jq -s .)
    }
}
EOF
)

    # Send POST request
    curl -X POST -H "Content-Type: application/json" \
    "http://localhost:8080/api/v1/ad" \
    --data "$json_data"

    echo ""
    # echo "$json_data"
done