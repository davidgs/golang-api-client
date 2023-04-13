/*
Pinot Controller API

Testing LeaderApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/davidgs/golang-api-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_LeaderApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LeaderApiService GetLeaderForTable", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.LeaderApi.GetLeaderForTable(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LeaderApiService GetLeadersForAllTables", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LeaderApi.GetLeadersForAllTables(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
