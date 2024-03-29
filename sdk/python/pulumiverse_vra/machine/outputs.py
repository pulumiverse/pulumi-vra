# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MachineBootConfig',
    'MachineConstraint',
    'MachineDisk',
    'MachineDisksList',
    'MachineImageDiskConstraint',
    'MachineLink',
    'MachineNic',
    'MachineTag',
    'GetMachineLinkResult',
    'GetMachineTagResult',
]

@pulumi.output_type
class MachineBootConfig(dict):
    def __init__(__self__, *,
                 content: Optional[str] = None):
        """
        :param str content: Calid cloud config data in json-escaped yaml syntax.
        """
        if content is not None:
            pulumi.set(__self__, "content", content)

    @property
    @pulumi.getter
    def content(self) -> Optional[str]:
        """
        Calid cloud config data in json-escaped yaml syntax.
        """
        return pulumi.get(self, "content")


@pulumi.output_type
class MachineConstraint(dict):
    def __init__(__self__, *,
                 expression: str,
                 mandatory: bool):
        """
        :param str expression: Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        :param bool mandatory: Indicates whether this constraint should be strictly enforced or not.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def mandatory(self) -> bool:
        """
        Indicates whether this constraint should be strictly enforced or not.
        """
        return pulumi.get(self, "mandatory")


@pulumi.output_type
class MachineDisk(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "blockDeviceId":
            suggest = "block_device_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MachineDisk. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MachineDisk.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MachineDisk.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 block_device_id: str,
                 description: Optional[str] = None,
                 name: Optional[str] = None):
        """
        :param str block_device_id: ID of the existing block device.
        :param str description: Human-friendly description.
        :param str name: Human-friendly name used as an identifier in APIs that support this option.
        """
        pulumi.set(__self__, "block_device_id", block_device_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="blockDeviceId")
    def block_device_id(self) -> str:
        """
        ID of the existing block device.
        """
        return pulumi.get(self, "block_device_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class MachineDisksList(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "blockDeviceId":
            suggest = "block_device_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MachineDisksList. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MachineDisksList.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MachineDisksList.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 block_device_id: str,
                 description: Optional[str] = None,
                 name: Optional[str] = None):
        """
        :param str block_device_id: ID of the existing block device.
        :param str description: Human-friendly description.
        :param str name: Human-friendly name used as an identifier in APIs that support this option.
        """
        pulumi.set(__self__, "block_device_id", block_device_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="blockDeviceId")
    def block_device_id(self) -> str:
        """
        ID of the existing block device.
        """
        return pulumi.get(self, "block_device_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class MachineImageDiskConstraint(dict):
    def __init__(__self__, *,
                 expression: str,
                 mandatory: bool):
        """
        :param str expression: Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        :param bool mandatory: Indicates whether this constraint should be strictly enforced or not.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def mandatory(self) -> bool:
        """
        Indicates whether this constraint should be strictly enforced or not.
        """
        return pulumi.get(self, "mandatory")


@pulumi.output_type
class MachineLink(dict):
    def __init__(__self__, *,
                 rel: str,
                 href: Optional[str] = None,
                 hrefs: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> str:
        return pulumi.get(self, "rel")

    @property
    @pulumi.getter
    def href(self) -> Optional[str]:
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "hrefs")


@pulumi.output_type
class MachineNic(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkId":
            suggest = "network_id"
        elif key == "customProperties":
            suggest = "custom_properties"
        elif key == "deviceIndex":
            suggest = "device_index"
        elif key == "securityGroupIds":
            suggest = "security_group_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MachineNic. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MachineNic.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MachineNic.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_id: str,
                 addresses: Optional[Sequence[str]] = None,
                 custom_properties: Optional[Mapping[str, Any]] = None,
                 description: Optional[str] = None,
                 device_index: Optional[int] = None,
                 name: Optional[str] = None,
                 security_group_ids: Optional[Sequence[str]] = None):
        """
        :param str network_id: ID of the network instance that this network interface plugs into.
        :param Sequence[str] addresses: List of IP addresses allocated or in use by this network interface.
               example:[ "10.1.2.190" ]
        :param Mapping[str, Any] custom_properties: Additional properties that may be used to extend the base type.
        :param str description: Human-friendly description.
        :param int device_index: The device index of this network interface.
        :param str name: Human-friendly name used as an identifier in APIs that support this option.
        :param Sequence[str] security_group_ids: List of security group ids which this network interface will be assigned to.
        """
        pulumi.set(__self__, "network_id", network_id)
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if custom_properties is not None:
            pulumi.set(__self__, "custom_properties", custom_properties)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_index is not None:
            pulumi.set(__self__, "device_index", device_index)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        ID of the network instance that this network interface plugs into.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def addresses(self) -> Optional[Sequence[str]]:
        """
        List of IP addresses allocated or in use by this network interface.
        example:[ "10.1.2.190" ]
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> Optional[Mapping[str, Any]]:
        """
        Additional properties that may be used to extend the base type.
        """
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceIndex")
    def device_index(self) -> Optional[int]:
        """
        The device index of this network interface.
        """
        return pulumi.get(self, "device_index")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        List of security group ids which this network interface will be assigned to.
        """
        return pulumi.get(self, "security_group_ids")


@pulumi.output_type
class MachineTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: Tag’s key.
        :param str value: Tag’s value.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Tag’s key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetMachineLinkResult(dict):
    def __init__(__self__, *,
                 rel: str,
                 href: Optional[str] = None,
                 hrefs: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> str:
        return pulumi.get(self, "rel")

    @property
    @pulumi.getter
    def href(self) -> Optional[str]:
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "hrefs")


@pulumi.output_type
class GetMachineTagResult(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


