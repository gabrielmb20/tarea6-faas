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

type Groups struct {
	Id          int       `json:"_id"`
	Code      string    `json:"code"`
	Location string    `json:"location"`
	Courses      []CoursesRef    `json:"name"`
}

var items []Groups

var jsonData string = `[
	{
		"_id": 1,
		"code": "ECSI-II",
		"location": "A101",
		"courses": [
			{
				"course_id": 1,
				"name": "English for Computer Science I"
			},
			{
				"course_id": 2,
				"name": "English for Computer Science II"
			}
		]
	},
	{
		"_id": 2,
		"code": "ECSIII-PI",
		"location": "A102",
		"courses": [
			{
				"course_id": 3,
				"name": "English for Computer Science III"
			},
			{
				"course_id": 4,
				"name": "Programming I"
			}
		]
	},
	{
		"_id": 3,
		"code": "PII-PIII",
		"location": "A103",
		"courses": [
			{
				"course_id": 5,
				"name": "Programming II"
			},
			{
				"course_id": 6,
				"name": "Programming III"
			}
		]
	},
	{
		"_id": 4,
		"code": "PI",
		"location": "A104",
		"courses": [
			{
				"course_id": 4,
				"name": "Programming I"
			}
		]
	}
]`

func FindItem(id int) *Groups {
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
