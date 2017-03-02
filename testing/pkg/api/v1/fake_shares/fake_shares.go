// Copyright (c) 2016 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
This module implements the entry into CRUD operation of shares.

*/

package shares

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/opensds/opensds/testing/pkg/api"
	"github.com/opensds/opensds/testing/pkg/api/grpcapi"
	pb "github.com/opensds/opensds/testing/pkg/grpc/fake_opensds"
)

type ShareRequestDeliver interface {
	createShare() *pb.Response

	getShare() *pb.Response

	listShares() *pb.Response

	deleteShare() *pb.Response
}

type FakeShareRequest struct {
	ResourceType string `json:"resourceType,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Size         int32  `json:"size"`
	ShareType    string `json:"shareType,omitempty"`
	ShareProto   string `json:"shareProto,omitempty"`
	AllowDetails bool   `json:"allowDetails"`
}

func (fsr FakeShareRequest) createShare() *pb.Response {
	return grpcapi.CreateShare(fsr.ResourceType, fsr.Name, fsr.ShareType, fsr.ShareProto, fsr.Size)
}

func (fsr FakeShareRequest) getShare() *pb.Response {
	return grpcapi.GetShare(fsr.ResourceType, fsr.Id)
}

func (fsr FakeShareRequest) listShares() *pb.Response {
	return grpcapi.ListShares(fsr.ResourceType, fsr.AllowDetails)
}

func (fsr FakeShareRequest) deleteShare() *pb.Response {
	return grpcapi.DeleteShare(fsr.ResourceType, fsr.Id)
}

func CreateShare(srd ShareRequestDeliver) (api.ShareResponse, error) {
	var nullResponse api.ShareResponse

	result := srd.createShare()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("Create file share error:", err)
		return nullResponse, err
	}

	var shareResponse api.ShareResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &shareResponse); err != nil {
		return nullResponse, err
	}
	return shareResponse, nil
}

func GetShare(srd ShareRequestDeliver) (api.ShareDetailResponse, error) {
	var nullResponse api.ShareDetailResponse

	result := srd.getShare()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("Get file share error:", err)
		return nullResponse, err
	}

	var shareDetailResponse api.ShareDetailResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &shareDetailResponse); err != nil {
		return nullResponse, err
	}
	return shareDetailResponse, nil
}

func ListShares(srd ShareRequestDeliver) ([]api.ShareResponse, error) {
	var nullResponses []api.ShareResponse

	result := srd.listShares()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("List all file shares error:", err)
		return nullResponses, err
	}

	var sharesResponse []api.ShareResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &sharesResponse); err != nil {
		return nullResponses, err
	}
	return sharesResponse, nil
}

func DeleteShare(srd ShareRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := srd.deleteShare()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Delete file share error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

var sampleShareCreateRequest = `{
	"resourceType":"manila",
	"name":"My_share",
	"shareType":"25747776-08e5-494f-ab40-a64b9d20d8f7",
	"shareProto":"NFS",
	"size":2
}`

var sampleShareGetRequest = `{
	"resourceType":"manila",
	"id":"d94a8548-2079-4be0-b21c-0a887acd31ca"
}`

var sampleShareListRequest = `{
	"resourceType":"manila",
	"allowDetails":false
}`

var sampleShareDeleteRequest = `{
	"resourceType":"manila",
	"id":"d94a8548-2079-4be0-b21c-0a887acd31ca"
}`

var sampleShareData = `{
    "id": "d94a8548-2079-4be0-b21c-0a887acd31ca",
    "links": [
		{
			"href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
			"rel": "self"
		},
		{
            "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
			"rel": "bookmark"
        }
    ],
    "name": "My_share"
}`

var sampleShareDetailData = `{
    "links": [
        {
            "href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
            "rel": "self"
        },
        {
            "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
            "rel": "bookmark"
        }
    ],
    "availability_zone": "nova",
    "share_network_id": "713df749-aac0-4a54-af52-10f6c991e80c",
    "export_locations": [],
    "share_server_id": "e268f4aa-d571-43dd-9ab3-f49ad06ffaef",
    "snapshot_id": null,
    "id": "d94a8548-2079-4be0-b21c-0a887acd31ca",
    "size": 1,
    "share_type": "25747776-08e5-494f-ab40-a64b9d20d8f7",
    "share_type_name": "default",
    "export_location": null,
    "consistency_group_id": "9397c191-8427-4661-a2e8-b23820dc01d4",
    "project_id": "16e1ab15c35a457e9c2b2aa189f544e1",
    "metadata": {
        "project": "my_app",
        "aim": "doc"
    },
    "status": "available",
    "description": "My custom share London",
    "host": "manila2@generic1#GENERIC1",
    "access_rules_status": "active",
    "has_replicas": false,
    "replication_type": null,
    "task_state": null,
    "is_public": true,
    "snapshot_support": true,
    "name": "My_share",
    "created_at": "2015-09-18T10:25:24.000000",
    "share_proto": "NFS",
    "volume_type": "default",
    "source_cgsnapshot_member_id": null
}`

var sampleSharesData = `[
    {
        "id": "d94a8548-2079-4be0-b21c-0a887acd31ca",
        "links": [
            {
                "href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
                "rel": "self"
            },
            {
                "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/d94a8548-2079-4be0-b21c-0a887acd31ca",
                "rel": "bookmark"
            }
        ],
        "name": "My_share"
    },
    {
        "id": "406ea93b-32e9-4907-a117-148b3945749f",
        "links": [
            {
                "href": "http://172.18.198.54:8786/v2/16e1ab15c35a457e9c2b2aa189f544e1/shares/406ea93b-32e9-4907-a117-148b3945749f",
                "rel": "self"
            },
            {
                "href": "http://172.18.198.54:8786/16e1ab15c35a457e9c2b2aa189f544e1/shares/406ea93b-32e9-4907-a117-148b3945749f",
                "rel": "bookmark"
            }
        ],
        "name": "Share1"
    }
]`
