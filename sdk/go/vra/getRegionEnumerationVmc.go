// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
//
// This is an example of how to lookup a region enumeration data source for VMC cloud account.
//
// **Region enumeration data source for VMC**
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.GetRegionEnumerationVmc(ctx, &GetRegionEnumerationVmcArgs{
//				AcceptSelfSignedCert: pulumi.BoolRef(true),
//				DcId:                 pulumi.StringRef(_var.Vra_data_collector_id),
//				ApiToken:             _var.Api_token,
//				SddcName:             _var.Sddc_name,
//				NsxHostname:          _var.Nsx_hostname,
//				VcenterHostname:      _var.Vcenter_hostname,
//				VcenterPassword:      _var.Vcenter_password,
//				VcenterUsername:      _var.Vcenter_username,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// The region enumeration data source for VMC cloud account supports the following arguments:
func GetRegionEnumerationVmc(ctx *pulumi.Context, args *GetRegionEnumerationVmcArgs, opts ...pulumi.InvokeOption) (*GetRegionEnumerationVmcResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRegionEnumerationVmcResult
	err := ctx.Invoke("vra:index/getRegionEnumerationVmc:getRegionEnumerationVmc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegionEnumerationVmc.
type GetRegionEnumerationVmcArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// API Token for the cloud account endpoint.
	ApiToken string `pulumi:"apiToken"`
	// ID of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	DcId *string `pulumi:"dcId"`
	// The IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname string `pulumi:"nsxHostname"`
	// Identifier of the on-premise SDDC to be used by this cloud account.
	SddcName string `pulumi:"sddcName"`
	// The IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname string `pulumi:"vcenterHostname"`
	// Password for the user used to authenticate with the cloud Account
	VcenterPassword string `pulumi:"vcenterPassword"`
	// vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername string `pulumi:"vcenterUsername"`
}

// A collection of values returned by getRegionEnumerationVmc.
type GetRegionEnumerationVmcResult struct {
	AcceptSelfSignedCert *bool   `pulumi:"acceptSelfSignedCert"`
	ApiToken             string  `pulumi:"apiToken"`
	DcId                 *string `pulumi:"dcId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	NsxHostname string `pulumi:"nsxHostname"`
	// A set of Region names to enable provisioning on. Example: northamerica-northeast1
	Regions         []string `pulumi:"regions"`
	SddcName        string   `pulumi:"sddcName"`
	VcenterHostname string   `pulumi:"vcenterHostname"`
	VcenterPassword string   `pulumi:"vcenterPassword"`
	VcenterUsername string   `pulumi:"vcenterUsername"`
}

func GetRegionEnumerationVmcOutput(ctx *pulumi.Context, args GetRegionEnumerationVmcOutputArgs, opts ...pulumi.InvokeOption) GetRegionEnumerationVmcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegionEnumerationVmcResult, error) {
			args := v.(GetRegionEnumerationVmcArgs)
			r, err := GetRegionEnumerationVmc(ctx, &args, opts...)
			var s GetRegionEnumerationVmcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegionEnumerationVmcResultOutput)
}

// A collection of arguments for invoking getRegionEnumerationVmc.
type GetRegionEnumerationVmcOutputArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert pulumi.BoolPtrInput `pulumi:"acceptSelfSignedCert"`
	// API Token for the cloud account endpoint.
	ApiToken pulumi.StringInput `pulumi:"apiToken"`
	// ID of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	DcId pulumi.StringPtrInput `pulumi:"dcId"`
	// The IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname pulumi.StringInput `pulumi:"nsxHostname"`
	// Identifier of the on-premise SDDC to be used by this cloud account.
	SddcName pulumi.StringInput `pulumi:"sddcName"`
	// The IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname pulumi.StringInput `pulumi:"vcenterHostname"`
	// Password for the user used to authenticate with the cloud Account
	VcenterPassword pulumi.StringInput `pulumi:"vcenterPassword"`
	// vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername pulumi.StringInput `pulumi:"vcenterUsername"`
}

func (GetRegionEnumerationVmcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationVmcArgs)(nil)).Elem()
}

// A collection of values returned by getRegionEnumerationVmc.
type GetRegionEnumerationVmcResultOutput struct{ *pulumi.OutputState }

func (GetRegionEnumerationVmcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionEnumerationVmcResult)(nil)).Elem()
}

func (o GetRegionEnumerationVmcResultOutput) ToGetRegionEnumerationVmcResultOutput() GetRegionEnumerationVmcResultOutput {
	return o
}

func (o GetRegionEnumerationVmcResultOutput) ToGetRegionEnumerationVmcResultOutputWithContext(ctx context.Context) GetRegionEnumerationVmcResultOutput {
	return o
}

func (o GetRegionEnumerationVmcResultOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) *bool { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

func (o GetRegionEnumerationVmcResultOutput) ApiToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.ApiToken }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationVmcResultOutput) DcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) *string { return v.DcId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegionEnumerationVmcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationVmcResultOutput) NsxHostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.NsxHostname }).(pulumi.StringOutput)
}

// A set of Region names to enable provisioning on. Example: northamerica-northeast1
func (o GetRegionEnumerationVmcResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func (o GetRegionEnumerationVmcResultOutput) SddcName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.SddcName }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationVmcResultOutput) VcenterHostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.VcenterHostname }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationVmcResultOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.VcenterPassword }).(pulumi.StringOutput)
}

func (o GetRegionEnumerationVmcResultOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionEnumerationVmcResult) string { return v.VcenterUsername }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegionEnumerationVmcResultOutput{})
}
