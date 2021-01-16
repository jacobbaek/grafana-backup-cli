package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Url      = "monitor.jacobbaek.com:3000"
	UserID   = "grafana-user"
	UserPass = "grafana-password"
)

// get dashboards using API from grafana
func apiCall(Url string, Token string, Method string) string {
	req, err := http.NewRequest(Method, Url, nil)
	if err != nil {
		panic(err)
	}

	bearer := "Bearer " + Token
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes)

	return str
}

// git control
// https://git-scm.com/book/ko/v2/Appendix-B%3A-%EC%95%A0%ED%94%8C%EB%A6%AC%EC%BC%80%EC%9D%B4%EC%85%98%EC%97%90-Git-%EB%84%A3%EA%B8%B0-go-git

func main() {
	configdata, err := ioutil.ReadFile(".config")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(configdata))

	tokendata, err := ioutil.ReadFile(".token")
	if err != nil {
		panic(err)
	}

	Token := strings.Trim(string(tokendata), "\n")
	fmt.Print(apiCall("http://"+UserID+":"+UserPass+"@"+Url, Token, "GET"))
}
