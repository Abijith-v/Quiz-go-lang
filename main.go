package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Question struct {
	Id       int
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
	Answer   string
}

type QuizResponse struct {
	Id       int
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
	Answer   string
	ErrorMsg string
	Status   int64
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/quiz", getRandom).Methods("GET")

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("Server listening to port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func getRandom(responseWriter http.ResponseWriter, request *http.Request) {

	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var allQuestions []Question = loadQuestions()
	len := len(allQuestions)

	randomTime := rand.NewSource(time.Now().UnixNano())
	randomNumberOutOfRange := rand.New(randomTime)
	randomIndex := randomNumberOutOfRange.Intn(200) % len

	rawResponse := QuizResponse{
		Id:       allQuestions[randomIndex].Id,
		Question: allQuestions[randomIndex].Question,
		OptionA:  allQuestions[randomIndex].OptionA,
		OptionB:  allQuestions[randomIndex].OptionB,
		OptionC:  allQuestions[randomIndex].OptionC,
		OptionD:  allQuestions[randomIndex].OptionD,
		Answer:   allQuestions[randomIndex].Answer,
		ErrorMsg: "",
		Status:   http.StatusOK,
	}
	jsonResponse, err := json.Marshal(rawResponse)

	if err != nil {
		log.Fatal(err)
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}

	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, string(jsonResponse)+"\n")
}

func loadQuestions() []Question {

	var questions = getQuestions()
	var allQuestions []Question
	for index, question := range questions {
		allQuestions = append(allQuestions, Question{
			Id:       index,
			Question: question[0],
			OptionA:  question[1],
			OptionB:  question[2],
			OptionC:  question[3],
			OptionD:  question[4],
			Answer:   question[5],
		})
	}

	return allQuestions
}
