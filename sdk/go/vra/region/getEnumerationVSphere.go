// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package region

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
func GetEnumerationVSphere(ctx *pulumi.Context, args *GetEnumerationVSphereArgs, opts ...pulumi.InvokeOption) (*GetEnumerationVSphereResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetEnumerationVSphereResult
	err := ctx.Invoke("vra:region/getEnumerationVSphere:getEnumerationVSphere", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnumerationVSphere.
type GetEnumerationVSphereArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
	Dcid *string `pulumi:"dcid"`
	// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
	Hostname string `pulumi:"hostname"`
	// Password for the user used to authenticate with the cloud Account
	Password string `pulumi:"password"`
	// Username to authenticate with the cloud account
	Username string `pulumi:"username"`
}

// A collection of values returned by getEnumerationVSphere.
type GetEnumerationVSphereResult struct {
	AcceptSelfSignedCert *bool   `pulumi:"acceptSelfSignedCert"`
	Dcid                 *string `pulumi:"dcid"`
	Hostname             string  `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Password string `pulumi:"password"`
	// A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
	Regions  []string `pulumi:"regions"`
	Username string   `pulumi:"username"`
}

func GetEnumerationVSphereOutput(ctx *pulumi.Context, args GetEnumerationVSphereOutputArgs, opts ...pulumi.InvokeOption) GetEnumerationVSphereResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEnumerationVSphereResult, error) {
			args := v.(GetEnumerationVSphereArgs)
			r, err := GetEnumerationVSphere(ctx, &args, opts...)
			var s GetEnumerationVSphereResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEnumerationVSphereResultOutput)
}

// A collection of arguments for invoking getEnumerationVSphere.
type GetEnumerationVSphereOutputArgs struct {
	// Accept self signed certificate when connecting to vSphere. Example: false
	AcceptSelfSignedCert pulumi.BoolPtrInput `pulumi:"acceptSelfSignedCert"`
	// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
	Dcid pulumi.StringPtrInput `pulumi:"dcid"`
	// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
	Hostname pulumi.StringInput `pulumi:"hostname"`
	// Password for the user used to authenticate with the cloud Account
	Password pulumi.StringInput `pulumi:"password"`
	// Username to authenticate with the cloud account
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetEnumerationVSphereOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnumerationVSphereArgs)(nil)).Elem()
}

// A collection of values returned by getEnumerationVSphere.
type GetEnumerationVSphereResultOutput struct{ *pulumi.OutputState }

func (GetEnumerationVSphereResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEnumerationVSphereResult)(nil)).Elem()
}

func (o GetEnumerationVSphereResultOutput) ToGetEnumerationVSphereResultOutput() GetEnumerationVSphereResultOutput {
	return o
}

func (o GetEnumerationVSphereResultOutput) ToGetEnumerationVSphereResultOutputWithContext(ctx context.Context) GetEnumerationVSphereResultOutput {
	return o
}

func (o GetEnumerationVSphereResultOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) *bool { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

func (o GetEnumerationVSphereResultOutput) Dcid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) *string { return v.Dcid }).(pulumi.StringPtrOutput)
}

func (o GetEnumerationVSphereResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEnumerationVSphereResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEnumerationVSphereResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) string { return v.Password }).(pulumi.StringOutput)
}

// A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
func (o GetEnumerationVSphereResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

func (o GetEnumerationVSphereResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetEnumerationVSphereResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEnumerationVSphereResultOutput{})
}
