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
    'AwsLinkArgs',
    'AwsTagArgs',
    'AzureLinkArgs',
    'AzureTagArgs',
    'GcpLinkArgs',
    'GcpTagArgs',
    'NsxtLinkArgs',
    'NsxtTagArgs',
    'NsxvLinkArgs',
    'NsxvTagArgs',
    'VSphereLinkArgs',
    'VSphereTagArgs',
    'VmcLinkArgs',
    'VmcTagArgs',
    'GetAwsTagArgs',
    'GetAzureTagArgs',
    'GetGcpTagArgs',
    'GetNsxtTagArgs',
    'GetNsxvTagArgs',
    'GetVSphereTagArgs',
    'GetVmcTagArgs',
]

@pulumi.input_type
class AwsLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class AwsTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class AzureLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class AzureTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GcpLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class GcpTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NsxtLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class NsxtTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NsxvLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class NsxvTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class VSphereLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class VSphereTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class VmcLinkArgs:
    def __init__(__self__, *,
                 rel: pulumi.Input[str],
                 href: Optional[pulumi.Input[str]] = None,
                 hrefs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "rel", rel)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if hrefs is not None:
            pulumi.set(__self__, "hrefs", hrefs)

    @property
    @pulumi.getter
    def rel(self) -> pulumi.Input[str]:
        return pulumi.get(self, "rel")

    @rel.setter
    def rel(self, value: pulumi.Input[str]):
        pulumi.set(self, "rel", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def hrefs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "hrefs")

    @hrefs.setter
    def hrefs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "hrefs", value)


@pulumi.input_type
class VmcTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetAwsTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetAzureTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetGcpTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetNsxtTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetNsxvTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetVSphereTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetVmcTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Tag’s value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


