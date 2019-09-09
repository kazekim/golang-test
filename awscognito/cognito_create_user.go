/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"fmt"
	"os"
)

func main() {

	emailID := "test@jirawat.kim"
	userPoolID := "PoolIDExample"
	userName := "kazekim"


	// Initialize a session in ap-southeast-1 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	},
	)
	if err != nil {
		fmt.Println("Got error creating session:", err)
		os.Exit(1)
	}

	// Create Cognito service client
	cognitoClient := cognitoidentityprovider.New(sess)

	newUserData := &cognitoidentityprovider.AdminCreateUserInput{
		DesiredDeliveryMediums: []*string{
			aws.String("EMAIL"),
		},
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(emailID),
			},
		},
	}

	newUserData.SetUserPoolId(userPoolID)
	newUserData.SetUsername(userName)

	_, err = cognitoClient.AdminCreateUser(newUserData)
	if err != nil {
		fmt.Println("Got error creating user:", err)
	}
}
