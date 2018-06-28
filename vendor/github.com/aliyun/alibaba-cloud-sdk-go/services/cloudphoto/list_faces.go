package cloudphoto

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

// ListFaces invokes the cloudphoto.ListFaces API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/listfaces.html
func (client *Client) ListFaces(request *ListFacesRequest) (response *ListFacesResponse, err error) {
	response = CreateListFacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFacesWithChan invokes the cloudphoto.ListFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listfaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFacesWithChan(request *ListFacesRequest) (<-chan *ListFacesResponse, <-chan error) {
	responseChan := make(chan *ListFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaces(request)
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

// ListFacesWithCallback invokes the cloudphoto.ListFaces API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/listfaces.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFacesWithCallback(request *ListFacesRequest, callback func(response *ListFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFacesResponse
		var err error
		defer close(result)
		response, err = client.ListFaces(request)
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

// ListFacesRequest is the request struct for api ListFaces
type ListFacesRequest struct {
	*requests.RpcRequest
	Direction   string           `position:"Query" name:"Direction"`
	Size        requests.Integer `position:"Query" name:"Size"`
	Cursor      string           `position:"Query" name:"Cursor"`
	State       string           `position:"Query" name:"State"`
	StoreName   string           `position:"Query" name:"StoreName"`
	LibraryId   string           `position:"Query" name:"LibraryId"`
	HasFaceName string           `position:"Query" name:"HasFaceName"`
}

// ListFacesResponse is the response struct for api ListFaces
type ListFacesResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	NextCursor string `json:"NextCursor" xml:"NextCursor"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Action     string `json:"Action" xml:"Action"`
	Faces      []Face `json:"Faces" xml:"Faces"`
}

// CreateListFacesRequest creates a request to invoke ListFaces API
func CreateListFacesRequest() (request *ListFacesRequest) {
	request = &ListFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListFaces", "cloudphoto", "openAPI")
	return
}

// CreateListFacesResponse creates a response to parse from ListFaces response
func CreateListFacesResponse() (response *ListFacesResponse) {
	response = &ListFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}