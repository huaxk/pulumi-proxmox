// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmox

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type QemuVM struct {
	pulumi.CustomResourceState

	AdditionalWait pulumi.IntPtrOutput    `pulumi:"additionalWait"`
	Agent          pulumi.IntPtrOutput    `pulumi:"agent"`
	Args           pulumi.StringPtrOutput `pulumi:"args"`
	Balloon        pulumi.IntPtrOutput    `pulumi:"balloon"`
	Bios           pulumi.StringPtrOutput `pulumi:"bios"`
	Boot           pulumi.StringPtrOutput `pulumi:"boot"`
	Bootdisk       pulumi.StringOutput    `pulumi:"bootdisk"`
	// Deprecated: Use `network.bridge` instead
	Bridge                pulumi.StringPtrOutput `pulumi:"bridge"`
	CiWait                pulumi.IntPtrOutput    `pulumi:"ciWait"`
	Cicustom              pulumi.StringPtrOutput `pulumi:"cicustom"`
	Cipassword            pulumi.StringPtrOutput `pulumi:"cipassword"`
	Ciuser                pulumi.StringPtrOutput `pulumi:"ciuser"`
	Clone                 pulumi.StringPtrOutput `pulumi:"clone"`
	CloneWait             pulumi.IntPtrOutput    `pulumi:"cloneWait"`
	CloudinitCdromStorage pulumi.StringPtrOutput `pulumi:"cloudinitCdromStorage"`
	Cores                 pulumi.IntPtrOutput    `pulumi:"cores"`
	Cpu                   pulumi.StringPtrOutput `pulumi:"cpu"`
	DefaultIpv4Address    pulumi.StringOutput    `pulumi:"defaultIpv4Address"`
	DefineConnectionInfo  pulumi.BoolPtrOutput   `pulumi:"defineConnectionInfo"`
	Desc                  pulumi.StringPtrOutput `pulumi:"desc"`
	// Deprecated: Use `disk.size` instead
	DiskGb                  pulumi.Float64PtrOutput `pulumi:"diskGb"`
	Disks                   QemuVMDiskArrayOutput   `pulumi:"disks"`
	ForceCreate             pulumi.BoolPtrOutput    `pulumi:"forceCreate"`
	ForceRecreateOnChangeOf pulumi.StringPtrOutput  `pulumi:"forceRecreateOnChangeOf"`
	FullClone               pulumi.BoolPtrOutput    `pulumi:"fullClone"`
	GuestAgentReadyTimeout  pulumi.IntPtrOutput     `pulumi:"guestAgentReadyTimeout"`
	Hastate                 pulumi.StringPtrOutput  `pulumi:"hastate"`
	Hotplug                 pulumi.StringPtrOutput  `pulumi:"hotplug"`
	Ipconfig0               pulumi.StringPtrOutput  `pulumi:"ipconfig0"`
	Ipconfig1               pulumi.StringPtrOutput  `pulumi:"ipconfig1"`
	Ipconfig2               pulumi.StringPtrOutput  `pulumi:"ipconfig2"`
	Iso                     pulumi.StringPtrOutput  `pulumi:"iso"`
	Kvm                     pulumi.BoolPtrOutput    `pulumi:"kvm"`
	// Deprecated: Use `network.macaddr` to access the auto generated MAC address
	Mac        pulumi.StringPtrOutput   `pulumi:"mac"`
	Memory     pulumi.IntPtrOutput      `pulumi:"memory"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Nameserver pulumi.StringOutput      `pulumi:"nameserver"`
	Networks   QemuVMNetworkArrayOutput `pulumi:"networks"`
	// Deprecated: Use `network` instead
	Nic             pulumi.StringPtrOutput `pulumi:"nic"`
	Numa            pulumi.BoolPtrOutput   `pulumi:"numa"`
	Onboot          pulumi.BoolPtrOutput   `pulumi:"onboot"`
	OsNetworkConfig pulumi.StringPtrOutput `pulumi:"osNetworkConfig"`
	OsType          pulumi.StringPtrOutput `pulumi:"osType"`
	Pool            pulumi.StringPtrOutput `pulumi:"pool"`
	Preprovision    pulumi.BoolPtrOutput   `pulumi:"preprovision"`
	QemuOs          pulumi.StringPtrOutput `pulumi:"qemuOs"`
	// Internal variable, true if any of the modified parameters require a reboot to take effect.
	RebootRequired pulumi.BoolOutput       `pulumi:"rebootRequired"`
	Scsihw         pulumi.StringOutput     `pulumi:"scsihw"`
	Searchdomain   pulumi.StringOutput     `pulumi:"searchdomain"`
	Serials        QemuVMSerialArrayOutput `pulumi:"serials"`
	Sockets        pulumi.IntPtrOutput     `pulumi:"sockets"`
	SshForwardIp   pulumi.StringPtrOutput  `pulumi:"sshForwardIp"`
	SshHost        pulumi.StringOutput     `pulumi:"sshHost"`
	SshPort        pulumi.StringOutput     `pulumi:"sshPort"`
	SshPrivateKey  pulumi.StringPtrOutput  `pulumi:"sshPrivateKey"`
	SshUser        pulumi.StringPtrOutput  `pulumi:"sshUser"`
	Sshkeys        pulumi.StringPtrOutput  `pulumi:"sshkeys"`
	// Deprecated: Use `disk.storage` instead
	Storage pulumi.StringPtrOutput `pulumi:"storage"`
	// Deprecated: Use `disk.type` instead
	StorageType pulumi.StringPtrOutput `pulumi:"storageType"`
	Tags        pulumi.StringPtrOutput `pulumi:"tags"`
	TargetNode  pulumi.StringOutput    `pulumi:"targetNode"`
	// Record unused disks in proxmox. This is intended to be read-only for now.
	UnusedDisks QemuVMUnusedDiskArrayOutput `pulumi:"unusedDisks"`
	Vcpus       pulumi.IntPtrOutput         `pulumi:"vcpus"`
	Vgas        QemuVMVgaArrayOutput        `pulumi:"vgas"`
	// Deprecated: Use `network.tag` instead
	Vlan pulumi.IntPtrOutput `pulumi:"vlan"`
	Vmid pulumi.IntOutput    `pulumi:"vmid"`
}

// NewQemuVM registers a new resource with the given unique name, arguments, and options.
func NewQemuVM(ctx *pulumi.Context,
	name string, args *QemuVMArgs, opts ...pulumi.ResourceOption) (*QemuVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetNode == nil {
		return nil, errors.New("invalid value for required argument 'TargetNode'")
	}
	var resource QemuVM
	err := ctx.RegisterResource("proxmox:index/qemuVM:QemuVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQemuVM gets an existing QemuVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQemuVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QemuVMState, opts ...pulumi.ResourceOption) (*QemuVM, error) {
	var resource QemuVM
	err := ctx.ReadResource("proxmox:index/qemuVM:QemuVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QemuVM resources.
type qemuVMState struct {
	AdditionalWait *int    `pulumi:"additionalWait"`
	Agent          *int    `pulumi:"agent"`
	Args           *string `pulumi:"args"`
	Balloon        *int    `pulumi:"balloon"`
	Bios           *string `pulumi:"bios"`
	Boot           *string `pulumi:"boot"`
	Bootdisk       *string `pulumi:"bootdisk"`
	// Deprecated: Use `network.bridge` instead
	Bridge                *string `pulumi:"bridge"`
	CiWait                *int    `pulumi:"ciWait"`
	Cicustom              *string `pulumi:"cicustom"`
	Cipassword            *string `pulumi:"cipassword"`
	Ciuser                *string `pulumi:"ciuser"`
	Clone                 *string `pulumi:"clone"`
	CloneWait             *int    `pulumi:"cloneWait"`
	CloudinitCdromStorage *string `pulumi:"cloudinitCdromStorage"`
	Cores                 *int    `pulumi:"cores"`
	Cpu                   *string `pulumi:"cpu"`
	DefaultIpv4Address    *string `pulumi:"defaultIpv4Address"`
	DefineConnectionInfo  *bool   `pulumi:"defineConnectionInfo"`
	Desc                  *string `pulumi:"desc"`
	// Deprecated: Use `disk.size` instead
	DiskGb                  *float64     `pulumi:"diskGb"`
	Disks                   []QemuVMDisk `pulumi:"disks"`
	ForceCreate             *bool        `pulumi:"forceCreate"`
	ForceRecreateOnChangeOf *string      `pulumi:"forceRecreateOnChangeOf"`
	FullClone               *bool        `pulumi:"fullClone"`
	GuestAgentReadyTimeout  *int         `pulumi:"guestAgentReadyTimeout"`
	Hastate                 *string      `pulumi:"hastate"`
	Hotplug                 *string      `pulumi:"hotplug"`
	Ipconfig0               *string      `pulumi:"ipconfig0"`
	Ipconfig1               *string      `pulumi:"ipconfig1"`
	Ipconfig2               *string      `pulumi:"ipconfig2"`
	Iso                     *string      `pulumi:"iso"`
	Kvm                     *bool        `pulumi:"kvm"`
	// Deprecated: Use `network.macaddr` to access the auto generated MAC address
	Mac        *string         `pulumi:"mac"`
	Memory     *int            `pulumi:"memory"`
	Name       *string         `pulumi:"name"`
	Nameserver *string         `pulumi:"nameserver"`
	Networks   []QemuVMNetwork `pulumi:"networks"`
	// Deprecated: Use `network` instead
	Nic             *string `pulumi:"nic"`
	Numa            *bool   `pulumi:"numa"`
	Onboot          *bool   `pulumi:"onboot"`
	OsNetworkConfig *string `pulumi:"osNetworkConfig"`
	OsType          *string `pulumi:"osType"`
	Pool            *string `pulumi:"pool"`
	Preprovision    *bool   `pulumi:"preprovision"`
	QemuOs          *string `pulumi:"qemuOs"`
	// Internal variable, true if any of the modified parameters require a reboot to take effect.
	RebootRequired *bool          `pulumi:"rebootRequired"`
	Scsihw         *string        `pulumi:"scsihw"`
	Searchdomain   *string        `pulumi:"searchdomain"`
	Serials        []QemuVMSerial `pulumi:"serials"`
	Sockets        *int           `pulumi:"sockets"`
	SshForwardIp   *string        `pulumi:"sshForwardIp"`
	SshHost        *string        `pulumi:"sshHost"`
	SshPort        *string        `pulumi:"sshPort"`
	SshPrivateKey  *string        `pulumi:"sshPrivateKey"`
	SshUser        *string        `pulumi:"sshUser"`
	Sshkeys        *string        `pulumi:"sshkeys"`
	// Deprecated: Use `disk.storage` instead
	Storage *string `pulumi:"storage"`
	// Deprecated: Use `disk.type` instead
	StorageType *string `pulumi:"storageType"`
	Tags        *string `pulumi:"tags"`
	TargetNode  *string `pulumi:"targetNode"`
	// Record unused disks in proxmox. This is intended to be read-only for now.
	UnusedDisks []QemuVMUnusedDisk `pulumi:"unusedDisks"`
	Vcpus       *int               `pulumi:"vcpus"`
	Vgas        []QemuVMVga        `pulumi:"vgas"`
	// Deprecated: Use `network.tag` instead
	Vlan *int `pulumi:"vlan"`
	Vmid *int `pulumi:"vmid"`
}

type QemuVMState struct {
	AdditionalWait pulumi.IntPtrInput
	Agent          pulumi.IntPtrInput
	Args           pulumi.StringPtrInput
	Balloon        pulumi.IntPtrInput
	Bios           pulumi.StringPtrInput
	Boot           pulumi.StringPtrInput
	Bootdisk       pulumi.StringPtrInput
	// Deprecated: Use `network.bridge` instead
	Bridge                pulumi.StringPtrInput
	CiWait                pulumi.IntPtrInput
	Cicustom              pulumi.StringPtrInput
	Cipassword            pulumi.StringPtrInput
	Ciuser                pulumi.StringPtrInput
	Clone                 pulumi.StringPtrInput
	CloneWait             pulumi.IntPtrInput
	CloudinitCdromStorage pulumi.StringPtrInput
	Cores                 pulumi.IntPtrInput
	Cpu                   pulumi.StringPtrInput
	DefaultIpv4Address    pulumi.StringPtrInput
	DefineConnectionInfo  pulumi.BoolPtrInput
	Desc                  pulumi.StringPtrInput
	// Deprecated: Use `disk.size` instead
	DiskGb                  pulumi.Float64PtrInput
	Disks                   QemuVMDiskArrayInput
	ForceCreate             pulumi.BoolPtrInput
	ForceRecreateOnChangeOf pulumi.StringPtrInput
	FullClone               pulumi.BoolPtrInput
	GuestAgentReadyTimeout  pulumi.IntPtrInput
	Hastate                 pulumi.StringPtrInput
	Hotplug                 pulumi.StringPtrInput
	Ipconfig0               pulumi.StringPtrInput
	Ipconfig1               pulumi.StringPtrInput
	Ipconfig2               pulumi.StringPtrInput
	Iso                     pulumi.StringPtrInput
	Kvm                     pulumi.BoolPtrInput
	// Deprecated: Use `network.macaddr` to access the auto generated MAC address
	Mac        pulumi.StringPtrInput
	Memory     pulumi.IntPtrInput
	Name       pulumi.StringPtrInput
	Nameserver pulumi.StringPtrInput
	Networks   QemuVMNetworkArrayInput
	// Deprecated: Use `network` instead
	Nic             pulumi.StringPtrInput
	Numa            pulumi.BoolPtrInput
	Onboot          pulumi.BoolPtrInput
	OsNetworkConfig pulumi.StringPtrInput
	OsType          pulumi.StringPtrInput
	Pool            pulumi.StringPtrInput
	Preprovision    pulumi.BoolPtrInput
	QemuOs          pulumi.StringPtrInput
	// Internal variable, true if any of the modified parameters require a reboot to take effect.
	RebootRequired pulumi.BoolPtrInput
	Scsihw         pulumi.StringPtrInput
	Searchdomain   pulumi.StringPtrInput
	Serials        QemuVMSerialArrayInput
	Sockets        pulumi.IntPtrInput
	SshForwardIp   pulumi.StringPtrInput
	SshHost        pulumi.StringPtrInput
	SshPort        pulumi.StringPtrInput
	SshPrivateKey  pulumi.StringPtrInput
	SshUser        pulumi.StringPtrInput
	Sshkeys        pulumi.StringPtrInput
	// Deprecated: Use `disk.storage` instead
	Storage pulumi.StringPtrInput
	// Deprecated: Use `disk.type` instead
	StorageType pulumi.StringPtrInput
	Tags        pulumi.StringPtrInput
	TargetNode  pulumi.StringPtrInput
	// Record unused disks in proxmox. This is intended to be read-only for now.
	UnusedDisks QemuVMUnusedDiskArrayInput
	Vcpus       pulumi.IntPtrInput
	Vgas        QemuVMVgaArrayInput
	// Deprecated: Use `network.tag` instead
	Vlan pulumi.IntPtrInput
	Vmid pulumi.IntPtrInput
}

func (QemuVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*qemuVMState)(nil)).Elem()
}

type qemuVMArgs struct {
	AdditionalWait *int    `pulumi:"additionalWait"`
	Agent          *int    `pulumi:"agent"`
	Args           *string `pulumi:"args"`
	Balloon        *int    `pulumi:"balloon"`
	Bios           *string `pulumi:"bios"`
	Boot           *string `pulumi:"boot"`
	Bootdisk       *string `pulumi:"bootdisk"`
	// Deprecated: Use `network.bridge` instead
	Bridge                *string `pulumi:"bridge"`
	CiWait                *int    `pulumi:"ciWait"`
	Cicustom              *string `pulumi:"cicustom"`
	Cipassword            *string `pulumi:"cipassword"`
	Ciuser                *string `pulumi:"ciuser"`
	Clone                 *string `pulumi:"clone"`
	CloneWait             *int    `pulumi:"cloneWait"`
	CloudinitCdromStorage *string `pulumi:"cloudinitCdromStorage"`
	Cores                 *int    `pulumi:"cores"`
	Cpu                   *string `pulumi:"cpu"`
	DefineConnectionInfo  *bool   `pulumi:"defineConnectionInfo"`
	Desc                  *string `pulumi:"desc"`
	// Deprecated: Use `disk.size` instead
	DiskGb                  *float64     `pulumi:"diskGb"`
	Disks                   []QemuVMDisk `pulumi:"disks"`
	ForceCreate             *bool        `pulumi:"forceCreate"`
	ForceRecreateOnChangeOf *string      `pulumi:"forceRecreateOnChangeOf"`
	FullClone               *bool        `pulumi:"fullClone"`
	GuestAgentReadyTimeout  *int         `pulumi:"guestAgentReadyTimeout"`
	Hastate                 *string      `pulumi:"hastate"`
	Hotplug                 *string      `pulumi:"hotplug"`
	Ipconfig0               *string      `pulumi:"ipconfig0"`
	Ipconfig1               *string      `pulumi:"ipconfig1"`
	Ipconfig2               *string      `pulumi:"ipconfig2"`
	Iso                     *string      `pulumi:"iso"`
	Kvm                     *bool        `pulumi:"kvm"`
	// Deprecated: Use `network.macaddr` to access the auto generated MAC address
	Mac        *string         `pulumi:"mac"`
	Memory     *int            `pulumi:"memory"`
	Name       *string         `pulumi:"name"`
	Nameserver *string         `pulumi:"nameserver"`
	Networks   []QemuVMNetwork `pulumi:"networks"`
	// Deprecated: Use `network` instead
	Nic             *string        `pulumi:"nic"`
	Numa            *bool          `pulumi:"numa"`
	Onboot          *bool          `pulumi:"onboot"`
	OsNetworkConfig *string        `pulumi:"osNetworkConfig"`
	OsType          *string        `pulumi:"osType"`
	Pool            *string        `pulumi:"pool"`
	Preprovision    *bool          `pulumi:"preprovision"`
	QemuOs          *string        `pulumi:"qemuOs"`
	Scsihw          *string        `pulumi:"scsihw"`
	Searchdomain    *string        `pulumi:"searchdomain"`
	Serials         []QemuVMSerial `pulumi:"serials"`
	Sockets         *int           `pulumi:"sockets"`
	SshForwardIp    *string        `pulumi:"sshForwardIp"`
	SshPrivateKey   *string        `pulumi:"sshPrivateKey"`
	SshUser         *string        `pulumi:"sshUser"`
	Sshkeys         *string        `pulumi:"sshkeys"`
	// Deprecated: Use `disk.storage` instead
	Storage *string `pulumi:"storage"`
	// Deprecated: Use `disk.type` instead
	StorageType *string     `pulumi:"storageType"`
	Tags        *string     `pulumi:"tags"`
	TargetNode  string      `pulumi:"targetNode"`
	Vcpus       *int        `pulumi:"vcpus"`
	Vgas        []QemuVMVga `pulumi:"vgas"`
	// Deprecated: Use `network.tag` instead
	Vlan *int `pulumi:"vlan"`
	Vmid *int `pulumi:"vmid"`
}

// The set of arguments for constructing a QemuVM resource.
type QemuVMArgs struct {
	AdditionalWait pulumi.IntPtrInput
	Agent          pulumi.IntPtrInput
	Args           pulumi.StringPtrInput
	Balloon        pulumi.IntPtrInput
	Bios           pulumi.StringPtrInput
	Boot           pulumi.StringPtrInput
	Bootdisk       pulumi.StringPtrInput
	// Deprecated: Use `network.bridge` instead
	Bridge                pulumi.StringPtrInput
	CiWait                pulumi.IntPtrInput
	Cicustom              pulumi.StringPtrInput
	Cipassword            pulumi.StringPtrInput
	Ciuser                pulumi.StringPtrInput
	Clone                 pulumi.StringPtrInput
	CloneWait             pulumi.IntPtrInput
	CloudinitCdromStorage pulumi.StringPtrInput
	Cores                 pulumi.IntPtrInput
	Cpu                   pulumi.StringPtrInput
	DefineConnectionInfo  pulumi.BoolPtrInput
	Desc                  pulumi.StringPtrInput
	// Deprecated: Use `disk.size` instead
	DiskGb                  pulumi.Float64PtrInput
	Disks                   QemuVMDiskArrayInput
	ForceCreate             pulumi.BoolPtrInput
	ForceRecreateOnChangeOf pulumi.StringPtrInput
	FullClone               pulumi.BoolPtrInput
	GuestAgentReadyTimeout  pulumi.IntPtrInput
	Hastate                 pulumi.StringPtrInput
	Hotplug                 pulumi.StringPtrInput
	Ipconfig0               pulumi.StringPtrInput
	Ipconfig1               pulumi.StringPtrInput
	Ipconfig2               pulumi.StringPtrInput
	Iso                     pulumi.StringPtrInput
	Kvm                     pulumi.BoolPtrInput
	// Deprecated: Use `network.macaddr` to access the auto generated MAC address
	Mac        pulumi.StringPtrInput
	Memory     pulumi.IntPtrInput
	Name       pulumi.StringPtrInput
	Nameserver pulumi.StringPtrInput
	Networks   QemuVMNetworkArrayInput
	// Deprecated: Use `network` instead
	Nic             pulumi.StringPtrInput
	Numa            pulumi.BoolPtrInput
	Onboot          pulumi.BoolPtrInput
	OsNetworkConfig pulumi.StringPtrInput
	OsType          pulumi.StringPtrInput
	Pool            pulumi.StringPtrInput
	Preprovision    pulumi.BoolPtrInput
	QemuOs          pulumi.StringPtrInput
	Scsihw          pulumi.StringPtrInput
	Searchdomain    pulumi.StringPtrInput
	Serials         QemuVMSerialArrayInput
	Sockets         pulumi.IntPtrInput
	SshForwardIp    pulumi.StringPtrInput
	SshPrivateKey   pulumi.StringPtrInput
	SshUser         pulumi.StringPtrInput
	Sshkeys         pulumi.StringPtrInput
	// Deprecated: Use `disk.storage` instead
	Storage pulumi.StringPtrInput
	// Deprecated: Use `disk.type` instead
	StorageType pulumi.StringPtrInput
	Tags        pulumi.StringPtrInput
	TargetNode  pulumi.StringInput
	Vcpus       pulumi.IntPtrInput
	Vgas        QemuVMVgaArrayInput
	// Deprecated: Use `network.tag` instead
	Vlan pulumi.IntPtrInput
	Vmid pulumi.IntPtrInput
}

func (QemuVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*qemuVMArgs)(nil)).Elem()
}

type QemuVMInput interface {
	pulumi.Input

	ToQemuVMOutput() QemuVMOutput
	ToQemuVMOutputWithContext(ctx context.Context) QemuVMOutput
}

func (*QemuVM) ElementType() reflect.Type {
	return reflect.TypeOf((*QemuVM)(nil))
}

func (i *QemuVM) ToQemuVMOutput() QemuVMOutput {
	return i.ToQemuVMOutputWithContext(context.Background())
}

func (i *QemuVM) ToQemuVMOutputWithContext(ctx context.Context) QemuVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QemuVMOutput)
}

func (i *QemuVM) ToQemuVMPtrOutput() QemuVMPtrOutput {
	return i.ToQemuVMPtrOutputWithContext(context.Background())
}

func (i *QemuVM) ToQemuVMPtrOutputWithContext(ctx context.Context) QemuVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QemuVMPtrOutput)
}

type QemuVMPtrInput interface {
	pulumi.Input

	ToQemuVMPtrOutput() QemuVMPtrOutput
	ToQemuVMPtrOutputWithContext(ctx context.Context) QemuVMPtrOutput
}

type qemuVMPtrType QemuVMArgs

func (*qemuVMPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QemuVM)(nil))
}

func (i *qemuVMPtrType) ToQemuVMPtrOutput() QemuVMPtrOutput {
	return i.ToQemuVMPtrOutputWithContext(context.Background())
}

func (i *qemuVMPtrType) ToQemuVMPtrOutputWithContext(ctx context.Context) QemuVMPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QemuVMPtrOutput)
}

// QemuVMArrayInput is an input type that accepts QemuVMArray and QemuVMArrayOutput values.
// You can construct a concrete instance of `QemuVMArrayInput` via:
//
//          QemuVMArray{ QemuVMArgs{...} }
type QemuVMArrayInput interface {
	pulumi.Input

	ToQemuVMArrayOutput() QemuVMArrayOutput
	ToQemuVMArrayOutputWithContext(context.Context) QemuVMArrayOutput
}

type QemuVMArray []QemuVMInput

func (QemuVMArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*QemuVM)(nil))
}

func (i QemuVMArray) ToQemuVMArrayOutput() QemuVMArrayOutput {
	return i.ToQemuVMArrayOutputWithContext(context.Background())
}

func (i QemuVMArray) ToQemuVMArrayOutputWithContext(ctx context.Context) QemuVMArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QemuVMArrayOutput)
}

// QemuVMMapInput is an input type that accepts QemuVMMap and QemuVMMapOutput values.
// You can construct a concrete instance of `QemuVMMapInput` via:
//
//          QemuVMMap{ "key": QemuVMArgs{...} }
type QemuVMMapInput interface {
	pulumi.Input

	ToQemuVMMapOutput() QemuVMMapOutput
	ToQemuVMMapOutputWithContext(context.Context) QemuVMMapOutput
}

type QemuVMMap map[string]QemuVMInput

func (QemuVMMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*QemuVM)(nil))
}

func (i QemuVMMap) ToQemuVMMapOutput() QemuVMMapOutput {
	return i.ToQemuVMMapOutputWithContext(context.Background())
}

func (i QemuVMMap) ToQemuVMMapOutputWithContext(ctx context.Context) QemuVMMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QemuVMMapOutput)
}

type QemuVMOutput struct {
	*pulumi.OutputState
}

func (QemuVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QemuVM)(nil))
}

func (o QemuVMOutput) ToQemuVMOutput() QemuVMOutput {
	return o
}

func (o QemuVMOutput) ToQemuVMOutputWithContext(ctx context.Context) QemuVMOutput {
	return o
}

func (o QemuVMOutput) ToQemuVMPtrOutput() QemuVMPtrOutput {
	return o.ToQemuVMPtrOutputWithContext(context.Background())
}

func (o QemuVMOutput) ToQemuVMPtrOutputWithContext(ctx context.Context) QemuVMPtrOutput {
	return o.ApplyT(func(v QemuVM) *QemuVM {
		return &v
	}).(QemuVMPtrOutput)
}

type QemuVMPtrOutput struct {
	*pulumi.OutputState
}

func (QemuVMPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QemuVM)(nil))
}

func (o QemuVMPtrOutput) ToQemuVMPtrOutput() QemuVMPtrOutput {
	return o
}

func (o QemuVMPtrOutput) ToQemuVMPtrOutputWithContext(ctx context.Context) QemuVMPtrOutput {
	return o
}

type QemuVMArrayOutput struct{ *pulumi.OutputState }

func (QemuVMArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QemuVM)(nil))
}

func (o QemuVMArrayOutput) ToQemuVMArrayOutput() QemuVMArrayOutput {
	return o
}

func (o QemuVMArrayOutput) ToQemuVMArrayOutputWithContext(ctx context.Context) QemuVMArrayOutput {
	return o
}

func (o QemuVMArrayOutput) Index(i pulumi.IntInput) QemuVMOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QemuVM {
		return vs[0].([]QemuVM)[vs[1].(int)]
	}).(QemuVMOutput)
}

type QemuVMMapOutput struct{ *pulumi.OutputState }

func (QemuVMMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QemuVM)(nil))
}

func (o QemuVMMapOutput) ToQemuVMMapOutput() QemuVMMapOutput {
	return o
}

func (o QemuVMMapOutput) ToQemuVMMapOutputWithContext(ctx context.Context) QemuVMMapOutput {
	return o
}

func (o QemuVMMapOutput) MapIndex(k pulumi.StringInput) QemuVMOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QemuVM {
		return vs[0].(map[string]QemuVM)[vs[1].(string)]
	}).(QemuVMOutput)
}

func init() {
	pulumi.RegisterOutputType(QemuVMOutput{})
	pulumi.RegisterOutputType(QemuVMPtrOutput{})
	pulumi.RegisterOutputType(QemuVMArrayOutput{})
	pulumi.RegisterOutputType(QemuVMMapOutput{})
}
