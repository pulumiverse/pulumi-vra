// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates a VMware vRealize Automation fabricNetworkVsphere resource.
//
// ## Example Usage
// ### S
//
// You cannot create a vSphere fabric network resource, however you can import using the command specified in the import section below.
// Once a resource is imported, you can update it as shown below:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/fabric"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/fabric"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fabric.NewNetworkVSphere(ctx, "simple", &fabric.NetworkVSphereArgs{
// 			Cidr:           pulumi.Any(_var.Cidr),
// 			DefaultGateway: pulumi.Any(_var.Gateway),
// 			Domain:         pulumi.Any(_var.Domain),
// 			Tags: fabric.NetworkVSphereTagArray{
// 				&fabric.NetworkVSphereTagArgs{
// 					Key:   pulumi.String("foo"),
// 					Value: pulumi.String("bar"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// To import the vSphere fabric network resource, use the ID as in the following example
//
// ```sh
//  $ pulumi import vra:fabric/networkVSphere:NetworkVSphere new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
// ```
type NetworkVSphere struct {
	pulumi.CustomResourceState

	// Network CIDR to be used.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayOutput `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        pulumi.StringOutput `pulumi:"createdAt"`
	CustomProperties pulumi.MapOutput    `pulumi:"customProperties"`
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringOutput `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringOutput `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayOutput `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayOutput `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// External entity Id on the provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrOutput `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links NetworkVSphereLinkArrayOutput `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  pulumi.StringOutput `pulumi:"name"`
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of organization that entity belongs to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NetworkVSphereTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewNetworkVSphere registers a new resource with the given unique name, arguments, and options.
func NewNetworkVSphere(ctx *pulumi.Context,
	name string, args *NetworkVSphereArgs, opts ...pulumi.ResourceOption) (*NetworkVSphere, error) {
	if args == nil {
		args = &NetworkVSphereArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource NetworkVSphere
	err := ctx.RegisterResource("vra:fabric/networkVSphere:NetworkVSphere", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkVSphere gets an existing NetworkVSphere resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkVSphere(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkVSphereState, opts ...pulumi.ResourceOption) (*NetworkVSphere, error) {
	var resource NetworkVSphere
	err := ctx.ReadResource("vra:fabric/networkVSphere:NetworkVSphere", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkVSphere resources.
type networkVSphereState struct {
	// Network CIDR to be used.
	Cidr *string `pulumi:"cidr"`
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds []string `pulumi:"cloudAccountIds"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        *string                `pulumi:"createdAt"`
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// IPv4 default gateway to be used.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway *string `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains []string `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses []string `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain *string `pulumi:"domain"`
	// External entity Id on the provider side.
	ExternalId *string `pulumi:"externalId"`
	// The id of the region for which this network is defined.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr *string `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `pulumi:"isPublic"`
	// HATEOAS of the entity
	Links []NetworkVSphereLink `pulumi:"links"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  *string `pulumi:"name"`
	OrgId *string `pulumi:"orgId"`
	// ID of organization that entity belongs to.
	OrganizationId *string `pulumi:"organizationId"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []NetworkVSphereTag `pulumi:"tags"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type NetworkVSphereState struct {
	// Network CIDR to be used.
	Cidr pulumi.StringPtrInput
	// Set of ids of the cloud accounts this entity belongs to.
	CloudAccountIds pulumi.StringArrayInput
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt        pulumi.StringPtrInput
	CustomProperties pulumi.MapInput
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringPtrInput
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringPtrInput
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayInput
	// Domain for the vSphere network.
	Domain pulumi.StringPtrInput
	// External entity Id on the provider side.
	ExternalId pulumi.StringPtrInput
	// The id of the region for which this network is defined.
	ExternalRegionId pulumi.StringPtrInput
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrInput
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrInput
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrInput
	// HATEOAS of the entity
	Links NetworkVSphereLinkArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	Name  pulumi.StringPtrInput
	OrgId pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrganizationId pulumi.StringPtrInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NetworkVSphereTagArrayInput
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (NetworkVSphereState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVSphereState)(nil)).Elem()
}

type networkVSphereArgs struct {
	// Network CIDR to be used.
	Cidr *string `pulumi:"cidr"`
	// IPv4 default gateway to be used.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway *string `pulumi:"defaultIpv6Gateway"`
	// List of dns search domains for the vSphere network.
	DnsSearchDomains []string `pulumi:"dnsSearchDomains"`
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses []string `pulumi:"dnsServerAddresses"`
	// Domain for the vSphere network.
	Domain *string `pulumi:"domain"`
	// Network IPv6 CIDR to be used.
	Ipv6Cidr *string `pulumi:"ipv6Cidr"`
	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `pulumi:"isDefault"`
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `pulumi:"isPublic"`
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []NetworkVSphereTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkVSphere resource.
type NetworkVSphereArgs struct {
	// Network CIDR to be used.
	Cidr pulumi.StringPtrInput
	// IPv4 default gateway to be used.
	DefaultGateway pulumi.StringPtrInput
	// IPv6 default gateway to be used.
	DefaultIpv6Gateway pulumi.StringPtrInput
	// List of dns search domains for the vSphere network.
	DnsSearchDomains pulumi.StringArrayInput
	// A human-friendly name used as an identifier in APIs that support this option.
	DnsServerAddresses pulumi.StringArrayInput
	// Domain for the vSphere network.
	Domain pulumi.StringPtrInput
	// Network IPv6 CIDR to be used.
	Ipv6Cidr pulumi.StringPtrInput
	// Indicates whether this is the default subnet for the zone.
	IsDefault pulumi.BoolPtrInput
	// Indicates whether the sub-network supports public IP assignment.
	IsPublic pulumi.BoolPtrInput
	// Set of tag keys and values to apply to the resource.
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NetworkVSphereTagArrayInput
}

func (NetworkVSphereArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkVSphereArgs)(nil)).Elem()
}

type NetworkVSphereInput interface {
	pulumi.Input

	ToNetworkVSphereOutput() NetworkVSphereOutput
	ToNetworkVSphereOutputWithContext(ctx context.Context) NetworkVSphereOutput
}

func (*NetworkVSphere) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVSphere)(nil)).Elem()
}

func (i *NetworkVSphere) ToNetworkVSphereOutput() NetworkVSphereOutput {
	return i.ToNetworkVSphereOutputWithContext(context.Background())
}

func (i *NetworkVSphere) ToNetworkVSphereOutputWithContext(ctx context.Context) NetworkVSphereOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVSphereOutput)
}

// NetworkVSphereArrayInput is an input type that accepts NetworkVSphereArray and NetworkVSphereArrayOutput values.
// You can construct a concrete instance of `NetworkVSphereArrayInput` via:
//
//          NetworkVSphereArray{ NetworkVSphereArgs{...} }
type NetworkVSphereArrayInput interface {
	pulumi.Input

	ToNetworkVSphereArrayOutput() NetworkVSphereArrayOutput
	ToNetworkVSphereArrayOutputWithContext(context.Context) NetworkVSphereArrayOutput
}

type NetworkVSphereArray []NetworkVSphereInput

func (NetworkVSphereArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkVSphere)(nil)).Elem()
}

func (i NetworkVSphereArray) ToNetworkVSphereArrayOutput() NetworkVSphereArrayOutput {
	return i.ToNetworkVSphereArrayOutputWithContext(context.Background())
}

func (i NetworkVSphereArray) ToNetworkVSphereArrayOutputWithContext(ctx context.Context) NetworkVSphereArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVSphereArrayOutput)
}

// NetworkVSphereMapInput is an input type that accepts NetworkVSphereMap and NetworkVSphereMapOutput values.
// You can construct a concrete instance of `NetworkVSphereMapInput` via:
//
//          NetworkVSphereMap{ "key": NetworkVSphereArgs{...} }
type NetworkVSphereMapInput interface {
	pulumi.Input

	ToNetworkVSphereMapOutput() NetworkVSphereMapOutput
	ToNetworkVSphereMapOutputWithContext(context.Context) NetworkVSphereMapOutput
}

type NetworkVSphereMap map[string]NetworkVSphereInput

func (NetworkVSphereMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkVSphere)(nil)).Elem()
}

func (i NetworkVSphereMap) ToNetworkVSphereMapOutput() NetworkVSphereMapOutput {
	return i.ToNetworkVSphereMapOutputWithContext(context.Background())
}

func (i NetworkVSphereMap) ToNetworkVSphereMapOutputWithContext(ctx context.Context) NetworkVSphereMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkVSphereMapOutput)
}

type NetworkVSphereOutput struct{ *pulumi.OutputState }

func (NetworkVSphereOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkVSphere)(nil)).Elem()
}

func (o NetworkVSphereOutput) ToNetworkVSphereOutput() NetworkVSphereOutput {
	return o
}

func (o NetworkVSphereOutput) ToNetworkVSphereOutputWithContext(ctx context.Context) NetworkVSphereOutput {
	return o
}

// Network CIDR to be used.
func (o NetworkVSphereOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// Set of ids of the cloud accounts this entity belongs to.
func (o NetworkVSphereOutput) CloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringArrayOutput { return v.CloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o NetworkVSphereOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o NetworkVSphereOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// IPv4 default gateway to be used.
func (o NetworkVSphereOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.DefaultGateway }).(pulumi.StringOutput)
}

// IPv6 default gateway to be used.
func (o NetworkVSphereOutput) DefaultIpv6Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.DefaultIpv6Gateway }).(pulumi.StringOutput)
}

// List of dns search domains for the vSphere network.
func (o NetworkVSphereOutput) DnsSearchDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringArrayOutput { return v.DnsSearchDomains }).(pulumi.StringArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o NetworkVSphereOutput) DnsServerAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringArrayOutput { return v.DnsServerAddresses }).(pulumi.StringArrayOutput)
}

// Domain for the vSphere network.
func (o NetworkVSphereOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// External entity Id on the provider side.
func (o NetworkVSphereOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// The id of the region for which this network is defined.
func (o NetworkVSphereOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// Network IPv6 CIDR to be used.
func (o NetworkVSphereOutput) Ipv6Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringPtrOutput { return v.Ipv6Cidr }).(pulumi.StringPtrOutput)
}

// Indicates whether this is the default subnet for the zone.
func (o NetworkVSphereOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// Indicates whether the sub-network supports public IP assignment.
func (o NetworkVSphereOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// HATEOAS of the entity
func (o NetworkVSphereOutput) Links() NetworkVSphereLinkArrayOutput {
	return o.ApplyT(func(v *NetworkVSphere) NetworkVSphereLinkArrayOutput { return v.Links }).(NetworkVSphereLinkArrayOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o NetworkVSphereOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkVSphereOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o NetworkVSphereOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource.
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o NetworkVSphereOutput) Tags() NetworkVSphereTagArrayOutput {
	return o.ApplyT(func(v *NetworkVSphere) NetworkVSphereTagArrayOutput { return v.Tags }).(NetworkVSphereTagArrayOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o NetworkVSphereOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkVSphere) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type NetworkVSphereArrayOutput struct{ *pulumi.OutputState }

func (NetworkVSphereArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkVSphere)(nil)).Elem()
}

func (o NetworkVSphereArrayOutput) ToNetworkVSphereArrayOutput() NetworkVSphereArrayOutput {
	return o
}

func (o NetworkVSphereArrayOutput) ToNetworkVSphereArrayOutputWithContext(ctx context.Context) NetworkVSphereArrayOutput {
	return o
}

func (o NetworkVSphereArrayOutput) Index(i pulumi.IntInput) NetworkVSphereOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkVSphere {
		return vs[0].([]*NetworkVSphere)[vs[1].(int)]
	}).(NetworkVSphereOutput)
}

type NetworkVSphereMapOutput struct{ *pulumi.OutputState }

func (NetworkVSphereMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkVSphere)(nil)).Elem()
}

func (o NetworkVSphereMapOutput) ToNetworkVSphereMapOutput() NetworkVSphereMapOutput {
	return o
}

func (o NetworkVSphereMapOutput) ToNetworkVSphereMapOutputWithContext(ctx context.Context) NetworkVSphereMapOutput {
	return o
}

func (o NetworkVSphereMapOutput) MapIndex(k pulumi.StringInput) NetworkVSphereOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkVSphere {
		return vs[0].(map[string]*NetworkVSphere)[vs[1].(string)]
	}).(NetworkVSphereOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkVSphereInput)(nil)).Elem(), &NetworkVSphere{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkVSphereArrayInput)(nil)).Elem(), NetworkVSphereArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkVSphereMapInput)(nil)).Elem(), NetworkVSphereMap{})
	pulumi.RegisterOutputType(NetworkVSphereOutput{})
	pulumi.RegisterOutputType(NetworkVSphereArrayOutput{})
	pulumi.RegisterOutputType(NetworkVSphereMapOutput{})
}
