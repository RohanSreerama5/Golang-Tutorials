package main

import "net/http"

func main() {
	//This is a very simple web application
	//HandleFunc registers a function lister that is gonna listen on every URL in our application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!")) //We're writing to the ResponseWriter which gives us access to the response to the web request
	}) //end of Handlefunc  //The whole func part is a callback that gets called everytime a URL comes in that matches this route. This prints Hello GO

	err := http.ListenAndServe(":8080", nil) //To see output open chrome and go to localhost:8080
	//ListenAndServe returns an optional err object
	if err != nil { //err can result if the tcp port 8080 is blocked.
		panic(err.Error()) //ListenAndServe won't panic on its own, it will jsut return an err. We have to manually tell Go
		//to panic if there is an err
	} //So once we get the pgm to panic, we want to get the application to recover. Panics don't have to be fatal unless
	//we make the pgm panic all the way up to the Go run-time. We manually panicked here, so no use of run-time

	//If let the Go app panic all the way to the Go run-time, it won't know what to do with a panicking app, so Go will kill it.
	//We want to panic on our own and then recover from it.

}
