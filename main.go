package main

import (
	"fmt"
	"log"

	"mag-stadistics-dna-processed-function/src/config/response"
	stadistics "mag-stadistics-dna-processed-function/src/controllers/stadisticsController"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type Request struct {
	ID float64 `json:"id"`
}

var (
	start = lambda.Start
)

func Handler(request Request) (*response.Response, error) {
	fmt.Println(request)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	resp := stadistics.GetStadisticsDnaProcessed()
	fmt.Println(resp)
	
	return &response.Response{
		Message:    "OK",
		StatusCode: 200,
		Body: response.BodyStruct{
			Count_mutant_dna: 0,
			Count_human_dna:  0,
			Ratio:            0,
		},
	}, nil
}

func main() {
	start(Handler)
}
