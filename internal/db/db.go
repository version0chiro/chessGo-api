// db for the chess-go-server
package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// This will be dyanmoDB

// create connection to the db

// create ddb connection here

func GetItem(client *dynamodb.Client) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: "johndoe"},
		},
	}

	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error getting item:")
		fmt.Println(err)
	}
	// Access and print the actual values instead of pointers
	username := result.Item["Username"].(*types.AttributeValueMemberS).Value
	home := result.Item["home"].(*types.AttributeValueMemberS).Value

	fmt.Println("Username:", username)
	fmt.Println("Home:", home)

}

func PutItem(client *dynamodb.Client) {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: "Sachin"},
			"home":     &types.AttributeValueMemberS{Value: "New York"},
		},
	}

	result, err := client.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error putting item:")
		fmt.Println(err)
	}
	fmt.Println("Item put successfully:", result)
}

func AddUser(client *dynamodb.Client, username, password string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: username},
			"Password": &types.AttributeValueMemberS{Value: password},
		},
	}
	result, err := client.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error adding user:")
		fmt.Println(err)
	}
	fmt.Println("User added successfully:", result)
	return nil
}

func GetUser(client *dynamodb.Client, username string) (string, string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"Username": &types.AttributeValueMemberS{Value: username},
		},
	}
	result, err := client.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error getting item:")
		fmt.Println(err)
	}
	username = result.Item["Username"].(*types.AttributeValueMemberS).Value
	password := result.Item["Password"].(*types.AttributeValueMemberS).Value
	return username, password, nil
}
