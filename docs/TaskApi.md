# \TaskApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanUpTasks**](TaskApi.md#CleanUpTasks) | **Put** /tasks/{taskType}/cleanup | Clean up finished tasks (COMPLETED, FAILED) for the given task type
[**CleanUpTasksDeprecated**](TaskApi.md#CleanUpTasksDeprecated) | **Put** /tasks/cleanuptasks/{taskType} | Clean up finished tasks (COMPLETED, FAILED) for the given task type (deprecated)
[**DeleteTask**](TaskApi.md#DeleteTask) | **Delete** /tasks/task/{taskName} | Delete a single task given its task name
[**DeleteTaskMetadataByTable**](TaskApi.md#DeleteTaskMetadataByTable) | **Delete** /tasks/{taskType}/{tableNameWithType}/metadata | Delete task metadata for the given task type and table
[**DeleteTaskQueue**](TaskApi.md#DeleteTaskQueue) | **Delete** /tasks/taskqueue/{taskType} | Delete a task queue (deprecated)
[**DeleteTasks**](TaskApi.md#DeleteTasks) | **Delete** /tasks/{taskType} | Delete all tasks (as well as the task queue) for the given task type
[**ExecuteAdhocTask**](TaskApi.md#ExecuteAdhocTask) | **Post** /tasks/execute | Execute a task on minion
[**GetCronSchedulerInformation**](TaskApi.md#GetCronSchedulerInformation) | **Get** /tasks/scheduler/information | Fetch cron scheduler information
[**GetCronSchedulerJobDetails**](TaskApi.md#GetCronSchedulerJobDetails) | **Get** /tasks/scheduler/jobDetails | Fetch cron scheduler job keys
[**GetCronSchedulerJobKeys**](TaskApi.md#GetCronSchedulerJobKeys) | **Get** /tasks/scheduler/jobKeys | Fetch cron scheduler job keys
[**GetSubtaskConfigs**](TaskApi.md#GetSubtaskConfigs) | **Get** /tasks/subtask/{taskName}/config | Get the configs of specified sub tasks for the given task
[**GetSubtaskOnWorkerProgress**](TaskApi.md#GetSubtaskOnWorkerProgress) | **Get** /tasks/subtask/workers/progress | Get progress of all subtasks with specified state tracked by minion worker in memory
[**GetSubtaskProgress**](TaskApi.md#GetSubtaskProgress) | **Get** /tasks/subtask/{taskName}/progress | Get progress of specified sub tasks for the given task tracked by minion worker in memory
[**GetSubtaskStates**](TaskApi.md#GetSubtaskStates) | **Get** /tasks/subtask/{taskName}/state | Get the states of all the sub tasks for the given task
[**GetTaskConfig**](TaskApi.md#GetTaskConfig) | **Get** /tasks/task/{taskName}/runtime/config | Get the task runtime config for the given task
[**GetTaskConfigs**](TaskApi.md#GetTaskConfigs) | **Get** /tasks/task/{taskName}/config | Get the task config (a list of child task configs) for the given task
[**GetTaskConfigsDeprecated**](TaskApi.md#GetTaskConfigsDeprecated) | **Get** /tasks/taskconfig/{taskName} | Get the task config (a list of child task configs) for the given task (deprecated)
[**GetTaskCounts**](TaskApi.md#GetTaskCounts) | **Get** /tasks/{taskType}/taskcounts | Fetch count of sub-tasks for each of the tasks for the given task type
[**GetTaskDebugInfo**](TaskApi.md#GetTaskDebugInfo) | **Get** /tasks/task/{taskName}/debug | Fetch information for the given task name
[**GetTaskGenerationDebugInto**](TaskApi.md#GetTaskGenerationDebugInto) | **Get** /tasks/generator/{tableNameWithType}/{taskType}/debug | Fetch task generation information for the recent runs of the given task for the given table
[**GetTaskMetadataByTable**](TaskApi.md#GetTaskMetadataByTable) | **Get** /tasks/{taskType}/{tableNameWithType}/metadata | Get task metadata for the given task type and table
[**GetTaskQueueState**](TaskApi.md#GetTaskQueueState) | **Get** /tasks/{taskType}/state | Get the state (task queue state) for the given task type
[**GetTaskQueueStateDeprecated**](TaskApi.md#GetTaskQueueStateDeprecated) | **Get** /tasks/taskqueuestate/{taskType} | Get the state (task queue state) for the given task type (deprecated)
[**GetTaskQueues**](TaskApi.md#GetTaskQueues) | **Get** /tasks/taskqueues | List all task queues (deprecated)
[**GetTaskState**](TaskApi.md#GetTaskState) | **Get** /tasks/task/{taskName}/state | Get the task state for the given task
[**GetTaskStateDeprecated**](TaskApi.md#GetTaskStateDeprecated) | **Get** /tasks/taskstate/{taskName} | Get the task state for the given task (deprecated)
[**GetTaskStates**](TaskApi.md#GetTaskStates) | **Get** /tasks/{taskType}/taskstates | Get a map from task to task state for the given task type
[**GetTaskStatesByTable**](TaskApi.md#GetTaskStatesByTable) | **Get** /tasks/{taskType}/{tableNameWithType}/state | List all tasks for the given task type
[**GetTaskStatesDeprecated**](TaskApi.md#GetTaskStatesDeprecated) | **Get** /tasks/taskstates/{taskType} | Get a map from task to task state for the given task type (deprecated)
[**GetTasks**](TaskApi.md#GetTasks) | **Get** /tasks/{taskType}/tasks | List all tasks for the given task type
[**GetTasksDebugInfo**](TaskApi.md#GetTasksDebugInfo) | **Get** /tasks/{taskType}/{tableNameWithType}/debug | Fetch information for all the tasks for the given task type and table
[**GetTasksDebugInfo1**](TaskApi.md#GetTasksDebugInfo1) | **Get** /tasks/{taskType}/debug | Fetch information for all the tasks for the given task type
[**GetTasksDeprecated**](TaskApi.md#GetTasksDeprecated) | **Get** /tasks/tasks/{taskType} | List all tasks for the given task type (deprecated)
[**ListTaskTypes**](TaskApi.md#ListTaskTypes) | **Get** /tasks/tasktypes | List all task types
[**ResumeTasks**](TaskApi.md#ResumeTasks) | **Put** /tasks/{taskType}/resume | Resume all stopped tasks (as well as the task queue) for the given task type
[**ScheduleTasks**](TaskApi.md#ScheduleTasks) | **Post** /tasks/schedule | Schedule tasks and return a map from task type to task name scheduled
[**ScheduleTasksDeprecated**](TaskApi.md#ScheduleTasksDeprecated) | **Put** /tasks/scheduletasks | Schedule tasks (deprecated)
[**StopTasks**](TaskApi.md#StopTasks) | **Put** /tasks/{taskType}/stop | Stop all running/pending tasks (as well as the task queue) for the given task type
[**ToggleTaskQueueState**](TaskApi.md#ToggleTaskQueueState) | **Put** /tasks/taskqueue/{taskType} | Stop/resume a task queue (deprecated)



## CleanUpTasks

> SuccessResponse CleanUpTasks(ctx, taskType).Execute()

Clean up finished tasks (COMPLETED, FAILED) for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.CleanUpTasks(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CleanUpTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CleanUpTasks`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.CleanUpTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiCleanUpTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CleanUpTasksDeprecated

> SuccessResponse CleanUpTasksDeprecated(ctx, taskType).Execute()

Clean up finished tasks (COMPLETED, FAILED) for the given task type (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.CleanUpTasksDeprecated(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CleanUpTasksDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CleanUpTasksDeprecated`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.CleanUpTasksDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiCleanUpTasksDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> SuccessResponse DeleteTask(ctx, taskName).ForceDelete(forceDelete).Execute()

Delete a single task given its task name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name
    forceDelete := true // bool | Whether to force deleting the task (expert only option, enable with cautious (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.DeleteTask(context.Background(), taskName).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.DeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTask`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.DeleteTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Whether to force deleting the task (expert only option, enable with cautious | [default to false]

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskMetadataByTable

> SuccessResponse DeleteTaskMetadataByTable(ctx, taskType, tableNameWithType).Execute()

Delete task metadata for the given task type and table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    tableNameWithType := "tableNameWithType_example" // string | Table name with type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.DeleteTaskMetadataByTable(context.Background(), taskType, tableNameWithType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.DeleteTaskMetadataByTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTaskMetadataByTable`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.DeleteTaskMetadataByTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskMetadataByTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskQueue

> SuccessResponse DeleteTaskQueue(ctx, taskType).ForceDelete(forceDelete).Execute()

Delete a task queue (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    forceDelete := true // bool | Whether to force delete the task queue (expert only option, enable with cautious (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.DeleteTaskQueue(context.Background(), taskType).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.DeleteTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTaskQueue`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.DeleteTaskQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Whether to force delete the task queue (expert only option, enable with cautious | [default to false]

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTasks

> SuccessResponse DeleteTasks(ctx, taskType).ForceDelete(forceDelete).Execute()

Delete all tasks (as well as the task queue) for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    forceDelete := true // bool | Whether to force deleting the tasks (expert only option, enable with cautious (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.DeleteTasks(context.Background(), taskType).ForceDelete(forceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.DeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTasks`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.DeleteTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **bool** | Whether to force deleting the tasks (expert only option, enable with cautious | [default to false]

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteAdhocTask

> ExecuteAdhocTask(ctx).Body(body).Execute()

Execute a task on minion

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    body := *openapiclient.NewAdhocTaskConfig("TaskType_example", "TableName_example") // AdhocTaskConfig |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskApi.ExecuteAdhocTask(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ExecuteAdhocTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteAdhocTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdhocTaskConfig**](AdhocTaskConfig.md) |  |

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronSchedulerInformation

> map[string]map[string]interface{} GetCronSchedulerInformation(ctx).Execute()

Fetch cron scheduler information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetCronSchedulerInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetCronSchedulerInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCronSchedulerInformation`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetCronSchedulerInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronSchedulerInformationRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronSchedulerJobDetails

> map[string]map[string]interface{} GetCronSchedulerJobDetails(ctx).TableName(tableName).TaskType(taskType).Execute()

Fetch cron scheduler job keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    tableName := "tableName_example" // string | Table name (with type suffix) (optional)
    taskType := "taskType_example" // string | Task type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetCronSchedulerJobDetails(context.Background()).TableName(tableName).TaskType(taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetCronSchedulerJobDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCronSchedulerJobDetails`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetCronSchedulerJobDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCronSchedulerJobDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableName** | **string** | Table name (with type suffix) |
 **taskType** | **string** | Task type |

### Return type

**map[string]map[string]interface{}**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCronSchedulerJobKeys

> []JobKey GetCronSchedulerJobKeys(ctx).Execute()

Fetch cron scheduler job keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetCronSchedulerJobKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetCronSchedulerJobKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCronSchedulerJobKeys`: []JobKey
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetCronSchedulerJobKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronSchedulerJobKeysRequest struct via the builder pattern


### Return type

[**[]JobKey**](JobKey.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtaskConfigs

> map[string]PinotTaskConfig GetSubtaskConfigs(ctx, taskName).SubtaskNames(subtaskNames).Execute()

Get the configs of specified sub tasks for the given task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name
    subtaskNames := "subtaskNames_example" // string | Sub task names separated by comma (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetSubtaskConfigs(context.Background(), taskName).SubtaskNames(subtaskNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetSubtaskConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubtaskConfigs`: map[string]PinotTaskConfig
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetSubtaskConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtaskConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subtaskNames** | **string** | Sub task names separated by comma |

### Return type

[**map[string]PinotTaskConfig**](PinotTaskConfig.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtaskOnWorkerProgress

> GetSubtaskOnWorkerProgress(ctx).SubTaskState(subTaskState).MinionWorkerIds(minionWorkerIds).Execute()

Get progress of all subtasks with specified state tracked by minion worker in memory

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    subTaskState := "subTaskState_example" // string | Subtask state (UNKNOWN,IN_PROGRESS,SUCCEEDED,CANCELLED,ERROR)
    minionWorkerIds := "minionWorkerIds_example" // string | Minion worker IDs separated by comma (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskApi.GetSubtaskOnWorkerProgress(context.Background()).SubTaskState(subTaskState).MinionWorkerIds(minionWorkerIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetSubtaskOnWorkerProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtaskOnWorkerProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subTaskState** | **string** | Subtask state (UNKNOWN,IN_PROGRESS,SUCCEEDED,CANCELLED,ERROR) |
 **minionWorkerIds** | **string** | Minion worker IDs separated by comma |

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtaskProgress

> string GetSubtaskProgress(ctx, taskName).SubtaskNames(subtaskNames).Execute()

Get progress of specified sub tasks for the given task tracked by minion worker in memory

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name
    subtaskNames := "subtaskNames_example" // string | Sub task names separated by comma (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetSubtaskProgress(context.Background(), taskName).SubtaskNames(subtaskNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetSubtaskProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubtaskProgress`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetSubtaskProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtaskProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subtaskNames** | **string** | Sub task names separated by comma |

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtaskStates

> map[string]string GetSubtaskStates(ctx, taskName).Execute()

Get the states of all the sub tasks for the given task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetSubtaskStates(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetSubtaskStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubtaskStates`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetSubtaskStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtaskStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskConfig

> map[string]string GetTaskConfig(ctx, taskName).Execute()

Get the task runtime config for the given task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskConfig(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskConfig`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskConfigs

> []PinotTaskConfig GetTaskConfigs(ctx, taskName).Execute()

Get the task config (a list of child task configs) for the given task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskConfigs(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskConfigs`: []PinotTaskConfig
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PinotTaskConfig**](PinotTaskConfig.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskConfigsDeprecated

> []PinotTaskConfig GetTaskConfigsDeprecated(ctx, taskName).Execute()

Get the task config (a list of child task configs) for the given task (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskConfigsDeprecated(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskConfigsDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskConfigsDeprecated`: []PinotTaskConfig
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskConfigsDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskConfigsDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PinotTaskConfig**](PinotTaskConfig.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskCounts

> map[string]TaskCount GetTaskCounts(ctx, taskType).Execute()

Fetch count of sub-tasks for each of the tasks for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskCounts(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskCounts`: map[string]TaskCount
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]TaskCount**](TaskCount.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskDebugInfo

> TaskDebugInfo GetTaskDebugInfo(ctx, taskName).Verbosity(verbosity).Execute()

Fetch information for the given task name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name
    verbosity := int32(56) // int32 | verbosity (Prints information for the given task name.By default, only prints subtask details for running and error tasks. Value of > 0 prints subtask details for all tasks) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskDebugInfo(context.Background(), taskName).Verbosity(verbosity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskDebugInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskDebugInfo`: TaskDebugInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskDebugInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDebugInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verbosity** | **int32** | verbosity (Prints information for the given task name.By default, only prints subtask details for running and error tasks. Value of &gt; 0 prints subtask details for all tasks) | [default to 0]

### Return type

[**TaskDebugInfo**](TaskDebugInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskGenerationDebugInto

> string GetTaskGenerationDebugInto(ctx, taskType, tableNameWithType).LocalOnly(localOnly).Execute()

Fetch task generation information for the recent runs of the given task for the given table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    tableNameWithType := "tableNameWithType_example" // string | Table name with type
    localOnly := true // bool | Whether to only lookup local cache for logs (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskGenerationDebugInto(context.Background(), taskType, tableNameWithType).LocalOnly(localOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskGenerationDebugInto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskGenerationDebugInto`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskGenerationDebugInto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskGenerationDebugIntoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **localOnly** | **bool** | Whether to only lookup local cache for logs | [default to false]

### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskMetadataByTable

> string GetTaskMetadataByTable(ctx, taskType, tableNameWithType).Execute()

Get task metadata for the given task type and table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    tableNameWithType := "tableNameWithType_example" // string | Table name with type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskMetadataByTable(context.Background(), taskType, tableNameWithType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskMetadataByTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskMetadataByTable`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskMetadataByTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskMetadataByTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskQueueState

> string GetTaskQueueState(ctx, taskType).Execute()

Get the state (task queue state) for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskQueueState(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskQueueState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskQueueState`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskQueueState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskQueueStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskQueueStateDeprecated

> StringResultResponse GetTaskQueueStateDeprecated(ctx, taskType).Execute()

Get the state (task queue state) for the given task type (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskQueueStateDeprecated(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskQueueStateDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskQueueStateDeprecated`: StringResultResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskQueueStateDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskQueueStateDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StringResultResponse**](StringResultResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskQueues

> []string GetTaskQueues(ctx).Execute()

List all task queues (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskQueues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskQueues`: []string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskQueues`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskQueuesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskState

> string GetTaskState(ctx, taskName).Execute()

Get the task state for the given task

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskState(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskState`: string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStateDeprecated

> StringResultResponse GetTaskStateDeprecated(ctx, taskName).Execute()

Get the task state for the given task (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskName := "taskName_example" // string | Task name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskStateDeprecated(context.Background(), taskName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskStateDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskStateDeprecated`: StringResultResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskStateDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskName** | **string** | Task name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStateDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StringResultResponse**](StringResultResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStates

> map[string]string GetTaskStates(ctx, taskType).Execute()

Get a map from task to task state for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskStates(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskStates`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStatesByTable

> map[string]string GetTaskStatesByTable(ctx, taskType, tableNameWithType).Execute()

List all tasks for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    tableNameWithType := "tableNameWithType_example" // string | Table name with type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskStatesByTable(context.Background(), taskType, tableNameWithType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskStatesByTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskStatesByTable`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskStatesByTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatesByTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStatesDeprecated

> map[string]string GetTaskStatesDeprecated(ctx, taskType).Execute()

Get a map from task to task state for the given task type (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTaskStatesDeprecated(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTaskStatesDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskStatesDeprecated`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTaskStatesDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatesDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> []string GetTasks(ctx, taskType).Execute()

List all tasks for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTasks(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasks`: []string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksDebugInfo

> map[string]TaskDebugInfo GetTasksDebugInfo(ctx, taskType, tableNameWithType).Verbosity(verbosity).Execute()

Fetch information for all the tasks for the given task type and table

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    tableNameWithType := "tableNameWithType_example" // string | Table name with type
    verbosity := int32(56) // int32 | verbosity (Prints information for all the tasks for the given task type and table.By default, only prints subtask details for running and error tasks. Value of > 0 prints subtask details for all tasks) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTasksDebugInfo(context.Background(), taskType, tableNameWithType).Verbosity(verbosity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTasksDebugInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasksDebugInfo`: map[string]TaskDebugInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTasksDebugInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |
**tableNameWithType** | **string** | Table name with type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksDebugInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verbosity** | **int32** | verbosity (Prints information for all the tasks for the given task type and table.By default, only prints subtask details for running and error tasks. Value of &gt; 0 prints subtask details for all tasks) | [default to 0]

### Return type

[**map[string]TaskDebugInfo**](TaskDebugInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksDebugInfo1

> map[string]TaskDebugInfo GetTasksDebugInfo1(ctx, taskType).Verbosity(verbosity).Execute()

Fetch information for all the tasks for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    verbosity := int32(56) // int32 | verbosity (Prints information for all the tasks for the given task type.By default, only prints subtask details for running and error tasks. Value of > 0 prints subtask details for all tasks) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTasksDebugInfo1(context.Background(), taskType).Verbosity(verbosity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTasksDebugInfo1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasksDebugInfo1`: map[string]TaskDebugInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTasksDebugInfo1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksDebugInfo1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verbosity** | **int32** | verbosity (Prints information for all the tasks for the given task type.By default, only prints subtask details for running and error tasks. Value of &gt; 0 prints subtask details for all tasks) | [default to 0]

### Return type

[**map[string]TaskDebugInfo**](TaskDebugInfo.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksDeprecated

> []string GetTasksDeprecated(ctx, taskType).Execute()

List all tasks for the given task type (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.GetTasksDeprecated(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTasksDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTasksDeprecated`: []string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTasksDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskTypes

> []string ListTaskTypes(ctx).Execute()

List all task types

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.ListTaskTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ListTaskTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.ListTaskTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTaskTypesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeTasks

> SuccessResponse ResumeTasks(ctx, taskType).Execute()

Resume all stopped tasks (as well as the task queue) for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.ResumeTasks(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ResumeTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeTasks`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.ResumeTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiResumeTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleTasks

> map[string]string ScheduleTasks(ctx).TaskType(taskType).TableName(tableName).Execute()

Schedule tasks and return a map from task type to task name scheduled

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type (optional)
    tableName := "tableName_example" // string | Table name (with type suffix) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.ScheduleTasks(context.Background()).TaskType(taskType).TableName(tableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ScheduleTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleTasks`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.ScheduleTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskType** | **string** | Task type |
 **tableName** | **string** | Table name (with type suffix) |

### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleTasksDeprecated

> map[string]string ScheduleTasksDeprecated(ctx).Execute()

Schedule tasks (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.ScheduleTasksDeprecated(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ScheduleTasksDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleTasksDeprecated`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.ScheduleTasksDeprecated`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleTasksDeprecatedRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopTasks

> SuccessResponse StopTasks(ctx, taskType).Execute()

Stop all running/pending tasks (as well as the task queue) for the given task type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.StopTasks(context.Background(), taskType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.StopTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopTasks`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.StopTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiStopTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleTaskQueueState

> SuccessResponse ToggleTaskQueueState(ctx, taskType).State(state).Execute()

Stop/resume a task queue (deprecated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/davidgs/golang-api-client"
)

func main() {
    taskType := "taskType_example" // string | Task type
    state := "state_example" // string | state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskApi.ToggleTaskQueueState(context.Background(), taskType).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.ToggleTaskQueueState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleTaskQueueState`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.ToggleTaskQueueState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskType** | **string** | Task type |

### Other Parameters

Other parameters are passed through a pointer to a apiToggleTaskQueueStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | state |

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

