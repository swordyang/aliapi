package v20140526

import "aliyuncs/core"

type StopInstanceRequest struct {
	core.Request
}

func NewStopInstanceRequest() *StopInstanceRequest {
	r := new(StopInstanceRequest)
	r.Init("Ecs", "2014-05-26", "StopInstance")
	return r
}

func (this *StopInstanceRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *StopInstanceRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *StopInstanceRequest) SetForceStop(v bool) {
	this.SetBool("ForceStop", v)
}

func (this *StopInstanceRequest) GetForceStop() bool {
	return this.GetBool("ForceStop")
}
