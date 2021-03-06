package v20140526

import "aliyuncs/core"

type JoinSecurityGroupRequest struct {
	core.Request
}

func NewJoinSecurityGroupRequest() *JoinSecurityGroupRequest {
	r := new(JoinSecurityGroupRequest)
	r.Init("Ecs", "2014-05-26", "JoinSecurityGroup")
	return r
}

func (this *JoinSecurityGroupRequest) SetInstanceId(v string) {
	this.Set("InstanceId", v)
}

func (this *JoinSecurityGroupRequest) GetInstanceId() string {
	return this.Get("InstanceId")
}

func (this *JoinSecurityGroupRequest) SetSecurityGroupId(v string) {
	this.Set("SecurityGroupId", v)
}

func (this *JoinSecurityGroupRequest) GetSecurityGroupId() string {
	return this.Get("SecurityGroupId")
}
