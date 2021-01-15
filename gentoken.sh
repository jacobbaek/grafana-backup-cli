#!/bin/bash

#
# generation for grafana token
# 

if [ -f .token ]; then
  echo "already generated token"
  exit
fi

source .config

TOKEN=$(curl -X POST -H "Content-Type: application/json" -d '{"name":"apiuser"}' http://$USERID:$USERPASS@$URL/api/orgs | jq .key)
echo $TOKEN
cat << EOF >> .token
$TOKEN
EOF

echo "generation completed"
