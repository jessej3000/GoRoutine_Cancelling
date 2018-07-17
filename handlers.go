package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

// invalidEndpoint handles access to root /
func invalidEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invalid endpoint please try /numbers."))
}

// processRequest handles request coming in from numbers/
func processRequest(w http.ResponseWriter, r *http.Request) {
	paramValues, _ := r.URL.Query()["u"] // Get url values of parameter 'u'

	if len(paramValues) > 0 { // If parameters exist then loop through it and make a request

		c := make(chan []byte, 10) // Declare a channel c where results will be returned back
		var nums map[string][]int
		var combinedNumbers []int

		ctx := context.Background() // Use context to cancel running child go routines after deadline of 500 miliseconds
		ctx, cancel := context.WithCancel(ctx)

		for _, value := range paramValues {
			go requestNumbers(ctx, value, c) // execute request on their own go routines and wait for result in c channel
		}

		timeout := time.After(450 * time.Millisecond) // NOTE: I set deadeline to 450 miliseconds before the for loop
		for {                                         // to give allowance of 50 miliseconds for the merging
			select { //                                    and sorting algorithm below
			case results := <-c: // Wait for the results here
				fmt.Println(string(results))
				err := json.Unmarshal(results, &nums)
				if err != nil {
					fmt.Println("Invalid result. ", err.Error(), " Continuing process.") // Report invalid result and continue
				} else {
					arrayNums := nums["numbers"]
					for counter := 0; counter < len(arrayNums); counter++ {
						if isItemInSlice(arrayNums[counter], combinedNumbers) == false {
							combinedNumbers = append(combinedNumbers, []int{arrayNums[counter]}...) // MERGING: Combine all unique integers into one integer slice
						}
					}
				}
			case <-timeout: // Deadline comes here and cancel remaining go routines requesting from urls
				fmt.Println("Time Up")
				sort.Slice(combinedNumbers, func(i, j int) bool { // SORTING: Sort slice
					return combinedNumbers[i] < combinedNumbers[j]
				})
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]interface{}{"numbers": combinedNumbers}) // Return result to client
				cancel()                                                                      // Cancel all unfinished go routines
				return
			}
		}
	} else { // URL has no u parameters
		w.Write([]byte("No valid parameters."))
	}
}
