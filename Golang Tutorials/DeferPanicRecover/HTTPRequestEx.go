package main

//Program that makes a resource request using the HTTP package
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt") //using Get func from HTTP package to request the robots.txt file from google
	if err != nil {                                          //So we will get a response (res) and an optional error (err)
		log.Fatal(err) //If we have an error, we log it and exit the app
	}

	robots, err := ioutil.ReadAll(res.Body) //Use the ReadAll func of ioutil package. THis will take a stream and parse that out to a slice of bytes to
	//work with
	defer res.Body.Close() //close the body of the response to the let the web request know we're done working with it, so resources can be freed
	//NOTE: we defer the closing of the body, so the body won't close till this func is done. This is useful so you don't have to know where to place
	//the body close Benefit: We clearly open the resource and close them neatly
	if err != nil { //if read operation above fails, we log that err
		log.Fatal(err)
	}

	fmt.Printf("%s", robots) //Note we used printf so that the []byte (robots) (stream of bytes) gets converted to its actual string representation from
	//Unicode
	//fmt.Println(robots) //If I just print like this then I only get numbers, the Unicode representation of the strings.

	//When looping, not best to use defer because if you loop over and open a buncha requests and defer all of hteir closures, you'll
	//have a bunch of requests open and taking up memory.
}
