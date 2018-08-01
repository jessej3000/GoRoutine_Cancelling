#  Go routine cancelling with context


# NOTE: 
The test servers or your api servers needs to be running before the code can be tested.

# To run the code
1. Download or clone my code "git clone https://github.com/jessej3000/travel_a_answer.git"
2. Then go into the directory and run "go run service.go handlers.go functions.go"

# Unit test
1. To test my code just go into the directory 
2. And run "go test -v"

# Files
Program files
- service.go        - contains func main()
- handlers.go       - contains request handlers
- functions.go      - contains supporting functions

Test file
- functions_test.go - testing for functions

# Some explanation of the functions and handlers
- processRequest: This handles request for the endpoint /numbers. It puts all values from url parameters into a slice. And calls and loops through it then calls each url with its individual go routine. The main routine then waits for results or wait for the timeout 450 miliseconds. (I set the timeout to 450 miliseconds to give some allowance for sorting and merging after all possible values are returned). After all possible values are returned or if the timer timed out. It will cancel all go routines that are still running using context.Withcancel. Then merge the unique values and sort the final slice before sending the result to the client.

- requestNumbers: This function manages individual request. While waiting for the result from individual endpoints it also waits for ctx.Done. If ctx.Done comes first, it will cancel it's operation.

- queryURL: The actual consumption of an endpoint.

- isItemInSlice: This function searches for an interger in the slice and returns true if found and false otherwise. It does this by dividing the search into 3 parts of the slice to somehow make the search have a higher rate of finding the equivalent integer in the slice quickly if there are any.
