// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware vRA CloudAccountGcp data source.
//
// ## Example Usage
// ### S
//
// **GCP cloud account data source by its id:**
//
// This is an example of how to create an GCP cloud account resource and read it as a data source using its id.
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
//			_, err := vra.LookupCloudAccountGcp(ctx, &GetCloudAccountGcpArgs{
//				Id: pulumi.StringRef(_var.Vra_cloud_account_gcp_id),
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
// **GCP cloud account data source by its name:**
//
// This is an example of how to read the cloud account data source using its name.
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
//			_, err := vra.LookupCloudAccountGcp(ctx, &GetCloudAccountGcpArgs{
//				Name: pulumi.StringRef(_var.Vra_cloud_account_gcp_name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCloudAccountGcp(ctx *pulumi.Context, args *LookupCloudAccountGcpArgs, opts ...pulumi.InvokeOption) (*LookupCloudAccountGcpResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCloudAccountGcpResult
	err := ctx.Invoke("vra:index/getCloudAccountGcp:getCloudAccountGcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudAccountGcp.
type LookupCloudAccountGcpArgs struct {
	// The id of this GCP cloud account.
	Id *string `pulumi:"id"`
	// The name of this GCP cloud account.
	Name *string `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountGcpTag `pulumi:"tags"`
}

// A collection of values returned by getCloudAccountGcp.
type LookupCloudAccountGcpResult struct {
	// GCP Client email.
	ClientEmail string `pulumi:"clientEmail"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// A human-friendly description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// HATEOAS of the entity.
	Links []GetCloudAccountGcpLink `pulumi:"links"`
	Name  string                   `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// GCP Private key ID.
	PrivateKeyId string `pulumi:"privateKeyId"`
	// GCP Project ID.
	ProjectId string `pulumi:"projectId"`
	// A set of region names that are enabled for this account.
	Regions []string `pulumi:"regions"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetCloudAccountGcpTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupCloudAccountGcpOutput(ctx *pulumi.Context, args LookupCloudAccountGcpOutputArgs, opts ...pulumi.InvokeOption) LookupCloudAccountGcpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudAccountGcpResult, error) {
			args := v.(LookupCloudAccountGcpArgs)
			r, err := LookupCloudAccountGcp(ctx, &args, opts...)
			var s LookupCloudAccountGcpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudAccountGcpResultOutput)
}

// A collection of arguments for invoking getCloudAccountGcp.
type LookupCloudAccountGcpOutputArgs struct {
	// The id of this GCP cloud account.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of this GCP cloud account.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A set of tag keys and optional values that were set on this resource.
	// example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetCloudAccountGcpTagArrayInput `pulumi:"tags"`
}

func (LookupCloudAccountGcpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountGcpArgs)(nil)).Elem()
}

// A collection of values returned by getCloudAccountGcp.
type LookupCloudAccountGcpResultOutput struct{ *pulumi.OutputState }

func (LookupCloudAccountGcpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudAccountGcpResult)(nil)).Elem()
}

func (o LookupCloudAccountGcpResultOutput) ToLookupCloudAccountGcpResultOutput() LookupCloudAccountGcpResultOutput {
	return o
}

func (o LookupCloudAccountGcpResultOutput) ToLookupCloudAccountGcpResultOutputWithContext(ctx context.Context) LookupCloudAccountGcpResultOutput {
	return o
}

// GCP Client email.
func (o LookupCloudAccountGcpResultOutput) ClientEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.ClientEmail }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o LookupCloudAccountGcpResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A human-friendly description.
func (o LookupCloudAccountGcpResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCloudAccountGcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.Id }).(pulumi.StringOutput)
}

// HATEOAS of the entity.
func (o LookupCloudAccountGcpResultOutput) Links() GetCloudAccountGcpLinkArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) []GetCloudAccountGcpLink { return v.Links }).(GetCloudAccountGcpLinkArrayOutput)
}

func (o LookupCloudAccountGcpResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o LookupCloudAccountGcpResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o LookupCloudAccountGcpResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.Owner }).(pulumi.StringOutput)
}

// GCP Private key ID.
func (o LookupCloudAccountGcpResultOutput) PrivateKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.PrivateKeyId }).(pulumi.StringOutput)
}

// GCP Project ID.
func (o LookupCloudAccountGcpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// A set of region names that are enabled for this account.
func (o LookupCloudAccountGcpResultOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) []string { return v.Regions }).(pulumi.StringArrayOutput)
}

// A set of tag keys and optional values that were set on this resource.
// example:[ { "key" : "vmware", "value": "provider" } ]
func (o LookupCloudAccountGcpResultOutput) Tags() GetCloudAccountGcpTagArrayOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) []GetCloudAccountGcpTag { return v.Tags }).(GetCloudAccountGcpTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o LookupCloudAccountGcpResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudAccountGcpResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudAccountGcpResultOutput{})
}
