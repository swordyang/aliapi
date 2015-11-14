package model

import "aliyuncs/ecs"

type CreateSnapshotRequest struct {
	ecs.Request
}

func NewCreateSnapshotRequest() *CreateSnapshotRequest{
	r := new(CreateSnapshotRequest)
	r.Set("Action", "CreateSnapshot")
	return r
}

func (this *CreateSnapshotRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *CreateSnapshotRequest) GetDiskId() string {
	return this.Get("DiskId")
}

func (this *CreateSnapshotRequest) SetSnapshotName(v string){
	this.Set("SnapshotName", v)
}

func (this *CreateSnapshotRequest) GetSnapshotName() string {
	return this.Get("SnapshotName")
}

func (this *CreateSnapshotRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *CreateSnapshotRequest) GetDescription() string {
	return this.Get("Description")
}

func (this *CreateSnapshotRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *CreateSnapshotRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
