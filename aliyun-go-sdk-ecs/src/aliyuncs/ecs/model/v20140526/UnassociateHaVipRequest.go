package v20140526

import "aliyuncs/core"

type UnassociateHaVipRequest struct {
	core.Request
}

func NewUnassociateHaVipRequest() *UnassociateHaVipRequest {
	r := new(UnassociateHaVipRequest)
	r.Init("Ecs", "2014-05-26", "UnassociateHaVip")
	return r
}

func (this *UnassociateHaVipRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *UnassociateHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *UnassociateHaVipRequest) SetHaVipId(v string) {
	this.Set("HaVipId", v)
}

func (this *UnassociateHaVipRequest) GetHaVipId() string {
	return this.Get("HaVipId")
}

func (this *UnassociateHaVipRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *UnassociateHaVipRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *UnassociateHaVipRequest) SetForce(v bool) {
	this.SetBool("Force", v)
}

func (this *UnassociateHaVipRequest) GetForce() bool {
	return this.GetBool("Force")
}
