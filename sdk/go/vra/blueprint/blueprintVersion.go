// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a VMware vRealize Automation cloud template (blueprint) version resource.
//
// ## Example Usage
// ### S
//
// The following example shows how to create a cloud template (blueprint) version resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/blueprint"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := blueprint.NewBlueprintVersion(ctx, "this", &blueprint.BlueprintVersionArgs{
//				BlueprintId: pulumi.Any(_var.Vra_blueprint_id),
//				ChangeLog:   pulumi.String("First version"),
//				Description: pulumi.String("Released from vRA terraform provider"),
//				Release:     pulumi.Bool(true),
//				Version:     random_integer.Suffix.Result / random_integer.Suffix.Result,
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
// # To import the cloud template (blueprint) version, use the ID as in the following example
//
// ```sh
//
//	$ pulumi import vra:blueprint/blueprintVersion:BlueprintVersion this 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type BlueprintVersion struct {
	pulumi.CustomResourceState

	// Description of cloud template (blueprint).
	BlueprintDescription pulumi.StringOutput `pulumi:"blueprintDescription"`
	// ID of the cloud template  (blueprint).
	BlueprintId pulumi.StringOutput `pulumi:"blueprintId"`
	// Cloud template  (blueprint) version log.
	ChangeLog pulumi.StringPtrOutput `pulumi:"changeLog"`
	// Blueprint YAML content.
	Content pulumi.StringOutput `pulumi:"content"`
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// User who created the entity.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Human-friendly description for the cloud template  (blueprint) version.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of cloud template (blueprint) version.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of project that entity belongs to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Name of project that entity belongs to.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// Flag to indicate whether to release the version.
	Release pulumi.BoolPtrOutput `pulumi:"release"`
	// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// User who last updated the entity.
	UpdatedBy pulumi.StringOutput `pulumi:"updatedBy"`
	// Flag to indicate if the current content of the cloud template (blueprint) is valid.
	Valid pulumi.StringOutput `pulumi:"valid"`
	// Cloud template  (blueprint) version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewBlueprintVersion registers a new resource with the given unique name, arguments, and options.
func NewBlueprintVersion(ctx *pulumi.Context,
	name string, args *BlueprintVersionArgs, opts ...pulumi.ResourceOption) (*BlueprintVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintId == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BlueprintVersion
	err := ctx.RegisterResource("vra:blueprint/blueprintVersion:BlueprintVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlueprintVersion gets an existing BlueprintVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlueprintVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlueprintVersionState, opts ...pulumi.ResourceOption) (*BlueprintVersion, error) {
	var resource BlueprintVersion
	err := ctx.ReadResource("vra:blueprint/blueprintVersion:BlueprintVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlueprintVersion resources.
type blueprintVersionState struct {
	// Description of cloud template (blueprint).
	BlueprintDescription *string `pulumi:"blueprintDescription"`
	// ID of the cloud template  (blueprint).
	BlueprintId *string `pulumi:"blueprintId"`
	// Cloud template  (blueprint) version log.
	ChangeLog *string `pulumi:"changeLog"`
	// Blueprint YAML content.
	Content *string `pulumi:"content"`
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt *string `pulumi:"createdAt"`
	// User who created the entity.
	CreatedBy *string `pulumi:"createdBy"`
	// Human-friendly description for the cloud template  (blueprint) version.
	Description *string `pulumi:"description"`
	// Name of cloud template (blueprint) version.
	Name *string `pulumi:"name"`
	// ID of organization that entity belongs to.
	OrgId *string `pulumi:"orgId"`
	// ID of project that entity belongs to.
	ProjectId *string `pulumi:"projectId"`
	// Name of project that entity belongs to.
	ProjectName *string `pulumi:"projectName"`
	// Flag to indicate whether to release the version.
	Release *bool `pulumi:"release"`
	// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
	Status *string `pulumi:"status"`
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt *string `pulumi:"updatedAt"`
	// User who last updated the entity.
	UpdatedBy *string `pulumi:"updatedBy"`
	// Flag to indicate if the current content of the cloud template (blueprint) is valid.
	Valid *string `pulumi:"valid"`
	// Cloud template  (blueprint) version.
	Version *string `pulumi:"version"`
}

type BlueprintVersionState struct {
	// Description of cloud template (blueprint).
	BlueprintDescription pulumi.StringPtrInput
	// ID of the cloud template  (blueprint).
	BlueprintId pulumi.StringPtrInput
	// Cloud template  (blueprint) version log.
	ChangeLog pulumi.StringPtrInput
	// Blueprint YAML content.
	Content pulumi.StringPtrInput
	// Date when the entity was created. Date and time format is ISO 8601 and UTC.
	CreatedAt pulumi.StringPtrInput
	// User who created the entity.
	CreatedBy pulumi.StringPtrInput
	// Human-friendly description for the cloud template  (blueprint) version.
	Description pulumi.StringPtrInput
	// Name of cloud template (blueprint) version.
	Name pulumi.StringPtrInput
	// ID of organization that entity belongs to.
	OrgId pulumi.StringPtrInput
	// ID of project that entity belongs to.
	ProjectId pulumi.StringPtrInput
	// Name of project that entity belongs to.
	ProjectName pulumi.StringPtrInput
	// Flag to indicate whether to release the version.
	Release pulumi.BoolPtrInput
	// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
	Status pulumi.StringPtrInput
	// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
	UpdatedAt pulumi.StringPtrInput
	// User who last updated the entity.
	UpdatedBy pulumi.StringPtrInput
	// Flag to indicate if the current content of the cloud template (blueprint) is valid.
	Valid pulumi.StringPtrInput
	// Cloud template  (blueprint) version.
	Version pulumi.StringPtrInput
}

func (BlueprintVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintVersionState)(nil)).Elem()
}

type blueprintVersionArgs struct {
	// ID of the cloud template  (blueprint).
	BlueprintId string `pulumi:"blueprintId"`
	// Cloud template  (blueprint) version log.
	ChangeLog *string `pulumi:"changeLog"`
	// Human-friendly description for the cloud template  (blueprint) version.
	Description *string `pulumi:"description"`
	// Flag to indicate whether to release the version.
	Release *bool `pulumi:"release"`
	// Cloud template  (blueprint) version.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a BlueprintVersion resource.
type BlueprintVersionArgs struct {
	// ID of the cloud template  (blueprint).
	BlueprintId pulumi.StringInput
	// Cloud template  (blueprint) version log.
	ChangeLog pulumi.StringPtrInput
	// Human-friendly description for the cloud template  (blueprint) version.
	Description pulumi.StringPtrInput
	// Flag to indicate whether to release the version.
	Release pulumi.BoolPtrInput
	// Cloud template  (blueprint) version.
	Version pulumi.StringInput
}

func (BlueprintVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blueprintVersionArgs)(nil)).Elem()
}

type BlueprintVersionInput interface {
	pulumi.Input

	ToBlueprintVersionOutput() BlueprintVersionOutput
	ToBlueprintVersionOutputWithContext(ctx context.Context) BlueprintVersionOutput
}

func (*BlueprintVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**BlueprintVersion)(nil)).Elem()
}

func (i *BlueprintVersion) ToBlueprintVersionOutput() BlueprintVersionOutput {
	return i.ToBlueprintVersionOutputWithContext(context.Background())
}

func (i *BlueprintVersion) ToBlueprintVersionOutputWithContext(ctx context.Context) BlueprintVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintVersionOutput)
}

// BlueprintVersionArrayInput is an input type that accepts BlueprintVersionArray and BlueprintVersionArrayOutput values.
// You can construct a concrete instance of `BlueprintVersionArrayInput` via:
//
//	BlueprintVersionArray{ BlueprintVersionArgs{...} }
type BlueprintVersionArrayInput interface {
	pulumi.Input

	ToBlueprintVersionArrayOutput() BlueprintVersionArrayOutput
	ToBlueprintVersionArrayOutputWithContext(context.Context) BlueprintVersionArrayOutput
}

type BlueprintVersionArray []BlueprintVersionInput

func (BlueprintVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlueprintVersion)(nil)).Elem()
}

func (i BlueprintVersionArray) ToBlueprintVersionArrayOutput() BlueprintVersionArrayOutput {
	return i.ToBlueprintVersionArrayOutputWithContext(context.Background())
}

func (i BlueprintVersionArray) ToBlueprintVersionArrayOutputWithContext(ctx context.Context) BlueprintVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintVersionArrayOutput)
}

// BlueprintVersionMapInput is an input type that accepts BlueprintVersionMap and BlueprintVersionMapOutput values.
// You can construct a concrete instance of `BlueprintVersionMapInput` via:
//
//	BlueprintVersionMap{ "key": BlueprintVersionArgs{...} }
type BlueprintVersionMapInput interface {
	pulumi.Input

	ToBlueprintVersionMapOutput() BlueprintVersionMapOutput
	ToBlueprintVersionMapOutputWithContext(context.Context) BlueprintVersionMapOutput
}

type BlueprintVersionMap map[string]BlueprintVersionInput

func (BlueprintVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlueprintVersion)(nil)).Elem()
}

func (i BlueprintVersionMap) ToBlueprintVersionMapOutput() BlueprintVersionMapOutput {
	return i.ToBlueprintVersionMapOutputWithContext(context.Background())
}

func (i BlueprintVersionMap) ToBlueprintVersionMapOutputWithContext(ctx context.Context) BlueprintVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlueprintVersionMapOutput)
}

type BlueprintVersionOutput struct{ *pulumi.OutputState }

func (BlueprintVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlueprintVersion)(nil)).Elem()
}

func (o BlueprintVersionOutput) ToBlueprintVersionOutput() BlueprintVersionOutput {
	return o
}

func (o BlueprintVersionOutput) ToBlueprintVersionOutputWithContext(ctx context.Context) BlueprintVersionOutput {
	return o
}

// Description of cloud template (blueprint).
func (o BlueprintVersionOutput) BlueprintDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.BlueprintDescription }).(pulumi.StringOutput)
}

// ID of the cloud template  (blueprint).
func (o BlueprintVersionOutput) BlueprintId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.BlueprintId }).(pulumi.StringOutput)
}

// Cloud template  (blueprint) version log.
func (o BlueprintVersionOutput) ChangeLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringPtrOutput { return v.ChangeLog }).(pulumi.StringPtrOutput)
}

// Blueprint YAML content.
func (o BlueprintVersionOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Date when the entity was created. Date and time format is ISO 8601 and UTC.
func (o BlueprintVersionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// User who created the entity.
func (o BlueprintVersionOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// Human-friendly description for the cloud template  (blueprint) version.
func (o BlueprintVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of cloud template (blueprint) version.
func (o BlueprintVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of organization that entity belongs to.
func (o BlueprintVersionOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of project that entity belongs to.
func (o BlueprintVersionOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Name of project that entity belongs to.
func (o BlueprintVersionOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// Flag to indicate whether to release the version.
func (o BlueprintVersionOutput) Release() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.BoolPtrOutput { return v.Release }).(pulumi.BoolPtrOutput)
}

// Status of the cloud template (blueprint). Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
func (o BlueprintVersionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
func (o BlueprintVersionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// User who last updated the entity.
func (o BlueprintVersionOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.UpdatedBy }).(pulumi.StringOutput)
}

// Flag to indicate if the current content of the cloud template (blueprint) is valid.
func (o BlueprintVersionOutput) Valid() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.Valid }).(pulumi.StringOutput)
}

// Cloud template  (blueprint) version.
func (o BlueprintVersionOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *BlueprintVersion) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type BlueprintVersionArrayOutput struct{ *pulumi.OutputState }

func (BlueprintVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlueprintVersion)(nil)).Elem()
}

func (o BlueprintVersionArrayOutput) ToBlueprintVersionArrayOutput() BlueprintVersionArrayOutput {
	return o
}

func (o BlueprintVersionArrayOutput) ToBlueprintVersionArrayOutputWithContext(ctx context.Context) BlueprintVersionArrayOutput {
	return o
}

func (o BlueprintVersionArrayOutput) Index(i pulumi.IntInput) BlueprintVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlueprintVersion {
		return vs[0].([]*BlueprintVersion)[vs[1].(int)]
	}).(BlueprintVersionOutput)
}

type BlueprintVersionMapOutput struct{ *pulumi.OutputState }

func (BlueprintVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlueprintVersion)(nil)).Elem()
}

func (o BlueprintVersionMapOutput) ToBlueprintVersionMapOutput() BlueprintVersionMapOutput {
	return o
}

func (o BlueprintVersionMapOutput) ToBlueprintVersionMapOutputWithContext(ctx context.Context) BlueprintVersionMapOutput {
	return o
}

func (o BlueprintVersionMapOutput) MapIndex(k pulumi.StringInput) BlueprintVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlueprintVersion {
		return vs[0].(map[string]*BlueprintVersion)[vs[1].(string)]
	}).(BlueprintVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintVersionInput)(nil)).Elem(), &BlueprintVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintVersionArrayInput)(nil)).Elem(), BlueprintVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlueprintVersionMapInput)(nil)).Elem(), BlueprintVersionMap{})
	pulumi.RegisterOutputType(BlueprintVersionOutput{})
	pulumi.RegisterOutputType(BlueprintVersionArrayOutput{})
	pulumi.RegisterOutputType(BlueprintVersionMapOutput{})
}
