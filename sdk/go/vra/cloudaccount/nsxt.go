// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation NSX-T cloud account resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create an NSX-T cloud account resource.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudaccount.NewNsxt(ctx, "this", &cloudaccount.NsxtArgs{
// 			Description:          pulumi.String("foobar"),
// 			Username:             pulumi.Any(_var.Username),
// 			Password:             pulumi.Any(_var.Password),
// 			Hostname:             pulumi.Any(_var.Hostname),
// 			DcId:                 pulumi.Any(_var.Vra_data_collector_id),
// 			AcceptSelfSignedCert: pulumi.Bool(true),
// 			Tags: cloudaccount.NsxtTagArray{
// 				&cloudaccount.NsxtTagArgs{
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
// To import the NSX-T cloud account, use the ID as in the following example
//
// ```sh
//  $ pulumi import vra:cloudaccount/nsxt:Nsxt new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15`
// ```
type Nsxt struct {
	pulumi.CustomResourceState

	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrOutput `pulumi:"acceptSelfSignedCert"`
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds pulumi.StringArrayOutput `pulumi:"associatedCloudAccountIds"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identifier of a data collector VM deployed in the on premise infrastructure.
	DcId pulumi.StringPtrOutput `pulumi:"dcId"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Host name for NSX-T cloud account.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// HATEOAS of entity.
	Links NsxtLinkArrayOutput `pulumi:"links"`
	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
	// Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode pulumi.BoolPtrOutput `pulumi:"managerMode"`
	// Name of NSX-T cloud account.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Password used to authenticate to the cloud Account.
	Password pulumi.StringOutput `pulumi:"password"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NsxtTagArrayOutput `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Username used to authenticate to the cloud account.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewNsxt registers a new resource with the given unique name, arguments, and options.
func NewNsxt(ctx *pulumi.Context,
	name string, args *NsxtArgs, opts ...pulumi.ResourceOption) (*Nsxt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Nsxt
	err := ctx.RegisterResource("vra:cloudaccount/nsxt:Nsxt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxt gets an existing Nsxt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtState, opts ...pulumi.ResourceOption) (*Nsxt, error) {
	var resource Nsxt
	err := ctx.ReadResource("vra:cloudaccount/nsxt:Nsxt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nsxt resources.
type nsxtState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds []string `pulumi:"associatedCloudAccountIds"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Identifier of a data collector VM deployed in the on premise infrastructure.
	DcId *string `pulumi:"dcId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Host name for NSX-T cloud account.
	Hostname *string `pulumi:"hostname"`
	// HATEOAS of entity.
	Links []NsxtLink `pulumi:"links"`
	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
	// Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode *bool `pulumi:"managerMode"`
	// Name of NSX-T cloud account.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// Password used to authenticate to the cloud Account.
	Password *string `pulumi:"password"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []NsxtTag `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Username used to authenticate to the cloud account.
	Username *string `pulumi:"username"`
}

type NsxtState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// Cloud accounts associated with the cloud account.
	AssociatedCloudAccountIds pulumi.StringArrayInput
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Identifier of a data collector VM deployed in the on premise infrastructure.
	DcId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Host name for NSX-T cloud account.
	Hostname pulumi.StringPtrInput
	// HATEOAS of entity.
	Links NsxtLinkArrayInput
	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
	// Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode pulumi.BoolPtrInput
	// Name of NSX-T cloud account.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// Password used to authenticate to the cloud Account.
	Password pulumi.StringPtrInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NsxtTagArrayInput
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
	// Username used to authenticate to the cloud account.
	Username pulumi.StringPtrInput
}

func (NsxtState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtState)(nil)).Elem()
}

type nsxtArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// Identifier of a data collector VM deployed in the on premise infrastructure.
	DcId *string `pulumi:"dcId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Host name for NSX-T cloud account.
	Hostname string `pulumi:"hostname"`
	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
	// Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode *bool `pulumi:"managerMode"`
	// Name of NSX-T cloud account.
	Name *string `pulumi:"name"`
	// Password used to authenticate to the cloud Account.
	Password string `pulumi:"password"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []NsxtTag `pulumi:"tags"`
	// Username used to authenticate to the cloud account.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Nsxt resource.
type NsxtArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// Identifier of a data collector VM deployed in the on premise infrastructure.
	DcId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Host name for NSX-T cloud account.
	Hostname pulumi.StringInput
	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
	// Mode cannot be changed after cloud account is created. Default value is false.
	ManagerMode pulumi.BoolPtrInput
	// Name of NSX-T cloud account.
	Name pulumi.StringPtrInput
	// Password used to authenticate to the cloud Account.
	Password pulumi.StringInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags NsxtTagArrayInput
	// Username used to authenticate to the cloud account.
	Username pulumi.StringInput
}

func (NsxtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtArgs)(nil)).Elem()
}

type NsxtInput interface {
	pulumi.Input

	ToNsxtOutput() NsxtOutput
	ToNsxtOutputWithContext(ctx context.Context) NsxtOutput
}

func (*Nsxt) ElementType() reflect.Type {
	return reflect.TypeOf((**Nsxt)(nil)).Elem()
}

func (i *Nsxt) ToNsxtOutput() NsxtOutput {
	return i.ToNsxtOutputWithContext(context.Background())
}

func (i *Nsxt) ToNsxtOutputWithContext(ctx context.Context) NsxtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtOutput)
}

// NsxtArrayInput is an input type that accepts NsxtArray and NsxtArrayOutput values.
// You can construct a concrete instance of `NsxtArrayInput` via:
//
//          NsxtArray{ NsxtArgs{...} }
type NsxtArrayInput interface {
	pulumi.Input

	ToNsxtArrayOutput() NsxtArrayOutput
	ToNsxtArrayOutputWithContext(context.Context) NsxtArrayOutput
}

type NsxtArray []NsxtInput

func (NsxtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nsxt)(nil)).Elem()
}

func (i NsxtArray) ToNsxtArrayOutput() NsxtArrayOutput {
	return i.ToNsxtArrayOutputWithContext(context.Background())
}

func (i NsxtArray) ToNsxtArrayOutputWithContext(ctx context.Context) NsxtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtArrayOutput)
}

// NsxtMapInput is an input type that accepts NsxtMap and NsxtMapOutput values.
// You can construct a concrete instance of `NsxtMapInput` via:
//
//          NsxtMap{ "key": NsxtArgs{...} }
type NsxtMapInput interface {
	pulumi.Input

	ToNsxtMapOutput() NsxtMapOutput
	ToNsxtMapOutputWithContext(context.Context) NsxtMapOutput
}

type NsxtMap map[string]NsxtInput

func (NsxtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nsxt)(nil)).Elem()
}

func (i NsxtMap) ToNsxtMapOutput() NsxtMapOutput {
	return i.ToNsxtMapOutputWithContext(context.Background())
}

func (i NsxtMap) ToNsxtMapOutputWithContext(ctx context.Context) NsxtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtMapOutput)
}

type NsxtOutput struct{ *pulumi.OutputState }

func (NsxtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nsxt)(nil)).Elem()
}

func (o NsxtOutput) ToNsxtOutput() NsxtOutput {
	return o
}

func (o NsxtOutput) ToNsxtOutputWithContext(ctx context.Context) NsxtOutput {
	return o
}

// Accept self-signed certificate when connecting to the cloud account.
func (o NsxtOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.BoolPtrOutput { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

// Cloud accounts associated with the cloud account.
func (o NsxtOutput) AssociatedCloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringArrayOutput { return v.AssociatedCloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o NsxtOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier of a data collector VM deployed in the on premise infrastructure.
func (o NsxtOutput) DcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringPtrOutput { return v.DcId }).(pulumi.StringPtrOutput)
}

// Human-friendly description.
func (o NsxtOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Host name for NSX-T cloud account.
func (o NsxtOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// HATEOAS of entity.
func (o NsxtOutput) Links() NsxtLinkArrayOutput {
	return o.ApplyT(func(v *Nsxt) NsxtLinkArrayOutput { return v.Links }).(NsxtLinkArrayOutput)
}

// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
// Mode cannot be changed after cloud account is created. Default value is false.
func (o NsxtOutput) ManagerMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.BoolPtrOutput { return v.ManagerMode }).(pulumi.BoolPtrOutput)
}

// Name of NSX-T cloud account.
func (o NsxtOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o NsxtOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o NsxtOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Password used to authenticate to the cloud Account.
func (o NsxtOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the cloud account.\
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o NsxtOutput) Tags() NsxtTagArrayOutput {
	return o.ApplyT(func(v *Nsxt) NsxtTagArrayOutput { return v.Tags }).(NsxtTagArrayOutput)
}

// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
func (o NsxtOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Username used to authenticate to the cloud account.
func (o NsxtOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Nsxt) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type NsxtArrayOutput struct{ *pulumi.OutputState }

func (NsxtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nsxt)(nil)).Elem()
}

func (o NsxtArrayOutput) ToNsxtArrayOutput() NsxtArrayOutput {
	return o
}

func (o NsxtArrayOutput) ToNsxtArrayOutputWithContext(ctx context.Context) NsxtArrayOutput {
	return o
}

func (o NsxtArrayOutput) Index(i pulumi.IntInput) NsxtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nsxt {
		return vs[0].([]*Nsxt)[vs[1].(int)]
	}).(NsxtOutput)
}

type NsxtMapOutput struct{ *pulumi.OutputState }

func (NsxtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nsxt)(nil)).Elem()
}

func (o NsxtMapOutput) ToNsxtMapOutput() NsxtMapOutput {
	return o
}

func (o NsxtMapOutput) ToNsxtMapOutputWithContext(ctx context.Context) NsxtMapOutput {
	return o
}

func (o NsxtMapOutput) MapIndex(k pulumi.StringInput) NsxtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nsxt {
		return vs[0].(map[string]*Nsxt)[vs[1].(string)]
	}).(NsxtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtInput)(nil)).Elem(), &Nsxt{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtArrayInput)(nil)).Elem(), NsxtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtMapInput)(nil)).Elem(), NsxtMap{})
	pulumi.RegisterOutputType(NsxtOutput{})
	pulumi.RegisterOutputType(NsxtArrayOutput{})
	pulumi.RegisterOutputType(NsxtMapOutput{})
}
