package main

// requestNumbers routine to call on servers and get the numbers
import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// requestNumbers function manages request to each url
func requestNumbers(ctx context.Context, url string, c chan []byte) {
	body := make(chan []byte)
	go queryURL(url, body)

	select {
	case values := <-body:
		if len(values) > 0 {
			fmt.Print(url, " :::> ")
			c <- values
		} else {
			fmt.Println("No values returned.")
		}

	case <-ctx.Done():
		fmt.Println("Sorry I am taking too long.")
	}
}

// queryURL function actual event when consuming a service
func queryURL(url string, values chan []byte) {
	result, err := http.Get(url)
	if err != nil {
		fmt.Println("Error requesting: ", url)
		return
	}
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println("Error requesting: ", url)
		return
	}
	values <- body
}

// isItemInSlice function to check if the a single int value
//     exists in the slice of integers
func isItemInSlice(item int, intSlice []int) bool {
	sliceLength := len(intSlice)  // Get max length
	maxSlice := (sliceLength / 3) // Devide slice by three to make searching rate faster
	if (sliceLength % 3) > 0 {    // if sliceLength modulus is greater than 0 add 1 to
		maxSlice++ //                  maxSlice so other items at the last part of the
	} //                             slice will be included in the search.

	for counter := 0; counter < maxSlice; counter++ {
		if intSlice[counter] == item {
			//Found item in slice 1
			return true
		}
		counterSlice2 := maxSlice + counter
		if (counterSlice2 < (maxSlice * 2)) && (counterSlice2 < sliceLength) {
			if intSlice[counterSlice2] == item {
				//Found item in slice 2
				return true
			}
		}
		counterSlice3 := (maxSlice * 2) + counter
		if counterSlice3 < sliceLength {
			if intSlice[counterSlice3] == item {
				//Found  item in slice 3
				return true
			}
		}
	}

	return false
}
