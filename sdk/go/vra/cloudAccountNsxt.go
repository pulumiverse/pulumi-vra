// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

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
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.NewCloudAccountNsxt(ctx, "this", &vra.CloudAccountNsxtArgs{
//				Description:          pulumi.String("foobar"),
//				Username:             pulumi.Any(_var.Username),
//				Password:             pulumi.Any(_var.Password),
//				Hostname:             pulumi.Any(_var.Hostname),
//				DcId:                 pulumi.Any(_var.Vra_data_collector_id),
//				AcceptSelfSignedCert: pulumi.Bool(true),
//				Tags: CloudAccountNsxtTagArray{
//					&CloudAccountNsxtTagArgs{
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
// ## Import
//
// # To import the NSX-T cloud account, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:index/cloudAccountNsxt:CloudAccountNsxt new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type CloudAccountNsxt struct {
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
	Links CloudAccountNsxtLinkArrayOutput `pulumi:"links"`
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
	Tags CloudAccountNsxtTagArrayOutput `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Username used to authenticate to the cloud account.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewCloudAccountNsxt registers a new resource with the given unique name, arguments, and options.
func NewCloudAccountNsxt(ctx *pulumi.Context,
	name string, args *CloudAccountNsxtArgs, opts ...pulumi.ResourceOption) (*CloudAccountNsxt, error) {
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
	var resource CloudAccountNsxt
	err := ctx.RegisterResource("vra:index/cloudAccountNsxt:CloudAccountNsxt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudAccountNsxt gets an existing CloudAccountNsxt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudAccountNsxt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudAccountNsxtState, opts ...pulumi.ResourceOption) (*CloudAccountNsxt, error) {
	var resource CloudAccountNsxt
	err := ctx.ReadResource("vra:index/cloudAccountNsxt:CloudAccountNsxt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudAccountNsxt resources.
type cloudAccountNsxtState struct {
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
	Links []CloudAccountNsxtLink `pulumi:"links"`
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
	Tags []CloudAccountNsxtTag `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Username used to authenticate to the cloud account.
	Username *string `pulumi:"username"`
}

type CloudAccountNsxtState struct {
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
	Links CloudAccountNsxtLinkArrayInput
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
	Tags CloudAccountNsxtTagArrayInput
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
	// Username used to authenticate to the cloud account.
	Username pulumi.StringPtrInput
}

func (CloudAccountNsxtState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountNsxtState)(nil)).Elem()
}

type cloudAccountNsxtArgs struct {
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
	Tags []CloudAccountNsxtTag `pulumi:"tags"`
	// Username used to authenticate to the cloud account.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a CloudAccountNsxt resource.
type CloudAccountNsxtArgs struct {
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
	Tags CloudAccountNsxtTagArrayInput
	// Username used to authenticate to the cloud account.
	Username pulumi.StringInput
}

func (CloudAccountNsxtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudAccountNsxtArgs)(nil)).Elem()
}

type CloudAccountNsxtInput interface {
	pulumi.Input

	ToCloudAccountNsxtOutput() CloudAccountNsxtOutput
	ToCloudAccountNsxtOutputWithContext(ctx context.Context) CloudAccountNsxtOutput
}

func (*CloudAccountNsxt) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccountNsxt)(nil)).Elem()
}

func (i *CloudAccountNsxt) ToCloudAccountNsxtOutput() CloudAccountNsxtOutput {
	return i.ToCloudAccountNsxtOutputWithContext(context.Background())
}

func (i *CloudAccountNsxt) ToCloudAccountNsxtOutputWithContext(ctx context.Context) CloudAccountNsxtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountNsxtOutput)
}

// CloudAccountNsxtArrayInput is an input type that accepts CloudAccountNsxtArray and CloudAccountNsxtArrayOutput values.
// You can construct a concrete instance of `CloudAccountNsxtArrayInput` via:
//
//	CloudAccountNsxtArray{ CloudAccountNsxtArgs{...} }
type CloudAccountNsxtArrayInput interface {
	pulumi.Input

	ToCloudAccountNsxtArrayOutput() CloudAccountNsxtArrayOutput
	ToCloudAccountNsxtArrayOutputWithContext(context.Context) CloudAccountNsxtArrayOutput
}

type CloudAccountNsxtArray []CloudAccountNsxtInput

func (CloudAccountNsxtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccountNsxt)(nil)).Elem()
}

func (i CloudAccountNsxtArray) ToCloudAccountNsxtArrayOutput() CloudAccountNsxtArrayOutput {
	return i.ToCloudAccountNsxtArrayOutputWithContext(context.Background())
}

func (i CloudAccountNsxtArray) ToCloudAccountNsxtArrayOutputWithContext(ctx context.Context) CloudAccountNsxtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountNsxtArrayOutput)
}

// CloudAccountNsxtMapInput is an input type that accepts CloudAccountNsxtMap and CloudAccountNsxtMapOutput values.
// You can construct a concrete instance of `CloudAccountNsxtMapInput` via:
//
//	CloudAccountNsxtMap{ "key": CloudAccountNsxtArgs{...} }
type CloudAccountNsxtMapInput interface {
	pulumi.Input

	ToCloudAccountNsxtMapOutput() CloudAccountNsxtMapOutput
	ToCloudAccountNsxtMapOutputWithContext(context.Context) CloudAccountNsxtMapOutput
}

type CloudAccountNsxtMap map[string]CloudAccountNsxtInput

func (CloudAccountNsxtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccountNsxt)(nil)).Elem()
}

func (i CloudAccountNsxtMap) ToCloudAccountNsxtMapOutput() CloudAccountNsxtMapOutput {
	return i.ToCloudAccountNsxtMapOutputWithContext(context.Background())
}

func (i CloudAccountNsxtMap) ToCloudAccountNsxtMapOutputWithContext(ctx context.Context) CloudAccountNsxtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudAccountNsxtMapOutput)
}

type CloudAccountNsxtOutput struct{ *pulumi.OutputState }

func (CloudAccountNsxtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudAccountNsxt)(nil)).Elem()
}

func (o CloudAccountNsxtOutput) ToCloudAccountNsxtOutput() CloudAccountNsxtOutput {
	return o
}

func (o CloudAccountNsxtOutput) ToCloudAccountNsxtOutputWithContext(ctx context.Context) CloudAccountNsxtOutput {
	return o
}

// Accept self-signed certificate when connecting to the cloud account.
func (o CloudAccountNsxtOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.BoolPtrOutput { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

// Cloud accounts associated with the cloud account.
func (o CloudAccountNsxtOutput) AssociatedCloudAccountIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringArrayOutput { return v.AssociatedCloudAccountIds }).(pulumi.StringArrayOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o CloudAccountNsxtOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier of a data collector VM deployed in the on premise infrastructure.
func (o CloudAccountNsxtOutput) DcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringPtrOutput { return v.DcId }).(pulumi.StringPtrOutput)
}

// Human-friendly description.
func (o CloudAccountNsxtOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Host name for NSX-T cloud account.
func (o CloudAccountNsxtOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// HATEOAS of entity.
func (o CloudAccountNsxtOutput) Links() CloudAccountNsxtLinkArrayOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) CloudAccountNsxtLinkArrayOutput { return v.Links }).(CloudAccountNsxtLinkArrayOutput)
}

// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
// Mode cannot be changed after cloud account is created. Default value is false.
func (o CloudAccountNsxtOutput) ManagerMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.BoolPtrOutput { return v.ManagerMode }).(pulumi.BoolPtrOutput)
}

// Name of NSX-T cloud account.
func (o CloudAccountNsxtOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o CloudAccountNsxtOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o CloudAccountNsxtOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Password used to authenticate to the cloud Account.
func (o CloudAccountNsxtOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the cloud account.\
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o CloudAccountNsxtOutput) Tags() CloudAccountNsxtTagArrayOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) CloudAccountNsxtTagArrayOutput { return v.Tags }).(CloudAccountNsxtTagArrayOutput)
}

// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
func (o CloudAccountNsxtOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Username used to authenticate to the cloud account.
func (o CloudAccountNsxtOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudAccountNsxt) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type CloudAccountNsxtArrayOutput struct{ *pulumi.OutputState }

func (CloudAccountNsxtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudAccountNsxt)(nil)).Elem()
}

func (o CloudAccountNsxtArrayOutput) ToCloudAccountNsxtArrayOutput() CloudAccountNsxtArrayOutput {
	return o
}

func (o CloudAccountNsxtArrayOutput) ToCloudAccountNsxtArrayOutputWithContext(ctx context.Context) CloudAccountNsxtArrayOutput {
	return o
}

func (o CloudAccountNsxtArrayOutput) Index(i pulumi.IntInput) CloudAccountNsxtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudAccountNsxt {
		return vs[0].([]*CloudAccountNsxt)[vs[1].(int)]
	}).(CloudAccountNsxtOutput)
}

type CloudAccountNsxtMapOutput struct{ *pulumi.OutputState }

func (CloudAccountNsxtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudAccountNsxt)(nil)).Elem()
}

func (o CloudAccountNsxtMapOutput) ToCloudAccountNsxtMapOutput() CloudAccountNsxtMapOutput {
	return o
}

func (o CloudAccountNsxtMapOutput) ToCloudAccountNsxtMapOutputWithContext(ctx context.Context) CloudAccountNsxtMapOutput {
	return o
}

func (o CloudAccountNsxtMapOutput) MapIndex(k pulumi.StringInput) CloudAccountNsxtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudAccountNsxt {
		return vs[0].(map[string]*CloudAccountNsxt)[vs[1].(string)]
	}).(CloudAccountNsxtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountNsxtInput)(nil)).Elem(), &CloudAccountNsxt{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountNsxtArrayInput)(nil)).Elem(), CloudAccountNsxtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudAccountNsxtMapInput)(nil)).Elem(), CloudAccountNsxtMap{})
	pulumi.RegisterOutputType(CloudAccountNsxtOutput{})
	pulumi.RegisterOutputType(CloudAccountNsxtArrayOutput{})
	pulumi.RegisterOutputType(CloudAccountNsxtMapOutput{})
}
