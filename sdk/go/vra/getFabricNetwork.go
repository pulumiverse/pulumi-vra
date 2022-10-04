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
// This is an example of how to lookup fabric networks.
//
// **Fabric network by filter query:**
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.GetFabricNetwork(ctx, &GetFabricNetworkArgs{
//				Filter: fmt.Sprintf("name eq '%v' and externalRegionId eq '%v'", _var.Name, _var.External_region_id),
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
// A fabric network data source supports the following arguments:
func GetFabricNetwork(ctx *pulumi.Context, args *GetFabricNetworkArgs, opts ...pulumi.InvokeOption) (*GetFabricNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFabricNetworkResult
	err := ctx.Invoke("vra:index/getFabricNetwork:getFabricNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFabricNetwork.
type GetFabricNetworkArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API.
	Filter string `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetFabricNetworkTag `pulumi:"tags"`
}

// A collection of values returned by getFabricNetwork.
type GetFabricNetworkResult struct {
	// Network CIDR to be used.
	Cidr string `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// Additional properties that may be used to extend the base resource.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// State object representing a network on a external cloud provider.
	Description string `pulumi:"description"`
	// External entity Id on the provider side.
	ExternalId string `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId string `pulumi:"externalRegionId"`
	Filter           string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic bool `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links []GetFabricNetworkLink `pulumi:"links"`
	// Name of the fabric network.
	Name string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrganizationId string `pulumi:"organizationId"`
	// Email of the user that owns the entity.
	Owner string `pulumi:"owner"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []GetFabricNetworkTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetFabricNetworkOutput(ctx *pulumi.Context, args GetFabricNetworkOutputArgs, opts ...pulumi.InvokeOption) GetFabricNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFabricNetworkResult, error) {
			args := v.(GetFabricNetworkArgs)
			r, err := GetFabricNetwork(ctx, &args, opts...)
			var s GetFabricNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFabricNetworkResultOutput)
}

// A collection of arguments for invoking getFabricNetwork.
type GetFabricNetworkOutputArgs struct {
	// Filter query string that is supported by vRA multi-cloud IaaS API.
	Filter pulumi.StringInput `pulumi:"filter"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags GetFabricNetworkTagArrayInput `pulumi:"tags"`
}

func (GetFabricNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFabricNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getFabricNetwork.
type GetFabricNetworkResultOutput struct{ *pulumi.OutputState }

func (GetFabricNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFabricNetworkResult)(nil)).Elem()
}

func (o GetFabricNetworkResultOutput) ToGetFabricNetworkResultOutput() GetFabricNetworkResultOutput {
	return o
}

func (o GetFabricNetworkResultOutput) ToGetFabricNetworkResultOutputWithContext(ctx context.Context) GetFabricNetworkResultOutput {
	return o
}

// Network CIDR to be used.
func (o GetFabricNetworkResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Cidr }).(pulumi.StringOutput)
}

// Set of ids of the cloud accounts this entity belongs to.
func (o GetFabricNetworkResultOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) []string { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetFabricNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Additional properties that may be used to extend the base resource.
func (o GetFabricNetworkResultOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) map[string]interface{} { return v.CustomProperties }).(pulumi.MapOutput)
}

// State object representing a network on a external cloud provider.
func (o GetFabricNetworkResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity Id on the provider side.
func (o GetFabricNetworkResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this network is defined.
func (o GetFabricNetworkResultOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.ExternalRegionId }).(pulumi.StringOutput)
}

func (o GetFabricNetworkResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFabricNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether this is the default subnet for the zone.
func (o GetFabricNetworkResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// Indicates whether the sub-network supports public IP assignment.
func (o GetFabricNetworkResultOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) bool { return v.IsPublic }).(pulumi.BoolOutput)
}

// HATEOAS of the entity
func (o GetFabricNetworkResultOutput) Links() GetFabricNetworkLinkArrayOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) []GetFabricNetworkLink { return v.Links }).(GetFabricNetworkLinkArrayOutput)
}

// Name of the fabric network.
func (o GetFabricNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o GetFabricNetworkResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o GetFabricNetworkResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o GetFabricNetworkResultOutput) Tags() GetFabricNetworkTagArrayOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) []GetFabricNetworkTag { return v.Tags }).(GetFabricNetworkTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetFabricNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetFabricNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFabricNetworkResultOutput{})
}
