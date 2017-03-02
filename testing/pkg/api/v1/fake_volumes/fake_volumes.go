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
This module implements the entry into CRUD operation of volumes.

*/

package volumes

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/opensds/opensds/testing/pkg/api"
	"github.com/opensds/opensds/testing/pkg/api/grpcapi"
	pb "github.com/opensds/opensds/testing/pkg/grpc/fake_opensds"
)

type VolumeRequestDeliver interface {
	createVolume() *pb.Response

	getVolume() *pb.Response

	listVolumes() *pb.Response

	deleteVolume() *pb.Response

	attachVolume() *pb.Response

	detachVolume() *pb.Response

	mountVolume() *pb.Response

	unmountVolume() *pb.Response
}

type FakeVolumeRequest struct {
	ResourceType string `json:"resourcetType,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Size         int32  `json:"size"`
	AllowDetails bool   `json:"allowDetails"`

	ActionType string `json:"actionType,omitempty"`
	Host       string `json:"host,omitempty"`
	Device     string `json:"device,omitempty"`
	Attachment string `json:"attachment,omitempty"`
	MountDir   string `json:"mountDir,omitempty"`
	FsType     string `json:"fsType,omitempty"`
}

func (fvr FakeVolumeRequest) createVolume() *pb.Response {
	return grpcapi.CreateVolume(fvr.ResourceType, fvr.Name, fvr.Size)
}

func (fvr FakeVolumeRequest) getVolume() *pb.Response {
	return grpcapi.GetVolume(fvr.ResourceType, fvr.Id)
}

func (fvr FakeVolumeRequest) listVolumes() *pb.Response {
	return grpcapi.ListVolumes(fvr.ResourceType, fvr.AllowDetails)
}

func (fvr FakeVolumeRequest) deleteVolume() *pb.Response {
	return grpcapi.DeleteVolume(fvr.ResourceType, fvr.Id)
}

func (fvr FakeVolumeRequest) attachVolume() *pb.Response {
	return grpcapi.AttachVolume(fvr.ResourceType, fvr.Id, fvr.Host, fvr.Device)
}

func (fvr FakeVolumeRequest) detachVolume() *pb.Response {
	return grpcapi.DetachVolume(fvr.ResourceType, fvr.Id, fvr.Attachment)
}

func (fvr FakeVolumeRequest) mountVolume() *pb.Response {
	return grpcapi.MountVolume(fvr.MountDir, fvr.Device, fvr.FsType)
}

func (fvr FakeVolumeRequest) unmountVolume() *pb.Response {
	return grpcapi.UnmountVolume(fvr.MountDir)
}

func CreateVolume(vrd VolumeRequestDeliver) (api.VolumeResponse, error) {
	var nullResponse api.VolumeResponse

	result := vrd.createVolume()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("Create volume error:", err)
		return nullResponse, err
	}

	var volumeResponse api.VolumeResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &volumeResponse); err != nil {
		return nullResponse, err
	}
	return volumeResponse, nil
}

func GetVolume(vrd VolumeRequestDeliver) (api.VolumeDetailResponse, error) {
	var nullResponse api.VolumeDetailResponse

	result := vrd.getVolume()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("Get volume error:", err)
		return nullResponse, err
	}

	var volumeDetailResponse api.VolumeDetailResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &volumeDetailResponse); err != nil {
		return nullResponse, err
	}
	return volumeDetailResponse, nil
}

func ListVolumes(vrd VolumeRequestDeliver) ([]api.VolumeResponse, error) {
	var nullResponses []api.VolumeResponse

	result := vrd.listVolumes()
	if result.GetStatus() == "Failure" {
		err := errors.New(result.GetError())
		log.Println("List all volumes error:", err)
		return nullResponses, err
	}

	var volumesResponse []api.VolumeResponse
	rbody := []byte(result.GetMessage())
	if err := json.Unmarshal(rbody, &volumesResponse); err != nil {
		return nullResponses, err
	}
	return volumesResponse, nil
}

func DeleteVolume(vrd VolumeRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := vrd.deleteVolume()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Delete volume error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

func AttachVolume(vrd VolumeRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := vrd.attachVolume()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Attach volume error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

func DetachVolume(vrd VolumeRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := vrd.detachVolume()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Detach volume error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

func MountVolume(vrd VolumeRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := vrd.mountVolume()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Mount volume error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

func UnmountVolume(vrd VolumeRequestDeliver) api.DefaultResponse {
	var defaultResponse api.DefaultResponse

	result := vrd.unmountVolume()
	if result.GetStatus() == "Failure" {
		defaultResponse.Status = "Failure"
		defaultResponse.Error = result.GetError()
		log.Println("Unmount volume error:", defaultResponse.Error)
		return defaultResponse
	}

	defaultResponse.Status = "Success"
	return defaultResponse
}

var sampleVolumeCreateRequest = `{
	"resourceType":"cinder",
	"name":"myvol1",
	"size":2
}`

var sampleVolumeGetRequest = `{
	"resourceType":"cinder",
	"id":"30becf77-63fe-4f5e-9507-a0578ffe0949"
}`

var sampleVolumeListRequest = `{
	"resourceType":"cinder",
	"allowDetails":false
}`

var sampleVolumeDeleteRequest = `{
	"resourceType":"cinder",
	"id":"f5fc9874-fc89-4814-a358-23ba83a6115f"
}`

var sampleVolumeAttachRequest = `{
	"resourceType":"cinder",
	"id":"f5fc9874-fc89-4814-a358-23ba83a6115f",
	"host":"localhost",
	"device":"/dev/vdc"
}`

var sampleVolumeDetachRequest = `{
	"resourceType":"cinder",
	"id":"f5fc9874-fc89-4814-a358-23ba83a6115f",
	"attachment":"ddb2ac07-ed62-49eb-93da-73b258dd9bec"
}`

var sampleVolumeMountRequest = `{
	"mountDir":"/mnt",
	"device":"/dev/vdc",
	"id":"f5fc9874-fc89-4814-a358-23ba83a6115f",
	"fsType":"ext4"
}`

var sampleVolumeUnmountRequest = `{
	"mountDir":"/mnt"
}`

var sampleVolumeData = `{
	"name":"myvol1",
	"id":"f5fc9874-fc89-4814-a358-23ba83a6115f",
	"status":"available",
	"size":2,
	"volume_type":"lvmdriver-1",
	"attachments":[]
}`

var sampleVolumeDetailData = `{
	"id":"30becf77-63fe-4f5e-9507-a0578ffe0949",
	"attachments":[
		{
			"attachment_id": "ddb2ac07-ed62-49eb-93da-73b258dd9bec",
			"host_name": "host_test",
			"volume_id": "30becf77-63fe-4f5e-9507-a0578ffe0949",
			"device": "/dev/vdb",
			"id": "30becf77-63fe-4f5e-9507-a0578ffe0949",
			"server_id": "0f081aae-1b0c-4b89-930c-5f2562460c72"
		}
	],
	"links":[
		{
			"href": "http://172.16.197.131:8776/v2/1d8837c5fcef4892951397df97661f97/volumes/30becf77-63fe-4f5e-9507-a0578ffe0949",
			"rel": "self"
		},
		{
			"href": "http://172.16.197.131:8776/1d8837c5fcef4892951397df97661f97/volumes/30becf77-63fe-4f5e-9507-a0578ffe0949",
			"rel": "bookmark"
		}
	],
	"metadata":{
		"readonly": "false",
		"attached_mode": "rw"
	},
	"protected":false,
	"status":"available",
	"migrationStatus":null,
	"user_id":"a971aa69-c61a-4a49-b392-b0e41609bc5d",
	"encrypted":false,
	"multiattach":false,
	"created_at":"2014-09-29T14:44:31",
	"description":"test volume",
	"volume_type":"test_type",
	"name":"test_volume",
	"source_volid":"4b58bbb8-3b00-4f87-8243-8c622707bbab",
	"snapshot_id":"cc488e4a-9649-4e5f-ad12-20ab37c683b5",
	"size":2,

	"availability_zone":"default_cluster",
	"replication_status":null,
	"consistencygroup_id":null
}`

var sampleVolumesData = `[
	{
		"name":"myvol1",
		"id":"f5fc9874-fc89-4814-a358-23ba83a6115f",
		"status":"in-use",
		"size":1,
		"volume_type":"lvmdriver-1",
		"attachments":[
			{
				"attached_at":"2017-02-11T14:08:17.000000",
				"attachment_id":"c7f84865-640c-44ea-94ab-379a27f0ff65",
				"device":"/dev/vdc",
				"host_name":"localhost",
				"id":"034af8c9-ef44-4855-8e70-d51dceed7fc4",
				"server_id":"",
				"volume_id":"034af8c9-ef44-4855-8e70-d51dceed7fc4"
			}
		]
	},
	{
		"name":"myvol2",
		"id":"60055a0a-2451-4d78-af9c-f2302150602f",
		"status":"available",
		"size":2,
		"volume_type":"lvmdriver-1",
		"attachments":[]
	}
]`
