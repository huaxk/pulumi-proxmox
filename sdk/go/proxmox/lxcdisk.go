// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmox

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LXCDisk struct {
	pulumi.CustomResourceState

	Acl          pulumi.BoolPtrOutput         `pulumi:"acl"`
	Backup       pulumi.BoolPtrOutput         `pulumi:"backup"`
	Container    pulumi.StringOutput          `pulumi:"container"`
	Mountoptions LXCDiskMountoptionsPtrOutput `pulumi:"mountoptions"`
	Mp           pulumi.StringOutput          `pulumi:"mp"`
	Quota        pulumi.BoolPtrOutput         `pulumi:"quota"`
	Replicate    pulumi.BoolPtrOutput         `pulumi:"replicate"`
	Shared       pulumi.BoolPtrOutput         `pulumi:"shared"`
	Size         pulumi.StringOutput          `pulumi:"size"`
	Slot         pulumi.IntOutput             `pulumi:"slot"`
	Storage      pulumi.StringOutput          `pulumi:"storage"`
	Volume       pulumi.StringOutput          `pulumi:"volume"`
}

// NewLXCDisk registers a new resource with the given unique name, arguments, and options.
func NewLXCDisk(ctx *pulumi.Context,
	name string, args *LXCDiskArgs, opts ...pulumi.ResourceOption) (*LXCDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Container == nil {
		return nil, errors.New("invalid value for required argument 'Container'")
	}
	if args.Mp == nil {
		return nil, errors.New("invalid value for required argument 'Mp'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	var resource LXCDisk
	err := ctx.RegisterResource("proxmox:index/lXCDisk:LXCDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLXCDisk gets an existing LXCDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLXCDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LXCDiskState, opts ...pulumi.ResourceOption) (*LXCDisk, error) {
	var resource LXCDisk
	err := ctx.ReadResource("proxmox:index/lXCDisk:LXCDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LXCDisk resources.
type lxcdiskState struct {
	Acl          *bool                `pulumi:"acl"`
	Backup       *bool                `pulumi:"backup"`
	Container    *string              `pulumi:"container"`
	Mountoptions *LXCDiskMountoptions `pulumi:"mountoptions"`
	Mp           *string              `pulumi:"mp"`
	Quota        *bool                `pulumi:"quota"`
	Replicate    *bool                `pulumi:"replicate"`
	Shared       *bool                `pulumi:"shared"`
	Size         *string              `pulumi:"size"`
	Slot         *int                 `pulumi:"slot"`
	Storage      *string              `pulumi:"storage"`
	Volume       *string              `pulumi:"volume"`
}

type LXCDiskState struct {
	Acl          pulumi.BoolPtrInput
	Backup       pulumi.BoolPtrInput
	Container    pulumi.StringPtrInput
	Mountoptions LXCDiskMountoptionsPtrInput
	Mp           pulumi.StringPtrInput
	Quota        pulumi.BoolPtrInput
	Replicate    pulumi.BoolPtrInput
	Shared       pulumi.BoolPtrInput
	Size         pulumi.StringPtrInput
	Slot         pulumi.IntPtrInput
	Storage      pulumi.StringPtrInput
	Volume       pulumi.StringPtrInput
}

func (LXCDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*lxcdiskState)(nil)).Elem()
}

type lxcdiskArgs struct {
	Acl          *bool                `pulumi:"acl"`
	Backup       *bool                `pulumi:"backup"`
	Container    string               `pulumi:"container"`
	Mountoptions *LXCDiskMountoptions `pulumi:"mountoptions"`
	Mp           string               `pulumi:"mp"`
	Quota        *bool                `pulumi:"quota"`
	Replicate    *bool                `pulumi:"replicate"`
	Shared       *bool                `pulumi:"shared"`
	Size         string               `pulumi:"size"`
	Slot         int                  `pulumi:"slot"`
	Storage      string               `pulumi:"storage"`
	Volume       *string              `pulumi:"volume"`
}

// The set of arguments for constructing a LXCDisk resource.
type LXCDiskArgs struct {
	Acl          pulumi.BoolPtrInput
	Backup       pulumi.BoolPtrInput
	Container    pulumi.StringInput
	Mountoptions LXCDiskMountoptionsPtrInput
	Mp           pulumi.StringInput
	Quota        pulumi.BoolPtrInput
	Replicate    pulumi.BoolPtrInput
	Shared       pulumi.BoolPtrInput
	Size         pulumi.StringInput
	Slot         pulumi.IntInput
	Storage      pulumi.StringInput
	Volume       pulumi.StringPtrInput
}

func (LXCDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lxcdiskArgs)(nil)).Elem()
}

type LXCDiskInput interface {
	pulumi.Input

	ToLXCDiskOutput() LXCDiskOutput
	ToLXCDiskOutputWithContext(ctx context.Context) LXCDiskOutput
}

func (*LXCDisk) ElementType() reflect.Type {
	return reflect.TypeOf((*LXCDisk)(nil))
}

func (i *LXCDisk) ToLXCDiskOutput() LXCDiskOutput {
	return i.ToLXCDiskOutputWithContext(context.Background())
}

func (i *LXCDisk) ToLXCDiskOutputWithContext(ctx context.Context) LXCDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCDiskOutput)
}

func (i *LXCDisk) ToLXCDiskPtrOutput() LXCDiskPtrOutput {
	return i.ToLXCDiskPtrOutputWithContext(context.Background())
}

func (i *LXCDisk) ToLXCDiskPtrOutputWithContext(ctx context.Context) LXCDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCDiskPtrOutput)
}

type LXCDiskPtrInput interface {
	pulumi.Input

	ToLXCDiskPtrOutput() LXCDiskPtrOutput
	ToLXCDiskPtrOutputWithContext(ctx context.Context) LXCDiskPtrOutput
}

type lxcdiskPtrType LXCDiskArgs

func (*lxcdiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LXCDisk)(nil))
}

func (i *lxcdiskPtrType) ToLXCDiskPtrOutput() LXCDiskPtrOutput {
	return i.ToLXCDiskPtrOutputWithContext(context.Background())
}

func (i *lxcdiskPtrType) ToLXCDiskPtrOutputWithContext(ctx context.Context) LXCDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCDiskPtrOutput)
}

// LXCDiskArrayInput is an input type that accepts LXCDiskArray and LXCDiskArrayOutput values.
// You can construct a concrete instance of `LXCDiskArrayInput` via:
//
//          LXCDiskArray{ LXCDiskArgs{...} }
type LXCDiskArrayInput interface {
	pulumi.Input

	ToLXCDiskArrayOutput() LXCDiskArrayOutput
	ToLXCDiskArrayOutputWithContext(context.Context) LXCDiskArrayOutput
}

type LXCDiskArray []LXCDiskInput

func (LXCDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LXCDisk)(nil))
}

func (i LXCDiskArray) ToLXCDiskArrayOutput() LXCDiskArrayOutput {
	return i.ToLXCDiskArrayOutputWithContext(context.Background())
}

func (i LXCDiskArray) ToLXCDiskArrayOutputWithContext(ctx context.Context) LXCDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCDiskArrayOutput)
}

// LXCDiskMapInput is an input type that accepts LXCDiskMap and LXCDiskMapOutput values.
// You can construct a concrete instance of `LXCDiskMapInput` via:
//
//          LXCDiskMap{ "key": LXCDiskArgs{...} }
type LXCDiskMapInput interface {
	pulumi.Input

	ToLXCDiskMapOutput() LXCDiskMapOutput
	ToLXCDiskMapOutputWithContext(context.Context) LXCDiskMapOutput
}

type LXCDiskMap map[string]LXCDiskInput

func (LXCDiskMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LXCDisk)(nil))
}

func (i LXCDiskMap) ToLXCDiskMapOutput() LXCDiskMapOutput {
	return i.ToLXCDiskMapOutputWithContext(context.Background())
}

func (i LXCDiskMap) ToLXCDiskMapOutputWithContext(ctx context.Context) LXCDiskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LXCDiskMapOutput)
}

type LXCDiskOutput struct {
	*pulumi.OutputState
}

func (LXCDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LXCDisk)(nil))
}

func (o LXCDiskOutput) ToLXCDiskOutput() LXCDiskOutput {
	return o
}

func (o LXCDiskOutput) ToLXCDiskOutputWithContext(ctx context.Context) LXCDiskOutput {
	return o
}

func (o LXCDiskOutput) ToLXCDiskPtrOutput() LXCDiskPtrOutput {
	return o.ToLXCDiskPtrOutputWithContext(context.Background())
}

func (o LXCDiskOutput) ToLXCDiskPtrOutputWithContext(ctx context.Context) LXCDiskPtrOutput {
	return o.ApplyT(func(v LXCDisk) *LXCDisk {
		return &v
	}).(LXCDiskPtrOutput)
}

type LXCDiskPtrOutput struct {
	*pulumi.OutputState
}

func (LXCDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LXCDisk)(nil))
}

func (o LXCDiskPtrOutput) ToLXCDiskPtrOutput() LXCDiskPtrOutput {
	return o
}

func (o LXCDiskPtrOutput) ToLXCDiskPtrOutputWithContext(ctx context.Context) LXCDiskPtrOutput {
	return o
}

type LXCDiskArrayOutput struct{ *pulumi.OutputState }

func (LXCDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LXCDisk)(nil))
}

func (o LXCDiskArrayOutput) ToLXCDiskArrayOutput() LXCDiskArrayOutput {
	return o
}

func (o LXCDiskArrayOutput) ToLXCDiskArrayOutputWithContext(ctx context.Context) LXCDiskArrayOutput {
	return o
}

func (o LXCDiskArrayOutput) Index(i pulumi.IntInput) LXCDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LXCDisk {
		return vs[0].([]LXCDisk)[vs[1].(int)]
	}).(LXCDiskOutput)
}

type LXCDiskMapOutput struct{ *pulumi.OutputState }

func (LXCDiskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LXCDisk)(nil))
}

func (o LXCDiskMapOutput) ToLXCDiskMapOutput() LXCDiskMapOutput {
	return o
}

func (o LXCDiskMapOutput) ToLXCDiskMapOutputWithContext(ctx context.Context) LXCDiskMapOutput {
	return o
}

func (o LXCDiskMapOutput) MapIndex(k pulumi.StringInput) LXCDiskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LXCDisk {
		return vs[0].(map[string]LXCDisk)[vs[1].(string)]
	}).(LXCDiskOutput)
}

func init() {
	pulumi.RegisterOutputType(LXCDiskOutput{})
	pulumi.RegisterOutputType(LXCDiskPtrOutput{})
	pulumi.RegisterOutputType(LXCDiskArrayOutput{})
	pulumi.RegisterOutputType(LXCDiskMapOutput{})
}