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
    'GetAwsResult',
    'AwaitableGetAwsResult',
    'get_aws',
    'get_aws_output',
]

@pulumi.output_type
class GetAwsResult:
    """
    A collection of values returned by getAws.
    """
    def __init__(__self__, cloud_account_id=None, created_at=None, default_item=None, description=None, device_type=None, external_region_id=None, filter=None, id=None, iops=None, links=None, name=None, org_id=None, owner=None, supports_encryption=None, tags=None, updated_at=None, volume_type=None):
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
        if device_type and not isinstance(device_type, str):
            raise TypeError("Expected argument 'device_type' to be a str")
        pulumi.set(__self__, "device_type", device_type)
        if external_region_id and not isinstance(external_region_id, str):
            raise TypeError("Expected argument 'external_region_id' to be a str")
        pulumi.set(__self__, "external_region_id", external_region_id)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iops and not isinstance(iops, str):
            raise TypeError("Expected argument 'iops' to be a str")
        pulumi.set(__self__, "iops", iops)
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
        if supports_encryption and not isinstance(supports_encryption, bool):
            raise TypeError("Expected argument 'supports_encryption' to be a bool")
        pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> str:
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
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> str:
        return pulumi.get(self, "device_type")

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
    @pulumi.getter
    def iops(self) -> str:
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter
    def links(self) -> Sequence['outputs.GetAwsLinkResult']:
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
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> bool:
        return pulumi.get(self, "supports_encryption")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetAwsTagResult']:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        return pulumi.get(self, "volume_type")


class AwaitableGetAwsResult(GetAwsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsResult(
            cloud_account_id=self.cloud_account_id,
            created_at=self.created_at,
            default_item=self.default_item,
            description=self.description,
            device_type=self.device_type,
            external_region_id=self.external_region_id,
            filter=self.filter,
            id=self.id,
            iops=self.iops,
            links=self.links,
            name=self.name,
            org_id=self.org_id,
            owner=self.owner,
            supports_encryption=self.supports_encryption,
            tags=self.tags,
            updated_at=self.updated_at,
            volume_type=self.volume_type)


def get_aws(filter: Optional[str] = None,
            id: Optional[str] = None,
            tags: Optional[Sequence[pulumi.InputType['GetAwsTagArgs']]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsResult:
    """
    ## Example Usage
    ### S
    This is an example of how to create a storage profile aws resource.

    **Storage profile aws data source by its id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_aws(id=vra_storage_profile_aws["this"]["id"])
    ```

    **Vra storage profile data source filter by external region id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_aws(filter="externalRegionId eq 'foobar'")
    ```

    A storage profile aws data source supports the following arguments:


    :param str filter: Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
    :param str id: The id of the image profile instance.
    :param Sequence[pulumi.InputType['GetAwsTagArgs']] tags: A set of tag keys and optional values that were set on this Network Profile.
           example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['id'] = id
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vra:storageprofile/getAws:getAws', __args__, opts=opts, typ=GetAwsResult).value

    return AwaitableGetAwsResult(
        cloud_account_id=__ret__.cloud_account_id,
        created_at=__ret__.created_at,
        default_item=__ret__.default_item,
        description=__ret__.description,
        device_type=__ret__.device_type,
        external_region_id=__ret__.external_region_id,
        filter=__ret__.filter,
        id=__ret__.id,
        iops=__ret__.iops,
        links=__ret__.links,
        name=__ret__.name,
        org_id=__ret__.org_id,
        owner=__ret__.owner,
        supports_encryption=__ret__.supports_encryption,
        tags=__ret__.tags,
        updated_at=__ret__.updated_at,
        volume_type=__ret__.volume_type)


@_utilities.lift_output_func(get_aws)
def get_aws_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                   id: Optional[pulumi.Input[Optional[str]]] = None,
                   tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetAwsTagArgs']]]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAwsResult]:
    """
    ## Example Usage
    ### S
    This is an example of how to create a storage profile aws resource.

    **Storage profile aws data source by its id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_aws(id=vra_storage_profile_aws["this"]["id"])
    ```

    **Vra storage profile data source filter by external region id:**

    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.storageprofile.get_aws(filter="externalRegionId eq 'foobar'")
    ```

    A storage profile aws data source supports the following arguments:


    :param str filter: Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
    :param str id: The id of the image profile instance.
    :param Sequence[pulumi.InputType['GetAwsTagArgs']] tags: A set of tag keys and optional values that were set on this Network Profile.
           example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
    """
    ...
