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

__all__ = ['StorageProfileAzureArgs', 'StorageProfileAzure']

@pulumi.input_type
class StorageProfileAzureArgs:
    def __init__(__self__, *,
                 default_item: pulumi.Input[bool],
                 region_id: pulumi.Input[str],
                 data_disk_caching: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_caching: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]] = None):
        """
        The set of arguments for constructing a StorageProfileAzure resource.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] region_id: A link to the region that is associated with the storage profile.
        :param pulumi.Input[str] data_disk_caching: Indicates the caching mechanism for additional disk.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[str] disk_type: Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] os_disk_caching: Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        :param pulumi.Input[str] storage_account_id: Id of a storage account where in the disk is placed.
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage policy should support encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        pulumi.set(__self__, "default_item", default_item)
        pulumi.set(__self__, "region_id", region_id)
        if data_disk_caching is not None:
            pulumi.set(__self__, "data_disk_caching", data_disk_caching)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if os_disk_caching is not None:
            pulumi.set(__self__, "os_disk_caching", os_disk_caching)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if supports_encryption is not None:
            pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> pulumi.Input[bool]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @default_item.setter
    def default_item(self, value: pulumi.Input[bool]):
        pulumi.set(self, "default_item", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Input[str]:
        """
        A link to the region that is associated with the storage profile.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="dataDiskCaching")
    def data_disk_caching(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the caching mechanism for additional disk.
        """
        return pulumi.get(self, "data_disk_caching")

    @data_disk_caching.setter
    def data_disk_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_disk_caching", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="osDiskCaching")
    def os_disk_caching(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        """
        return pulumi.get(self, "os_disk_caching")

    @os_disk_caching.setter
    def os_disk_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_disk_caching", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of a storage account where in the disk is placed.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this storage policy should support encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @supports_encryption.setter
    def supports_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "supports_encryption", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _StorageProfileAzureState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 data_disk_caching: Optional[pulumi.Input[str]] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 external_region_id: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureLinkArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 os_disk_caching: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StorageProfileAzure resources.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[str] data_disk_caching: Indicates the caching mechanism for additional disk.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[str] disk_type: Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        :param pulumi.Input[str] external_region_id: The id of the region as seen in the cloud provider for which this profile is defined.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureLinkArgs']]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] organization_id: The id of the organization this entity belongs to.
        :param pulumi.Input[str] os_disk_caching: Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        :param pulumi.Input[str] owner: Email of the user that owns the entity.
        :param pulumi.Input[str] region_id: A link to the region that is associated with the storage profile.
        :param pulumi.Input[str] storage_account_id: Id of a storage account where in the disk is placed.
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage policy should support encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if data_disk_caching is not None:
            pulumi.set(__self__, "data_disk_caching", data_disk_caching)
        if default_item is not None:
            pulumi.set(__self__, "default_item", default_item)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_type is not None:
            pulumi.set(__self__, "disk_type", disk_type)
        if external_region_id is not None:
            pulumi.set(__self__, "external_region_id", external_region_id)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if os_disk_caching is not None:
            pulumi.set(__self__, "os_disk_caching", os_disk_caching)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if storage_account_id is not None:
            pulumi.set(__self__, "storage_account_id", storage_account_id)
        if supports_encryption is not None:
            pulumi.set(__self__, "supports_encryption", supports_encryption)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dataDiskCaching")
    def data_disk_caching(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the caching mechanism for additional disk.
        """
        return pulumi.get(self, "data_disk_caching")

    @data_disk_caching.setter
    def data_disk_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_disk_caching", value)

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @default_item.setter
    def default_item(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_item", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        """
        return pulumi.get(self, "disk_type")

    @disk_type.setter
    def disk_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_type", value)

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the region as seen in the cloud provider for which this profile is defined.
        """
        return pulumi.get(self, "external_region_id")

    @external_region_id.setter
    def external_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_region_id", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureLinkArgs']]]]:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the organization this entity belongs to.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="osDiskCaching")
    def os_disk_caching(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        """
        return pulumi.get(self, "os_disk_caching")

    @os_disk_caching.setter
    def os_disk_caching(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "os_disk_caching", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        A link to the region that is associated with the storage profile.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of a storage account where in the disk is placed.
        """
        return pulumi.get(self, "storage_account_id")

    @storage_account_id.setter
    def storage_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_account_id", value)

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this storage policy should support encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @supports_encryption.setter
    def supports_encryption(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "supports_encryption", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StorageProfileAzureTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class StorageProfileAzure(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_disk_caching: Optional[pulumi.Input[str]] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_caching: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureTagArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### S
        This is an example of how to create a storage profile azure resource.

        **Vra storage profile azure:**

        ```python
        import pulumi
        import schmidtw_vra as vra

        # Azure storage profile using vra_storage_profile_azure resource with managed disk.
        this_storage_profile_azure = vra.StorageProfileAzure("thisStorageProfileAzure",
            description="Azure Storage Profile with managed disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            data_disk_caching="None",
            disk_type="Standard_LRS",
            os_disk_caching="None",
            tags=[vra.StorageProfileAzureTagArgs(
                key="foo",
                value="bar",
            )])
        # Azure storage profile using vra_storage_profile_azure resource with unmanaged disk.
        this_index_storage_profile_azure_storage_profile_azure = vra.StorageProfileAzure("thisIndex/storageProfileAzureStorageProfileAzure",
            description="Azure Storage Profile with unmanaged disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            data_disk_caching="None",
            os_disk_caching="None",
            tags=[vra.StorageProfileAzureTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        A storage profile azure resource supports the following arguments:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_disk_caching: Indicates the caching mechanism for additional disk.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[str] disk_type: Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] os_disk_caching: Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        :param pulumi.Input[str] region_id: A link to the region that is associated with the storage profile.
        :param pulumi.Input[str] storage_account_id: Id of a storage account where in the disk is placed.
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage policy should support encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureTagArgs']]]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StorageProfileAzureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### S
        This is an example of how to create a storage profile azure resource.

        **Vra storage profile azure:**

        ```python
        import pulumi
        import schmidtw_vra as vra

        # Azure storage profile using vra_storage_profile_azure resource with managed disk.
        this_storage_profile_azure = vra.StorageProfileAzure("thisStorageProfileAzure",
            description="Azure Storage Profile with managed disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            data_disk_caching="None",
            disk_type="Standard_LRS",
            os_disk_caching="None",
            tags=[vra.StorageProfileAzureTagArgs(
                key="foo",
                value="bar",
            )])
        # Azure storage profile using vra_storage_profile_azure resource with unmanaged disk.
        this_index_storage_profile_azure_storage_profile_azure = vra.StorageProfileAzure("thisIndex/storageProfileAzureStorageProfileAzure",
            description="Azure Storage Profile with unmanaged disks.",
            region_id=data["vra_region"]["this"]["id"],
            default_item=False,
            supports_encryption=False,
            data_disk_caching="None",
            os_disk_caching="None",
            tags=[vra.StorageProfileAzureTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        A storage profile azure resource supports the following arguments:

        :param str resource_name: The name of the resource.
        :param StorageProfileAzureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StorageProfileAzureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_disk_caching: Optional[pulumi.Input[str]] = None,
                 default_item: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 os_disk_caching: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 supports_encryption: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StorageProfileAzureArgs.__new__(StorageProfileAzureArgs)

            __props__.__dict__["data_disk_caching"] = data_disk_caching
            if default_item is None and not opts.urn:
                raise TypeError("Missing required property 'default_item'")
            __props__.__dict__["default_item"] = default_item
            __props__.__dict__["description"] = description
            __props__.__dict__["disk_type"] = disk_type
            __props__.__dict__["name"] = name
            __props__.__dict__["os_disk_caching"] = os_disk_caching
            if region_id is None and not opts.urn:
                raise TypeError("Missing required property 'region_id'")
            __props__.__dict__["region_id"] = region_id
            __props__.__dict__["storage_account_id"] = storage_account_id
            __props__.__dict__["supports_encryption"] = supports_encryption
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["external_region_id"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["updated_at"] = None
        super(StorageProfileAzure, __self__).__init__(
            'vra:index/storageProfileAzure:StorageProfileAzure',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            data_disk_caching: Optional[pulumi.Input[str]] = None,
            default_item: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disk_type: Optional[pulumi.Input[str]] = None,
            external_region_id: Optional[pulumi.Input[str]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureLinkArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            os_disk_caching: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            region_id: Optional[pulumi.Input[str]] = None,
            storage_account_id: Optional[pulumi.Input[str]] = None,
            supports_encryption: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureTagArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'StorageProfileAzure':
        """
        Get an existing StorageProfileAzure resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[str] data_disk_caching: Indicates the caching mechanism for additional disk.
        :param pulumi.Input[bool] default_item: Indicates if this storage profile is a default profile.
        :param pulumi.Input[str] description: A human-friendly description.
        :param pulumi.Input[str] disk_type: Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        :param pulumi.Input[str] external_region_id: The id of the region as seen in the cloud provider for which this profile is defined.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureLinkArgs']]]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] organization_id: The id of the organization this entity belongs to.
        :param pulumi.Input[str] os_disk_caching: Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        :param pulumi.Input[str] owner: Email of the user that owns the entity.
        :param pulumi.Input[str] region_id: A link to the region that is associated with the storage profile.
        :param pulumi.Input[str] storage_account_id: Id of a storage account where in the disk is placed.
        :param pulumi.Input[bool] supports_encryption: Indicates whether this storage policy should support encryption or not.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StorageProfileAzureTagArgs']]]] tags: A set of tag keys and optional values that were set on this Network Profile.
               example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StorageProfileAzureState.__new__(_StorageProfileAzureState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["data_disk_caching"] = data_disk_caching
        __props__.__dict__["default_item"] = default_item
        __props__.__dict__["description"] = description
        __props__.__dict__["disk_type"] = disk_type
        __props__.__dict__["external_region_id"] = external_region_id
        __props__.__dict__["links"] = links
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["os_disk_caching"] = os_disk_caching
        __props__.__dict__["owner"] = owner
        __props__.__dict__["region_id"] = region_id
        __props__.__dict__["storage_account_id"] = storage_account_id
        __props__.__dict__["supports_encryption"] = supports_encryption
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return StorageProfileAzure(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dataDiskCaching")
    def data_disk_caching(self) -> pulumi.Output[str]:
        """
        Indicates the caching mechanism for additional disk.
        """
        return pulumi.get(self, "data_disk_caching")

    @property
    @pulumi.getter(name="defaultItem")
    def default_item(self) -> pulumi.Output[bool]:
        """
        Indicates if this storage profile is a default profile.
        """
        return pulumi.get(self, "default_item")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A human-friendly description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> pulumi.Output[str]:
        """
        Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> pulumi.Output[str]:
        """
        The id of the region as seen in the cloud provider for which this profile is defined.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.StorageProfileAzureLink']]:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The id of the organization this entity belongs to.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="osDiskCaching")
    def os_disk_caching(self) -> pulumi.Output[str]:
        """
        Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
        """
        return pulumi.get(self, "os_disk_caching")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Email of the user that owns the entity.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Output[str]:
        """
        A link to the region that is associated with the storage profile.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[str]:
        """
        Id of a storage account where in the disk is placed.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> pulumi.Output[bool]:
        """
        Indicates whether this storage policy should support encryption or not.
        """
        return pulumi.get(self, "supports_encryption")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.StorageProfileAzureTag']]:
        """
        A set of tag keys and optional values that were set on this Network Profile.
        example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

