#!/bin/bash

source .config

TOKEN=$(cat .token)
URLLST=$(curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/search | jq -r '.[] | select( .type == "dash-db" ) | .uid')

#IFS=' ' read -r -a array <<< $URLLST
dashuid=($URLLST)
echo ${dashuid[@]}
for num in ${dashuid[@]}; do
  echo ${dashuid[num]}
  #SUBURL="dashboards/uid/${dashuid[num]}"
  #curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/$SUBURL | jq
done
