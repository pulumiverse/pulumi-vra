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
// This is an example of how to create a storage profile vsphere data source.
//
// **Storage profile vsphere data source by its id:**
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupStorageProfileVsphere(ctx, &GetStorageProfileVsphereArgs{
//				Id: pulumi.StringRef(vra_storage_profile_vsphere.This.Id),
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
// **Vra storage profile data source filter by external region id:**
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.LookupStorageProfileVsphere(ctx, &GetStorageProfileVsphereArgs{
//				Filter: pulumi.StringRef("externalRegionId eq 'foobar'"),
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
// A storage profile vsphere data source supports the following arguments:
func LookupStorageProfileVsphere(ctx *pulumi.Context, args *LookupStorageProfileVsphereArgs, opts ...pulumi.InvokeOption) (*LookupStorageProfileVsphereResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupStorageProfileVsphereResult
	err := ctx.Invoke("vra:index/getStorageProfileVsphere:getStorageProfileVsphere", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageProfileVsphere.
type LookupStorageProfileVsphereArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter *string `pulumi:"filter"`
	// The id of the image profile instance.
	Id *string `pulumi:"id"`
	// Indicates whether this storage profile supports encryption or not.
	SharesLevel *string `pulumi:"sharesLevel"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetStorageProfileVsphereTag `pulumi:"tags"`
}

// A collection of values returned by getStorageProfileVsphere.
type LookupStorageProfileVsphereResult struct {
	// Id of the cloud account this storage profile belongs to.
	CloudAccountId string `pulumi:"cloudAccountId"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Indicates if this storage profile is a default profile.
	DefaultItem bool `pulumi:"defaultItem"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// Type of mode for the disk.
	DiskMode string `pulumi:"diskMode"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType string `pulumi:"diskType"`
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	Filter           *string `pulumi:"filter"`
	Id               string  `pulumi:"id"`
	// The upper bound for the I/O operations per second allocated for each virtual disk.
	LimitIops string `pulumi:"limitIops"`
	// HATEOAS of the entity
	Links []GetStorageProfileVsphereLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Type of provisioning policy for the disk.
	ProvisioningType string `pulumi:"provisioningType"`
	// A specific number of shares assigned to each virtual machine.
	Shares      string  `pulumi:"shares"`
	SharesLevel *string `pulumi:"sharesLevel"`
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption bool `pulumi:"supportsEncryption"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetStorageProfileVsphereTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupStorageProfileVsphereOutput(ctx *pulumi.Context, args LookupStorageProfileVsphereOutputArgs, opts ...pulumi.InvokeOption) LookupStorageProfileVsphereResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageProfileVsphereResult, error) {
			args := v.(LookupStorageProfileVsphereArgs)
			r, err := LookupStorageProfileVsphere(ctx, &args, opts...)
			var s LookupStorageProfileVsphereResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageProfileVsphereResultOutput)
}

// A collection of arguments for invoking getStorageProfileVsphere.
type LookupStorageProfileVsphereOutputArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the image profile instance.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Indicates whether this storage profile supports encryption or not.
	SharesLevel pulumi.StringPtrInput `pulumi:"sharesLevel"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags GetStorageProfileVsphereTagArrayInput `pulumi:"tags"`
}

func (LookupStorageProfileVsphereOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageProfileVsphereArgs)(nil)).Elem()
}

// A collection of values returned by getStorageProfileVsphere.
type LookupStorageProfileVsphereResultOutput struct{ *pulumi.OutputState }

func (LookupStorageProfileVsphereResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageProfileVsphereResult)(nil)).Elem()
}

func (o LookupStorageProfileVsphereResultOutput) ToLookupStorageProfileVsphereResultOutput() LookupStorageProfileVsphereResultOutput {
	return o
}

func (o LookupStorageProfileVsphereResultOutput) ToLookupStorageProfileVsphereResultOutputWithContext(ctx context.Context) LookupStorageProfileVsphereResultOutput {
	return o
}

// Id of the cloud account this storage profile belongs to.
func (o LookupStorageProfileVsphereResultOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.CloudAccountId }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupStorageProfileVsphereResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Indicates if this storage profile is a default profile.
func (o LookupStorageProfileVsphereResultOutput) DefaultItem() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) bool { return v.DefaultItem }).(pulumi.BoolOutput)
}

// A human-friendly description.
func (o LookupStorageProfileVsphereResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.Description }).(pulumi.StringOutput)
}

// Type of mode for the disk.
func (o LookupStorageProfileVsphereResultOutput) DiskMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.DiskMode }).(pulumi.StringOutput)
}

// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
func (o LookupStorageProfileVsphereResultOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.DiskType }).(pulumi.StringOutput)
}

// The id of the region as seen in the cloud provider for which this profile is defined.
func (o LookupStorageProfileVsphereResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o LookupStorageProfileVsphereResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupStorageProfileVsphereResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.Id }).(pulumi.StringOutput)
}

// The upper bound for the I/O operations per second allocated for each virtual disk.
func (o LookupStorageProfileVsphereResultOutput) LimitIops() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.LimitIops }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o LookupStorageProfileVsphereResultOutput) Links() GetStorageProfileVsphereLinkArrayOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) []GetStorageProfileVsphereLink { return v.Links }).(GetStorageProfileVsphereLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o LookupStorageProfileVsphereResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupStorageProfileVsphereResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupStorageProfileVsphereResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Type of provisioning policy for the disk.
func (o LookupStorageProfileVsphereResultOutput) ProvisioningType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.ProvisioningType }).(pulumi.StringOutput)
}

// A specific number of shares assigned to each virtual machine.
func (o LookupStorageProfileVsphereResultOutput) Shares() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.Shares }).(pulumi.StringOutput)
}

func (o LookupStorageProfileVsphereResultOutput) SharesLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) *string { return v.SharesLevel }).(pulumi.StringPtrOutput)
}

// Indicates whether this storage policy should support encryption or not.
func (o LookupStorageProfileVsphereResultOutput) SupportsEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) bool { return v.SupportsEncryption }).(pulumi.BoolOutput)
}

// A set of tag keys and optional values that were set on this Network Profile.
// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
func (o LookupStorageProfileVsphereResultOutput) Tags() GetStorageProfileVsphereTagArrayOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) []GetStorageProfileVsphereTag { return v.Tags }).(GetStorageProfileVsphereTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupStorageProfileVsphereResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageProfileVsphereResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageProfileVsphereResultOutput{})
}
