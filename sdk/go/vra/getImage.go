// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getImage` data source can be used to discover the lookup machine images with cloud accounts. This can then be used with resource that require an image. For example, to create an image profile using the `ImageProfile` resource.
//
// ## Example Usage
//
// This is an example of how to lookup images from a vSphere cloud account.
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-vra/sdk/go/vra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			thisCloudAccountVsphere, err := vra.LookupCloudAccountVsphere(ctx, &GetCloudAccountVsphereArgs{
//				Name: pulumi.StringRef(_var.Cloud_account),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			thisRegion, err := vra.GetRegion(ctx, &GetRegionArgs{
//				CloudAccountId: pulumi.StringRef(thisCloudAccountVsphere.Id),
//				Region:         pulumi.StringRef(_var.Region),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			image0, err := vra.GetImage(ctx, &GetImageArgs{
//				Filter: fmt.Sprintf("name eq '%v' and cloudAccountId eq '%v' and externalRegionId eq '%v'", _var.Image_name_0, thisCloudAccountVsphere.Id, _var.Region),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			image1, err := vra.GetImage(ctx, &GetImageArgs{
//				Filter: fmt.Sprintf("name eq '%v' and cloudAccountId eq '%v' and externalRegionId eq '%v'", _var.Image_name_1, thisCloudAccountVsphere.Id, _var.Region),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vra.NewImageProfile(ctx, "thisImageProfile", &vra.ImageProfileArgs{
//				Description: pulumi.Any(_var.Image_profile_description),
//				RegionId:    pulumi.String(thisRegion.Id),
//				ImageMappings: ImageProfileImageMappingArray{
//					&ImageProfileImageMappingArgs{
//						Name:    pulumi.Any(_var.Image_name_0),
//						ImageId: pulumi.String(image0.Id),
//					},
//					&ImageProfileImageMappingArgs{
//						Name:    pulumi.Any(_var.Image_name_1),
//						ImageId: pulumi.String(image1.Id),
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
func GetImage(ctx *pulumi.Context, args *GetImageArgs, opts ...pulumi.InvokeOption) (*GetImageResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetImageResult
	err := ctx.Invoke("vra:index/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type GetImageArgs struct {
	// Search criteria to narrow down the image discovery.
	Filter string `pulumi:"filter"`
}

// A collection of values returned by getImage.
type GetImageResult struct {
	// A human-friendly description of the image.
	Description string `pulumi:"description"`
	// External entity id on the provider side.
	ExternalId string `pulumi:"externalId"`
	Filter     string `pulumi:"filter"`
	// The id of the image.
	Id string `pulumi:"id"`
	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `pulumi:"name"`
	// Indicates whether this image is private. For vSphere, private images are templates and snapshots and public images are content library items.
	Private bool `pulumi:"private"`
	// The regionId of the image. For a vSphere cloud account, it is the `externalRegionId` such as `Datacenter:datacenter-2` and for an AWS cloud account, it is region name such as `us-east-1`, etc.
	Region string `pulumi:"region"`
}

func GetImageOutput(ctx *pulumi.Context, args GetImageOutputArgs, opts ...pulumi.InvokeOption) GetImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageResult, error) {
			args := v.(GetImageArgs)
			r, err := GetImage(ctx, &args, opts...)
			var s GetImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageResultOutput)
}

// A collection of arguments for invoking getImage.
type GetImageOutputArgs struct {
	// Search criteria to narrow down the image discovery.
	Filter pulumi.StringInput `pulumi:"filter"`
}

func (GetImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageArgs)(nil)).Elem()
}

// A collection of values returned by getImage.
type GetImageResultOutput struct{ *pulumi.OutputState }

func (GetImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageResult)(nil)).Elem()
}

func (o GetImageResultOutput) ToGetImageResultOutput() GetImageResultOutput {
	return o
}

func (o GetImageResultOutput) ToGetImageResultOutputWithContext(ctx context.Context) GetImageResultOutput {
	return o
}

// A human-friendly description of the image.
func (o GetImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Description }).(pulumi.StringOutput)
}

// External entity id on the provider side.
func (o GetImageResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

func (o GetImageResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Filter }).(pulumi.StringOutput)
}

// The id of the image.
func (o GetImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// A human-friendly name used as an identifier in APIs that support this option.
func (o GetImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether this image is private. For vSphere, private images are templates and snapshots and public images are content library items.
func (o GetImageResultOutput) Private() pulumi.BoolOutput {
	return o.ApplyT(func(v GetImageResult) bool { return v.Private }).(pulumi.BoolOutput)
}

// The regionId of the image. For a vSphere cloud account, it is the `externalRegionId` such as `Datacenter:datacenter-2` and for an AWS cloud account, it is region name such as `us-east-1`, etc.
func (o GetImageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageResultOutput{})
}
