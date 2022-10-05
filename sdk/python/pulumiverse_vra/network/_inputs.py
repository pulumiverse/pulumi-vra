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
    'IpRangeLinkArgs',
    'IpRangeTagArgs',
    'NetworkConstraintArgs',
    'NetworkLinkArgs',
    'NetworkTagArgs',
    'ProfileLinkArgs',
    'ProfileTagArgs',
    'GetDomainTagArgs',
    'GetNetworkConstraintArgs',
    'GetNetworkTagArgs',
    'GetProfileTagArgs',
]

@pulumi.input_type
class IpRangeLinkArgs:
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
class IpRangeTagArgs:
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
class NetworkConstraintArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 mandatory: pulumi.Input[bool]):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def mandatory(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "mandatory")

    @mandatory.setter
    def mandatory(self, value: pulumi.Input[bool]):
        pulumi.set(self, "mandatory", value)


@pulumi.input_type
class NetworkLinkArgs:
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
class NetworkTagArgs:
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
class ProfileLinkArgs:
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
class ProfileTagArgs:
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
class GetDomainTagArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetNetworkConstraintArgs:
    def __init__(__self__, *,
                 expression: str,
                 mandatory: bool):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "mandatory", mandatory)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: str):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def mandatory(self) -> bool:
        return pulumi.get(self, "mandatory")

    @mandatory.setter
    def mandatory(self, value: bool):
        pulumi.set(self, "mandatory", value)


@pulumi.input_type
class GetNetworkTagArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetProfileTagArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


