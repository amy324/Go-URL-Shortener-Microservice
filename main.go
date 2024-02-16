package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the Redis client
	var err error
	redisClient, err = initRedisClient()
	if err != nil {
		log.Fatalf("Failed to initialize Redis client: %v", err)
	}
}

func initRedisClient() (*redis.Client, error) {
	// Initialize the Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_STRING"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// Test the connection to Redis
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis")
	return client, nil
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Load environment variables
	os.Getenv("REDIS_STRING")
	 os.Getenv("REDIS_PASSWORD")

	// Initialize the Gorilla Mux router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/shorten", shortenURLHandler).Methods("POST")
	router.HandleFunc("/{shortURL}", redirectHandler).Methods("GET")

	//Handler for checking if the microservice is live
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL shortener microservice is live")
	})


	// Start the HTTP server
	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}



func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the JSON request body to get the long URL
    var requestBody struct {
        URL string `json:"url"`
    }
    err := json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        log.Printf("Error parsing request body: %v", err)
        return
    }

    // Get the long URL from the request body
    longURL := requestBody.URL

    // Generate short URL
    shortURL := generateShortURL(longURL)

    // Store the mapping in Redis
    err = storeURLMapping(redisClient, shortURL, longURL)
    if err != nil {
        http.Error(w, "Failed to store URL mapping", http.StatusInternalServerError)
        log.Printf("Error storing URL mapping in Redis: %v", err)
        return
    }

    // Return both short URL and long URL in the response
    jsonResponse := struct {
        ShortURL string `json:"short_url"`
        LongURL  string `json:"long_url"`
    }{
        ShortURL: shortURL,
        LongURL:  longURL,
    }
    jsonResp, err := json.Marshal(jsonResponse)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        log.Printf("Error marshalling JSON response: %v", err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)

    log.Printf("Shortened URL created: %s -> %s", longURL, shortURL)
}


func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the request path
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]
	log.Printf("Redirect requested for short URL: %s", shortURL)

	// Retrieve the long URL from Redis
	longURL, err := getLongURL(redisClient, shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Printf("Error retrieving long URL from Redis for short URL %s: %v", shortURL, err)
		return
	}

	// Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusFound)
	log.Printf("Redirecting to long URL: %s", longURL)
}


