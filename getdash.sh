#!/bin/bash

source .config

if [ ! -f ".token" ]; then
  echo "First, should run gentoken.sh"
  exit 1
fi

BACKUPDIR="BackupData"
TOKEN=$(cat .token)
URLLST=$(curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/search | jq -r '.[] | select( .type == "dash-db" ) | .uid')

#IFS=' ' read -r -a array <<< $URLLST
dashuid=($URLLST)
Num=1

if [ ! -d $BACKUPDIR ]; then
  mkdir $BACKUPDIR
fi

for dash in ${dashuid[@]}; do
  SUBURL="dashboards/uid/$dash"
  echo "Do "$SUBURL
  curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/$SUBURL | jq .dashboard > $BACKUPDIR/$Num".json"
  Num=$(($Num+1))
done
