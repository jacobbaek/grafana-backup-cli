#!/bin/bash

source .config

if [ ! -f ".token" ]; then
  echo "First, should run gentoken.sh"
  exit 1
fi

BACKUPDIR="BackupData"
TOKEN=$(cat .token)
URLLST=$(curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/search | jq -r '.[] | select( .type == "dash-db" ) | .uid')

dashuid=($URLLST)
Num=1

if [ ! -d $BACKUPDIR ]; then
  mkdir $BACKUPDIR
fi

for dash in ${dashuid[@]}; do
  URLPATH="dashboards/uid/$dash"
  echo "Do "$URLPATH
  curl -s -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://$USERID:$UESRPASS@$URL/api/$URLPATH | jq .dashboard > $BACKUPDIR/$Num".json"
  Num=$(($Num+1))
done
