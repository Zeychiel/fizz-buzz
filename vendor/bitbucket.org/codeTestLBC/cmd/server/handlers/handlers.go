package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"bitbucket.org/codeTestLBC/internal/fizzBuzz"
	"github.com/gorilla/mux"
)

const (
	maxLimit = 1000
)

var (
	requestCounter = make(map[string]int)
)

// GetFizzBuzzHandler handles the calls for the main function of this service
func GetFizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Authentication...

	// Get the request parameters
	vars := mux.Vars(r)
	paramInt1, _ := vars["int1"]   // Existence of the param ensured by the routing
	paramInt2, _ := vars["int2"]   // Existence of the param ensured by the routing
	paramLimit, _ := vars["limit"] // Existence of the param ensured by the routing
	str1, _ := vars["str1"]        // Existence of the param ensured by the routing
	str2, _ := vars["str2"]        // Existence of the param ensured by the routing

	int1, _ := strconv.Atoi(paramInt1)   // Integer validation has already been done in mux
	int2, _ := strconv.Atoi(paramInt2)   // Integer validation has already been done in mux
	limit, _ := strconv.Atoi(paramLimit) // Integer validation has already been done in mux

	// Bonus questions
	// Prometheus counter is more elegant than the package-wide global variale requestCounter
	// TODO: Could be moved in a middleware to be more elegant.
	TotalRequests.WithLabelValues(fmt.Sprintf("%s-%s-%s-%s-%s", paramInt1, paramInt2, paramLimit, str1, str2)).Inc()
	requestCounter[fmt.Sprintf("%s-%s-%s-%s-%s", paramInt1, paramInt2, paramLimit, str1, str2)] += 1

	// Inputs validation
	if int1 < 1 || int2 < 1 {
		NewAPIError(ErrBadRequest.Code, "Nice try !").Throw(w) // This error would be standardized in a larger project
		return
	}
	if limit > maxLimit {
		NewAPIError(ErrBadRequest.Code, "Let's say that the limit is "+strconv.Itoa(maxLimit)+", ok ?").Throw(w) // This error would be standardized in a larger project
		return
	}
	// Edge case
	if int1 > int2 {
		int1, int2 = int2, int1
	}

	// Call the internal function
	rslt, err := fizzBuzz.GetFizzBuzz(int1, int2, limit, str1, str2)
	if err != nil {
		NewAPIError(ErrBadRequest.Code, err.Error()).Throw(w)
		return
	}
	WriteResponse(w, rslt)
}

// GetMaxHandler return the most call fizz-buzz set of values
// requestCounter, storing the history of all the calls, is sorted by value and the highest is returned.
// The underneath logic would be moved in a "utils" package in a more complexe project.
func GetMaxHandler(w http.ResponseWriter, r *http.Request) {
	type kv struct {
		Key   string
		Value int
	}

	// Golang implements slice sorting. First, the map is transformed into a slice.
	var ss []kv
	for k, v := range requestCounter {
		ss = append(ss, kv{k, v})
	}
	// Then this slice can be sorted
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// And max value can be obtained
	if len(ss) > 0 {
		WriteResponse(w, fmt.Sprintf("Item %s called %d times", ss[0].Key, ss[0].Value))
		return
	}
	WriteResponse(w, "No stats yet")
}
