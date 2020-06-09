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
	"strings"

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

//  get the length of object for our parser
func getObjectLen(value []byte) (int, error) {
	ret := 0
	objectCallback := func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		ret++
		return nil
	}

	if err := jsonparser.ObjectEach(value, objectCallback); err != nil {
		return 0, fmt.Errorf("getObjectLen ObjectEach error: %v", err)
	}

	return ret, nil
}

// build our sources list for redditbooru
func build_redditbooru_sources() {
	// the url for the sources from redditbooru
	url := "https://redditbooru.com/sources/"

	// set 5 second timeout on request
	client := http.Client {
		Timeout: 5 * time.Second,
	}

	// get the content of the page
	resp, err := client.Get(url)
	if err != nil {
		fmt.Print("Error getting redditbooru sources, ")
		panic(err)
	}
	defer resp.Body.Close()

	// read response
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error reading response from redditbooru, ")
		panic(err)
	}

	// get the length of the sources
	outlen, err := getArrayLen(out)

	// set our slice to the appropriate size
	redditbooru_sources = make([]string, outlen)

	// set our sources into the slice
	for i := 0; i < outlen; i++ {
		// pull out the title
		title,_ := jsonparser.GetString(out, "[" + strconv.Itoa(i) + "]", "title")

		// set to lowercase then into our slice
		redditbooru_sources[i] = strings.ToLower(title)
	}
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
			fmt.Println("Error reading response from redditbooru, ", err)
			return
		}

		// randomize the seed for the random int
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

// imgur request
func get_imgur_image(sub string) <-chan string {
	// make channel
	ret := make(chan string)

	go func() {
		defer close(ret)

		// create the proper url with the subreddit
		url := "https://api.imgur.com/3/gallery/r/" + sub + "/time/1"

		// set 5 second timeout on request
		client := http.Client {
			Timeout: 5 * time.Second,
		}

		// create the request
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Client-ID " + auth.imgur)

		// get the content of the page
		resp, err := client.Do(req)
		defer resp.Body.Close()

		// read response
		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response from imgur, ", err)
			return
		}

		// randomize the seed for the random int
		rand.Seed(time.Now().UnixNano())

		// get a random number for the image
		//outlen, _ := getArrayLen(out)
		//random_img := rand.Intn(outlen)

		// parse the data (fix this)
		img_url, _ := jsonparser.GetString(out, "[0]", "[0]", "link")
		fmt.Println(string(img_url))

		ret <- ""
	}()

	return ret
}

// subreddit request
func get_subreddit_image(sub string) <-chan string {
	ret := make(chan string)

	go func() {
		defer close(ret)

		// create the proper url with the subreddit
		url := "https://www.reddit.com/r/" + sub + "/.json?show=all&limit=100"

		// set 5 second timeout on request
		client := http.Client {
			Timeout: 5 * time.Second,
		}

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("User-Agent", "BunnyBot")

		// get the content of the page
		resp, err := client.Do(req)
		defer resp.Body.Close()

		// read response
		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response from reddit, ", err)
			return
		}

		// randomize the seed for the random int
		rand.Seed(time.Now().UnixNano())

		// make sure we aren't grabbing a text post by cylcing through looking for an image
		limit64, _ := jsonparser.GetInt(out, "data", "dist")
		limit := int(limit64) // convert from int64 to int

		// loop through and try to find a post that isn't a text post
		for i := 0; i < limit; i++ {
			// get a random number
			random_img := rand.Intn(limit)

			// check the post hint to see what type of post it is
			hint, _ := jsonparser.GetString(out, "data", "children", "[" + strconv.Itoa(random_img) + "]", "data", "post_hint")

			// make sure that it is an image, or at least a gif
			if hint == "image" || hint == "link" || hint == "rich:video" {
				image, _ := jsonparser.GetString(out, "data", "children", "[" + strconv.Itoa(random_img) + "]", "data", "url")
				ret <- image
				break
			}
		}
	}()

	return ret
}
