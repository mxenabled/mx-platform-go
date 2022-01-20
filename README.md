*This project is currently in **Beta**. Please open up an issue [here](https://github.com/mxenabled/mx-platform-go/issues) to report issues using the MX Platform Go library.*

# MX Platform Go

The [MX Platform API](https://www.mx.com/products/platform-api) is a powerful, fully-featured API designed to make aggregating and enhancing financial data easy and reliable. It can seamlessly connect your app or website to tens of thousands of financial institutions.

## Documentation

Examples for the API endpoints can be found [here.](docs/MxPlatformApi.md)

## Requirements

- Go >= 1.13

## Installation

Make sure your project has a `go.mod` file. (it is using Go Modules)

Then `go get` the latest version of this library with the following command.
```shell
go get github.com/mxenabled/mx-platform-go
```

## Getting Started

In order to make requests, you will need to [sign up](https://dashboard.mx.com/sign_up) for the MX Platform API and get a `Client ID` and `API Key`.

Please follow the [installation](#installation) procedure and then run the following code to create your first User:
```go
package main

import (
	"context"
	"fmt"
	"github.com/mxenabled/mx-platform-go"
	"os"
)

func main() {
	configuration := mxplatformgo.NewConfiguration()
	api_client := mxplatformgo.NewAPIClient(configuration)

	// Configure environment. 0 for production, 1 for development
	ctx := context.WithValue(context.Background(), mxplatformgo.ContextServerIndex, 1)

	// Configure with your Client ID/API Key from https://dashboard.mx.com
	ctx = context.WithValue(ctx, mxplatformgo.ContextBasicAuth, mxplatformgo.BasicAuth{
		UserName: "Your Client ID",
		Password: "Your API Key",
	})

	userCreateRequest := *mxplatformgo.NewUserCreateRequest()
	userCreateRequest.SetMetadata("Creating a user!")
	userCreateRequestBody := *mxplatformgo.NewUserCreateRequestBody()
	userCreateRequestBody.SetUser(userCreateRequest)

	resp, r, err := api_client.MxPlatformApi.CreateUser(ctx).UserCreateRequestBody(userCreateRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MxPlatformApi.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	fmt.Printf("Response from `MxPlatformApi.CreateUser`: %#v\n", resp)
}
```

## Development

This project was generated by the [OpenAPI Generator](https://openapi-generator.tech). To generate this library, verify you have the latest version of the `openapi-generator-cli` found [here.](https://github.com/OpenAPITools/openapi-generator#17---npm)

Running the following command in this repo's directory will generate this library using the [MX Platform API OpenAPI spec](https://github.com/mxenabled/openapi/blob/master/openapi/mx_platform_api_beta.yml) with our [configuration and templates.](https://github.com/mxenabled/mx-platform-go/tree/master/openapi)
```shell
openapi-generator-cli generate \
-i https://raw.githubusercontent.com/mxenabled/openapi/master/openapi/mx_platform_api_beta.yml \
-g go \
-c ./openapi/config.yml \
-t ./openapi/templates
```

## Contributing

Please [open an issue](https://github.com/mxenabled/mx-platform-go/issues) or [submit a pull request.](https://github.com/mxenabled/mx-platform-go/pulls)