/*
MX Platform API

Testing TransactionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mxplatformgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mxenabled/mx-platform-go"
)

func Test_mxplatformgo_TransactionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransactionsAPIService UsersUserGuidAccountsAccountGuidTransactionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGuid string
		var accountGuid string

		resp, httpRes, err := apiClient.TransactionsAPI.UsersUserGuidAccountsAccountGuidTransactionsPost(context.Background(), userGuid, accountGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}