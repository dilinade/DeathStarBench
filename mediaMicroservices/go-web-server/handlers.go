package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/apache/thrift/tutorial/go/gen-go/media"
	"log"
	"net/http"
	"strconv"
)

var (
	defaultCtx = context.Background()
	cfg        = &thrift.TConfiguration{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	protocol   = thrift.NewTBinaryProtocolFactoryConf(cfg)
	k8s_suffix = ".default..default.10.110.187.38.sslip.io"
	port       = ":80"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.Form.Get() method to retrieve the relevant data fields
	// from the r.Form map.
	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	userName := r.Form.Get("username")
	password := r.Form.Get("password")

	if firstName == "" || lastName == "" || userName == "" || password == "" {
		fmt.Fprintf(w, "apiRegisterUser %s: %s \n", firstName, "Missing required fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Registering user: " + firstName + " " + lastName)
	transport, err := getConnectionForUrl("http://user-service-knative" + k8s_suffix + port)
	client := media.NewUserServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.RegisterUser(defaultCtx, 0, firstName, lastName, userName, password, carrier)
	fmt.Fprintf(w, "apiRegisterUser %s: %s \n", firstName, err)
}

func registerMovie(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.Form.Get() method to retrieve the relevant data fields
	// from the r.Form map.
	movieId := r.Form.Get("movie_id")
	title := r.Form.Get("title")

	if movieId == "" || title == "" {
		fmt.Fprintf(w, "apiRegisterMovie %s: %s \n", movieId, "Missing required fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Registering movie: %s %s\n", movieId, title)
	transport, err := getConnectionForUrl("http://movie-id-service-knative" + k8s_suffix + port)
	client := media.NewMovieIdServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.RegisterMovieId(defaultCtx, 0, title, movieId, carrier)
	fmt.Fprintf(w, "apiRegisterMovie %s: %s \n", movieId, err)
}

type plot struct {
	PlotId int64  `json:"plot_id"`
	Plot   string `json:"plot"`
}

func plotWrite(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	var p plot

	err := decodeJSONBody(w, r, &p)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	fmt.Println("Creating thttp client")
	fmt.Println("Creating thttp client")
	transport, err := getConnectionForUrl("http://plot-service-knative" + k8s_suffix + port)
	client := media.NewPlotServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.WritePlot(defaultCtx, 0, p.PlotId, p.Plot, carrier)
	fmt.Fprintf(w, "apiPlotWrite %s: %s \n", p, err)
}

func movieInfoWrite(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	var m media.MovieInfo
	err := decodeJSONBody(w, r, &m)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	fmt.Println("Creating thttp client")
	transport, err := getConnectionForUrl("http://movie-info-service-knative" + k8s_suffix + port)
	client := media.NewMovieInfoServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	// convert m.AvgRating from float64 to string
	avgRatingStr := strconv.FormatFloat(m.AvgRating, 'f', 2, 64)
	err = client.WriteMovieInfo(defaultCtx, 0, m.MovieID, m.Title, m.Casts, m.PlotID, m.ThumbnailIds, m.PhotoIds, m.VideoIds, avgRatingStr, m.NumRating, carrier)
	fmt.Fprintf(w, "apiMovieInfoWrite %s: %s \n", m, err)
}

func castInfoWrite(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	var c media.CastInfo

	err := decodeJSONBody(w, r, &c)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	fmt.Println("Creating thttp client")
	transport, err := getConnectionForUrl("http://cast-info-service-knative" + k8s_suffix + port)
	client := media.NewCastInfoServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.WriteCastInfo(defaultCtx, 0, c.CastInfoID, c.Name, c.Gender, c.Intro, carrier)
	fmt.Fprintf(w, "apiCastInfoWrite %s: %s \n", c, err)
}

func reviewCompose(w http.ResponseWriter, r *http.Request) {
	// ParseForm parses the raw query from the URL and updates r.Form.
	// ParseForm parses the raw query from the URL and updates r.Form.
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		return
	}

	// Use the r.Form.Get() method to retrieve the relevant data fields
	// from the r.Form map.
	title := r.Form.Get("title")
	text := r.Form.Get("text")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	rating := r.Form.Get("rating")

	if title == "" || text == "" || username == "" || password == "" || rating == "" {
		fmt.Fprintf(w, "apiReviewCompose %s: %s \n", title, "Missing required fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// convert rating from string to int32
	ratingInt, err := strconv.ParseInt(rating, 10, 32)
	if err != nil {
		fmt.Fprintf(w, "apiReviewCompose %s: %s \n", title, err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	// error chan to receive error from goroutines
	errChan := make(chan error, 4)
	// 4 concurrent calls
	go uploadUserId(username, errChan)
	go uploadText(text, errChan)
	go uploadMovieId(title, int32(ratingInt), errChan)
	go uploadUniqueId(errChan)

	// wait for all goroutines to finish
	for i := 0; i < 4; i++ {
		err := <-errChan
		if err != nil {
			fmt.Fprintf(w, "apiReviewCompose %s: %s \n", title, err)
			//w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	fmt.Fprintf(w, "apiReviewCompose %s: %s \n", title, "success")
}
func uploadUserId(username string, errChan chan error) {
	transport, err := getConnectionForUrl("http://user-service-knative" + k8s_suffix + port)
	client := media.NewUserServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.UploadUserWithUsername(defaultCtx, 0, username, carrier)
	fmt.Printf("uploadUserId %s err %s\n", username, err)
	errChan <- err
}

func uploadText(text string, errChan chan error) {
	transport, err := getConnectionForUrl("http://text-service-knative" + k8s_suffix + port)
	client := media.NewTextServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.UploadText(defaultCtx, 0, text, carrier)
	fmt.Printf("uploadText %s err %s\n", text, err)
	errChan <- err
}

func uploadMovieId(title string, rating int32, errChan chan error) {
	transport, err := getConnectionForUrl("http://movie-id-service-knative" + k8s_suffix + port)
	client := media.NewMovieIdServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.UploadMovieId(defaultCtx, 0, title, rating, carrier)
	fmt.Printf("uploadMovieId %s rating %d err %s\n", title, rating, err)
	errChan <- err
}

func uploadUniqueId(errChan chan error) {
	transport, err := getConnectionForUrl("http://unique-id-service-knative" + k8s_suffix + port)
	client := media.NewUniqueIdServiceClientFactory(transport, protocol)
	defer transport.Close()
	var carrier map[string]string
	err = client.UploadUniqueId(defaultCtx, 0, carrier)
	fmt.Printf("uploadUniqueId err %s\n", err)
	errChan <- err
}
