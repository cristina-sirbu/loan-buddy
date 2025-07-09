#!/bin/bash

set -e

BASE_URL="http://localhost:8080"

echo "> Get available loan offers"
curl -s ${BASE_URL}/offers | jq .
echo -e "\n> Offers listed above"

echo "> Choose one offer by its ID (e.g. offer2)"
SELECTED_OFFER_ID="offer3"
USER_ID="cristina"

echo "> Checkout selected offer (loan_id=${SELECTED_OFFER_ID}, user_id=${USER_ID})"
RESPONSE=$(curl -s -X POST ${BASE_URL}/checkout \
     -H "Content-Type: application/json" \
     -d "{\"user_id\":\"${USER_ID}\", \"loan_id\":\"${SELECTED_OFFER_ID}\"}")

echo "${RESPONSE}" | jq .
echo -e "\n> Checkout completed"
