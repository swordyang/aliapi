package v20140526

import "aliyuncs/core"

type AllocateEipAddressRequest struct {
	core.Request
}

func NewAllocateEipAddressRequest() *AllocateEipAddressRequest {
	r := new(AllocateEipAddressRequest)
	r.Init("Ecs", "2014-05-26", "AllocateEipAddress")
	return r
}

func (this *AllocateEipAddressRequest) SetRegionId(v string) {
	this.Set("RegionId", v)
}

func (this *AllocateEipAddressRequest) GetRegionId() string {
	return this.Get("RegionId")
}

func (this *AllocateEipAddressRequest) SetBandwidth(v string) {
	this.Set("Bandwidth", v)
}

func (this *AllocateEipAddressRequest) GetBandwidth() string {
	return this.Get("Bandwidth")
}

func (this *AllocateEipAddressRequest) SetInternetChargeType(v string) {
	this.Set("InternetChargeType", v)
}

func (this *AllocateEipAddressRequest) GetInternetChargeType() string {
	return this.Get("InternetChargeType")
}

func (this *AllocateEipAddressRequest) SetClientToken(v string) {
	this.Set("ClientToken", v)
}

func (this *AllocateEipAddressRequest) GetClientToken() string {
	return this.Get("ClientToken")
}
