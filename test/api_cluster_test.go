/*
Pinot Controller API

Testing ClusterApiService

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

func Test_openapi_ClusterApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterApiService DeleteClusterConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var configName string

		httpRes, err := apiClient.ClusterApi.DeleteClusterConfig(context.Background(), configName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterApiService GetClusterInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ClusterApi.GetClusterInfo(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterApiService GetSegmentDebugInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tableName string
		var segmentName string

		httpRes, err := apiClient.ClusterApi.GetSegmentDebugInfo(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterApiService GetTableDebugInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tableName string

		httpRes, err := apiClient.ClusterApi.GetTableDebugInfo(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterApiService ListClusterConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ClusterApi.ListClusterConfigs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClusterApiService UpdateClusterConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ClusterApi.UpdateClusterConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}