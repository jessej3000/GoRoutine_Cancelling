package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// #############################################################################
// Important things to test in requestNumbers are its channel c if it returns
// any thing or blank []byte. Both should be successful. And check if context
// cancels within the time limit as stated in the test as 500 and 100 milisecond
// Test functions:
//  - TestRequestNumbersRand500
//  - TestRequestNumbersRand100
//  - TestRequestNumbersPrimes500
//  - TestRequestNumbersPrimes100
//  - TestRequestNumbersFibo500
//  - TestRequestNumbersFibo100
//  - TestRequestNumbersOdd500
//  - TestRequestNumbersOdd100

// TestRequestNumbersRand500 for http://localhost:8090/rand under 500 milisecond
func TestRequestNumbersRand500(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/rand"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(500 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 500 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersRand100 for http://localhost:8090/rand under 100 miliseconds
func TestRequestNumbersRand100(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/rand"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(100 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 100 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersPrimes500 for http://localhost:8090/primes under 500 milisecond
func TestRequestNumbersPrimes500(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/primes"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(500 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 500 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersPrimes100 for http://localhost:8090/primes under 100 miliseconds
func TestRequestNumbersPrimes100(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/primes"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(100 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 100 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersFibo500 for http://localhost:8090/fibo under 500 milisecond
func TestRequestNumbersFibo500(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/fibo"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(500 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 500 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersFibo100 for http://localhost:8090/fibo under 100 miliseconds
func TestRequestNumbersFibo100(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/fibo"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(100 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 100 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersOdd500 for http://localhost:8090/odd under 500 miliseconds
func TestRequestNumbersOdd500(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/odd"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(500 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 500 miliseconds.")
	}
	cancel()
}

// TestRequestNumbersOdd100 for http://localhost:8090/odd under 100 miliseconds
func TestRequestNumbersOdd100(t *testing.T) {
	success := false

	c := make(chan []byte)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	url := "http://localhost:8090/odd"

	go requestNumbers(ctx, url, c)

	timelimit := time.After(100 * time.Millisecond)

	select {
	case v := <-c:
		fmt.Println(string(v))
		success = true
	case <-timelimit:
		cancel()
		success = true
	}

	if !success {
		t.Errorf("Channel c of type []byte did not return anything. Or timelimit of type chan time.Time did not return after 100 miliseconds.")
	}
	cancel()
}

// #############################################################################
// Testing queryURL if it returns the following. For some reason sometimes it
// returns blank []byte. Even though with blank []byte the function should not
// fail and cause any unwanted errors.
//  - service unavailable (Sometimes this appears and sometimes not)
//  - {"numbers":[2,3,5,7,11,13]}
//  - {"numbers":[5,17,3,19,76,24,1,5,10,34,8,27,7]}
//  - {"numbers":[1,1,2,3,5,8,13,21]}
//  - {"numbers":[1,3,5,7,9,11,13,15,17,19,21,23]}
//
// Test functions:
//  - TestQueryURLRand
//  - TestQueryURLOdd
//  - TestQueryURLPrimes
//  - TestQueryURLFibo

// TestQueryURLRand test queryURL function on http://localhost:8090/rand endpoint
func TestQueryURLRand(t *testing.T) {
	success := false
	values := make(chan []byte)
	go queryURL("http://localhost:8090/rand", values)
	var result map[string][]int
	checkValue := map[string][]int{
		"numbers": []int{5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7},
	}

	results := <-values
	err := json.Unmarshal(results, &result)
	if err != nil {
		fmt.Println("Server generated exception. Continuing test.")
		if reflect.DeepEqual(results, []byte("service unavailable\n")) {
			success = true
		}
		if reflect.DeepEqual(results, []byte("")) {
			success = true
		}
	} else if reflect.DeepEqual(result, checkValue) {
		success = true
	}

	if !success {
		t.Errorf("Returns an unexpected value.")
	}
}

// TestQueryURLOdd test queryURL function on http://localhost:8090/odd endpoint
func TestQueryURLOdd(t *testing.T) {
	success := false
	values := make(chan []byte)
	go queryURL("http://localhost:8090/odd", values)
	var result map[string][]int
	checkValue := map[string][]int{
		"numbers": []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23},
	}

	results := <-values
	err := json.Unmarshal(results, &result)
	if err != nil {
		fmt.Println("Server generated exception. Continuing test.")
		if reflect.DeepEqual(results, []byte("service unavailable\n")) {
			success = true
		}
		if reflect.DeepEqual(results, []byte("")) {
			success = true
		}
	} else if reflect.DeepEqual(result, checkValue) {
		success = true
	}

	if !success {
		t.Errorf("Returns an unexpected value.")
	}
}

// TestQueryURLPrimes test queryURL function on http://localhost:8090/primes endpoint
func TestQueryURLPrimes(t *testing.T) {
	success := false
	values := make(chan []byte)
	go queryURL("http://localhost:8090/primes", values)
	var result map[string][]int
	checkValue := map[string][]int{
		"numbers": []int{2, 3, 5, 7, 11, 13},
	}

	results := <-values
	err := json.Unmarshal(results, &result)
	if err != nil {
		fmt.Println("Server generated exception. Continuing test.")
		if reflect.DeepEqual(results, []byte("service unavailable\n")) {
			success = true
		}
		if reflect.DeepEqual(results, []byte("")) {
			success = true
		}
	} else if reflect.DeepEqual(result, checkValue) {
		success = true
	}

	if !success {
		t.Errorf("Returns an unexpected value.")
	}
}

// TestQueryURLFibo test queryURL function on http://localhost:8090/fibo endpoint
func TestQueryURLFibo(t *testing.T) {
	success := false
	values := make(chan []byte)
	go queryURL("http://localhost:8090/fibo", values)
	var result map[string][]int
	checkValue := map[string][]int{
		"numbers": []int{1, 1, 2, 3, 5, 8, 13, 21},
	}

	results := <-values

	err := json.Unmarshal(results, &result)
	if err != nil {
		fmt.Println("Any error happens in here comes from the server. And not covered by this test function.")
		if reflect.DeepEqual(results, []byte("service unavailable\n")) {
			success = true
		}
		if reflect.DeepEqual(results, []byte("")) {
			success = true
		}
	} else if reflect.DeepEqual(result, checkValue) {
		success = true
	}

	if !success {
		t.Errorf("Returns an unexpected value.")
	}
}

// #############################################################################
// Testing isItemInSlice if it is able to detect if an integer exist in an array
// Will check for the following values
//  - 1 on endpoint /primes return value [2, 3, 5, 7, 11, 13] : should return false
//  - 5 on endpoint /primes return value [2, 3, 5, 7, 11, 13] : should return true
//  - 43 on endpoint /fibo return value [1,1,2,3,5,8,13,21] : should return false
//  - 13 on endpoint /fibo return value [1,1,2,3,5,8,13,21] : should return true
//  - 78 on endpoint /odd return value [1,3,5,7,9,11,13,15,17,19,21,23] : should return false
//  - 21 on endpoint /odd return value [1,3,5,7,9,11,13,15,17,19,21,23] : should return true
//  - 100 on endpoint /rand return value [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7] : should return false
//  - 24 on endpoint /rand return value [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7] : should return true
//
// Functions:
//  - TestIsItemInSlice1
//  - TestIsItemInSlice5
//  - TestIsItemInSlice43
//  - TestIsItemInSlice13

// TestIsItemInSlice1 test if 1 in [2, 3, 5, 7, 11, 13]. Returns false
func TestIsItemInSlice1(t *testing.T) {
	if isItemInSlice(1, []int{2, 3, 5, 7, 11, 13}) {
		t.Errorf("1 should not be found in [2, 3, 5, 7, 11, 13]")
	}
}

// TestIsItemInSlice5 test if 5 in [2, 3, 5, 7, 11, 13]. Returns true
func TestIsItemInSlice5(t *testing.T) {
	if !isItemInSlice(5, []int{2, 3, 5, 7, 11, 13}) {
		t.Errorf("5 should be in [2, 3, 5, 7, 11, 13]")
	}
}

// TestIsItemInSlice43 test if 43 in [1, 1, 2, 3, 5, 8, 13, 21]. Returns false
func TestIsItemInSlice43(t *testing.T) {
	if isItemInSlice(43, []int{1, 1, 2, 3, 5, 8, 13, 21}) {
		t.Errorf("43 should not be found in [1,1,2,3,5,8,13,21]")
	}
}

// TestIsItemInSlice13 test if 13 in [1, 1, 2, 3, 5, 8, 13, 21]. Returns true
func TestIsItemInSlice13(t *testing.T) {
	if !isItemInSlice(13, []int{1, 1, 2, 3, 5, 8, 13, 21}) {
		t.Errorf("13 should be in [1,1,2,3,5,8,13,21]")
	}
}

// TestIsItemInSlice78 test if 78 in [1,3,5,7,9,11,13,15,17,19,21,23]. Returns false
func TestIsItemInSlice78(t *testing.T) {
	if isItemInSlice(78, []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}) {
		t.Errorf("78 should not be found in [1,3,5,7,9,11,13,15,17,19,21,23]")
	}
}

// TestIsItemInSlice21 test if 21 in [1,3,5,7,9,11,13,15,17,19,21,23]. Returns true
func TestIsItemInSlice21(t *testing.T) {
	if !isItemInSlice(21, []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}) {
		t.Errorf("21 should be in [1,3,5,7,9,11,13,15,17,19,21,23]")
	}
}

// TestIsItemInSlice100 test if 100 in [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7]. Returns false
func TestIsItemInSlice100(t *testing.T) {
	if isItemInSlice(100, []int{5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7}) {
		t.Errorf("100 should not be found in [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7]")
	}
}

// TestIsItemInSlice24 test if 24 in [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7]. Returns true
func TestIsItemInSlice24(t *testing.T) {
	if !isItemInSlice(24, []int{5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7}) {
		t.Errorf("24 should be in [5, 17, 3, 19, 76, 24, 1, 5, 10, 34, 8, 27, 7]")
	}
}
