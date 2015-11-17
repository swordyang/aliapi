package model

import "aliyuncs/ecs"

type UnassociateEipAddressRequest struct {
	ecs.Request
}

func NewUnassociateEipAddressRequest() *UnassociateEipAddressRequest{
	r := new(UnassociateEipAddressRequest)
	r.Set("Action", "UnassociateEipAddress")
	return r
}

func (this *UnassociateEipAddressRequest) SetInstanceId(v string){
	this.Set("InstanceId", v)
}

func (this *UnassociateEipAddressRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *UnassociateEipAddressRequest) SetAllocationId(v string){
	this.Set("AllocationId", v)
}

func (this *UnassociateEipAddressRequest) GetAllocationId() string {
	return this.Get("AllocationId")
}

func (this *UnassociateEipAddressRequest) SetInstanceType(v string){
	this.Set("InstanceType", v)
}

func (this *UnassociateEipAddressRequest) GetInstanceType() string {
	return this.Get("InstanceType")
}
