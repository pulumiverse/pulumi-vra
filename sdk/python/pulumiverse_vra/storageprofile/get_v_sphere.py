# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetVSphereResult',
    'AwaitableGetVSphereResult',
    'get_v_sphere',
    'get_v_sphere_output',
]

@pulumi.output_type
class GetVSphereResult:
    """
    A collection of values returned by getVSphere.
    """
    def __init__(__self__, cloud_account_id=None, created_at=None, default_item=None, description=None, disk_mode=None, disk_type=None, external_region_id=None, filter=None, id=None, limit_iops=None, links=None, name=None, org_id=None, owner=None, provisioning_type=None, shares=None, shares_level=None, supports_encryption=None, tags=None, updated_at=None):
        if cloud_account_id and not isinstance(cloud_account_id, str):
            raise TypeError("Expected argument 'cloud_account_id' to be a str")
        pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if default_item and not isinstance(default_item, bool):
            raise TypeError("Expected argument 'default_item' to be a bool")
        pulumi.set(__self__, "default_item", default_item)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_mode and not isinstance(disk_mode, str):
            raise TypeError("Expected argument 'disk_mode' to be a str")
        pulumi.set(__self__, "disk_mode", disk_mode)
        if disk_type and not isinstance(disk_type, str):
            raise TypeError("Expected argument 'disk_type' to be a str")
        pulumi.set(__self__, "disk_type", disk_type)
        if external_region_id and not isinstance(external_region_id, str):
            raise TypeError("Expected argument 'external_region_id' to be a str")
        pulumi.set(__self__, "external_region_id", external_region_id)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limit_iops and not isinstance(limit_iops, str):
            raise TypeError("Expected argument 'limit_iops' to be a str")
        pulumi.set(__self__, "limit_iops", limit_iops)
        if links and not isinstance(links, list):
            raise TypeError("Expected argument 'links' to be a list")
        pulumi.set(__self__, "links", links)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if provisioning_type and not isinstance(provisioning_type, str):
            raise TypeError("Expected argument 'provisioning_type' to be a str")
        pulumi.set(__self__, "provisioning_type", provisioning_type)
        if shares and not isinstance(shares, str):
            raise TypeError("Expected argument 'shares' to be a str")
        pulumi.set(__self__, "shares", shares)
        if shares_level and not isinstance(shares_level, str):
            raise TypeError("Expected argument 'shares_level' to be a str")
        pulumi.set(__self__, "shares_level", shares_level)
        if supports_encryption and not isinstance(supports_encryption, bool):
            raise TypeError("Expected argument 'supports_encryption' to be a bool")
        pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> str:
        """
        Id of the cloud account this storage profile belongs to.
        """
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> bool:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskMode")
    def disk_mode(self) -> str:
        """
        Type of mode for the disk.
        """
        return pulumi.get(self, "disk_mode")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> str:
        """
        Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> str:
        """
        The id of the region as seen in the cloud provider for which this profile is defined.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="limitIops")
    def limit_iops(self) -> str:
        """
        The upper bound for the I/O operations per second allocated for each virtual disk.
        """
        return pulumi.get(self, "limit_iops")

    @property
    @pulumi.getter
    def links(self) -> Sequence['outputs.GetVSphereLinkResult']:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The id of the organization this entity belongs to.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="provisioningType")
    def provisioning_type(self) -> str:
        """
        Type of provisioning policy for the disk.
        """
        return pulumi.get(self, "provisioning_type")

    @property
    @pulumi.getter
    def shares(self) -> str:
        """
        A specific number of shares assigned to each virtual machine.
        """
        return pulumi.get(self, "shares")

    @property
    @pulumi.getter(name="sharesLevel")
    def shares_level(self) -> Optional[str]:
        return pulumi.get(self, "shares_level")

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> bool:
        """
        Indicates whether this storage policy should support encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetVSphereTagResult']:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetVSphereResult(GetVSphereResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVSphereResult(
            cloud_account_id=self.cloud_account_id,
            created_at=self.created_at,
            default_item=self.default_item,
            description=self.description,
            disk_mode=self.disk_mode,
            disk_type=self.disk_type,
            external_region_id=self.external_region_id,
            filter=self.filter,
            id=self.id,
            limit_iops=self.limit_iops,
            links=self.links,
            name=self.name,
            org_id=self.org_id,
            owner=self.owner,
            provisioning_type=self.provisioning_type,
            shares=self.shares,
            shares_level=self.shares_level,
            supports_encryption=self.supports_encryption,
            tags=self.tags,
            updated_at=self.updated_at)


def get_v_sphere(filter: Optional[str] = None,
                 id: Optional[str] = None,
                 shares_level: Optional[str] = None,
                 tags: Optional[Sequence[pulumi.InputType['GetVSphereTagArgs']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVSphereResult:
    """
    ## Example Usage
    ### S
    This is an example of how to create a storage profile vsphere data source.

    **Storage profile vsphere data source by its id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_v_sphere(id=vra_storage_profile_vsphere["this"]["id"])
    ```

    **Vra storage profile data source filter by external region id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_v_sphere(filter="externalRegionId eq 'foobar'")
    ```

    A storage profile vsphere data source supports the following arguments:


    :param str filter: Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
    :param str id: The id of the image profile instance.
    :param str shares_level: Indicates whether this storage profile supports encryption or not.
    :param Sequence[pulumi.InputType['GetVSphereTagArgs']] tags: A set of tag keys and optional values that were set on this Network Profile.
           example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['id'] = id
    __args__['sharesLevel'] = shares_level
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vra:storageprofile/getVSphere:getVSphere', __args__, opts=opts, typ=GetVSphereResult).value

    return AwaitableGetVSphereResult(
        cloud_account_id=__ret__.cloud_account_id,
        created_at=__ret__.created_at,
        default_item=__ret__.default_item,
        description=__ret__.description,
        disk_mode=__ret__.disk_mode,
        disk_type=__ret__.disk_type,
        external_region_id=__ret__.external_region_id,
        filter=__ret__.filter,
        id=__ret__.id,
        limit_iops=__ret__.limit_iops,
        links=__ret__.links,
        name=__ret__.name,
        org_id=__ret__.org_id,
        owner=__ret__.owner,
        provisioning_type=__ret__.provisioning_type,
        shares=__ret__.shares,
        shares_level=__ret__.shares_level,
        supports_encryption=__ret__.supports_encryption,
        tags=__ret__.tags,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_v_sphere)
def get_v_sphere_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                        id: Optional[pulumi.Input[Optional[str]]] = None,
                        shares_level: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVSphereTagArgs']]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVSphereResult]:
    """
    ## Example Usage
    ### S
    This is an example of how to create a storage profile vsphere data source.

    **Storage profile vsphere data source by its id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_v_sphere(id=vra_storage_profile_vsphere["this"]["id"])
    ```

    **Vra storage profile data source filter by external region id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_v_sphere(filter="externalRegionId eq 'foobar'")
    ```

    A storage profile vsphere data source supports the following arguments:


    :param str filter: Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
    :param str id: The id of the image profile instance.
    :param str shares_level: Indicates whether this storage profile supports encryption or not.
    :param Sequence[pulumi.InputType['GetVSphereTagArgs']] tags: A set of tag keys and optional values that were set on this Network Profile.
           example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
    """
    ...
