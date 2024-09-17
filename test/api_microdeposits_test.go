/*
MX Platform API

Testing MicrodepositsAPIService

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

func Test_mxplatformgo_MicrodepositsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MicrodepositsAPIService MicroDepositsMicrodepositGuidVerifyPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var microdepositGuid string

		resp, httpRes, err := apiClient.MicrodepositsAPI.MicroDepositsMicrodepositGuidVerifyPut(context.Background(), microdepositGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicrodepositsAPIService UsersUserGuidMicroDepositsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGuid string

		resp, httpRes, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsGet(context.Background(), userGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicrodepositsAPIService UsersUserGuidMicroDepositsMicroDepositGuidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var microDepositGuid string
		var userGuid string

		httpRes, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidDelete(context.Background(), microDepositGuid, userGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicrodepositsAPIService UsersUserGuidMicroDepositsMicroDepositGuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGuid string
		var microDepositGuid string

		resp, httpRes, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsMicroDepositGuidGet(context.Background(), userGuid, microDepositGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MicrodepositsAPIService UsersUserGuidMicroDepositsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userGuid string

		resp, httpRes, err := apiClient.MicrodepositsAPI.UsersUserGuidMicroDepositsPost(context.Background(), userGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
