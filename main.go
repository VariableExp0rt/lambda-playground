package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type NewFunction func(c *CreateFunctionInput)

type CreateFunctionInput struct {
	Code         lambda.FunctionCode
	Description  string
	FunctionName string
	Role         string
	Runtime      string
}

func printLambdaConfig(c *CreateFunctionInput) {
	fmt.Printf("%v\n%v\n%v", c.Runtime, c.Description, c.FunctionName)
}

func createNewFunction(c *CreateFunctionInput, n NewFunction) {
	n(c)
}

// main executes the program, specifically it will ensure the following;
// - Setup a new client with the appropriate credentials and logger
// - Create a new s3 bucket to host a deployment package (referenced in CreateNewFunction struct)
// - Create the Lambda function from the given deployment package
// - Create an APIGateway Service with a POST method for triggering the Lambda function
// - Lambda function will create an EKS cluster
func main() {

	sess := session.Must(session.NewSession())

	newLambda := lambda.New(sess)

	_, err := newLambda.CreateFunction(//TODO)
	//must be zipped deployment package, add check here
	pkg := bufio.NewScanner(os.Stdin)
	if !pkg.Scan() {
		fmt.Errorf("Error reading file: ", pkg.Err())
		return
	}

	newpkg := pkg.Bytes()

	a := &CreateFunctionInput{
		Description:  "Testing some new Golang tricks",
		FunctionName: "my-new-function",
		Runtime:      "go1.x",
		Code:         lambda.FunctionCode{ZipFile: newpkg},
	}
	createNewFunction(a, printLambdaConfig)

}
