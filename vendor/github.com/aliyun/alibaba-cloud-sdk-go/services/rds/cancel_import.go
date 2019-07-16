package rds

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

// CancelImport invokes the rds.CancelImport API synchronously
// api document: https://help.aliyun.com/api/rds/cancelimport.html
func (client *Client) CancelImport(request *CancelImportRequest) (response *CancelImportResponse, err error) {
	response = CreateCancelImportResponse()
	err = client.DoAction(request, response)
	return
}

// CancelImportWithChan invokes the rds.CancelImport API asynchronously
// api document: https://help.aliyun.com/api/rds/cancelimport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelImportWithChan(request *CancelImportRequest) (<-chan *CancelImportResponse, <-chan error) {
	responseChan := make(chan *CancelImportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelImport(request)
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

// CancelImportWithCallback invokes the rds.CancelImport API asynchronously
// api document: https://help.aliyun.com/api/rds/cancelimport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelImportWithCallback(request *CancelImportRequest, callback func(response *CancelImportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelImportResponse
		var err error
		defer close(result)
		response, err = client.CancelImport(request)
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

// CancelImportRequest is the request struct for api CancelImport
type CancelImportRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImportId             requests.Integer `position:"Query" name:"ImportId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CancelImportResponse is the response struct for api CancelImport
type CancelImportResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelImportRequest creates a request to invoke CancelImport API
func CreateCancelImportRequest() (request *CancelImportRequest) {
	request = &CancelImportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CancelImport", "rds", "openAPI")
	return
}

// CreateCancelImportResponse creates a response to parse from CancelImport response
func CreateCancelImportResponse() (response *CancelImportResponse) {
	response = &CancelImportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
