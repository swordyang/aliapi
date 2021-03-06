package v20140526

import "aliyuncs/core"

type DescribeVpcsRequest struct {
	core.Request
}

func NewDescribeVpcsRequest() *DescribeVpcsRequest {
	r := new(DescribeVpcsRequest)
	r.Init("Ecs", "2014-05-26", "DescribeVpcs")
	return r
}

func (this *DescribeVpcsRequest) SetVpcId(v string) {
	this.Set("VpcId", v)
}

func (this *DescribeVpcsRequest) GetVpcId() string {
	return this.Get("VpcId")
}

func (this *DescribeVpcsRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *DescribeVpcsRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *DescribeVpcsRequest) SetPageNumber(v string) {
	this.Set("PageNumber", v)
}

func (this *DescribeVpcsRequest) GetPageNumber() string {
	return this.Get("PageNumber")
}

func (this *DescribeVpcsRequest) SetPageSize(v string) {
	this.Set("PageSize", v)
}

func (this *DescribeVpcsRequest) GetPageSize() string {
	return this.Get("PageSize")
}
