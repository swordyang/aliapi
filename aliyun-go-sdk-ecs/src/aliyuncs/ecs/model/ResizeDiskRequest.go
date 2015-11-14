package model

import "aliyuncs/ecs"

type ResizeDiskRequest struct {
	ecs.Request
}

func NewResizeDiskRequest() *ResizeDiskRequest{
	r := new(ResizeDiskRequest)
	r.Set("Action", "ResizeDisk")
	return r
}

func (this *ResizeDiskRequest) SetDiskId(v string){
	this.Set("DiskId", v)
}

func (this *ResizeDiskRequest) GetDiskId() string {
	return this.Get("DiskId")
}

/*
希望扩容到的磁盘大小。
以 GB 为单位，取值范围为 
Cloud：5 ~ 2000
*/
func (this *ResizeDiskRequest) SetNewSize(v int){
	this.SetInt("NewSize", v)
}

func (this *ResizeDiskRequest) GetNewSize() int {
	return this.GetInt("NewSize")
}

func (this *ResizeDiskRequest) SetClientToken(v string){
	this.Set("ClientToken", v)
}

func (this *ResizeDiskRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
