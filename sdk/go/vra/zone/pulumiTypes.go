// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ZoneLink struct {
	Href  *string  `pulumi:"href"`
	Hrefs []string `pulumi:"hrefs"`
	Rel   string   `pulumi:"rel"`
}

// ZoneLinkInput is an input type that accepts ZoneLinkArgs and ZoneLinkOutput values.
// You can construct a concrete instance of `ZoneLinkInput` via:
//
//          ZoneLinkArgs{...}
type ZoneLinkInput interface {
	pulumi.Input

	ToZoneLinkOutput() ZoneLinkOutput
	ToZoneLinkOutputWithContext(context.Context) ZoneLinkOutput
}

type ZoneLinkArgs struct {
	Href  pulumi.StringPtrInput   `pulumi:"href"`
	Hrefs pulumi.StringArrayInput `pulumi:"hrefs"`
	Rel   pulumi.StringInput      `pulumi:"rel"`
}

func (ZoneLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneLink)(nil)).Elem()
}

func (i ZoneLinkArgs) ToZoneLinkOutput() ZoneLinkOutput {
	return i.ToZoneLinkOutputWithContext(context.Background())
}

func (i ZoneLinkArgs) ToZoneLinkOutputWithContext(ctx context.Context) ZoneLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneLinkOutput)
}

// ZoneLinkArrayInput is an input type that accepts ZoneLinkArray and ZoneLinkArrayOutput values.
// You can construct a concrete instance of `ZoneLinkArrayInput` via:
//
//          ZoneLinkArray{ ZoneLinkArgs{...} }
type ZoneLinkArrayInput interface {
	pulumi.Input

	ToZoneLinkArrayOutput() ZoneLinkArrayOutput
	ToZoneLinkArrayOutputWithContext(context.Context) ZoneLinkArrayOutput
}

type ZoneLinkArray []ZoneLinkInput

func (ZoneLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneLink)(nil)).Elem()
}

func (i ZoneLinkArray) ToZoneLinkArrayOutput() ZoneLinkArrayOutput {
	return i.ToZoneLinkArrayOutputWithContext(context.Background())
}

func (i ZoneLinkArray) ToZoneLinkArrayOutputWithContext(ctx context.Context) ZoneLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneLinkArrayOutput)
}

type ZoneLinkOutput struct{ *pulumi.OutputState }

func (ZoneLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneLink)(nil)).Elem()
}

func (o ZoneLinkOutput) ToZoneLinkOutput() ZoneLinkOutput {
	return o
}

func (o ZoneLinkOutput) ToZoneLinkOutputWithContext(ctx context.Context) ZoneLinkOutput {
	return o
}

func (o ZoneLinkOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZoneLink) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o ZoneLinkOutput) Hrefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ZoneLink) []string { return v.Hrefs }).(pulumi.StringArrayOutput)
}

func (o ZoneLinkOutput) Rel() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneLink) string { return v.Rel }).(pulumi.StringOutput)
}

type ZoneLinkArrayOutput struct{ *pulumi.OutputState }

func (ZoneLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneLink)(nil)).Elem()
}

func (o ZoneLinkArrayOutput) ToZoneLinkArrayOutput() ZoneLinkArrayOutput {
	return o
}

func (o ZoneLinkArrayOutput) ToZoneLinkArrayOutputWithContext(ctx context.Context) ZoneLinkArrayOutput {
	return o
}

func (o ZoneLinkArrayOutput) Index(i pulumi.IntInput) ZoneLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneLink {
		return vs[0].([]ZoneLink)[vs[1].(int)]
	}).(ZoneLinkOutput)
}

type ZoneTag struct {
	// Tag’s key.
	Key string `pulumi:"key"`
	// Tag’s value.
	Value string `pulumi:"value"`
}

// ZoneTagInput is an input type that accepts ZoneTagArgs and ZoneTagOutput values.
// You can construct a concrete instance of `ZoneTagInput` via:
//
//          ZoneTagArgs{...}
type ZoneTagInput interface {
	pulumi.Input

	ToZoneTagOutput() ZoneTagOutput
	ToZoneTagOutputWithContext(context.Context) ZoneTagOutput
}

type ZoneTagArgs struct {
	// Tag’s key.
	Key pulumi.StringInput `pulumi:"key"`
	// Tag’s value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ZoneTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneTag)(nil)).Elem()
}

func (i ZoneTagArgs) ToZoneTagOutput() ZoneTagOutput {
	return i.ToZoneTagOutputWithContext(context.Background())
}

func (i ZoneTagArgs) ToZoneTagOutputWithContext(ctx context.Context) ZoneTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneTagOutput)
}

// ZoneTagArrayInput is an input type that accepts ZoneTagArray and ZoneTagArrayOutput values.
// You can construct a concrete instance of `ZoneTagArrayInput` via:
//
//          ZoneTagArray{ ZoneTagArgs{...} }
type ZoneTagArrayInput interface {
	pulumi.Input

	ToZoneTagArrayOutput() ZoneTagArrayOutput
	ToZoneTagArrayOutputWithContext(context.Context) ZoneTagArrayOutput
}

type ZoneTagArray []ZoneTagInput

func (ZoneTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneTag)(nil)).Elem()
}

func (i ZoneTagArray) ToZoneTagArrayOutput() ZoneTagArrayOutput {
	return i.ToZoneTagArrayOutputWithContext(context.Background())
}

func (i ZoneTagArray) ToZoneTagArrayOutputWithContext(ctx context.Context) ZoneTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneTagArrayOutput)
}

type ZoneTagOutput struct{ *pulumi.OutputState }

func (ZoneTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneTag)(nil)).Elem()
}

func (o ZoneTagOutput) ToZoneTagOutput() ZoneTagOutput {
	return o
}

func (o ZoneTagOutput) ToZoneTagOutputWithContext(ctx context.Context) ZoneTagOutput {
	return o
}

// Tag’s key.
func (o ZoneTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneTag) string { return v.Key }).(pulumi.StringOutput)
}

// Tag’s value.
func (o ZoneTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneTag) string { return v.Value }).(pulumi.StringOutput)
}

type ZoneTagArrayOutput struct{ *pulumi.OutputState }

func (ZoneTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneTag)(nil)).Elem()
}

func (o ZoneTagArrayOutput) ToZoneTagArrayOutput() ZoneTagArrayOutput {
	return o
}

func (o ZoneTagArrayOutput) ToZoneTagArrayOutputWithContext(ctx context.Context) ZoneTagArrayOutput {
	return o
}

func (o ZoneTagArrayOutput) Index(i pulumi.IntInput) ZoneTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneTag {
		return vs[0].([]ZoneTag)[vs[1].(int)]
	}).(ZoneTagOutput)
}

type ZoneTagsToMatch struct {
	// Tag’s key.
	Key string `pulumi:"key"`
	// Tag’s value.
	Value string `pulumi:"value"`
}

// ZoneTagsToMatchInput is an input type that accepts ZoneTagsToMatchArgs and ZoneTagsToMatchOutput values.
// You can construct a concrete instance of `ZoneTagsToMatchInput` via:
//
//          ZoneTagsToMatchArgs{...}
type ZoneTagsToMatchInput interface {
	pulumi.Input

	ToZoneTagsToMatchOutput() ZoneTagsToMatchOutput
	ToZoneTagsToMatchOutputWithContext(context.Context) ZoneTagsToMatchOutput
}

type ZoneTagsToMatchArgs struct {
	// Tag’s key.
	Key pulumi.StringInput `pulumi:"key"`
	// Tag’s value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ZoneTagsToMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneTagsToMatch)(nil)).Elem()
}

func (i ZoneTagsToMatchArgs) ToZoneTagsToMatchOutput() ZoneTagsToMatchOutput {
	return i.ToZoneTagsToMatchOutputWithContext(context.Background())
}

func (i ZoneTagsToMatchArgs) ToZoneTagsToMatchOutputWithContext(ctx context.Context) ZoneTagsToMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneTagsToMatchOutput)
}

// ZoneTagsToMatchArrayInput is an input type that accepts ZoneTagsToMatchArray and ZoneTagsToMatchArrayOutput values.
// You can construct a concrete instance of `ZoneTagsToMatchArrayInput` via:
//
//          ZoneTagsToMatchArray{ ZoneTagsToMatchArgs{...} }
type ZoneTagsToMatchArrayInput interface {
	pulumi.Input

	ToZoneTagsToMatchArrayOutput() ZoneTagsToMatchArrayOutput
	ToZoneTagsToMatchArrayOutputWithContext(context.Context) ZoneTagsToMatchArrayOutput
}

type ZoneTagsToMatchArray []ZoneTagsToMatchInput

func (ZoneTagsToMatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneTagsToMatch)(nil)).Elem()
}

func (i ZoneTagsToMatchArray) ToZoneTagsToMatchArrayOutput() ZoneTagsToMatchArrayOutput {
	return i.ToZoneTagsToMatchArrayOutputWithContext(context.Background())
}

func (i ZoneTagsToMatchArray) ToZoneTagsToMatchArrayOutputWithContext(ctx context.Context) ZoneTagsToMatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneTagsToMatchArrayOutput)
}

type ZoneTagsToMatchOutput struct{ *pulumi.OutputState }

func (ZoneTagsToMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZoneTagsToMatch)(nil)).Elem()
}

func (o ZoneTagsToMatchOutput) ToZoneTagsToMatchOutput() ZoneTagsToMatchOutput {
	return o
}

func (o ZoneTagsToMatchOutput) ToZoneTagsToMatchOutputWithContext(ctx context.Context) ZoneTagsToMatchOutput {
	return o
}

// Tag’s key.
func (o ZoneTagsToMatchOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneTagsToMatch) string { return v.Key }).(pulumi.StringOutput)
}

// Tag’s value.
func (o ZoneTagsToMatchOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ZoneTagsToMatch) string { return v.Value }).(pulumi.StringOutput)
}

type ZoneTagsToMatchArrayOutput struct{ *pulumi.OutputState }

func (ZoneTagsToMatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZoneTagsToMatch)(nil)).Elem()
}

func (o ZoneTagsToMatchArrayOutput) ToZoneTagsToMatchArrayOutput() ZoneTagsToMatchArrayOutput {
	return o
}

func (o ZoneTagsToMatchArrayOutput) ToZoneTagsToMatchArrayOutputWithContext(ctx context.Context) ZoneTagsToMatchArrayOutput {
	return o
}

func (o ZoneTagsToMatchArrayOutput) Index(i pulumi.IntInput) ZoneTagsToMatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZoneTagsToMatch {
		return vs[0].([]ZoneTagsToMatch)[vs[1].(int)]
	}).(ZoneTagsToMatchOutput)
}

type GetZoneLink struct {
	Href  *string  `pulumi:"href"`
	Hrefs []string `pulumi:"hrefs"`
	Rel   string   `pulumi:"rel"`
}

// GetZoneLinkInput is an input type that accepts GetZoneLinkArgs and GetZoneLinkOutput values.
// You can construct a concrete instance of `GetZoneLinkInput` via:
//
//          GetZoneLinkArgs{...}
type GetZoneLinkInput interface {
	pulumi.Input

	ToGetZoneLinkOutput() GetZoneLinkOutput
	ToGetZoneLinkOutputWithContext(context.Context) GetZoneLinkOutput
}

type GetZoneLinkArgs struct {
	Href  pulumi.StringPtrInput   `pulumi:"href"`
	Hrefs pulumi.StringArrayInput `pulumi:"hrefs"`
	Rel   pulumi.StringInput      `pulumi:"rel"`
}

func (GetZoneLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneLink)(nil)).Elem()
}

func (i GetZoneLinkArgs) ToGetZoneLinkOutput() GetZoneLinkOutput {
	return i.ToGetZoneLinkOutputWithContext(context.Background())
}

func (i GetZoneLinkArgs) ToGetZoneLinkOutputWithContext(ctx context.Context) GetZoneLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneLinkOutput)
}

// GetZoneLinkArrayInput is an input type that accepts GetZoneLinkArray and GetZoneLinkArrayOutput values.
// You can construct a concrete instance of `GetZoneLinkArrayInput` via:
//
//          GetZoneLinkArray{ GetZoneLinkArgs{...} }
type GetZoneLinkArrayInput interface {
	pulumi.Input

	ToGetZoneLinkArrayOutput() GetZoneLinkArrayOutput
	ToGetZoneLinkArrayOutputWithContext(context.Context) GetZoneLinkArrayOutput
}

type GetZoneLinkArray []GetZoneLinkInput

func (GetZoneLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneLink)(nil)).Elem()
}

func (i GetZoneLinkArray) ToGetZoneLinkArrayOutput() GetZoneLinkArrayOutput {
	return i.ToGetZoneLinkArrayOutputWithContext(context.Background())
}

func (i GetZoneLinkArray) ToGetZoneLinkArrayOutputWithContext(ctx context.Context) GetZoneLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneLinkArrayOutput)
}

type GetZoneLinkOutput struct{ *pulumi.OutputState }

func (GetZoneLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneLink)(nil)).Elem()
}

func (o GetZoneLinkOutput) ToGetZoneLinkOutput() GetZoneLinkOutput {
	return o
}

func (o GetZoneLinkOutput) ToGetZoneLinkOutputWithContext(ctx context.Context) GetZoneLinkOutput {
	return o
}

func (o GetZoneLinkOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneLink) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o GetZoneLinkOutput) Hrefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetZoneLink) []string { return v.Hrefs }).(pulumi.StringArrayOutput)
}

func (o GetZoneLinkOutput) Rel() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneLink) string { return v.Rel }).(pulumi.StringOutput)
}

type GetZoneLinkArrayOutput struct{ *pulumi.OutputState }

func (GetZoneLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneLink)(nil)).Elem()
}

func (o GetZoneLinkArrayOutput) ToGetZoneLinkArrayOutput() GetZoneLinkArrayOutput {
	return o
}

func (o GetZoneLinkArrayOutput) ToGetZoneLinkArrayOutputWithContext(ctx context.Context) GetZoneLinkArrayOutput {
	return o
}

func (o GetZoneLinkArrayOutput) Index(i pulumi.IntInput) GetZoneLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetZoneLink {
		return vs[0].([]GetZoneLink)[vs[1].(int)]
	}).(GetZoneLinkOutput)
}

type GetZoneTag struct {
	// Tag’s key.
	Key string `pulumi:"key"`
	// Tag’s value.
	Value string `pulumi:"value"`
}

// GetZoneTagInput is an input type that accepts GetZoneTagArgs and GetZoneTagOutput values.
// You can construct a concrete instance of `GetZoneTagInput` via:
//
//          GetZoneTagArgs{...}
type GetZoneTagInput interface {
	pulumi.Input

	ToGetZoneTagOutput() GetZoneTagOutput
	ToGetZoneTagOutputWithContext(context.Context) GetZoneTagOutput
}

type GetZoneTagArgs struct {
	// Tag’s key.
	Key pulumi.StringInput `pulumi:"key"`
	// Tag’s value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetZoneTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneTag)(nil)).Elem()
}

func (i GetZoneTagArgs) ToGetZoneTagOutput() GetZoneTagOutput {
	return i.ToGetZoneTagOutputWithContext(context.Background())
}

func (i GetZoneTagArgs) ToGetZoneTagOutputWithContext(ctx context.Context) GetZoneTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneTagOutput)
}

// GetZoneTagArrayInput is an input type that accepts GetZoneTagArray and GetZoneTagArrayOutput values.
// You can construct a concrete instance of `GetZoneTagArrayInput` via:
//
//          GetZoneTagArray{ GetZoneTagArgs{...} }
type GetZoneTagArrayInput interface {
	pulumi.Input

	ToGetZoneTagArrayOutput() GetZoneTagArrayOutput
	ToGetZoneTagArrayOutputWithContext(context.Context) GetZoneTagArrayOutput
}

type GetZoneTagArray []GetZoneTagInput

func (GetZoneTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneTag)(nil)).Elem()
}

func (i GetZoneTagArray) ToGetZoneTagArrayOutput() GetZoneTagArrayOutput {
	return i.ToGetZoneTagArrayOutputWithContext(context.Background())
}

func (i GetZoneTagArray) ToGetZoneTagArrayOutputWithContext(ctx context.Context) GetZoneTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneTagArrayOutput)
}

type GetZoneTagOutput struct{ *pulumi.OutputState }

func (GetZoneTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneTag)(nil)).Elem()
}

func (o GetZoneTagOutput) ToGetZoneTagOutput() GetZoneTagOutput {
	return o
}

func (o GetZoneTagOutput) ToGetZoneTagOutputWithContext(ctx context.Context) GetZoneTagOutput {
	return o
}

// Tag’s key.
func (o GetZoneTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneTag) string { return v.Key }).(pulumi.StringOutput)
}

// Tag’s value.
func (o GetZoneTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneTag) string { return v.Value }).(pulumi.StringOutput)
}

type GetZoneTagArrayOutput struct{ *pulumi.OutputState }

func (GetZoneTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneTag)(nil)).Elem()
}

func (o GetZoneTagArrayOutput) ToGetZoneTagArrayOutput() GetZoneTagArrayOutput {
	return o
}

func (o GetZoneTagArrayOutput) ToGetZoneTagArrayOutputWithContext(ctx context.Context) GetZoneTagArrayOutput {
	return o
}

func (o GetZoneTagArrayOutput) Index(i pulumi.IntInput) GetZoneTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetZoneTag {
		return vs[0].([]GetZoneTag)[vs[1].(int)]
	}).(GetZoneTagOutput)
}

type GetZoneTagsToMatch struct {
	// Tag’s key.
	Key string `pulumi:"key"`
	// Tag’s value.
	Value string `pulumi:"value"`
}

// GetZoneTagsToMatchInput is an input type that accepts GetZoneTagsToMatchArgs and GetZoneTagsToMatchOutput values.
// You can construct a concrete instance of `GetZoneTagsToMatchInput` via:
//
//          GetZoneTagsToMatchArgs{...}
type GetZoneTagsToMatchInput interface {
	pulumi.Input

	ToGetZoneTagsToMatchOutput() GetZoneTagsToMatchOutput
	ToGetZoneTagsToMatchOutputWithContext(context.Context) GetZoneTagsToMatchOutput
}

type GetZoneTagsToMatchArgs struct {
	// Tag’s key.
	Key pulumi.StringInput `pulumi:"key"`
	// Tag’s value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (GetZoneTagsToMatchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneTagsToMatch)(nil)).Elem()
}

func (i GetZoneTagsToMatchArgs) ToGetZoneTagsToMatchOutput() GetZoneTagsToMatchOutput {
	return i.ToGetZoneTagsToMatchOutputWithContext(context.Background())
}

func (i GetZoneTagsToMatchArgs) ToGetZoneTagsToMatchOutputWithContext(ctx context.Context) GetZoneTagsToMatchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneTagsToMatchOutput)
}

// GetZoneTagsToMatchArrayInput is an input type that accepts GetZoneTagsToMatchArray and GetZoneTagsToMatchArrayOutput values.
// You can construct a concrete instance of `GetZoneTagsToMatchArrayInput` via:
//
//          GetZoneTagsToMatchArray{ GetZoneTagsToMatchArgs{...} }
type GetZoneTagsToMatchArrayInput interface {
	pulumi.Input

	ToGetZoneTagsToMatchArrayOutput() GetZoneTagsToMatchArrayOutput
	ToGetZoneTagsToMatchArrayOutputWithContext(context.Context) GetZoneTagsToMatchArrayOutput
}

type GetZoneTagsToMatchArray []GetZoneTagsToMatchInput

func (GetZoneTagsToMatchArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneTagsToMatch)(nil)).Elem()
}

func (i GetZoneTagsToMatchArray) ToGetZoneTagsToMatchArrayOutput() GetZoneTagsToMatchArrayOutput {
	return i.ToGetZoneTagsToMatchArrayOutputWithContext(context.Background())
}

func (i GetZoneTagsToMatchArray) ToGetZoneTagsToMatchArrayOutputWithContext(ctx context.Context) GetZoneTagsToMatchArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetZoneTagsToMatchArrayOutput)
}

type GetZoneTagsToMatchOutput struct{ *pulumi.OutputState }

func (GetZoneTagsToMatchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneTagsToMatch)(nil)).Elem()
}

func (o GetZoneTagsToMatchOutput) ToGetZoneTagsToMatchOutput() GetZoneTagsToMatchOutput {
	return o
}

func (o GetZoneTagsToMatchOutput) ToGetZoneTagsToMatchOutputWithContext(ctx context.Context) GetZoneTagsToMatchOutput {
	return o
}

// Tag’s key.
func (o GetZoneTagsToMatchOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneTagsToMatch) string { return v.Key }).(pulumi.StringOutput)
}

// Tag’s value.
func (o GetZoneTagsToMatchOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneTagsToMatch) string { return v.Value }).(pulumi.StringOutput)
}

type GetZoneTagsToMatchArrayOutput struct{ *pulumi.OutputState }

func (GetZoneTagsToMatchArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetZoneTagsToMatch)(nil)).Elem()
}

func (o GetZoneTagsToMatchArrayOutput) ToGetZoneTagsToMatchArrayOutput() GetZoneTagsToMatchArrayOutput {
	return o
}

func (o GetZoneTagsToMatchArrayOutput) ToGetZoneTagsToMatchArrayOutputWithContext(ctx context.Context) GetZoneTagsToMatchArrayOutput {
	return o
}

func (o GetZoneTagsToMatchArrayOutput) Index(i pulumi.IntInput) GetZoneTagsToMatchOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetZoneTagsToMatch {
		return vs[0].([]GetZoneTagsToMatch)[vs[1].(int)]
	}).(GetZoneTagsToMatchOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneLinkInput)(nil)).Elem(), ZoneLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneLinkArrayInput)(nil)).Elem(), ZoneLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneTagInput)(nil)).Elem(), ZoneTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneTagArrayInput)(nil)).Elem(), ZoneTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneTagsToMatchInput)(nil)).Elem(), ZoneTagsToMatchArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneTagsToMatchArrayInput)(nil)).Elem(), ZoneTagsToMatchArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneLinkInput)(nil)).Elem(), GetZoneLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneLinkArrayInput)(nil)).Elem(), GetZoneLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneTagInput)(nil)).Elem(), GetZoneTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneTagArrayInput)(nil)).Elem(), GetZoneTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneTagsToMatchInput)(nil)).Elem(), GetZoneTagsToMatchArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetZoneTagsToMatchArrayInput)(nil)).Elem(), GetZoneTagsToMatchArray{})
	pulumi.RegisterOutputType(ZoneLinkOutput{})
	pulumi.RegisterOutputType(ZoneLinkArrayOutput{})
	pulumi.RegisterOutputType(ZoneTagOutput{})
	pulumi.RegisterOutputType(ZoneTagArrayOutput{})
	pulumi.RegisterOutputType(ZoneTagsToMatchOutput{})
	pulumi.RegisterOutputType(ZoneTagsToMatchArrayOutput{})
	pulumi.RegisterOutputType(GetZoneLinkOutput{})
	pulumi.RegisterOutputType(GetZoneLinkArrayOutput{})
	pulumi.RegisterOutputType(GetZoneTagOutput{})
	pulumi.RegisterOutputType(GetZoneTagArrayOutput{})
	pulumi.RegisterOutputType(GetZoneTagsToMatchOutput{})
	pulumi.RegisterOutputType(GetZoneTagsToMatchArrayOutput{})
}
