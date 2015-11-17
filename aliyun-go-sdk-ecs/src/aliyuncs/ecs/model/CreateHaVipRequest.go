package model

import "aliyuncs/ecs"

type CreateHaVipRequest struct {
	ecs.Request
}

func NewCreateHaVipRequest() *CreateHaVipRequest{
	r := new(CreateHaVipRequest)
	r.Set("Action", "CreateHaVip")
	return r
}

func (this *CreateHaVipRequest) SetRegionId(v string){
	this.Set("RegionId", v)
}

func (this *CreateHaVipRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *CreateHaVipRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *CreateHaVipRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *CreateHaVipRequest) SetIpAddress(v string){
	this.Set("IpAddress", v)
}

func (this *CreateHaVipRequest) GetIpAddress() string {
	return this.Get("IpAddress")
}

func (this *CreateHaVipRequest) SetDescription(v string){
	this.Set("Description", v)
}

func (this *CreateHaVipRequest) GetDescription() string {
	return this.Get("Description")
}
