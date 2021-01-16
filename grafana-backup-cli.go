package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// global variables

var Url string
var UserID string
var UserPass string

type Dash struct {
	Title string `json:"title"`
	Uid   string `json:"uid"`
	Type  string `json:"type"`
}

func init() {
	// default vaules
	Url = "monitor.jacobbaek.com:3000"
	UserID = "grafana-user"
	UserPass = "grafana-password"
}

func error_check(err error) {
	if err != nil {
		panic(err)
	}
}

// get dashboards using API from grafana
func apiCall(url string, token string) string {
	req, err := http.NewRequest("GET", url, nil)
	error_check(err)

	bearer := "Bearer " + token
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	error_check(err)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)

	return str
}

// git control
// https://git-scm.com/book/ko/v2/Appendix-B%3A-%EC%95%A0%ED%94%8C%EB%A6%AC%EC%BC%80%EC%9D%B4%EC%85%98%EC%97%90-Git-%EB%84%A3%EA%B8%B0-go-git

func save_bygit() {
	fmt.Print("save dashboards into git repo by git command")
}

// main function

func main() {

	if _, err := os.Stat(".config"); os.IsNotExist(err) {
		panic(err)
	}

	configdata, err := ioutil.ReadFile(".config")
	error_check(err)

	splitdata := strings.Split(strings.TrimRight(string(configdata), "\n"), "\n")
	for _, config := range splitdata {
		line := strings.Split(config, "=")
		if line[0] == "DOMAIN" {
			Url = strings.Trim(line[1], "\"")
		} else if line[0] == "USERID" {
			UserID = strings.Trim(line[1], "\"")
		} else if line[0] == "USERPASS" {
			UserPass = strings.Trim(line[1], "\"")
		} else {
			fmt.Println("don't use " + strings.Trim(line[1], "\"") + " Config line")
		}
	}

	if _, err := os.Stat(".token"); os.IsNotExist(err) {
		panic(err)
	}

	tokendata, err := ioutil.ReadFile(".token")
	error_check(err)

	Token := strings.Trim(string(tokendata), "\n")
	raw_data := apiCall("http://"+UserID+":"+UserPass+"@"+Url+"/api/search", Token)

	json_data := Dash{}

	err_marshal := json.Unmarshal([]byte(raw_data), &json_data)
	error_check(err_marshal)

	save_bygit()
}
