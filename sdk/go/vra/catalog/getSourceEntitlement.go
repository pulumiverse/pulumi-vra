// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package catalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about a catalog source entitlement in vRA.
//
// ## Example Usage
// ### S
//
// This is an example of how to get a vRA catalog source entitlement by its id:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/catalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/catalog"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := catalog.LookupSourceEntitlement(ctx, &catalog.LookupSourceEntitlementArgs{
//				Id:        pulumi.StringRef(_var.Catalog_source_entitlement_id),
//				ProjectId: _var.Project_id,
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
// This is an example of how to get a vRA catalog source entitlement by its catalog source id:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra/catalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/catalog"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := catalog.LookupSourceEntitlement(ctx, &catalog.LookupSourceEntitlementArgs{
//				CatalogSourceId: pulumi.StringRef(_var.Catalog_source_id),
//				ProjectId:       _var.Project_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSourceEntitlement(ctx *pulumi.Context, args *LookupSourceEntitlementArgs, opts ...pulumi.InvokeOption) (*LookupSourceEntitlementResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSourceEntitlementResult
	err := ctx.Invoke("vra:catalog/getSourceEntitlement:getSourceEntitlement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceEntitlement.
type LookupSourceEntitlementArgs struct {
	// The id of the catalog source to find the entitlement. One of `catalogSourceId` or `id` must be provided.
	CatalogSourceId *string `pulumi:"catalogSourceId"`
	// The id of entitlement. One of `catalogSourceId` or `id` must be provided.
	Id *string `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getSourceEntitlement.
type LookupSourceEntitlementResult struct {
	CatalogSourceId *string `pulumi:"catalogSourceId"`
	// Represents a catalog source that is linked to a project via an entitlement.
	Definitions []GetSourceEntitlementDefinition `pulumi:"definitions"`
	// Id of the catalog source.
	Id        *string `pulumi:"id"`
	ProjectId string  `pulumi:"projectId"`
}

func LookupSourceEntitlementOutput(ctx *pulumi.Context, args LookupSourceEntitlementOutputArgs, opts ...pulumi.InvokeOption) LookupSourceEntitlementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceEntitlementResult, error) {
			args := v.(LookupSourceEntitlementArgs)
			r, err := LookupSourceEntitlement(ctx, &args, opts...)
			var s LookupSourceEntitlementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceEntitlementResultOutput)
}

// A collection of arguments for invoking getSourceEntitlement.
type LookupSourceEntitlementOutputArgs struct {
	// The id of the catalog source to find the entitlement. One of `catalogSourceId` or `id` must be provided.
	CatalogSourceId pulumi.StringPtrInput `pulumi:"catalogSourceId"`
	// The id of entitlement. One of `catalogSourceId` or `id` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The id of the project that this entitlement belongs to.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupSourceEntitlementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceEntitlementArgs)(nil)).Elem()
}

// A collection of values returned by getSourceEntitlement.
type LookupSourceEntitlementResultOutput struct{ *pulumi.OutputState }

func (LookupSourceEntitlementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceEntitlementResult)(nil)).Elem()
}

func (o LookupSourceEntitlementResultOutput) ToLookupSourceEntitlementResultOutput() LookupSourceEntitlementResultOutput {
	return o
}

func (o LookupSourceEntitlementResultOutput) ToLookupSourceEntitlementResultOutputWithContext(ctx context.Context) LookupSourceEntitlementResultOutput {
	return o
}

func (o LookupSourceEntitlementResultOutput) CatalogSourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceEntitlementResult) *string { return v.CatalogSourceId }).(pulumi.StringPtrOutput)
}

// Represents a catalog source that is linked to a project via an entitlement.
func (o LookupSourceEntitlementResultOutput) Definitions() GetSourceEntitlementDefinitionArrayOutput {
	return o.ApplyT(func(v LookupSourceEntitlementResult) []GetSourceEntitlementDefinition { return v.Definitions }).(GetSourceEntitlementDefinitionArrayOutput)
}

// Id of the catalog source.
func (o LookupSourceEntitlementResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceEntitlementResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSourceEntitlementResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceEntitlementResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceEntitlementResultOutput{})
}
