# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BlockDeviceSnapshotInitArgs', 'BlockDeviceSnapshot']

@pulumi.input_type
class BlockDeviceSnapshotInitArgs:
    def __init__(__self__, *,
                 block_device_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BlockDeviceSnapshot resource.
        :param pulumi.Input[str] block_device_id: ID of block device.
        :param pulumi.Input[str] description: Human-friendly description.
        """
        pulumi.set(__self__, "block_device_id", block_device_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="blockDeviceId")
    def block_device_id(self) -> pulumi.Input[str]:
        """
        ID of block device.
        """
        return pulumi.get(self, "block_device_id")

    @block_device_id.setter
    def block_device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "block_device_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _BlockDeviceSnapshotState:
    def __init__(__self__, *,
                 block_device_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_current: Optional[pulumi.Input[bool]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['BlockDeviceSnapshotLinkArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BlockDeviceSnapshot resources.
        :param pulumi.Input[str] block_device_id: ID of block device.
        :param pulumi.Input[str] created_at: Date when entity was created. Date and time format is ISO 8601 and UTC.
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[bool] is_current: Indicates whether snapshot on block device is current.
        :param pulumi.Input[Sequence[pulumi.Input['BlockDeviceSnapshotLinkArgs']]] links: HATEOAS of entity
        :param pulumi.Input[str] name: Human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] org_id: ID of organization that entity belongs to.
        :param pulumi.Input[str] owner: Email of entity owner.
        """
        if block_device_id is not None:
            pulumi.set(__self__, "block_device_id", block_device_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_current is not None:
            pulumi.set(__self__, "is_current", is_current)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="blockDeviceId")
    def block_device_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of block device.
        """
        return pulumi.get(self, "block_device_id")

    @block_device_id.setter
    def block_device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block_device_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when entity was created. Date and time format is ISO 8601 and UTC.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isCurrent")
    def is_current(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether snapshot on block device is current.
        """
        return pulumi.get(self, "is_current")

    @is_current.setter
    def is_current(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_current", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BlockDeviceSnapshotLinkArgs']]]]:
        """
        HATEOAS of entity
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BlockDeviceSnapshotLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of organization that entity belongs to.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Email of entity owner.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class BlockDeviceSnapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_device_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a VMware vRealize Automation block device snapshot resource.

        ## Example Usage
        ### S

        The following example shows how to create a block device snapshot resource.

        ```python
        import pulumi
        import schmidtw_vra as vra

        snapshot1 = vra.BlockDeviceSnapshot("snapshot1",
            block_device_id=var["block_device_id"],
            description="terraform fcd snapshot")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block_device_id: ID of block device.
        :param pulumi.Input[str] description: Human-friendly description.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlockDeviceSnapshotInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a VMware vRealize Automation block device snapshot resource.

        ## Example Usage
        ### S

        The following example shows how to create a block device snapshot resource.

        ```python
        import pulumi
        import schmidtw_vra as vra

        snapshot1 = vra.BlockDeviceSnapshot("snapshot1",
            block_device_id=var["block_device_id"],
            description="terraform fcd snapshot")
        ```

        :param str resource_name: The name of the resource.
        :param BlockDeviceSnapshotInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlockDeviceSnapshotInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_device_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BlockDeviceSnapshotInitArgs.__new__(BlockDeviceSnapshotInitArgs)

            if block_device_id is None and not opts.urn:
                raise TypeError("Missing required property 'block_device_id'")
            __props__.__dict__["block_device_id"] = block_device_id
            __props__.__dict__["description"] = description
            __props__.__dict__["created_at"] = None
            __props__.__dict__["is_current"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["org_id"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["updated_at"] = None
        super(BlockDeviceSnapshot, __self__).__init__(
            'vra:index/blockDeviceSnapshot:BlockDeviceSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block_device_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_current: Optional[pulumi.Input[bool]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BlockDeviceSnapshotLinkArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'BlockDeviceSnapshot':
        """
        Get an existing BlockDeviceSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block_device_id: ID of block device.
        :param pulumi.Input[str] created_at: Date when entity was created. Date and time format is ISO 8601 and UTC.
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[bool] is_current: Indicates whether snapshot on block device is current.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BlockDeviceSnapshotLinkArgs']]]] links: HATEOAS of entity
        :param pulumi.Input[str] name: Human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] org_id: ID of organization that entity belongs to.
        :param pulumi.Input[str] owner: Email of entity owner.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlockDeviceSnapshotState.__new__(_BlockDeviceSnapshotState)

        __props__.__dict__["block_device_id"] = block_device_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["is_current"] = is_current
        __props__.__dict__["links"] = links
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["owner"] = owner
        __props__.__dict__["updated_at"] = updated_at
        return BlockDeviceSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockDeviceId")
    def block_device_id(self) -> pulumi.Output[str]:
        """
        ID of block device.
        """
        return pulumi.get(self, "block_device_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date when entity was created. Date and time format is ISO 8601 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isCurrent")
    def is_current(self) -> pulumi.Output[bool]:
        """
        Indicates whether snapshot on block device is current.
        """
        return pulumi.get(self, "is_current")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.BlockDeviceSnapshotLink']]:
        """
        HATEOAS of entity
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        ID of organization that entity belongs to.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Email of entity owner.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

