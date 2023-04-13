/*
Pinot Controller API

Testing SegmentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/golang-api-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_SegmentApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SegmentApiService DeleteAllSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.DeleteAllSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService DeleteSegment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.DeleteSegment(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService DeleteSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.DeleteSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService DownloadSegment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		httpRes, err := apiClient.SegmentApi.DownloadSegment(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService EndReplaceSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		httpRes, err := apiClient.SegmentApi.EndReplaceSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetReloadJobStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SegmentApi.GetReloadJobStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegmentMetadata(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentMetadataDeprecated1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegmentMetadataDeprecated1(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentMetadataDeprecated2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegmentMetadataDeprecated2(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		httpRes, err := apiClient.SegmentApi.GetSegmentTiers(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentToCrcMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegmentToCrcMap(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegmentToCrcMapDeprecated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegmentToCrcMapDeprecated(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetSelectedSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetSelectedSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetServerMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetServerMetadata(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetServerToSegmentsMap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetServerToSegmentsMap(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetServerToSegmentsMapDeprecated1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetServerToSegmentsMapDeprecated1(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetServerToSegmentsMapDeprecated2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetServerToSegmentsMapDeprecated2(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetTableTiers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		httpRes, err := apiClient.SegmentApi.GetTableTiers(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService GetZookeeperMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.GetZookeeperMetadata(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ListSegmentLineage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		httpRes, err := apiClient.SegmentApi.ListSegmentLineage(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadAllSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadAllSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadAllSegmentsDeprecated1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadAllSegmentsDeprecated1(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadAllSegmentsDeprecated2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadAllSegmentsDeprecated2(context.Background(), tableName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadSegment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadSegment(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadSegmentDeprecated1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadSegmentDeprecated1(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ReloadSegmentDeprecated2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.ReloadSegmentDeprecated2(context.Background(), tableName, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ResetSegment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableNameWithType string
		var segmentName string

		resp, httpRes, err := apiClient.SegmentApi.ResetSegment(context.Background(), tableNameWithType, segmentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService ResetSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableNameWithType string

		resp, httpRes, err := apiClient.SegmentApi.ResetSegments(context.Background(), tableNameWithType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService RevertReplaceSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		httpRes, err := apiClient.SegmentApi.RevertReplaceSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService StartReplaceSegments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableName string

		httpRes, err := apiClient.SegmentApi.StartReplaceSegments(context.Background(), tableName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService UpdateTimeIntervalZK", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tableNameWithType string

		httpRes, err := apiClient.SegmentApi.UpdateTimeIntervalZK(context.Background(), tableNameWithType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService UploadSegmentAsMultiPart", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SegmentApi.UploadSegmentAsMultiPart(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentApiService UploadSegmentAsMultiPartV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SegmentApi.UploadSegmentAsMultiPartV2(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
