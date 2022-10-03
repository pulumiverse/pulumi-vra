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

__all__ = [
    'GetZoneResult',
    'AwaitableGetZoneResult',
    'get_zone',
    'get_zone_output',
]

@pulumi.output_type
class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, cloud_account_id=None, compute_ids=None, created_at=None, custom_properties=None, description=None, external_region_id=None, folder=None, id=None, links=None, name=None, org_id=None, owner=None, placement_policy=None, tags=None, tags_to_matches=None, updated_at=None):
        if cloud_account_id and not isinstance(cloud_account_id, str):
            raise TypeError("Expected argument 'cloud_account_id' to be a str")
        pulumi.set(__self__, "cloud_account_id", cloud_account_id)
        if compute_ids and not isinstance(compute_ids, list):
            raise TypeError("Expected argument 'compute_ids' to be a list")
        pulumi.set(__self__, "compute_ids", compute_ids)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if custom_properties and not isinstance(custom_properties, dict):
            raise TypeError("Expected argument 'custom_properties' to be a dict")
        pulumi.set(__self__, "custom_properties", custom_properties)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if external_region_id and not isinstance(external_region_id, str):
            raise TypeError("Expected argument 'external_region_id' to be a str")
        pulumi.set(__self__, "external_region_id", external_region_id)
        if folder and not isinstance(folder, str):
            raise TypeError("Expected argument 'folder' to be a str")
        pulumi.set(__self__, "folder", folder)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
        if placement_policy and not isinstance(placement_policy, str):
            raise TypeError("Expected argument 'placement_policy' to be a str")
        pulumi.set(__self__, "placement_policy", placement_policy)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tags_to_matches and not isinstance(tags_to_matches, list):
            raise TypeError("Expected argument 'tags_to_matches' to be a list")
        pulumi.set(__self__, "tags_to_matches", tags_to_matches)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="cloudAccountId")
    def cloud_account_id(self) -> str:
        """
        The ID of the cloud account this zone belongs to.
        """
        return pulumi.get(self, "cloud_account_id")

    @property
    @pulumi.getter(name="computeIds")
    def compute_ids(self) -> Sequence[str]:
        """
        The ids of the compute resources that has been explicitly assigned to this zone.
        """
        return pulumi.get(self, "compute_ids")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date when the entity was created. The date is in ISO 8601 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> Mapping[str, Any]:
        """
        A list of key value pair of properties that will be used.
        """
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> str:
        """
        The id of the region for which this zone is defined.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter
    def folder(self) -> str:
        """
        The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def links(self) -> Sequence['outputs.GetZoneLinkResult']:
        """
        HATEOAS of the entity.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> str:
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
    @pulumi.getter(name="placementPolicy")
    def placement_policy(self) -> str:
        """
        The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`.
        """
        return pulumi.get(self, "placement_policy")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetZoneTagResult']:
        """
        A set of tag keys and optional values that were set on this resource:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsToMatches")
    def tags_to_matches(self) -> Sequence['outputs.GetZoneTagsToMatchResult']:
        """
        A set of tag keys and optional values for compute resource filtering:
        """
        return pulumi.get(self, "tags_to_matches")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            cloud_account_id=self.cloud_account_id,
            compute_ids=self.compute_ids,
            created_at=self.created_at,
            custom_properties=self.custom_properties,
            description=self.description,
            external_region_id=self.external_region_id,
            folder=self.folder,
            id=self.id,
            links=self.links,
            name=self.name,
            org_id=self.org_id,
            owner=self.owner,
            placement_policy=self.placement_policy,
            tags=self.tags,
            tags_to_matches=self.tags_to_matches,
            updated_at=self.updated_at)


def get_zone(id: Optional[str] = None,
             name: Optional[str] = None,
             tags: Optional[Sequence[pulumi.InputType['GetZoneTagArgs']]] = None,
             tags_to_matches: Optional[Sequence[pulumi.InputType['GetZoneTagsToMatchArgs']]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneResult:
    """
    ## Example Usage
    ### S

    This is an example of how to read a zone data source.

    ```python
    import pulumi
    import pulumi_vra as vra

    test_zone = vra.get_zone(name=var["zone_name"])
    ```

    A zone data source supports the following arguments:


    :param str id: The id of the zone resource instance.
    :param str name: A human-friendly name used as an identifier for the zone resource instance.
    :param Sequence[pulumi.InputType['GetZoneTagArgs']] tags: A set of tag keys and optional values that were set on this resource:
    :param Sequence[pulumi.InputType['GetZoneTagsToMatchArgs']] tags_to_matches: A set of tag keys and optional values for compute resource filtering:
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['tags'] = tags
    __args__['tagsToMatches'] = tags_to_matches
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vra:index/getZone:getZone', __args__, opts=opts, typ=GetZoneResult).value

    return AwaitableGetZoneResult(
        cloud_account_id=__ret__.cloud_account_id,
        compute_ids=__ret__.compute_ids,
        created_at=__ret__.created_at,
        custom_properties=__ret__.custom_properties,
        description=__ret__.description,
        external_region_id=__ret__.external_region_id,
        folder=__ret__.folder,
        id=__ret__.id,
        links=__ret__.links,
        name=__ret__.name,
        org_id=__ret__.org_id,
        owner=__ret__.owner,
        placement_policy=__ret__.placement_policy,
        tags=__ret__.tags,
        tags_to_matches=__ret__.tags_to_matches,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_zone)
def get_zone_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetZoneTagArgs']]]]] = None,
                    tags_to_matches: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetZoneTagsToMatchArgs']]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneResult]:
    """
    ## Example Usage
    ### S

    This is an example of how to read a zone data source.

    ```python
    import pulumi
    import pulumi_vra as vra

    test_zone = vra.get_zone(name=var["zone_name"])
    ```

    A zone data source supports the following arguments:


    :param str id: The id of the zone resource instance.
    :param str name: A human-friendly name used as an identifier for the zone resource instance.
    :param Sequence[pulumi.InputType['GetZoneTagArgs']] tags: A set of tag keys and optional values that were set on this resource:
    :param Sequence[pulumi.InputType['GetZoneTagsToMatchArgs']] tags_to_matches: A set of tag keys and optional values for compute resource filtering:
    """
    ...
