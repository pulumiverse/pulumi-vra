// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides a way to create a catalog item entitlement in VMware vRealize Automation.
//
// ## Example Usage
// ### S
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/schmidtw/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vra.NewCatalogItemEntitlement(ctx, "this", &vra.CatalogItemEntitlementArgs{
//				CatalogItemId: pulumi.Any(_var.Catalog_item_id),
//				ProjectId:     pulumi.Any(_var.Project_id),
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
// Catalog item entitlement can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import vra:index/catalogItemEntitlement:CatalogItemEntitlement this 05956583-6488-4e7d-84c9-92a7b7219a15`
//
// ```
type CatalogItemEntitlement struct {
	pulumi.CustomResourceState

	// The id of the catalog item to create the entitlement.
	CatalogItemId pulumi.StringOutput `pulumi:"catalogItemId"`
	// Represents a catalog item that is linked to a project via an entitlement.
	Definitions CatalogItemEntitlementDefinitionArrayOutput `pulumi:"definitions"`
	// The id of the project this entity belongs to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
}

// NewCatalogItemEntitlement registers a new resource with the given unique name, arguments, and options.
func NewCatalogItemEntitlement(ctx *pulumi.Context,
	name string, args *CatalogItemEntitlementArgs, opts ...pulumi.ResourceOption) (*CatalogItemEntitlement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogItemId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogItemId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CatalogItemEntitlement
	err := ctx.RegisterResource("vra:index/catalogItemEntitlement:CatalogItemEntitlement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalogItemEntitlement gets an existing CatalogItemEntitlement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalogItemEntitlement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogItemEntitlementState, opts ...pulumi.ResourceOption) (*CatalogItemEntitlement, error) {
	var resource CatalogItemEntitlement
	err := ctx.ReadResource("vra:index/catalogItemEntitlement:CatalogItemEntitlement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CatalogItemEntitlement resources.
type catalogItemEntitlementState struct {
	// The id of the catalog item to create the entitlement.
	CatalogItemId *string `pulumi:"catalogItemId"`
	// Represents a catalog item that is linked to a project via an entitlement.
	Definitions []CatalogItemEntitlementDefinition `pulumi:"definitions"`
	// The id of the project this entity belongs to.
	ProjectId *string `pulumi:"projectId"`
}

type CatalogItemEntitlementState struct {
	// The id of the catalog item to create the entitlement.
	CatalogItemId pulumi.StringPtrInput
	// Represents a catalog item that is linked to a project via an entitlement.
	Definitions CatalogItemEntitlementDefinitionArrayInput
	// The id of the project this entity belongs to.
	ProjectId pulumi.StringPtrInput
}

func (CatalogItemEntitlementState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemEntitlementState)(nil)).Elem()
}

type catalogItemEntitlementArgs struct {
	// The id of the catalog item to create the entitlement.
	CatalogItemId string `pulumi:"catalogItemId"`
	// The id of the project this entity belongs to.
	ProjectId string `pulumi:"projectId"`
}

// The set of arguments for constructing a CatalogItemEntitlement resource.
type CatalogItemEntitlementArgs struct {
	// The id of the catalog item to create the entitlement.
	CatalogItemId pulumi.StringInput
	// The id of the project this entity belongs to.
	ProjectId pulumi.StringInput
}

func (CatalogItemEntitlementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogItemEntitlementArgs)(nil)).Elem()
}

type CatalogItemEntitlementInput interface {
	pulumi.Input

	ToCatalogItemEntitlementOutput() CatalogItemEntitlementOutput
	ToCatalogItemEntitlementOutputWithContext(ctx context.Context) CatalogItemEntitlementOutput
}

func (*CatalogItemEntitlement) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogItemEntitlement)(nil)).Elem()
}

func (i *CatalogItemEntitlement) ToCatalogItemEntitlementOutput() CatalogItemEntitlementOutput {
	return i.ToCatalogItemEntitlementOutputWithContext(context.Background())
}

func (i *CatalogItemEntitlement) ToCatalogItemEntitlementOutputWithContext(ctx context.Context) CatalogItemEntitlementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemEntitlementOutput)
}

// CatalogItemEntitlementArrayInput is an input type that accepts CatalogItemEntitlementArray and CatalogItemEntitlementArrayOutput values.
// You can construct a concrete instance of `CatalogItemEntitlementArrayInput` via:
//
//	CatalogItemEntitlementArray{ CatalogItemEntitlementArgs{...} }
type CatalogItemEntitlementArrayInput interface {
	pulumi.Input

	ToCatalogItemEntitlementArrayOutput() CatalogItemEntitlementArrayOutput
	ToCatalogItemEntitlementArrayOutputWithContext(context.Context) CatalogItemEntitlementArrayOutput
}

type CatalogItemEntitlementArray []CatalogItemEntitlementInput

func (CatalogItemEntitlementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CatalogItemEntitlement)(nil)).Elem()
}

func (i CatalogItemEntitlementArray) ToCatalogItemEntitlementArrayOutput() CatalogItemEntitlementArrayOutput {
	return i.ToCatalogItemEntitlementArrayOutputWithContext(context.Background())
}

func (i CatalogItemEntitlementArray) ToCatalogItemEntitlementArrayOutputWithContext(ctx context.Context) CatalogItemEntitlementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemEntitlementArrayOutput)
}

// CatalogItemEntitlementMapInput is an input type that accepts CatalogItemEntitlementMap and CatalogItemEntitlementMapOutput values.
// You can construct a concrete instance of `CatalogItemEntitlementMapInput` via:
//
//	CatalogItemEntitlementMap{ "key": CatalogItemEntitlementArgs{...} }
type CatalogItemEntitlementMapInput interface {
	pulumi.Input

	ToCatalogItemEntitlementMapOutput() CatalogItemEntitlementMapOutput
	ToCatalogItemEntitlementMapOutputWithContext(context.Context) CatalogItemEntitlementMapOutput
}

type CatalogItemEntitlementMap map[string]CatalogItemEntitlementInput

func (CatalogItemEntitlementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CatalogItemEntitlement)(nil)).Elem()
}

func (i CatalogItemEntitlementMap) ToCatalogItemEntitlementMapOutput() CatalogItemEntitlementMapOutput {
	return i.ToCatalogItemEntitlementMapOutputWithContext(context.Background())
}

func (i CatalogItemEntitlementMap) ToCatalogItemEntitlementMapOutputWithContext(ctx context.Context) CatalogItemEntitlementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogItemEntitlementMapOutput)
}

type CatalogItemEntitlementOutput struct{ *pulumi.OutputState }

func (CatalogItemEntitlementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CatalogItemEntitlement)(nil)).Elem()
}

func (o CatalogItemEntitlementOutput) ToCatalogItemEntitlementOutput() CatalogItemEntitlementOutput {
	return o
}

func (o CatalogItemEntitlementOutput) ToCatalogItemEntitlementOutputWithContext(ctx context.Context) CatalogItemEntitlementOutput {
	return o
}

// The id of the catalog item to create the entitlement.
func (o CatalogItemEntitlementOutput) CatalogItemId() pulumi.StringOutput {
	return o.ApplyT(func(v *CatalogItemEntitlement) pulumi.StringOutput { return v.CatalogItemId }).(pulumi.StringOutput)
}

// Represents a catalog item that is linked to a project via an entitlement.
func (o CatalogItemEntitlementOutput) Definitions() CatalogItemEntitlementDefinitionArrayOutput {
	return o.ApplyT(func(v *CatalogItemEntitlement) CatalogItemEntitlementDefinitionArrayOutput { return v.Definitions }).(CatalogItemEntitlementDefinitionArrayOutput)
}

// The id of the project this entity belongs to.
func (o CatalogItemEntitlementOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CatalogItemEntitlement) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

type CatalogItemEntitlementArrayOutput struct{ *pulumi.OutputState }

func (CatalogItemEntitlementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CatalogItemEntitlement)(nil)).Elem()
}

func (o CatalogItemEntitlementArrayOutput) ToCatalogItemEntitlementArrayOutput() CatalogItemEntitlementArrayOutput {
	return o
}

func (o CatalogItemEntitlementArrayOutput) ToCatalogItemEntitlementArrayOutputWithContext(ctx context.Context) CatalogItemEntitlementArrayOutput {
	return o
}

func (o CatalogItemEntitlementArrayOutput) Index(i pulumi.IntInput) CatalogItemEntitlementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CatalogItemEntitlement {
		return vs[0].([]*CatalogItemEntitlement)[vs[1].(int)]
	}).(CatalogItemEntitlementOutput)
}

type CatalogItemEntitlementMapOutput struct{ *pulumi.OutputState }

func (CatalogItemEntitlementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CatalogItemEntitlement)(nil)).Elem()
}

func (o CatalogItemEntitlementMapOutput) ToCatalogItemEntitlementMapOutput() CatalogItemEntitlementMapOutput {
	return o
}

func (o CatalogItemEntitlementMapOutput) ToCatalogItemEntitlementMapOutputWithContext(ctx context.Context) CatalogItemEntitlementMapOutput {
	return o
}

func (o CatalogItemEntitlementMapOutput) MapIndex(k pulumi.StringInput) CatalogItemEntitlementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CatalogItemEntitlement {
		return vs[0].(map[string]*CatalogItemEntitlement)[vs[1].(string)]
	}).(CatalogItemEntitlementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemEntitlementInput)(nil)).Elem(), &CatalogItemEntitlement{})
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemEntitlementArrayInput)(nil)).Elem(), CatalogItemEntitlementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CatalogItemEntitlementMapInput)(nil)).Elem(), CatalogItemEntitlementMap{})
	pulumi.RegisterOutputType(CatalogItemEntitlementOutput{})
	pulumi.RegisterOutputType(CatalogItemEntitlementArrayOutput{})
	pulumi.RegisterOutputType(CatalogItemEntitlementMapOutput{})
}
