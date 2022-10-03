// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### S
// This is an example of how to create a network profile resource.
//
// **Network profile:**
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
//			_, err := vra.NewNetworkProfile(ctx, "simple", &vra.NetworkProfileArgs{
//				Description: pulumi.String("Simple Network Profile with no isolation."),
//				RegionId:    pulumi.Any(data.Vra_region.This.Id),
//				FabricNetworkIds: pulumi.StringArray{
//					pulumi.Any(data.Vra_fabric_network.Subnet_1.Id),
//					pulumi.Any(data.Vra_fabric_network.Subnet_2.Id),
//				},
//				IsolationType: pulumi.String("NONE"),
//				Tags: NetworkProfileTagArray{
//					&NetworkProfileTagArgs{
//						Key:   pulumi.String("foo"),
//						Value: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// A network profile resource supports the following arguments:
type NetworkProfile struct {
	pulumi.CustomResourceState

	// The ID of the cloud account this entity belongs to.
	CloudAccountId pulumi.StringOutput `pulumi:"cloudAccountId"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	CustomProperties pulumi.MapOutput `pulumi:"customProperties"`
	// A human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The external regionId of the resource.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// A list of fabric network Ids which are assigned to the network profile.
	// example:[ "6543" ]
	FabricNetworkIds pulumi.StringArrayOutput `pulumi:"fabricNetworkIds"`
	// The id of the fabric network used for outbound access.
	IsolatedExternalFabricNetworkId pulumi.StringPtrOutput `pulumi:"isolatedExternalFabricNetworkId"`
	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCidrPrefix pulumi.IntPtrOutput `pulumi:"isolatedNetworkCidrPrefix"`
	// CIDR of the isolation network domain.
	IsolatedNetworkDomainCidr pulumi.StringPtrOutput `pulumi:"isolatedNetworkDomainCidr"`
	// The id of the network domain used for creating isolated networks.
	IsolatedNetworkDomainId pulumi.StringPtrOutput `pulumi:"isolatedNetworkDomainId"`
	// Specifies the isolation type e.g. none, subnet or security group
	IsolationType pulumi.StringPtrOutput `pulumi:"isolationType"`
	// HATEOAS of the entity
	Links NetworkProfileLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// The id of the organization this entity belongs to. Deprecated, refer to orgId instead.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Email of the user that owns the entity.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// A list of security group Ids which are assigned to the network profile.
	// example:[ "6545" ]
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags NetworkProfileTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewNetworkProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NetworkProfile
	err := ctx.RegisterResource("vra:index/networkProfile:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkProfile gets an existing NetworkProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("vra:index/networkProfile:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkProfile resources.
type networkProfileState struct {
	// The ID of the cloud account this entity belongs to.
	CloudAccountId *string `pulumi:"cloudAccountId"`
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// The external regionId of the resource.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// A list of fabric network Ids which are assigned to the network profile.
	// example:[ "6543" ]
	FabricNetworkIds []string `pulumi:"fabricNetworkIds"`
	// The id of the fabric network used for outbound access.
	IsolatedExternalFabricNetworkId *string `pulumi:"isolatedExternalFabricNetworkId"`
	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCidrPrefix *int `pulumi:"isolatedNetworkCidrPrefix"`
	// CIDR of the isolation network domain.
	IsolatedNetworkDomainCidr *string `pulumi:"isolatedNetworkDomainCidr"`
	// The id of the network domain used for creating isolated networks.
	IsolatedNetworkDomainId *string `pulumi:"isolatedNetworkDomainId"`
	// Specifies the isolation type e.g. none, subnet or security group
	IsolationType *string `pulumi:"isolationType"`
	// HATEOAS of the entity
	Links []NetworkProfileLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// The id of the organization this entity belongs to. Deprecated, refer to orgId instead.
	OrganizationId *string `pulumi:"organizationId"`
	// Email of the user that owns the entity.
	Owner *string `pulumi:"owner"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId *string `pulumi:"regionId"`
	// A list of security group Ids which are assigned to the network profile.
	// example:[ "6545" ]
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []NetworkProfileTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type NetworkProfileState struct {
	// The ID of the cloud account this entity belongs to.
	CloudAccountId pulumi.StringPtrInput
	// Date when  entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	CustomProperties pulumi.MapInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// The external regionId of the resource.
	ExternalRegionId pulumi.StringPtrInput
	// A list of fabric network Ids which are assigned to the network profile.
	// example:[ "6543" ]
	FabricNetworkIds pulumi.StringArrayInput
	// The id of the fabric network used for outbound access.
	IsolatedExternalFabricNetworkId pulumi.StringPtrInput
	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCidrPrefix pulumi.IntPtrInput
	// CIDR of the isolation network domain.
	IsolatedNetworkDomainCidr pulumi.StringPtrInput
	// The id of the network domain used for creating isolated networks.
	IsolatedNetworkDomainId pulumi.StringPtrInput
	// Specifies the isolation type e.g. none, subnet or security group
	IsolationType pulumi.StringPtrInput
	// HATEOAS of the entity
	Links NetworkProfileLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// The id of the organization this entity belongs to. Deprecated, refer to orgId instead.
	OrganizationId pulumi.StringPtrInput
	// Email of the user that owns the entity.
	Owner pulumi.StringPtrInput
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringPtrInput
	// A list of security group Ids which are assigned to the network profile.
	// example:[ "6545" ]
	SecurityGroupIds pulumi.StringArrayInput
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags NetworkProfileTagArrayInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// A human-friendly description.
	Description *string `pulumi:"description"`
	// A list of fabric network Ids which are assigned to the network profile.
	// example:[ "6543" ]
	FabricNetworkIds []string `pulumi:"fabricNetworkIds"`
	// The id of the fabric network used for outbound access.
	IsolatedExternalFabricNetworkId *string `pulumi:"isolatedExternalFabricNetworkId"`
	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCidrPrefix *int `pulumi:"isolatedNetworkCidrPrefix"`
	// CIDR of the isolation network domain.
	IsolatedNetworkDomainCidr *string `pulumi:"isolatedNetworkDomainCidr"`
	// The id of the network domain used for creating isolated networks.
	IsolatedNetworkDomainId *string `pulumi:"isolatedNetworkDomainId"`
	// Specifies the isolation type e.g. none, subnet or security group
	IsolationType *string `pulumi:"isolationType"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId string `pulumi:"regionId"`
	// A list of security group Ids which are assigned to the network profile.
	// example:[ "6545" ]
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags []NetworkProfileTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkProfile resource.
type NetworkProfileArgs struct {
	// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	CustomProperties pulumi.MapInput
	// A human-friendly description.
	Description pulumi.StringPtrInput
	// A list of fabric network Ids which are assigned to the network profile.
	// example:[ "6543" ]
	FabricNetworkIds pulumi.StringArrayInput
	// The id of the fabric network used for outbound access.
	IsolatedExternalFabricNetworkId pulumi.StringPtrInput
	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCidrPrefix pulumi.IntPtrInput
	// CIDR of the isolation network domain.
	IsolatedNetworkDomainCidr pulumi.StringPtrInput
	// The id of the network domain used for creating isolated networks.
	IsolatedNetworkDomainId pulumi.StringPtrInput
	// Specifies the isolation type e.g. none, subnet or security group
	IsolationType pulumi.StringPtrInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// The id of the region for which this profile is defined as in vRealize Automation(vRA).
	RegionId pulumi.StringInput
	// A list of security group Ids which are assigned to the network profile.
	// example:[ "6545" ]
	SecurityGroupIds pulumi.StringArrayInput
	// A set of tag keys and optional values that were set on this Network Profile.
	// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
	Tags NetworkProfileTagArrayInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}

type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput
}

func (*NetworkProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *NetworkProfile) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i *NetworkProfile) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

// NetworkProfileArrayInput is an input type that accepts NetworkProfileArray and NetworkProfileArrayOutput values.
// You can construct a concrete instance of `NetworkProfileArrayInput` via:
//
//	NetworkProfileArray{ NetworkProfileArgs{...} }
type NetworkProfileArrayInput interface {
	pulumi.Input

	ToNetworkProfileArrayOutput() NetworkProfileArrayOutput
	ToNetworkProfileArrayOutputWithContext(context.Context) NetworkProfileArrayOutput
}

type NetworkProfileArray []NetworkProfileInput

func (NetworkProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArray) ToNetworkProfileArrayOutput() NetworkProfileArrayOutput {
	return i.ToNetworkProfileArrayOutputWithContext(context.Background())
}

func (i NetworkProfileArray) ToNetworkProfileArrayOutputWithContext(ctx context.Context) NetworkProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileArrayOutput)
}

// NetworkProfileMapInput is an input type that accepts NetworkProfileMap and NetworkProfileMapOutput values.
// You can construct a concrete instance of `NetworkProfileMapInput` via:
//
//	NetworkProfileMap{ "key": NetworkProfileArgs{...} }
type NetworkProfileMapInput interface {
	pulumi.Input

	ToNetworkProfileMapOutput() NetworkProfileMapOutput
	ToNetworkProfileMapOutputWithContext(context.Context) NetworkProfileMapOutput
}

type NetworkProfileMap map[string]NetworkProfileInput

func (NetworkProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileMap) ToNetworkProfileMapOutput() NetworkProfileMapOutput {
	return i.ToNetworkProfileMapOutputWithContext(context.Background())
}

func (i NetworkProfileMap) ToNetworkProfileMapOutputWithContext(ctx context.Context) NetworkProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileMapOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

// The ID of the cloud account this entity belongs to.
func (o NetworkProfileOutput) CloudAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.CloudAccountId }).(pulumi.StringOutput)
}

// Date when  entity was created. Date and time format is ISO 8601 and UTC.
func (o NetworkProfileOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
func (o NetworkProfileOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// A human-friendly description.
func (o NetworkProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The external regionId of the resource.
func (o NetworkProfileOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// A list of fabric network Ids which are assigned to the network profile.
// example:[ "6543" ]
func (o NetworkProfileOutput) FabricNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringArrayOutput { return v.FabricNetworkIds }).(pulumi.StringArrayOutput)
}

// The id of the fabric network used for outbound access.
func (o NetworkProfileOutput) IsolatedExternalFabricNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.IsolatedExternalFabricNetworkId }).(pulumi.StringPtrOutput)
}

// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
func (o NetworkProfileOutput) IsolatedNetworkCidrPrefix() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.IsolatedNetworkCidrPrefix }).(pulumi.IntPtrOutput)
}

// CIDR of the isolation network domain.
func (o NetworkProfileOutput) IsolatedNetworkDomainCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.IsolatedNetworkDomainCidr }).(pulumi.StringPtrOutput)
}

// The id of the network domain used for creating isolated networks.
func (o NetworkProfileOutput) IsolatedNetworkDomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.IsolatedNetworkDomainId }).(pulumi.StringPtrOutput)
}

// Specifies the isolation type e.g. none, subnet or security group
func (o NetworkProfileOutput) IsolationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.IsolationType }).(pulumi.StringPtrOutput)
}

// HATEOAS of the entity
func (o NetworkProfileOutput) Links() NetworkProfileLinkArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfileLinkArrayOutput { return v.Links }).(NetworkProfileLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o NetworkProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o NetworkProfileOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to. Deprecated, refer to orgId instead.
func (o NetworkProfileOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Email of the user that owns the entity.
func (o NetworkProfileOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The id of the region for which this profile is defined as in vRealize Automation(vRA).
func (o NetworkProfileOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// A list of security group Ids which are assigned to the network profile.
// example:[ "6545" ]
func (o NetworkProfileOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A set of tag keys and optional values that were set on this Network Profile.
// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
func (o NetworkProfileOutput) Tags() NetworkProfileTagArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfileTagArrayOutput { return v.Tags }).(NetworkProfileTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o NetworkProfileOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type NetworkProfileArrayOutput struct{ *pulumi.OutputState }

func (NetworkProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileArrayOutput) ToNetworkProfileArrayOutput() NetworkProfileArrayOutput {
	return o
}

func (o NetworkProfileArrayOutput) ToNetworkProfileArrayOutputWithContext(ctx context.Context) NetworkProfileArrayOutput {
	return o
}

func (o NetworkProfileArrayOutput) Index(i pulumi.IntInput) NetworkProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkProfile {
		return vs[0].([]*NetworkProfile)[vs[1].(int)]
	}).(NetworkProfileOutput)
}

type NetworkProfileMapOutput struct{ *pulumi.OutputState }

func (NetworkProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileMapOutput) ToNetworkProfileMapOutput() NetworkProfileMapOutput {
	return o
}

func (o NetworkProfileMapOutput) ToNetworkProfileMapOutputWithContext(ctx context.Context) NetworkProfileMapOutput {
	return o
}

func (o NetworkProfileMapOutput) MapIndex(k pulumi.StringInput) NetworkProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkProfile {
		return vs[0].(map[string]*NetworkProfile)[vs[1].(string)]
	}).(NetworkProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileInput)(nil)).Elem(), &NetworkProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileArrayInput)(nil)).Elem(), NetworkProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileMapInput)(nil)).Elem(), NetworkProfileMap{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfileArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileMapOutput{})
}
