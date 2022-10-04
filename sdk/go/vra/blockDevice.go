// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation block device resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create a block device resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.NewBlockDevice(ctx, "disk1", &vra.BlockDeviceArgs{
//				CapacityInGb: pulumi.Int(10),
//				ProjectId:    pulumi.Any(_var.Project_id),
//				Persistent:   pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type BlockDevice struct {
	pulumi.CustomResourceState

	// Capacity of block device in GB.
	CapacityInGb pulumi.IntOutput `pulumi:"capacityInGb"`
	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	Constraints BlockDeviceConstraintArrayOutput `pulumi:"constraints"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Additional custom properties that may be used to extend the machine.
	CustomProperties pulumi.MapOutput `pulumi:"customProperties"`
	// ID of deployment associated with resource.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Content of a disk, base64 encoded.
	DiskContentBase64 pulumi.StringPtrOutput `pulumi:"diskContentBase64"`
	// Indicates whether block device should be encrypted or not.
	Encrypted pulumi.BoolPtrOutput `pulumi:"encrypted"`
	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	ExpandSnapshots pulumi.BoolPtrOutput `pulumi:"expandSnapshots"`
	// External entity ID on provider side.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// External regionId of resource.
	ExternalRegionId pulumi.StringOutput `pulumi:"externalRegionId"`
	// External zoneId of resource.
	ExternalZoneId pulumi.StringOutput `pulumi:"externalZoneId"`
	// HATEOAS of the entity
	Links BlockDeviceLinkArrayOutput `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that block device snapshot belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Email of block device snapshot owner.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Indicates whether block device survives a delete action.
	Persistent pulumi.BoolPtrOutput `pulumi:"persistent"`
	// ID of project that current user belongs to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	Purge pulumi.BoolPtrOutput `pulumi:"purge"`
	// Represents a machine snapshot.
	Snapshots BlockDeviceSnapshotTypeArrayOutput `pulumi:"snapshots"`
	// URI to use for block device. Example: ami-0d4cfd66
	SourceReference pulumi.StringPtrOutput `pulumi:"sourceReference"`
	// Status of block device.
	Status pulumi.StringOutput `pulumi:"status"`
	// Set of tag keys and values to apply to the resource instance.\
	// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
	Tags BlockDeviceTagArrayOutput `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewBlockDevice registers a new resource with the given unique name, arguments, and options.
func NewBlockDevice(ctx *pulumi.Context,
	name string, args *BlockDeviceArgs, opts ...pulumi.ResourceOption) (*BlockDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityInGb == nil {
		return nil, errors.New("invalid value for required argument 'CapacityInGb'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BlockDevice
	err := ctx.RegisterResource("vra:index/blockDevice:BlockDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockDevice gets an existing BlockDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockDeviceState, opts ...pulumi.ResourceOption) (*BlockDevice, error) {
	var resource BlockDevice
	err := ctx.ReadResource("vra:index/blockDevice:BlockDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockDevice resources.
type blockDeviceState struct {
	// Capacity of block device in GB.
	CapacityInGb *int `pulumi:"capacityInGb"`
	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	Constraints []BlockDeviceConstraint `pulumi:"constraints"`
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// Additional custom properties that may be used to extend the machine.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// ID of deployment associated with resource.
	DeploymentId *string `pulumi:"deploymentId"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description *string `pulumi:"description"`
	// Content of a disk, base64 encoded.
	DiskContentBase64 *string `pulumi:"diskContentBase64"`
	// Indicates whether block device should be encrypted or not.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	ExpandSnapshots *bool `pulumi:"expandSnapshots"`
	// External entity ID on provider side.
	ExternalId *string `pulumi:"externalId"`
	// External regionId of resource.
	ExternalRegionId *string `pulumi:"externalRegionId"`
	// External zoneId of resource.
	ExternalZoneId *string `pulumi:"externalZoneId"`
	// HATEOAS of the entity
	Links []BlockDeviceLink `pulumi:"links"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// ID of organization that block device snapshot belongs to.
	OrgId *string `pulumi:"orgId"`
	// Email of block device snapshot owner.
	Owner *string `pulumi:"owner"`
	// Indicates whether block device survives a delete action.
	Persistent *bool `pulumi:"persistent"`
	// ID of project that current user belongs to.
	ProjectId *string `pulumi:"projectId"`
	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	Purge *bool `pulumi:"purge"`
	// Represents a machine snapshot.
	Snapshots []BlockDeviceSnapshotType `pulumi:"snapshots"`
	// URI to use for block device. Example: ami-0d4cfd66
	SourceReference *string `pulumi:"sourceReference"`
	// Status of block device.
	Status *string `pulumi:"status"`
	// Set of tag keys and values to apply to the resource instance.\
	// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
	Tags []BlockDeviceTag `pulumi:"tags"`
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type BlockDeviceState struct {
	// Capacity of block device in GB.
	CapacityInGb pulumi.IntPtrInput
	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	Constraints BlockDeviceConstraintArrayInput
	// Date when entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// Additional custom properties that may be used to extend the machine.
	CustomProperties pulumi.MapInput
	// ID of deployment associated with resource.
	DeploymentId pulumi.StringPtrInput
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description pulumi.StringPtrInput
	// Content of a disk, base64 encoded.
	DiskContentBase64 pulumi.StringPtrInput
	// Indicates whether block device should be encrypted or not.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	ExpandSnapshots pulumi.BoolPtrInput
	// External entity ID on provider side.
	ExternalId pulumi.StringPtrInput
	// External regionId of resource.
	ExternalRegionId pulumi.StringPtrInput
	// External zoneId of resource.
	ExternalZoneId pulumi.StringPtrInput
	// HATEOAS of the entity
	Links BlockDeviceLinkArrayInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// ID of organization that block device snapshot belongs to.
	OrgId pulumi.StringPtrInput
	// Email of block device snapshot owner.
	Owner pulumi.StringPtrInput
	// Indicates whether block device survives a delete action.
	Persistent pulumi.BoolPtrInput
	// ID of project that current user belongs to.
	ProjectId pulumi.StringPtrInput
	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	Purge pulumi.BoolPtrInput
	// Represents a machine snapshot.
	Snapshots BlockDeviceSnapshotTypeArrayInput
	// URI to use for block device. Example: ami-0d4cfd66
	SourceReference pulumi.StringPtrInput
	// Status of block device.
	Status pulumi.StringPtrInput
	// Set of tag keys and values to apply to the resource instance.\
	// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
	Tags BlockDeviceTagArrayInput
	// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
}

func (BlockDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockDeviceState)(nil)).Elem()
}

type blockDeviceArgs struct {
	// Capacity of block device in GB.
	CapacityInGb int `pulumi:"capacityInGb"`
	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	Constraints []BlockDeviceConstraint `pulumi:"constraints"`
	// Additional custom properties that may be used to extend the machine.
	CustomProperties map[string]interface{} `pulumi:"customProperties"`
	// ID of deployment associated with resource.
	DeploymentId *string `pulumi:"deploymentId"`
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description *string `pulumi:"description"`
	// Content of a disk, base64 encoded.
	DiskContentBase64 *string `pulumi:"diskContentBase64"`
	// Indicates whether block device should be encrypted or not.
	Encrypted *bool `pulumi:"encrypted"`
	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	ExpandSnapshots *bool `pulumi:"expandSnapshots"`
	// Human-friendly name used as an identifier in APIs that support this option.
	Name *string `pulumi:"name"`
	// Indicates whether block device survives a delete action.
	Persistent *bool `pulumi:"persistent"`
	// ID of project that current user belongs to.
	ProjectId string `pulumi:"projectId"`
	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	Purge *bool `pulumi:"purge"`
	// URI to use for block device. Example: ami-0d4cfd66
	SourceReference *string `pulumi:"sourceReference"`
	// Set of tag keys and values to apply to the resource instance.\
	// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
	Tags []BlockDeviceTag `pulumi:"tags"`
}

// The set of arguments for constructing a BlockDevice resource.
type BlockDeviceArgs struct {
	// Capacity of block device in GB.
	CapacityInGb pulumi.IntInput
	// Storage, network, and extensibility constraints to be applied when provisioning through the project.
	Constraints BlockDeviceConstraintArrayInput
	// Additional custom properties that may be used to extend the machine.
	CustomProperties pulumi.MapInput
	// ID of deployment associated with resource.
	DeploymentId pulumi.StringPtrInput
	// Describes machine within the scope of your organization and is not propagated to the cloud.
	Description pulumi.StringPtrInput
	// Content of a disk, base64 encoded.
	DiskContentBase64 pulumi.StringPtrInput
	// Indicates whether block device should be encrypted or not.
	Encrypted pulumi.BoolPtrInput
	// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
	ExpandSnapshots pulumi.BoolPtrInput
	// Human-friendly name used as an identifier in APIs that support this option.
	Name pulumi.StringPtrInput
	// Indicates whether block device survives a delete action.
	Persistent pulumi.BoolPtrInput
	// ID of project that current user belongs to.
	ProjectId pulumi.StringInput
	// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
	Purge pulumi.BoolPtrInput
	// URI to use for block device. Example: ami-0d4cfd66
	SourceReference pulumi.StringPtrInput
	// Set of tag keys and values to apply to the resource instance.\
	// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
	Tags BlockDeviceTagArrayInput
}

func (BlockDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockDeviceArgs)(nil)).Elem()
}

type BlockDeviceInput interface {
	pulumi.Input

	ToBlockDeviceOutput() BlockDeviceOutput
	ToBlockDeviceOutputWithContext(ctx context.Context) BlockDeviceOutput
}

func (*BlockDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockDevice)(nil)).Elem()
}

func (i *BlockDevice) ToBlockDeviceOutput() BlockDeviceOutput {
	return i.ToBlockDeviceOutputWithContext(context.Background())
}

func (i *BlockDevice) ToBlockDeviceOutputWithContext(ctx context.Context) BlockDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockDeviceOutput)
}

// BlockDeviceArrayInput is an input type that accepts BlockDeviceArray and BlockDeviceArrayOutput values.
// You can construct a concrete instance of `BlockDeviceArrayInput` via:
//
//	BlockDeviceArray{ BlockDeviceArgs{...} }
type BlockDeviceArrayInput interface {
	pulumi.Input

	ToBlockDeviceArrayOutput() BlockDeviceArrayOutput
	ToBlockDeviceArrayOutputWithContext(context.Context) BlockDeviceArrayOutput
}

type BlockDeviceArray []BlockDeviceInput

func (BlockDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockDevice)(nil)).Elem()
}

func (i BlockDeviceArray) ToBlockDeviceArrayOutput() BlockDeviceArrayOutput {
	return i.ToBlockDeviceArrayOutputWithContext(context.Background())
}

func (i BlockDeviceArray) ToBlockDeviceArrayOutputWithContext(ctx context.Context) BlockDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockDeviceArrayOutput)
}

// BlockDeviceMapInput is an input type that accepts BlockDeviceMap and BlockDeviceMapOutput values.
// You can construct a concrete instance of `BlockDeviceMapInput` via:
//
//	BlockDeviceMap{ "key": BlockDeviceArgs{...} }
type BlockDeviceMapInput interface {
	pulumi.Input

	ToBlockDeviceMapOutput() BlockDeviceMapOutput
	ToBlockDeviceMapOutputWithContext(context.Context) BlockDeviceMapOutput
}

type BlockDeviceMap map[string]BlockDeviceInput

func (BlockDeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockDevice)(nil)).Elem()
}

func (i BlockDeviceMap) ToBlockDeviceMapOutput() BlockDeviceMapOutput {
	return i.ToBlockDeviceMapOutputWithContext(context.Background())
}

func (i BlockDeviceMap) ToBlockDeviceMapOutputWithContext(ctx context.Context) BlockDeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockDeviceMapOutput)
}

type BlockDeviceOutput struct{ *pulumi.OutputState }

func (BlockDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockDevice)(nil)).Elem()
}

func (o BlockDeviceOutput) ToBlockDeviceOutput() BlockDeviceOutput {
	return o
}

func (o BlockDeviceOutput) ToBlockDeviceOutputWithContext(ctx context.Context) BlockDeviceOutput {
	return o
}

// Capacity of block device in GB.
func (o BlockDeviceOutput) CapacityInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.IntOutput { return v.CapacityInGb }).(pulumi.IntOutput)
}

// Storage, network, and extensibility constraints to be applied when provisioning through the project.
func (o BlockDeviceOutput) Constraints() BlockDeviceConstraintArrayOutput {
	return o.ApplyT(func(v *BlockDevice) BlockDeviceConstraintArrayOutput { return v.Constraints }).(BlockDeviceConstraintArrayOutput)
}

// Date when entity was created. Date and time format is ISO 8601 and UTC.
func (o BlockDeviceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Additional custom properties that may be used to extend the machine.
func (o BlockDeviceOutput) CustomProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.MapOutput { return v.CustomProperties }).(pulumi.MapOutput)
}

// ID of deployment associated with resource.
func (o BlockDeviceOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// Describes machine within the scope of your organization and is not propagated to the cloud.
func (o BlockDeviceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Content of a disk, base64 encoded.
func (o BlockDeviceOutput) DiskContentBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringPtrOutput { return v.DiskContentBase64 }).(pulumi.StringPtrOutput)
}

// Indicates whether block device should be encrypted or not.
func (o BlockDeviceOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.BoolPtrOutput { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

// Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.
func (o BlockDeviceOutput) ExpandSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.BoolPtrOutput { return v.ExpandSnapshots }).(pulumi.BoolPtrOutput)
}

// External entity ID on provider side.
func (o BlockDeviceOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// External regionId of resource.
func (o BlockDeviceOutput) ExternalRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.ExternalRegionId }).(pulumi.StringOutput)
}

// External zoneId of resource.
func (o BlockDeviceOutput) ExternalZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.ExternalZoneId }).(pulumi.StringOutput)
}

// HATEOAS of the entity
func (o BlockDeviceOutput) Links() BlockDeviceLinkArrayOutput {
	return o.ApplyT(func(v *BlockDevice) BlockDeviceLinkArrayOutput { return v.Links }).(BlockDeviceLinkArrayOutput)
}

// Human-friendly name used as an identifier in APIs that support this option.
func (o BlockDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that block device snapshot belongs to.
func (o BlockDeviceOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Email of block device snapshot owner.
func (o BlockDeviceOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Indicates whether block device survives a delete action.
func (o BlockDeviceOutput) Persistent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.BoolPtrOutput { return v.Persistent }).(pulumi.BoolPtrOutput)
}

// ID of project that current user belongs to.
func (o BlockDeviceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.
func (o BlockDeviceOutput) Purge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.BoolPtrOutput { return v.Purge }).(pulumi.BoolPtrOutput)
}

// Represents a machine snapshot.
func (o BlockDeviceOutput) Snapshots() BlockDeviceSnapshotTypeArrayOutput {
	return o.ApplyT(func(v *BlockDevice) BlockDeviceSnapshotTypeArrayOutput { return v.Snapshots }).(BlockDeviceSnapshotTypeArrayOutput)
}

// URI to use for block device. Example: ami-0d4cfd66
func (o BlockDeviceOutput) SourceReference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringPtrOutput { return v.SourceReference }).(pulumi.StringPtrOutput)
}

// Status of block device.
func (o BlockDeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Set of tag keys and values to apply to the resource instance.\
// Example:[ { "key" : "vmware.enumeration.type", "value": "nebsBlock" } ]
func (o BlockDeviceOutput) Tags() BlockDeviceTagArrayOutput {
	return o.ApplyT(func(v *BlockDevice) BlockDeviceTagArrayOutput { return v.Tags }).(BlockDeviceTagArrayOutput)
}

// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
func (o BlockDeviceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockDevice) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type BlockDeviceArrayOutput struct{ *pulumi.OutputState }

func (BlockDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockDevice)(nil)).Elem()
}

func (o BlockDeviceArrayOutput) ToBlockDeviceArrayOutput() BlockDeviceArrayOutput {
	return o
}

func (o BlockDeviceArrayOutput) ToBlockDeviceArrayOutputWithContext(ctx context.Context) BlockDeviceArrayOutput {
	return o
}

func (o BlockDeviceArrayOutput) Index(i pulumi.IntInput) BlockDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlockDevice {
		return vs[0].([]*BlockDevice)[vs[1].(int)]
	}).(BlockDeviceOutput)
}

type BlockDeviceMapOutput struct{ *pulumi.OutputState }

func (BlockDeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockDevice)(nil)).Elem()
}

func (o BlockDeviceMapOutput) ToBlockDeviceMapOutput() BlockDeviceMapOutput {
	return o
}

func (o BlockDeviceMapOutput) ToBlockDeviceMapOutputWithContext(ctx context.Context) BlockDeviceMapOutput {
	return o
}

func (o BlockDeviceMapOutput) MapIndex(k pulumi.StringInput) BlockDeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlockDevice {
		return vs[0].(map[string]*BlockDevice)[vs[1].(string)]
	}).(BlockDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockDeviceInput)(nil)).Elem(), &BlockDevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockDeviceArrayInput)(nil)).Elem(), BlockDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockDeviceMapInput)(nil)).Elem(), BlockDeviceMap{})
	pulumi.RegisterOutputType(BlockDeviceOutput{})
	pulumi.RegisterOutputType(BlockDeviceArrayOutput{})
	pulumi.RegisterOutputType(BlockDeviceMapOutput{})
}
