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
    'ProjectAdministratorRoleArgs',
    'ProjectConstraintsArgs',
    'ProjectConstraintsExtensibilityArgs',
    'ProjectConstraintsNetworkArgs',
    'ProjectConstraintsStorageArgs',
    'ProjectMemberRoleArgs',
    'ProjectViewerRoleArgs',
    'ProjectZoneAssignmentArgs',
    'GetProjectAdministratorRoleArgs',
    'GetProjectConstraintsArgs',
    'GetProjectConstraintsExtensibilityArgs',
    'GetProjectConstraintsNetworkArgs',
    'GetProjectConstraintsStorageArgs',
    'GetProjectMemberRoleArgs',
    'GetProjectViewerRoleArgs',
    'GetProjectZoneAssignmentArgs',
]

@pulumi.input_type
class ProjectAdministratorRoleArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProjectConstraintsArgs:
    def __init__(__self__, *,
                 extensibilities: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsExtensibilityArgs']]]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsNetworkArgs']]]] = None,
                 storages: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsStorageArgs']]]] = None):
        if extensibilities is not None:
            pulumi.set(__self__, "extensibilities", extensibilities)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if storages is not None:
            pulumi.set(__self__, "storages", storages)

    @property
    @pulumi.getter
    def extensibilities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsExtensibilityArgs']]]]:
        return pulumi.get(self, "extensibilities")

    @extensibilities.setter
    def extensibilities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsExtensibilityArgs']]]]):
        pulumi.set(self, "extensibilities", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsNetworkArgs']]]]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsNetworkArgs']]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def storages(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsStorageArgs']]]]:
        return pulumi.get(self, "storages")

    @storages.setter
    def storages(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectConstraintsStorageArgs']]]]):
        pulumi.set(self, "storages", value)


@pulumi.input_type
class ProjectConstraintsExtensibilityArgs:
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
class ProjectConstraintsNetworkArgs:
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
class ProjectConstraintsStorageArgs:
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
class ProjectMemberRoleArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProjectViewerRoleArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProjectZoneAssignmentArgs:
    def __init__(__self__, *,
                 zone_id: pulumi.Input[str],
                 cpu_limit: Optional[pulumi.Input[int]] = None,
                 max_instances: Optional[pulumi.Input[int]] = None,
                 memory_limit_mb: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 storage_limit_gb: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "zone_id", zone_id)
        if cpu_limit is not None:
            pulumi.set(__self__, "cpu_limit", cpu_limit)
        if max_instances is not None:
            pulumi.set(__self__, "max_instances", max_instances)
        if memory_limit_mb is not None:
            pulumi.set(__self__, "memory_limit_mb", memory_limit_mb)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if storage_limit_gb is not None:
            pulumi.set(__self__, "storage_limit_gb", storage_limit_gb)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cpu_limit")

    @cpu_limit.setter
    def cpu_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_limit", value)

    @property
    @pulumi.getter(name="maxInstances")
    def max_instances(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_instances")

    @max_instances.setter
    def max_instances(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_instances", value)

    @property
    @pulumi.getter(name="memoryLimitMb")
    def memory_limit_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory_limit_mb")

    @memory_limit_mb.setter
    def memory_limit_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_limit_mb", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="storageLimitGb")
    def storage_limit_gb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "storage_limit_gb")

    @storage_limit_gb.setter
    def storage_limit_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "storage_limit_gb", value)


@pulumi.input_type
class GetProjectAdministratorRoleArgs:
    def __init__(__self__, *,
                 email: str,
                 type: Optional[str] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: str):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GetProjectConstraintsArgs:
    def __init__(__self__, *,
                 extensibilities: Optional[Sequence['GetProjectConstraintsExtensibilityArgs']] = None,
                 networks: Optional[Sequence['GetProjectConstraintsNetworkArgs']] = None,
                 storages: Optional[Sequence['GetProjectConstraintsStorageArgs']] = None):
        if extensibilities is not None:
            pulumi.set(__self__, "extensibilities", extensibilities)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if storages is not None:
            pulumi.set(__self__, "storages", storages)

    @property
    @pulumi.getter
    def extensibilities(self) -> Optional[Sequence['GetProjectConstraintsExtensibilityArgs']]:
        return pulumi.get(self, "extensibilities")

    @extensibilities.setter
    def extensibilities(self, value: Optional[Sequence['GetProjectConstraintsExtensibilityArgs']]):
        pulumi.set(self, "extensibilities", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[Sequence['GetProjectConstraintsNetworkArgs']]:
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[Sequence['GetProjectConstraintsNetworkArgs']]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def storages(self) -> Optional[Sequence['GetProjectConstraintsStorageArgs']]:
        return pulumi.get(self, "storages")

    @storages.setter
    def storages(self, value: Optional[Sequence['GetProjectConstraintsStorageArgs']]):
        pulumi.set(self, "storages", value)


@pulumi.input_type
class GetProjectConstraintsExtensibilityArgs:
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
class GetProjectConstraintsNetworkArgs:
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
class GetProjectConstraintsStorageArgs:
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
class GetProjectMemberRoleArgs:
    def __init__(__self__, *,
                 email: str,
                 type: Optional[str] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: str):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GetProjectViewerRoleArgs:
    def __init__(__self__, *,
                 email: str,
                 type: Optional[str] = None):
        pulumi.set(__self__, "email", email)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: str):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class GetProjectZoneAssignmentArgs:
    def __init__(__self__, *,
                 cpu_limit: int,
                 max_instances: int,
                 memory_limit_mb: int,
                 priority: int,
                 storage_limit_gb: int,
                 zone_id: str):
        pulumi.set(__self__, "cpu_limit", cpu_limit)
        pulumi.set(__self__, "max_instances", max_instances)
        pulumi.set(__self__, "memory_limit_mb", memory_limit_mb)
        pulumi.set(__self__, "priority", priority)
        pulumi.set(__self__, "storage_limit_gb", storage_limit_gb)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> int:
        return pulumi.get(self, "cpu_limit")

    @cpu_limit.setter
    def cpu_limit(self, value: int):
        pulumi.set(self, "cpu_limit", value)

    @property
    @pulumi.getter(name="maxInstances")
    def max_instances(self) -> int:
        return pulumi.get(self, "max_instances")

    @max_instances.setter
    def max_instances(self, value: int):
        pulumi.set(self, "max_instances", value)

    @property
    @pulumi.getter(name="memoryLimitMb")
    def memory_limit_mb(self) -> int:
        return pulumi.get(self, "memory_limit_mb")

    @memory_limit_mb.setter
    def memory_limit_mb(self, value: int):
        pulumi.set(self, "memory_limit_mb", value)

    @property
    @pulumi.getter
    def priority(self) -> int:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: int):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="storageLimitGb")
    def storage_limit_gb(self) -> int:
        return pulumi.get(self, "storage_limit_gb")

    @storage_limit_gb.setter
    def storage_limit_gb(self, value: int):
        pulumi.set(self, "storage_limit_gb", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: str):
        pulumi.set(self, "zone_id", value)


