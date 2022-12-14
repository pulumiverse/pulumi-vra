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

__all__ = ['AwsArgs', 'Aws']

@pulumi.input_type
class AwsArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 regions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 secret_key: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]] = None):
        """
        The set of arguments for constructing a Aws resource.
        :param pulumi.Input[str] access_key: Access key ID for AWS.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Set of region names enabled for the cloud account.
        :param pulumi.Input[str] secret_key: AWS Secret Access Key
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[str] name: Name of AWS cloud account.
        :param pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]] tags: Set of tag keys and values to apply to the cloud account.  
               Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "regions", regions)
        pulumi.set(__self__, "secret_key", secret_key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        Access key ID for AWS.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Set of region names enabled for the cloud account.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        AWS Secret Access Key
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of AWS cloud account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]]:
        """
        Set of tag keys and values to apply to the cloud account.  
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AwsState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['AwsLinkArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Aws resources.
        :param pulumi.Input[str] access_key: Access key ID for AWS.
        :param pulumi.Input[str] created_at: Date when entity was created. Date and time format is ISO 8601 and UTC.
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[Sequence[pulumi.Input['AwsLinkArgs']]] links: HATEOAS of entity.
        :param pulumi.Input[str] name: Name of AWS cloud account.
        :param pulumi.Input[str] org_id: ID of organization that entity belongs to.
        :param pulumi.Input[str] owner: Email of entity owner.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Set of region names enabled for the cloud account.
        :param pulumi.Input[str] secret_key: AWS Secret Access Key
        :param pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]] tags: Set of tag keys and values to apply to the cloud account.  
               Example:[ { "key" : "vmware", "value": "provider" } ]
        :param pulumi.Input[str] updated_at: Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Access key ID for AWS.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

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
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsLinkArgs']]]]:
        """
        HATEOAS of entity.
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of AWS cloud account.
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
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of region names enabled for the cloud account.
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        AWS Secret Access Key
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]]:
        """
        Set of tag keys and values to apply to the cloud account.  
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AwsTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Aws(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsTagArgs']]]]] = None,
                 __props__=None):
        """
        Creates a VMware vRealize Automation AWS cloud account resource.

        ## Example Usage
        ### S

        The following example shows how to create an AWS cloud account resource.

        ```python
        import pulumi
        import pulumiverse_vra as vra

        this = vra.cloudaccount.Aws("this",
            description="terraform test cloud account aws",
            access_key=var["access_key"],
            secret_key=var["secret_key"],
            regions=[
                "us-east-1",
                "us-west-1",
            ],
            tags=[vra.cloudaccount.AwsTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        ## Import

        To import the AWS cloud account, use the ID as in the following example

        ```sh
         $ pulumi import vra:cloudaccount/aws:Aws new_aws 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access key ID for AWS.
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[str] name: Name of AWS cloud account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Set of region names enabled for the cloud account.
        :param pulumi.Input[str] secret_key: AWS Secret Access Key
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsTagArgs']]]] tags: Set of tag keys and values to apply to the cloud account.  
               Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AwsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a VMware vRealize Automation AWS cloud account resource.

        ## Example Usage
        ### S

        The following example shows how to create an AWS cloud account resource.

        ```python
        import pulumi
        import pulumiverse_vra as vra

        this = vra.cloudaccount.Aws("this",
            description="terraform test cloud account aws",
            access_key=var["access_key"],
            secret_key=var["secret_key"],
            regions=[
                "us-east-1",
                "us-west-1",
            ],
            tags=[vra.cloudaccount.AwsTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        ## Import

        To import the AWS cloud account, use the ID as in the following example

        ```sh
         $ pulumi import vra:cloudaccount/aws:Aws new_aws 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param AwsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AwsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AwsArgs.__new__(AwsArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if regions is None and not opts.urn:
                raise TypeError("Missing required property 'regions'")
            __props__.__dict__["regions"] = regions
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["org_id"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["updated_at"] = None
        super(Aws, __self__).__init__(
            'vra:cloudaccount/aws:Aws',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsLinkArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsTagArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Aws':
        """
        Get an existing Aws resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access key ID for AWS.
        :param pulumi.Input[str] created_at: Date when entity was created. Date and time format is ISO 8601 and UTC.
        :param pulumi.Input[str] description: Human-friendly description.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsLinkArgs']]]] links: HATEOAS of entity.
        :param pulumi.Input[str] name: Name of AWS cloud account.
        :param pulumi.Input[str] org_id: ID of organization that entity belongs to.
        :param pulumi.Input[str] owner: Email of entity owner.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: Set of region names enabled for the cloud account.
        :param pulumi.Input[str] secret_key: AWS Secret Access Key
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AwsTagArgs']]]] tags: Set of tag keys and values to apply to the cloud account.  
               Example:[ { "key" : "vmware", "value": "provider" } ]
        :param pulumi.Input[str] updated_at: Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AwsState.__new__(_AwsState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["links"] = links
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["owner"] = owner
        __props__.__dict__["regions"] = regions
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return Aws(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        Access key ID for AWS.
        """
        return pulumi.get(self, "access_key")

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
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.AwsLink']]:
        """
        HATEOAS of entity.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of AWS cloud account.
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
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[str]]:
        """
        Set of region names enabled for the cloud account.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        AWS Secret Access Key
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.AwsTag']]:
        """
        Set of tag keys and values to apply to the cloud account.  
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

