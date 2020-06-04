package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"math/rand"
)

func get_redditbooru(sub string) string{
	// create the proper url with the subreddit
	url := "https://" + sub + ".redditbooru.com/images/?limit=1000"
	
	// set 5 second timeout on request
	client := http.Client {
		Timeout: 5 * time.Second,
	}

	// get the content of the page
	resp, err := client.Get(url)

	defer resp.Body.Close()

	// read html as slice of bytes
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// convert output to string
	jsonout := string(out)

	// convert to proper json
	var results []map[string] interface {}
	json.Unmarshal([]byte(jsonout), &results)

	// randomize the seed
	rand.Seed(time.Now().UnixNano())
	
	// select a random image to return
	return results[rand.Intn(len(results))]["cdnUrl"].(string)
}

