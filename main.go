package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func main() {
    lambda.Start(meshiteroGetUserPostOutlines)
}

type user struct {
    Id string `json:"userId"`
}

type userPostOutline struct {
    PostedTime string `json:"postedTime"`
    SmallImageUrl string `json:"smallImageUrl"`
    PostId string `json:"postId"`
}

func meshiteroGetUserPostOutlines(user user) ([]userPostOutline, error) {
    db := dynamo.New(
        session.New(), 
        &aws.Config{
            Region: aws.String(os.Getenv("DYNAMO_REGION")),
        },
    )
    table := db.Table(os.Getenv("TABLE_NAME"))

    var outlines []userPostOutline
    err := table.Get(os.Getenv("PARTITION_KEY_NAME"), user.Id).All(&outlines)
    return outlines, err
}
