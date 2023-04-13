/*
Pinot Controller API

Testing TunerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_TunerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TunerApiService TuneTable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.TunerApi.TuneTable(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TunerApiService TuneTable1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.TunerApi.TuneTable1(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}