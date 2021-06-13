# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LXCDiskArgs', 'LXCDisk']

@pulumi.input_type
class LXCDiskArgs:
    def __init__(__self__, *,
                 container: pulumi.Input[str],
                 mp: pulumi.Input[str],
                 size: pulumi.Input[str],
                 slot: pulumi.Input[int],
                 storage: pulumi.Input[str],
                 acl: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 mountoptions: Optional[pulumi.Input['LXCDiskMountoptionsArgs']] = None,
                 quota: Optional[pulumi.Input[bool]] = None,
                 replicate: Optional[pulumi.Input[bool]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 volume: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LXCDisk resource.
        """
        pulumi.set(__self__, "container", container)
        pulumi.set(__self__, "mp", mp)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "slot", slot)
        pulumi.set(__self__, "storage", storage)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if mountoptions is not None:
            pulumi.set(__self__, "mountoptions", mountoptions)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)
        if replicate is not None:
            pulumi.set(__self__, "replicate", replicate)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if volume is not None:
            pulumi.set(__self__, "volume", volume)

    @property
    @pulumi.getter
    def container(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: pulumi.Input[str]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def mp(self) -> pulumi.Input[str]:
        return pulumi.get(self, "mp")

    @mp.setter
    def mp(self, value: pulumi.Input[str]):
        pulumi.set(self, "mp", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[str]:
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[str]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def slot(self) -> pulumi.Input[int]:
        return pulumi.get(self, "slot")

    @slot.setter
    def slot(self, value: pulumi.Input[int]):
        pulumi.set(self, "slot", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input[str]:
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def mountoptions(self) -> Optional[pulumi.Input['LXCDiskMountoptionsArgs']]:
        return pulumi.get(self, "mountoptions")

    @mountoptions.setter
    def mountoptions(self, value: Optional[pulumi.Input['LXCDiskMountoptionsArgs']]):
        pulumi.set(self, "mountoptions", value)

    @property
    @pulumi.getter
    def quota(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "quota")

    @quota.setter
    def quota(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "quota", value)

    @property
    @pulumi.getter
    def replicate(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "replicate")

    @replicate.setter
    def replicate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replicate", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def volume(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "volume")

    @volume.setter
    def volume(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume", value)


@pulumi.input_type
class _LXCDiskState:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 mountoptions: Optional[pulumi.Input['LXCDiskMountoptionsArgs']] = None,
                 mp: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[bool]] = None,
                 replicate: Optional[pulumi.Input[bool]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 slot: Optional[pulumi.Input[int]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 volume: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LXCDisk resources.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if container is not None:
            pulumi.set(__self__, "container", container)
        if mountoptions is not None:
            pulumi.set(__self__, "mountoptions", mountoptions)
        if mp is not None:
            pulumi.set(__self__, "mp", mp)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)
        if replicate is not None:
            pulumi.set(__self__, "replicate", replicate)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if slot is not None:
            pulumi.set(__self__, "slot", slot)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)
        if volume is not None:
            pulumi.set(__self__, "volume", volume)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def mountoptions(self) -> Optional[pulumi.Input['LXCDiskMountoptionsArgs']]:
        return pulumi.get(self, "mountoptions")

    @mountoptions.setter
    def mountoptions(self, value: Optional[pulumi.Input['LXCDiskMountoptionsArgs']]):
        pulumi.set(self, "mountoptions", value)

    @property
    @pulumi.getter
    def mp(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mp")

    @mp.setter
    def mp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mp", value)

    @property
    @pulumi.getter
    def quota(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "quota")

    @quota.setter
    def quota(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "quota", value)

    @property
    @pulumi.getter
    def replicate(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "replicate")

    @replicate.setter
    def replicate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "replicate", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def slot(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "slot")

    @slot.setter
    def slot(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "slot", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter
    def volume(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "volume")

    @volume.setter
    def volume(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume", value)


class LXCDisk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 mountoptions: Optional[pulumi.Input[pulumi.InputType['LXCDiskMountoptionsArgs']]] = None,
                 mp: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[bool]] = None,
                 replicate: Optional[pulumi.Input[bool]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 slot: Optional[pulumi.Input[int]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 volume: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LXCDisk resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LXCDiskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LXCDisk resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LXCDiskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LXCDiskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 container: Optional[pulumi.Input[str]] = None,
                 mountoptions: Optional[pulumi.Input[pulumi.InputType['LXCDiskMountoptionsArgs']]] = None,
                 mp: Optional[pulumi.Input[str]] = None,
                 quota: Optional[pulumi.Input[bool]] = None,
                 replicate: Optional[pulumi.Input[bool]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 slot: Optional[pulumi.Input[int]] = None,
                 storage: Optional[pulumi.Input[str]] = None,
                 volume: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LXCDiskArgs.__new__(LXCDiskArgs)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["backup"] = backup
            if container is None and not opts.urn:
                raise TypeError("Missing required property 'container'")
            __props__.__dict__["container"] = container
            __props__.__dict__["mountoptions"] = mountoptions
            if mp is None and not opts.urn:
                raise TypeError("Missing required property 'mp'")
            __props__.__dict__["mp"] = mp
            __props__.__dict__["quota"] = quota
            __props__.__dict__["replicate"] = replicate
            __props__.__dict__["shared"] = shared
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            if slot is None and not opts.urn:
                raise TypeError("Missing required property 'slot'")
            __props__.__dict__["slot"] = slot
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__.__dict__["storage"] = storage
            __props__.__dict__["volume"] = volume
        super(LXCDisk, __self__).__init__(
            'proxmox:index/lXCDisk:LXCDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[bool]] = None,
            backup: Optional[pulumi.Input[bool]] = None,
            container: Optional[pulumi.Input[str]] = None,
            mountoptions: Optional[pulumi.Input[pulumi.InputType['LXCDiskMountoptionsArgs']]] = None,
            mp: Optional[pulumi.Input[str]] = None,
            quota: Optional[pulumi.Input[bool]] = None,
            replicate: Optional[pulumi.Input[bool]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            size: Optional[pulumi.Input[str]] = None,
            slot: Optional[pulumi.Input[int]] = None,
            storage: Optional[pulumi.Input[str]] = None,
            volume: Optional[pulumi.Input[str]] = None) -> 'LXCDisk':
        """
        Get an existing LXCDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LXCDiskState.__new__(_LXCDiskState)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["backup"] = backup
        __props__.__dict__["container"] = container
        __props__.__dict__["mountoptions"] = mountoptions
        __props__.__dict__["mp"] = mp
        __props__.__dict__["quota"] = quota
        __props__.__dict__["replicate"] = replicate
        __props__.__dict__["shared"] = shared
        __props__.__dict__["size"] = size
        __props__.__dict__["slot"] = slot
        __props__.__dict__["storage"] = storage
        __props__.__dict__["volume"] = volume
        return LXCDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter
    def container(self) -> pulumi.Output[str]:
        return pulumi.get(self, "container")

    @property
    @pulumi.getter
    def mountoptions(self) -> pulumi.Output[Optional['outputs.LXCDiskMountoptions']]:
        return pulumi.get(self, "mountoptions")

    @property
    @pulumi.getter
    def mp(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mp")

    @property
    @pulumi.getter
    def quota(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "quota")

    @property
    @pulumi.getter
    def replicate(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "replicate")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[str]:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def slot(self) -> pulumi.Output[int]:
        return pulumi.get(self, "slot")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output[str]:
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def volume(self) -> pulumi.Output[str]:
        return pulumi.get(self, "volume")

