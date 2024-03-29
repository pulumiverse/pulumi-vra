// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storageprofile

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to create a storage profile azure resource.
//
// **Storage profile azure data source by its id:**
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/storageprofile"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/storageprofile"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storageprofile.LookupAzure(ctx, &storageprofile.LookupAzureArgs{
//				Id: pulumi.StringRef(vra_storage_profile_azure.This.Id),
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
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/storageprofile"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/storageprofile"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storageprofile.LookupAzure(ctx, &storageprofile.LookupAzureArgs{
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
// A storage profile azure data source supports the following arguments:
func LookupAzure(ctx *pulumi.Context, args *LookupAzureArgs, opts ...pulumi.InvokeOption) (*LookupAzureResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAzureResult
	err := ctx.Invoke("vra:storageprofile/getAzure:getAzure", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzure.
type LookupAzureArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter *string `pulumi:"filter"`
	// The id of the image profile instance.
	Id *string `pulumi:"id"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetAzureTag `pulumi:"tags"`
}

// A collection of values returned by getAzure.
type LookupAzureResult struct {
	CloudAccountId string `pulumi:"cloudAccountId"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Indicates the caching mechanism for additional disk.
	DataDiskCaching string `pulumi:"dataDiskCaching"`
	// Indicates if this storage profile is a default profile.
	DefaultItem bool `pulumi:"defaultItem"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	DiskType string `pulumi:"diskType"`
	// The id of the region as seen in the cloud provider for which this profile is defined.
	ExternalRegionId string  `pulumi:"externalRegionId"`
	Filter           *string `pulumi:"filter"`
	Id               string  `pulumi:"id"`
	// HATEOAS of the entity
	Links []GetAzureLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  string `pulumi:"name"`
	OrgId string `pulumi:"orgId"`
	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	OsDiskCaching string `pulumi:"osDiskCaching"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Indicates whether this storage policy should support encryption or not.
	SupportsEncryption bool `pulumi:"supportsEncryption"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []GetAzureTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupAzureOutput(ctx *pulumi.Context, args LookupAzureOutputArgs, opts ...pulumi.InvokeOption) LookupAzureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureResult, error) {
			args := v.(LookupAzureArgs)
			r, err := LookupAzure(ctx, &args, opts...)
			var s LookupAzureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureResultOutput)
}

// A collection of arguments for invoking getAzure.
type LookupAzureOutputArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The id of the image profile instance.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags GetAzureTagArrayInput `pulumi:"tags"`
}

func (LookupAzureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureArgs)(nil)).Elem()
}

// A collection of values returned by getAzure.
type LookupAzureResultOutput struct{ *pulumi.OutputState }

func (LookupAzureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureResult)(nil)).Elem()
}

func (o LookupAzureResultOutput) ToLookupAzureResultOutput() LookupAzureResultOutput {
	return o
}

func (o LookupAzureResultOutput) ToLookupAzureResultOutputWithContext(ctx context.Context) LookupAzureResultOutput {
	return o
}

func (o LookupAzureResultOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.CloudAccountId }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupAzureResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Indicates the caching mechanism for additional disk.
func (o LookupAzureResultOutput) DataDiskCaching() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.DataDiskCaching }).(pulumi.StringOutput)
}

// Indicates if this storage profile is a default profile.
func (o LookupAzureResultOutput) DefaultItem() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAzureResult) bool { return v.DefaultItem }).(pulumi.BoolOutput)
}

// A human-friendly description.
func (o LookupAzureResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
func (o LookupAzureResultOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.DiskType }).(pulumi.StringOutput)
}

// The id of the region as seen in the cloud provider for which this profile is defined.
func (o LookupAzureResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o LookupAzureResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupAzureResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o LookupAzureResultOutput) Links() GetAzureLinkArrayOutput {
	return o.ApplyT(func(v LookupAzureResult) []GetAzureLink { return v.Links }).(GetAzureLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o LookupAzureResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAzureResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
func (o LookupAzureResultOutput) OsDiskCaching() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.OsDiskCaching }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupAzureResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Indicates whether this storage policy should support encryption or not.
func (o LookupAzureResultOutput) SupportsEncryption() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAzureResult) bool { return v.SupportsEncryption }).(pulumi.BoolOutput)
}

// A set of tag keys and optional values that were set on this Network Profile.
// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
func (o LookupAzureResultOutput) Tags() GetAzureTagArrayOutput {
	return o.ApplyT(func(v LookupAzureResult) []GetAzureTag { return v.Tags }).(GetAzureTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupAzureResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureResultOutput{})
}
