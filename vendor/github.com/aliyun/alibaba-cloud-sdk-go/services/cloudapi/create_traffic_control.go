package cloudapi

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

// CreateTrafficControl invokes the cloudapi.CreateTrafficControl API synchronously
// api document: https://help.aliyun.com/api/cloudapi/createtrafficcontrol.html
func (client *Client) CreateTrafficControl(request *CreateTrafficControlRequest) (response *CreateTrafficControlResponse, err error) {
	response = CreateCreateTrafficControlResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTrafficControlWithChan invokes the cloudapi.CreateTrafficControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createtrafficcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrafficControlWithChan(request *CreateTrafficControlRequest) (<-chan *CreateTrafficControlResponse, <-chan error) {
	responseChan := make(chan *CreateTrafficControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTrafficControl(request)
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

// CreateTrafficControlWithCallback invokes the cloudapi.CreateTrafficControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/createtrafficcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTrafficControlWithCallback(request *CreateTrafficControlRequest, callback func(response *CreateTrafficControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTrafficControlResponse
		var err error
		defer close(result)
		response, err = client.CreateTrafficControl(request)
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

// CreateTrafficControlRequest is the request struct for api CreateTrafficControl
type CreateTrafficControlRequest struct {
	*requests.RpcRequest
	ApiDefault         requests.Integer `position:"Query" name:"ApiDefault"`
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	TrafficControlName string           `position:"Query" name:"TrafficControlName"`
	TrafficControlUnit string           `position:"Query" name:"TrafficControlUnit"`
	Description        string           `position:"Query" name:"Description"`
	UserDefault        requests.Integer `position:"Query" name:"UserDefault"`
	AppDefault         requests.Integer `position:"Query" name:"AppDefault"`
}

// CreateTrafficControlResponse is the response struct for api CreateTrafficControl
type CreateTrafficControlResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TrafficControlId string `json:"TrafficControlId" xml:"TrafficControlId"`
}

// CreateCreateTrafficControlRequest creates a request to invoke CreateTrafficControl API
func CreateCreateTrafficControlRequest() (request *CreateTrafficControlRequest) {
	request = &CreateTrafficControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateTrafficControl", "apigateway", "openAPI")
	return
}

// CreateCreateTrafficControlResponse creates a response to parse from CreateTrafficControl response
func CreateCreateTrafficControlResponse() (response *CreateTrafficControlResponse) {
	response = &CreateTrafficControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
