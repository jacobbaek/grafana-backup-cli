# Grafana dashboard json files Backup cli tool


# What is in
* collect dashboard json files from grafana using API
* commit and push with specific message into git server

## components
* gentoken.sh : make grafana token if you don't have .token file
* getdash.sh : collect dashboards using cURL from grafana
* grafana-backup-cli.go : 
  * collect dashboards using golang program from grafana 
  * git commit and push dashboards that should backup

# How to use
1. run gettoken.sh
2. will be wrote .token file in current directory
3. run the grafana-backup-cli (or run the getdash.sh script.)

# Recommended API urls
* https://grafana.com/docs/grafana/latest/http_api/create-api-tokens-for-org/
* 