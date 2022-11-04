/*
Exercise: Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples
of 5 by "buzz", and all multiples of 15 by "fizzbuzz".
The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

Your goal is to implement a web server that will expose a REST API endpoint that:

Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.

Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2
are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:

- Ready for production
- Easy to maintain by other developers

Bonus Question :
- Add a statistics endpoint allowing users to know what the most frequent request has been.

This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request
*/

package main

import (
	"log"
	"net/http"
	"os"

	"testLBC/cmd/server/handlers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	APIVersion    = "1.0.0"
	debugEnvPath  = "../../.env"
	errLoadingEnv = "Error loading .env file"
)

type MyServer struct {
	r *mux.Router
}

func main() {
	log.Println("API Version : ", APIVersion)

	// Load env
	setEnv()

	// Init server
	app := MyServer{
		r: mux.NewRouter(), // INIT mux router
	}

	// init Middlewares
	app.r.Use(handlers.PrometheusMiddleware)
	prometheus.MustRegister(handlers.TotalRequests)
	common := negroni.New(
		negroni.HandlerFunc(handlers.CorsMiddleware),
	)

	// Routes can be moved in another file, but not necessary for such a tiny module
	/* -- Fizz-Buzz endpoints -- */
	app.r.Handle("/fizz-buzz", common.With(negroni.Wrap(http.HandlerFunc(handlers.GetFizzBuzzHandler)))).
		Methods(http.MethodGet, http.MethodOptions).
		Queries("int1", "{int1:[[:digit:]]+}", "int2", "{int2:[[:digit:]]+}").
		Queries("limit", "{limit:[[:digit:]]+}").
		Queries("str1", "{str1:[[:word:]]+}", "str2", "{str2:[[:word:]]+}") // A more complex regexp could be added here to extend possibilities

	app.r.Handle("/statistics", common.With(negroni.Wrap(http.HandlerFunc(handlers.GetMaxHandler)))).
		Methods(http.MethodGet, http.MethodOptions)

	/* -- Health check endpoints -- */
	app.r.Handle("/metrics", promhttp.Handler())
	app.r.Handle("/healthz", common.With(negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})))).Methods("GET", "OPTIONS")
	app.r.Handle("/liveness", common.With(negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))).Methods("GET", "OPTIONS")

	http.Handle("/V1/", app.r)

	port := ":" + os.Getenv("PORT")
	log.Println("\U0001f37a\tListening on port", port)
	if err := http.ListenAndServe(port, app.r); err != nil {
		panic("Error: " + err.Error())
	}
}

// setEnv Loads environment variables from .env file
func setEnv() {
	// Check if the env vars are already set
	env := os.Getenv("PORT")
	if env == "" {
		// LOAD env
		if err := godotenv.Load(); err != nil {
			err := godotenv.Load(debugEnvPath) // Try the debugging path for the env file
			if err != nil {
				log.Fatal(errLoadingEnv)
			}
		}
	}
}
