// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// add the following for a short random delay - still far greater than the comparison cost
	//now := time.Now().Nanosecond()
	//for i:=0;i<now%100;i++ {
	//	time.Now()
	//}

	guess := r.URL.Query().Get("guess")
	if "youwillneverguess" == guess {
		fmt.Fprintln(w, "You guessed it!")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
