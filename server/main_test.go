package main

import (
	"fmt"
	"github.com/mrjonze/goexpert/rate-limiter/server/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestLimitIpOk(t *testing.T) {
	configs, err := config.LoadConfig()
	if err != nil {
		println("Error loading config")
		panic(err)
	}

	ipLimit := configs.RequestLimitIp
	var mapOfResponses = make(map[int]int)

	for i := 1; i <= ipLimit; i++ {
		doRequest(false, &mapOfResponses)
	}

	assert.Equal(t, ipLimit, mapOfResponses[200], "Failed to reach ip limit")
}

func TestLimitIpFail(t *testing.T) {

	configs, err := config.LoadConfig()

	if err != nil {
		println("Error loading config")
		panic(err)
	}

	ipLimit := configs.RequestLimitIp
	blockTimeIp := configs.BlockTimeIp

	time.Sleep(time.Second * time.Duration(blockTimeIp))

	var mapOfResponses = make(map[int]int)

	for i := 1; i <= 2*ipLimit; i++ {
		doRequest(false, &mapOfResponses)
	}

	assert.True(t, mapOfResponses[429] == ipLimit, fmt.Sprint(mapOfResponses[429])+" ip requests were blocked")
	assert.True(t, mapOfResponses[200] == ipLimit, fmt.Sprint(mapOfResponses[200])+" ip requests were successful")

}

func TestLimitTokenOk(t *testing.T) {
	configs, err := config.LoadConfig()
	if err != nil {
		println("Error loading config")
		panic(err)
	}

	tokenLimit := configs.RequestLimitToken
	var mapOfResponses = make(map[int]int)

	for i := 1; i <= tokenLimit; i++ {
		doRequest(true, &mapOfResponses)
	}

	assert.Equal(t, tokenLimit, mapOfResponses[200], "Failed to reach token limit")
}

func TestLimitTokenFail(t *testing.T) {

	configs, err := config.LoadConfig()

	if err != nil {
		println("Error loading config")
		panic(err)
	}

	tokenLimit := configs.RequestLimitToken
	blockTimeToken := configs.BlockTimeToken

	time.Sleep(time.Second * time.Duration(blockTimeToken))

	var mapOfResponses = make(map[int]int)

	for i := 1; i <= 2*tokenLimit; i++ {
		doRequest(true, &mapOfResponses)
	}

	assert.True(t, mapOfResponses[429] == tokenLimit, fmt.Sprint(mapOfResponses[429])+" token requests were blocked")
	assert.True(t, mapOfResponses[200] == tokenLimit, fmt.Sprint(mapOfResponses[200])+" token requests were successful")

}

func doRequest(includeHeader bool, mapOfResponses *map[int]int) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	if includeHeader {
		req.Header.Add("API_KEY", "abc123")
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode

	value, ok := (*mapOfResponses)[code]
	if ok {
		(*mapOfResponses)[code] = value + 1
	} else {
		(*mapOfResponses)[code] = 1
	}
}
