package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type URL struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}

/*
	     quedndh -> URL {
			   	id:quedndh ,
					url: "https://www.google.com",
					short_url: "http://localhost:8080/quedndh",
					created_at: time.Now()
			 }
*/
var DB = make(map[string]URL)

func generateShortURL(originalString string) string {
	//1: converted to byte stream to original array
	hasher := md5.New()
	hasher.Write([]byte(originalString))
	//2: get sum of  hash
	data := hasher.Sum(nil)
	//3: convert to string
	hash := hex.EncodeToString(data)
	saveURL(originalString, hash[:8])
	return hash[:8]
}

func getShortURL(w http.ResponseWriter, r *http.Request) {
	type request struct {
		OrgUrl string `json:"orgurl"`
	}
	var req request
	error := json.NewDecoder(r.Body).Decode(&req)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}
	shortURL := generateShortURL(req.OrgUrl)

	response := map[string]string{"short_url": shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func saveURL(originalurl string, short_url string) string {
	DB[short_url] = URL{
		ID:        short_url,
		URL:       originalurl,
		ShortURL:  short_url,
		CreatedAt: time.Now(),
	}
	return short_url
}

func getURL(shortURL string) (URL, error) {
	url, error := DB[shortURL]
	if !error {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	short_url := strings.TrimPrefix(r.URL.Path, "/redirect/")
	if short_url == "" {
		http.Error(w, "Missing URL ID", http.StatusBadRequest)
		return
	}
	url, error := getURL(short_url)
	if error != nil {
		http.Error(w, error.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.URL, http.StatusSeeOther)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to URL Shortner")
}

func main() {
	fmt.Println("<<<<<<<< Starting Server At 1911 >>>>>>>>")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/short_url", getShortURL)
	http.HandleFunc("/redirect/", redirectHandler)
	error := http.ListenAndServe(":1911", nil)
	if error != nil {
		fmt.Println("Error for starting the server")
	}
	fmt.Println("Server Started")
}
