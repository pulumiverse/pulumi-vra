// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides information about a cloud template (blueprint) version in vRA.
//
// ## Example Usage
// ### S
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vra/sdk/go/vra/blueprint"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/blueprint"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := blueprint.GetVersion(ctx, &blueprint.GetVersionArgs{
// 			BlueprintId: _var.Blueprint_id,
// 			Id:          _var.Blueprint_version_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetVersion(ctx *pulumi.Context, args *GetVersionArgs, opts ...pulumi.InvokeOption) (*GetVersionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetVersionResult
	err := ctx.Invoke("vra:blueprint/getVersion:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVersion.
type GetVersionArgs struct {
	// Name of the cloud template. One of `id` or `name` must be provided.
	BlueprintId string `pulumi:"blueprintId"`
	// The id of the cloud template version.
	Id string `pulumi:"id"`
}

// A collection of values returned by getVersion.
type GetVersionResult struct {
	// Description of the cloud template.
	BlueprintDescription string `pulumi:"blueprintDescription"`
	BlueprintId          string `pulumi:"blueprintId"`
	// Blueprint YAML content.
	Content string `pulumi:"content"`
	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `pulumi:"createdAt"`
	// The user the entity was created by.
	CreatedBy string `pulumi:"createdBy"`
	// (Optional) Cloud template version description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// Name of the cloud template version.
	Name string `pulumi:"name"`
	// The id of the organization this entity belongs to.
	OrgId string `pulumi:"orgId"`
	// The id of the project this entity belongs to.
	ProjectId string `pulumi:"projectId"`
	// The name of the project the entity belongs to.
	ProjectName string `pulumi:"projectName"`
	// Status of the cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
	Status string `pulumi:"status"`
	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `pulumi:"updatedAt"`
	// The user the entity was last updated by.
	UpdatedBy string `pulumi:"updatedBy"`
	// Flag to indicate if the current content of the cloud template is valid.
	Valid string `pulumi:"valid"`
	// Cloud template version.
	Version string `pulumi:"version"`
	// Cloud template version change log.
	VersionChangeLog string `pulumi:"versionChangeLog"`
}

func GetVersionOutput(ctx *pulumi.Context, args GetVersionOutputArgs, opts ...pulumi.InvokeOption) GetVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVersionResult, error) {
			args := v.(GetVersionArgs)
			r, err := GetVersion(ctx, &args, opts...)
			var s GetVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVersionResultOutput)
}

// A collection of arguments for invoking getVersion.
type GetVersionOutputArgs struct {
	// Name of the cloud template. One of `id` or `name` must be provided.
	BlueprintId pulumi.StringInput `pulumi:"blueprintId"`
	// The id of the cloud template version.
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVersionArgs)(nil)).Elem()
}

// A collection of values returned by getVersion.
type GetVersionResultOutput struct{ *pulumi.OutputState }

func (GetVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVersionResult)(nil)).Elem()
}

func (o GetVersionResultOutput) ToGetVersionResultOutput() GetVersionResultOutput {
	return o
}

func (o GetVersionResultOutput) ToGetVersionResultOutputWithContext(ctx context.Context) GetVersionResultOutput {
	return o
}

// Description of the cloud template.
func (o GetVersionResultOutput) BlueprintDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.BlueprintDescription }).(pulumi.StringOutput)
}

func (o GetVersionResultOutput) BlueprintId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.BlueprintId }).(pulumi.StringOutput)
}

// Blueprint YAML content.
func (o GetVersionResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Content }).(pulumi.StringOutput)
}

// Date when the entity was created. The date is in ISO 6801 and UTC.
func (o GetVersionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The user the entity was created by.
func (o GetVersionResultOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.CreatedBy }).(pulumi.StringOutput)
}

// (Optional) Cloud template version description.
func (o GetVersionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the cloud template version.
func (o GetVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the organization this entity belongs to.
func (o GetVersionResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// The id of the project this entity belongs to.
func (o GetVersionResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The name of the project the entity belongs to.
func (o GetVersionResultOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.ProjectName }).(pulumi.StringOutput)
}

// Status of the cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
func (o GetVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Date when the entity was last updated. The date is ISO 8601 and UTC.
func (o GetVersionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The user the entity was last updated by.
func (o GetVersionResultOutput) UpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.UpdatedBy }).(pulumi.StringOutput)
}

// Flag to indicate if the current content of the cloud template is valid.
func (o GetVersionResultOutput) Valid() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Valid }).(pulumi.StringOutput)
}

// Cloud template version.
func (o GetVersionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.Version }).(pulumi.StringOutput)
}

// Cloud template version change log.
func (o GetVersionResultOutput) VersionChangeLog() pulumi.StringOutput {
	return o.ApplyT(func(v GetVersionResult) string { return v.VersionChangeLog }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVersionResultOutput{})
}
