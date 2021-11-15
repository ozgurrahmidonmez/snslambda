package main

import (
	"errors"
	"fmt"
"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
)

type request struct {
	CustomerId int `json:"customerId" validate:"required"`
	Id         int `json:"id" validate:"required"`
}


func HandleLambdaEvent(r request) error {
	if rand.Intn(2) == 0 {
		fmt.Printf("customer with id %d with message id %d processed", r.CustomerId,r.Id)
		return nil
	}else{
		return errors.New("retry again, you are unlucky")
	}

}

func main() {
	lambda.Start(HandleLambdaEvent)
}
