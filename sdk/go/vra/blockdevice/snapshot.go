// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockdevice

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation block device snapshot resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create a block device snapshot resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/blockdevice"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blockdevice.NewSnapshot(ctx, "snapshot1", &blockdevice.SnapshotArgs{
//				BlockDeviceId: pulumi.Any(_var.Block_device_id),
//				Description:   pulumi.String("terraform fcd snapshot"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// ID of block device.
	BlockDeviceId pulumi.StringOutput `pulumi:"blockDeviceId"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Human-friendly description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether snapshot on block device is current.
	IsCurrent pulumi.BoolOutput `pulumi:"isCurrent"`
	// HATEOAS of entity
	Links SnapshotLinkArrayOutput `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of entity owner.
	Owner     pulumi.StringOutput `pulumi:"owner"`
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlockDeviceId == nil {
		return nil, errors.New("invalid value for required argument 'BlockDeviceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("vra:blockdevice/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("vra:blockdevice/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// ID of block device.
	BlockDeviceId *string `pulumi:"blockDeviceId"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
	// Indicates whether snapshot on block device is current.
	IsCurrent *bool `pulumi:"isCurrent"`
	// HATEOAS of entity
	Links []SnapshotLink `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of entity owner.
	Owner     *string `pulumi:"owner"`
	UpdatedAt *string `pulumi:"updatedAt"`
}

type SnapshotState struct {
	// ID of block device.
	BlockDeviceId pulumi.StringPtrInput
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
	// Indicates whether snapshot on block device is current.
	IsCurrent pulumi.BoolPtrInput
	// HATEOAS of entity
	Links SnapshotLinkArrayInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// Email of entity owner.
	Owner     pulumi.StringPtrInput
	UpdatedAt pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// ID of block device.
	BlockDeviceId string `pulumi:"blockDeviceId"`
	// Human-friendly description.
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// ID of block device.
	BlockDeviceId pulumi.StringInput
	// Human-friendly description.
	Description pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// ID of block device.
func (o SnapshotOutput) BlockDeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.BlockDeviceId }).(pulumi.StringOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o SnapshotOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Human-friendly description.
func (o SnapshotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether snapshot on block device is current.
func (o SnapshotOutput) IsCurrent() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.IsCurrent }).(pulumi.BoolOutput)
}

// HATEOAS of entity
func (o SnapshotOutput) Links() SnapshotLinkArrayOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotLinkArrayOutput { return v.Links }).(SnapshotLinkArrayOutput)
}

// Human-friendly name used as an identifier in APIs that support this option.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o SnapshotOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of entity owner.
func (o SnapshotOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

func (o SnapshotOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
