// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation VMC cloud account resource.
//
// ## Example Usage
// ### S
//
// **Create VMC cloud account:**
//
// The following example shows how to create a VMC cloud account resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/cloudaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/cloudaccount"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudaccount.NewVmc(ctx, "this", &cloudaccount.VmcArgs{
//				Description:          pulumi.String("tf test vmc cloud account"),
//				ApiToken:             pulumi.Any(_var.Api_token),
//				SddcName:             pulumi.Any(_var.Sddc_name),
//				VcenterHostname:      pulumi.Any(_var.Vcenter_hostname),
//				VcenterPassword:      pulumi.Any(_var.Vcenter_password),
//				VcenterUsername:      pulumi.Any(_var.Vcenter_username),
//				NsxHostname:          pulumi.Any(_var.Nsx_hostname),
//				DcId:                 pulumi.Any(_var.Data_collector_id),
//				Regions:              pulumi.Any(_var.Regions),
//				AcceptSelfSignedCert: pulumi.Bool(true),
//				Tags: cloudaccount.VmcTagArray{
//					&cloudaccount.VmcTagArgs{
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
// # To import the VMC cloud account, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:cloudaccount/vmc:Vmc new_vmc 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type Vmc struct {
	pulumi.CustomResourceState

	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrOutput `pulumi:"acceptSelfSignedCert"`
	// VMC API access key.
	ApiToken pulumi.StringOutput `pulumi:"apiToken"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
	DcId pulumi.StringPtrOutput `pulumi:"dcId"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// HATEOAS of entity.
	Links VmcLinkArrayOutput `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname pulumi.StringOutput `pulumi:"nsxHostname"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
	SddcName pulumi.StringOutput `pulumi:"sddcName"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags VmcTagArrayOutput `pulumi:"tags"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname pulumi.StringOutput `pulumi:"vcenterHostname"`
	// Password used to authenticate to the cloud Account.
	VcenterPassword pulumi.StringOutput `pulumi:"vcenterPassword"`
	// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername pulumi.StringOutput `pulumi:"vcenterUsername"`
}

// NewVmc registers a new resource with the given unique name, arguments, and options.
func NewVmc(ctx *pulumi.Context,
	name string, args *VmcArgs, opts ...pulumi.ResourceOption) (*Vmc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiToken == nil {
		return nil, errors.New("invalid value for required argument 'ApiToken'")
	}
	if args.NsxHostname == nil {
		return nil, errors.New("invalid value for required argument 'NsxHostname'")
	}
	if args.Regions == nil {
		return nil, errors.New("invalid value for required argument 'Regions'")
	}
	if args.SddcName == nil {
		return nil, errors.New("invalid value for required argument 'SddcName'")
	}
	if args.VcenterHostname == nil {
		return nil, errors.New("invalid value for required argument 'VcenterHostname'")
	}
	if args.VcenterPassword == nil {
		return nil, errors.New("invalid value for required argument 'VcenterPassword'")
	}
	if args.VcenterUsername == nil {
		return nil, errors.New("invalid value for required argument 'VcenterUsername'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Vmc
	err := ctx.RegisterResource("vra:cloudaccount/vmc:Vmc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmc gets an existing Vmc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmcState, opts ...pulumi.ResourceOption) (*Vmc, error) {
	var resource Vmc
	err := ctx.ReadResource("vra:cloudaccount/vmc:Vmc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vmc resources.
type vmcState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// VMC API access key.
	ApiToken *string `pulumi:"apiToken"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
	DcId *string `pulumi:"dcId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// HATEOAS of entity.
	Links []VmcLink `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname *string `pulumi:"nsxHostname"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner *string `pulumi:"owner"`
	// Set of region names enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
	SddcName *string `pulumi:"sddcName"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []VmcTag `pulumi:"tags"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
	// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname *string `pulumi:"vcenterHostname"`
	// Password used to authenticate to the cloud Account.
	VcenterPassword *string `pulumi:"vcenterPassword"`
	// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername *string `pulumi:"vcenterUsername"`
}

type VmcState struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// VMC API access key.
	ApiToken pulumi.StringPtrInput
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
	DcId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// HATEOAS of entity.
	Links VmcLinkArrayInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner pulumi.StringPtrInput
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
	SddcName pulumi.StringPtrInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags VmcTagArrayInput
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
	// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname pulumi.StringPtrInput
	// Password used to authenticate to the cloud Account.
	VcenterPassword pulumi.StringPtrInput
	// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername pulumi.StringPtrInput
}

func (VmcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmcState)(nil)).Elem()
}

type vmcArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert *bool `pulumi:"acceptSelfSignedCert"`
	// VMC API access key.
	ApiToken string `pulumi:"apiToken"`
	// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
	DcId *string `pulumi:"dcId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname string `pulumi:"nsxHostname"`
	// Set of region names enabled for the cloud account.
	Regions []string `pulumi:"regions"`
	// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
	SddcName string `pulumi:"sddcName"`
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags []VmcTag `pulumi:"tags"`
	// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname string `pulumi:"vcenterHostname"`
	// Password used to authenticate to the cloud Account.
	VcenterPassword string `pulumi:"vcenterPassword"`
	// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername string `pulumi:"vcenterUsername"`
}

// The set of arguments for constructing a Vmc resource.
type VmcArgs struct {
	// Accept self-signed certificate when connecting to the cloud account.
	AcceptSelfSignedCert pulumi.BoolPtrInput
	// VMC API access key.
	ApiToken pulumi.StringInput
	// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
	DcId pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// IP address of the NSX Manager server in the specified SDDC / FQDN.
	NsxHostname pulumi.StringInput
	// Set of region names enabled for the cloud account.
	Regions pulumi.StringArrayInput
	// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
	SddcName pulumi.StringInput
	// Set of tag keys and values to apply to the cloud account.\
	// Example:[ { "key" : "vmware", "value": "provider" } ]
	Tags VmcTagArrayInput
	// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	VcenterHostname pulumi.StringInput
	// Password used to authenticate to the cloud Account.
	VcenterPassword pulumi.StringInput
	// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	VcenterUsername pulumi.StringInput
}

func (VmcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmcArgs)(nil)).Elem()
}

type VmcInput interface {
	pulumi.Input

	ToVmcOutput() VmcOutput
	ToVmcOutputWithContext(ctx context.Context) VmcOutput
}

func (*Vmc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vmc)(nil)).Elem()
}

func (i *Vmc) ToVmcOutput() VmcOutput {
	return i.ToVmcOutputWithContext(context.Background())
}

func (i *Vmc) ToVmcOutputWithContext(ctx context.Context) VmcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmcOutput)
}

// VmcArrayInput is an input type that accepts VmcArray and VmcArrayOutput values.
// You can construct a concrete instance of `VmcArrayInput` via:
//
//	VmcArray{ VmcArgs{...} }
type VmcArrayInput interface {
	pulumi.Input

	ToVmcArrayOutput() VmcArrayOutput
	ToVmcArrayOutputWithContext(context.Context) VmcArrayOutput
}

type VmcArray []VmcInput

func (VmcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vmc)(nil)).Elem()
}

func (i VmcArray) ToVmcArrayOutput() VmcArrayOutput {
	return i.ToVmcArrayOutputWithContext(context.Background())
}

func (i VmcArray) ToVmcArrayOutputWithContext(ctx context.Context) VmcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmcArrayOutput)
}

// VmcMapInput is an input type that accepts VmcMap and VmcMapOutput values.
// You can construct a concrete instance of `VmcMapInput` via:
//
//	VmcMap{ "key": VmcArgs{...} }
type VmcMapInput interface {
	pulumi.Input

	ToVmcMapOutput() VmcMapOutput
	ToVmcMapOutputWithContext(context.Context) VmcMapOutput
}

type VmcMap map[string]VmcInput

func (VmcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vmc)(nil)).Elem()
}

func (i VmcMap) ToVmcMapOutput() VmcMapOutput {
	return i.ToVmcMapOutputWithContext(context.Background())
}

func (i VmcMap) ToVmcMapOutputWithContext(ctx context.Context) VmcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmcMapOutput)
}

type VmcOutput struct{ *pulumi.OutputState }

func (VmcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vmc)(nil)).Elem()
}

func (o VmcOutput) ToVmcOutput() VmcOutput {
	return o
}

func (o VmcOutput) ToVmcOutputWithContext(ctx context.Context) VmcOutput {
	return o
}

// Accept self-signed certificate when connecting to the cloud account.
func (o VmcOutput) AcceptSelfSignedCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vmc) pulumi.BoolPtrOutput { return v.AcceptSelfSignedCert }).(pulumi.BoolPtrOutput)
}

// VMC API access key.
func (o VmcOutput) ApiToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.ApiToken }).(pulumi.StringOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o VmcOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.
func (o VmcOutput) DcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringPtrOutput { return v.DcId }).(pulumi.StringPtrOutput)
}

// Human-friendly description.
func (o VmcOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// HATEOAS of entity.
func (o VmcOutput) Links() VmcLinkArrayOutput {
	return o.ApplyT(func(v *Vmc) VmcLinkArrayOutput { return v.Links }).(VmcLinkArrayOutput)
}

// Human-friendly name used as an identifier in APIs that support this option.
func (o VmcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IP address of the NSX Manager server in the specified SDDC / FQDN.
func (o VmcOutput) NsxHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.NsxHostname }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o VmcOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o VmcOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Set of region names enabled for the cloud account.
func (o VmcOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.
func (o VmcOutput) SddcName() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.SddcName }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the cloud account.\
// Example:[ { "key" : "vmware", "value": "provider" } ]
func (o VmcOutput) Tags() VmcTagArrayOutput {
	return o.ApplyT(func(v *Vmc) VmcTagArrayOutput { return v.Tags }).(VmcTagArrayOutput)
}

// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
func (o VmcOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
func (o VmcOutput) VcenterHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.VcenterHostname }).(pulumi.StringOutput)
}

// Password used to authenticate to the cloud Account.
func (o VmcOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.VcenterPassword }).(pulumi.StringOutput)
}

// vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
func (o VmcOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *Vmc) pulumi.StringOutput { return v.VcenterUsername }).(pulumi.StringOutput)
}

type VmcArrayOutput struct{ *pulumi.OutputState }

func (VmcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vmc)(nil)).Elem()
}

func (o VmcArrayOutput) ToVmcArrayOutput() VmcArrayOutput {
	return o
}

func (o VmcArrayOutput) ToVmcArrayOutputWithContext(ctx context.Context) VmcArrayOutput {
	return o
}

func (o VmcArrayOutput) Index(i pulumi.IntInput) VmcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vmc {
		return vs[0].([]*Vmc)[vs[1].(int)]
	}).(VmcOutput)
}

type VmcMapOutput struct{ *pulumi.OutputState }

func (VmcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vmc)(nil)).Elem()
}

func (o VmcMapOutput) ToVmcMapOutput() VmcMapOutput {
	return o
}

func (o VmcMapOutput) ToVmcMapOutputWithContext(ctx context.Context) VmcMapOutput {
	return o
}

func (o VmcMapOutput) MapIndex(k pulumi.StringInput) VmcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vmc {
		return vs[0].(map[string]*Vmc)[vs[1].(string)]
	}).(VmcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmcInput)(nil)).Elem(), &Vmc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmcArrayInput)(nil)).Elem(), VmcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmcMapInput)(nil)).Elem(), VmcMap{})
	pulumi.RegisterOutputType(VmcOutput{})
	pulumi.RegisterOutputType(VmcArrayOutput{})
	pulumi.RegisterOutputType(VmcMapOutput{})
}
