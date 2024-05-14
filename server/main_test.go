package main

import (
	"fmt"
	"github.com/mrjonze/goexpert/rate-limiter/server/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestLimitIpAllOk(t *testing.T) {
	configs, err := config.LoadConfig()
	if err != nil {
		println("Error loading config")
		panic(err)
	}

	ipLimit := configs.RequestLimitIp
	var mapOfResponses = make(map[int]int)

	for i := 1; i <= ipLimit; i++ {
		doRequest(false, &mapOfResponses, false)
	}

	assert.Equal(t, ipLimit, mapOfResponses[200], "Failed to reach ip limit")
}

func TestLimitIpHalfOkHalfFail(t *testing.T) {

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
		doRequest(false, &mapOfResponses, false)
	}

	assert.True(t, mapOfResponses[200] == ipLimit, fmt.Sprint(mapOfResponses[200])+" ip requests were successful")
	assert.True(t, mapOfResponses[429] == ipLimit, fmt.Sprint(mapOfResponses[429])+" ip requests were blocked")
}

func TestLimitTokenAllOk(t *testing.T) {
	configs, err := config.LoadConfig()
	if err != nil {
		println("Error loading config")
		panic(err)
	}

	tokenLimit := configs.RequestLimitToken
	var mapOfResponses = make(map[int]int)

	for i := 1; i <= tokenLimit; i++ {
		doRequest(true, &mapOfResponses, false)
	}

	assert.Equal(t, tokenLimit, mapOfResponses[200], "Failed to reach token limit")
}

func TestLimitTokenHalfOkHalfFail(t *testing.T) {

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
		doRequest(true, &mapOfResponses, false)
	}

	assert.True(t, mapOfResponses[200] == tokenLimit, fmt.Sprint(mapOfResponses[200])+" token requests were successful")
	assert.True(t, mapOfResponses[429] == tokenLimit, fmt.Sprint(mapOfResponses[429])+" token requests were blocked")
}

func TestLimitInvalidTokenHAllFail(t *testing.T) {

	configs, err := config.LoadConfig()

	if err != nil {
		println("Error loading config")
		panic(err)
	}

	tokenLimit := configs.RequestLimitToken
	blockTimeToken := configs.BlockTimeToken

	time.Sleep(time.Second * time.Duration(blockTimeToken))

	var mapOfResponses = make(map[int]int)

	for i := 1; i <= tokenLimit; i++ {
		doRequest(true, &mapOfResponses, true)
	}

	assert.True(t, mapOfResponses[401] == tokenLimit, fmt.Sprint(mapOfResponses[401])+" token requests were unauthorized")
}

func doRequest(includeHeader bool, mapOfResponses *map[int]int, invalidHeader bool) {
	client := &http.Client{}

	configs, err := config.LoadConfig()

	if err != nil {
		println("Error loading config")
		panic(err)
	}

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	if includeHeader {
		if invalidHeader {
			req.Header.Add("API_KEY", "invalidHeader")
		} else {
			req.Header.Add("API_KEY", configs.TokenName)
		}
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
