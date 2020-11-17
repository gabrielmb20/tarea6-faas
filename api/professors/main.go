package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strconv"
)


type CoursesRef struct {
	CourseId int    `json:"course_id"`
	Name  string `json:"name"`
}

type Professors struct {
	Id          int       `json:"_id"`
	Name      string    `json:"profname"`
	Country string    `json:"country"`
	Employee string    `json:"employee"`
	Courses      []CoursesRef    `json:"courses"`
}

var items []Professors

var jsonData string = `[
	{
		"_id": 1,
		"profname": "John Wile",
		"country": "United States",
		"employee": 1807,
		"courses": [
			{
				"course_id": 1,
				"name": "English for Computer Science I"
			},
			{
				"course_id": 2,
				"name": "English for Computer Science II"
			},
			{
				"course_id": 3,
				"name": "English for Computer Science III"
			}
		]
	},
	{
		"_id": 2,
		"profname": "Mary Smith",
		"country": "UK",
		"employee": 2017,
		"courses": [
			{
				"course_id": 4,
				"name": "Programming I"
			},
			{
				"course_id": 5,
				"name": "Programming II"
			}
		]
	},
	{
		"_id": 3,
		"profname": "Joshua Marley",
		"country": "CR",
		"employee": 1203,
		"courses": [
			{
				"course_id": 6,
				"name": "Programming III"
			}
		]
	},
	{
		"_id": 4,
		"profname": "Caroline Andrews",
		"country": "UR",
		"employee": 4392,
		"courses": [
			{
				"course_id": 1,
				"name": "English for Computer Science I"
			}
		]
	}
]`

func FindItem(id int) *Professors {
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
