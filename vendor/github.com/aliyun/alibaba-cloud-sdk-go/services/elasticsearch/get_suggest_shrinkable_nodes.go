package elasticsearch

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetSuggestShrinkableNodes invokes the elasticsearch.GetSuggestShrinkableNodes API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/getsuggestshrinkablenodes.html
func (client *Client) GetSuggestShrinkableNodes(request *GetSuggestShrinkableNodesRequest) (response *GetSuggestShrinkableNodesResponse, err error) {
	response = CreateGetSuggestShrinkableNodesResponse()
	err = client.DoAction(request, response)
	return
}

// GetSuggestShrinkableNodesWithChan invokes the elasticsearch.GetSuggestShrinkableNodes API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getsuggestshrinkablenodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSuggestShrinkableNodesWithChan(request *GetSuggestShrinkableNodesRequest) (<-chan *GetSuggestShrinkableNodesResponse, <-chan error) {
	responseChan := make(chan *GetSuggestShrinkableNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSuggestShrinkableNodes(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetSuggestShrinkableNodesWithCallback invokes the elasticsearch.GetSuggestShrinkableNodes API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/getsuggestshrinkablenodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSuggestShrinkableNodesWithCallback(request *GetSuggestShrinkableNodesRequest, callback func(response *GetSuggestShrinkableNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSuggestShrinkableNodesResponse
		var err error
		defer close(result)
		response, err = client.GetSuggestShrinkableNodes(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetSuggestShrinkableNodesRequest is the request struct for api GetSuggestShrinkableNodes
type GetSuggestShrinkableNodesRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	NodeType   string           `position:"Query" name:"nodeType"`
	Count      requests.Integer `position:"Query" name:"count"`
}

// GetSuggestShrinkableNodesResponse is the response struct for api GetSuggestShrinkableNodes
type GetSuggestShrinkableNodesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Code      string       `json:"Code" xml:"Code"`
	Message   string       `json:"Message" xml:"Message"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateGetSuggestShrinkableNodesRequest creates a request to invoke GetSuggestShrinkableNodes API
func CreateGetSuggestShrinkableNodesRequest() (request *GetSuggestShrinkableNodesRequest) {
	request = &GetSuggestShrinkableNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "GetSuggestShrinkableNodes", "/openapi/instances/[InstanceId]/suggest-shrinkable-nodes", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetSuggestShrinkableNodesResponse creates a response to parse from GetSuggestShrinkableNodes response
func CreateGetSuggestShrinkableNodesResponse() (response *GetSuggestShrinkableNodesResponse) {
	response = &GetSuggestShrinkableNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
