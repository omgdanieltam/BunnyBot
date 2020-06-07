/*
This contains our functions used for image searches
*/

package main

import (
	"io/ioutil"
	"net/http"
	"time"
	"math/rand"
	"fmt"
	"strconv"

	"github.com/buger/jsonparser"
)

// get the length of array for our parser
func getArrayLen(value []byte) (int, error) {
	ret := 0
	arrayCallback := func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ret++
	}

	if _, err := jsonparser.ArrayEach(value, arrayCallback); err != nil {
		return 0, fmt.Errorf("getArrayLen ArrayEach error: %v", err)
	}
	return ret, nil
}

// redditbooru request
func get_redditbooru_image(sub string) <-chan string{
	// make the channel
	ret := make(chan string)

	go func() {
		defer close(ret)

		// create the proper url with the subreddit
		url := "https://" + sub + ".redditbooru.com/images/?limit=1000"

		// set 5 second timeout on request
		client := http.Client {
			Timeout: 5 * time.Second,
		}

		// get the content of the page
		resp, err := client.Get(url)
		defer resp.Body.Close()

		// read response
		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// randomize the seed
		rand.Seed(time.Now().UnixNano())

		// get a random number for the image
		outlen,err := getArrayLen(out)
		random_img := rand.Intn(outlen)

		// select a random url from our list
		img_url,err := jsonparser.GetString(out, "[" + strconv.Itoa(random_img) + "]", "cdnUrl")

		// set the return value
		ret <- img_url
		}()

	return ret

}
