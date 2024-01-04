package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"strconv"
	//"github.com/go-chi/chi/v5/middleware"
)

func getConnectionForUrl(url string) (thrift.TTransport, error) {
	transport, err := thrift.NewTHttpClient(url)
	if err != nil {
		fmt.Print("Error opening socket:", err)
	}
	if err := transport.Open(); err != nil {
		fmt.Print("Error opening transport:", err)
	}
	return transport, err
}

func debug(w http.ResponseWriter, r *http.Request) {
	// get the environment variable fqdn_suffix
	fqdnSuffix := os.Getenv("fqdn_suffix")
	fmt.Fprintf(w, "debug %s\n", fqdnSuffix)
}

func main() {
	fqdnSuffix := os.Getenv("fqdn_suffix")
	if fqdnSuffix != "" {
		k8s_suffix = fqdnSuffix
		fmt.Printf("fqdn_suffix set to %s\n", fqdnSuffix)
	}
	portEnv := os.Getenv("port")
	if portEnv != "" {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			fmt.Printf("port %s is not a valid integer\n", portEnv)
		} else {
			port = portInt
			fmt.Printf("port set to %s\n", portEnv)
		}
	}
	prefixEnv, ok := os.LookupEnv("prefix")
	if ok {
		prefix = prefixEnv
		fmt.Printf("prefix set to %s\n", prefixEnv)
	}
	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Post("/go-api/user/register", registerUser)
	r.Post("/go-api/plot/write", plotWrite)            //done
	r.Post("/go-api/movie-info/write", movieInfoWrite) //done
	r.Post("/go-api/movie/register", registerMovie)    //done
	r.Post("/go-api/cast-info/write", castInfoWrite)   //done
	r.Post("/go-api/review/compose", reviewCompose)
	r.Get("/debug", debug)
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
