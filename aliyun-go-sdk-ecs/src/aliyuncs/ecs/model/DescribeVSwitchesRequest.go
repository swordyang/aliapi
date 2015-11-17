package model

import "aliyuncs/ecs"

type DescribeVSwitchesRequest struct {
	ecs.Request
}

func NewDescribeVSwitchesRequest() *DescribeVSwitchesRequest{
	r := new(DescribeVSwitchesRequest)
	r.Set("Action", "DescribeVSwitches")
	return r
}

func (this *DescribeVSwitchesRequest) SetVpcId(v string){
	this.Set("VpcId", v)
}

func (this *DescribeVSwitchesRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *DescribeVSwitchesRequest) SetVSwitchId(v string){
	this.Set("VSwitchId", v)
}

func (this *DescribeVSwitchesRequest) GetVSwitchId() string {
	return this.Get("VSwitchId")
}

func (this *DescribeVSwitchesRequest) SetZoneId(v string){
	this.Set("ZoneId", v)
}

func (this *DescribeVSwitchesRequest) GetZoneId() string {
	return this.Get("ZoneId")
}

func (this *DescribeVSwitchesRequest) SetPageNumber(v int){
	this.SetInt("PageNumber", v)
}

func (this *DescribeVSwitchesRequest) GetPageNumber() int {
	return this.GetInt("PageNumber")
}

func (this *DescribeVSwitchesRequest) SetPageSize(v int){
	this.SetInt("PageSize", v)
}

func (this *DescribeVSwitchesRequest) GetPageSize() int {
	return this.GetInt("PageSize")
}
