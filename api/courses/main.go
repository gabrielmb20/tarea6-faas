package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)

type GroupRef struct {
	GroupId int    `json:"group_id"`
	Code  string `json:"code"`
}

type ProfessorRef struct {
	ProfessorId int    `json:"professor_id"`
	Name  string `json:"profame"`
}

type Courses struct {
	Id          int       `json:"_id"`
	Name      string    `json:"name"`
	Period string    `json:"period"`
	Year   int       `json:"year"`
	Group_Id      ]GroupRef    `json:"group_id"`
	Code       ]GroupRef `json:"code"`
	Professor_Id      ]ProfessorRef    `json:"professor_id"`
	Professor_Name      ]ProfessorRef    `json:"profname"`
}

var items []Courses

var jsonData string = `[
	{
		"_id": 1,
		"name": "English for Computer Science I",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "ECSI-II",
		"professor_id": 4,
		"profname": "Caroline Andrews"
	},
	{
		"_id": 2,
		"name": "English for Computer Science II",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "ECSI-II",
		"professor_id": 1,
		"profname": "John Wile"
	},
	{
		"_id": 3,
		"name": "English for Computer Science III",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "ECSIII-PI",
		"professor_id": 1,
		"profname": "John Wile"
	},
	{
		"_id": 4,
		"name": "Programming I",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "ECSIII-PI",
		"professor_id": 2,
		"profname": "Mary Smith"
	},
	{
		"_id": 5,
		"name": "Programming II",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "PII-PIII",
		"professor_id": 2,
		"profname": "Mary Smith"
	},
	{
		"_id": 6,
		"name": "Programming III",
		"period": "II",
		"year": 2020,
		"group_id": 1,
		"code": "PII-PIII",
		"professor_id": 3,
		"profname": "Joshua Marley"
	}
]`

func FindItem(id int) *Courses {
	for _, item := range items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := req.QueryStringParameters["id"]
	var data []byte
	if id == "" {
		data, _ = json.Marshal(items)
	} else {
		param, err := strconv.Atoi(id)
		if err == nil {
			item := FindItem(param)
			if item != nil {
				data, _ = json.Marshal(*item)
			} else {
				data = []byte("error\n")
			}
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(data),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	_ = json.Unmarshal([]byte(jsonData), &items)
	lambda.Start(handler)
}
