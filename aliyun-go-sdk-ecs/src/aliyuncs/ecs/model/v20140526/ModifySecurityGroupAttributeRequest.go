package v20140526

import "aliyuncs/core"

type ModifySecurityGroupAttributeRequest struct {
	core.Request
}

func NewModifySecurityGroupAttributeRequest() *ModifySecurityGroupAttributeRequest {
	r := new(ModifySecurityGroupAttributeRequest)
	r.Init("Ecs", "2014-05-26", "ModifySecurityGroupAttribute")
	return r
}

func (this *ModifySecurityGroupAttributeRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *ModifySecurityGroupAttributeRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *ModifySecurityGroupAttributeRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *ModifySecurityGroupAttributeRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}

func (this *ModifySecurityGroupAttributeRequest) SetSecurityGroupName(v string) {
	this.Set("SecurityGroupName", v)
}

func (this *ModifySecurityGroupAttributeRequest) GetSecurityGroupName() string {
	return this.Get("SecurityGroupName")
}

func (this *ModifySecurityGroupAttributeRequest) SetDescription(v string) {
	this.Set("Description", v)
}

func (this *ModifySecurityGroupAttributeRequest) GetDescription() string {
	return this.Get("Description")
}
