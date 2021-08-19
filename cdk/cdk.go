package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdago"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	bundlingOptions := &awslambdago.BundlingOptions{
		GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w"`)},
	}

	welcome := awslambdago.NewGoFunction(stack, jsii.String("cdk-demo-lambda"), &awslambdago.GoFunctionProps{
		Runtime:  awslambda.Runtime_GO_1_X(),
		Entry:    jsii.String("../api/handlers/welcome/post"),
		Bundling: bundlingOptions,
		Tracing:  awslambda.Tracing_ACTIVE,
	})

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("cdk-demo-api"), &awsapigateway.LambdaRestApiProps{
		Handler: welcome,
		Proxy:   jsii.Bool(false),
	})
	apiResourceOpts := &awsapigateway.ResourceOptions{}
	apiLambdaOpts := &awsapigateway.LambdaIntegrationOptions{}
	// /welcome
	welcomeRestURL := api.Root().AddResource(jsii.String("welcome"), apiResourceOpts)
	welcomeIntegration := awsapigateway.NewLambdaIntegration(welcome, apiLambdaOpts)
	welcomeRestURL.AddMethod(jsii.String("POST"), welcomeIntegration, &awsapigateway.MethodOptions{})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
