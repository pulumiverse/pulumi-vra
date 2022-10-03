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

__all__ = ['CatalogSourceEntitlementArgs', 'CatalogSourceEntitlement']

@pulumi.input_type
class CatalogSourceEntitlementArgs:
    def __init__(__self__, *,
                 catalog_source_id: pulumi.Input[str],
                 project_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a CatalogSourceEntitlement resource.
        :param pulumi.Input[str] catalog_source_id: The id of the catalog source to create the entitlement.
        :param pulumi.Input[str] project_id: The id of the project this entity belongs to.
        """
        pulumi.set(__self__, "catalog_source_id", catalog_source_id)
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="catalogSourceId")
    def catalog_source_id(self) -> pulumi.Input[str]:
        """
        The id of the catalog source to create the entitlement.
        """
        return pulumi.get(self, "catalog_source_id")

    @catalog_source_id.setter
    def catalog_source_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_source_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The id of the project this entity belongs to.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _CatalogSourceEntitlementState:
    def __init__(__self__, *,
                 catalog_source_id: Optional[pulumi.Input[str]] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogSourceEntitlementDefinitionArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CatalogSourceEntitlement resources.
        :param pulumi.Input[str] catalog_source_id: The id of the catalog source to create the entitlement.
        :param pulumi.Input[str] project_id: The id of the project this entity belongs to.
        """
        if catalog_source_id is not None:
            pulumi.set(__self__, "catalog_source_id", catalog_source_id)
        if definitions is not None:
            pulumi.set(__self__, "definitions", definitions)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="catalogSourceId")
    def catalog_source_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the catalog source to create the entitlement.
        """
        return pulumi.get(self, "catalog_source_id")

    @catalog_source_id.setter
    def catalog_source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_source_id", value)

    @property
    @pulumi.getter
    def definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CatalogSourceEntitlementDefinitionArgs']]]]:
        return pulumi.get(self, "definitions")

    @definitions.setter
    def definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogSourceEntitlementDefinitionArgs']]]]):
        pulumi.set(self, "definitions", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the project this entity belongs to.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


class CatalogSourceEntitlement(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_source_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource provides a way to create a catalog source entitlement in VMware vRealize Automation.

        ## Example Usage
        ### S

        ```python
        import pulumi
        import schmidtw_vra as vra

        this = vra.CatalogSourceEntitlement("this",
            catalog_source_id=var["catalog_source_blueprint_id"],
            project_id=var["project_id"])
        ```
        ## Attribute Reference

        * `definition` - Represents a catalog source that is linked to a project via an entitlement.
          
            * `description` - Description of the catalog source.
          
            * `icon_id` - Icon id of associated catalog source.
          
            * `id` - Id of the catalog source.
          
            * `name` - Name of the catalog source.
          
            * `number_of_items` - Number of items in the associated catalog source.
          
            * `source_name` - Catalog source name.
          
            * `source_type` - Catalog source type.
          
            * `type` - Content definition type.

        ## Import

        Catalog source entitlement can be imported using the id, e.g.

        ```sh
         $ pulumi import vra:index/catalogSourceEntitlement:CatalogSourceEntitlement this 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_source_id: The id of the catalog source to create the entitlement.
        :param pulumi.Input[str] project_id: The id of the project this entity belongs to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CatalogSourceEntitlementArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource provides a way to create a catalog source entitlement in VMware vRealize Automation.

        ## Example Usage
        ### S

        ```python
        import pulumi
        import schmidtw_vra as vra

        this = vra.CatalogSourceEntitlement("this",
            catalog_source_id=var["catalog_source_blueprint_id"],
            project_id=var["project_id"])
        ```
        ## Attribute Reference

        * `definition` - Represents a catalog source that is linked to a project via an entitlement.
          
            * `description` - Description of the catalog source.
          
            * `icon_id` - Icon id of associated catalog source.
          
            * `id` - Id of the catalog source.
          
            * `name` - Name of the catalog source.
          
            * `number_of_items` - Number of items in the associated catalog source.
          
            * `source_name` - Catalog source name.
          
            * `source_type` - Catalog source type.
          
            * `type` - Content definition type.

        ## Import

        Catalog source entitlement can be imported using the id, e.g.

        ```sh
         $ pulumi import vra:index/catalogSourceEntitlement:CatalogSourceEntitlement this 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param CatalogSourceEntitlementArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CatalogSourceEntitlementArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_source_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CatalogSourceEntitlementArgs.__new__(CatalogSourceEntitlementArgs)

            if catalog_source_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_source_id'")
            __props__.__dict__["catalog_source_id"] = catalog_source_id
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["definitions"] = None
        super(CatalogSourceEntitlement, __self__).__init__(
            'vra:index/catalogSourceEntitlement:CatalogSourceEntitlement',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_source_id: Optional[pulumi.Input[str]] = None,
            definitions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CatalogSourceEntitlementDefinitionArgs']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None) -> 'CatalogSourceEntitlement':
        """
        Get an existing CatalogSourceEntitlement resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_source_id: The id of the catalog source to create the entitlement.
        :param pulumi.Input[str] project_id: The id of the project this entity belongs to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CatalogSourceEntitlementState.__new__(_CatalogSourceEntitlementState)

        __props__.__dict__["catalog_source_id"] = catalog_source_id
        __props__.__dict__["definitions"] = definitions
        __props__.__dict__["project_id"] = project_id
        return CatalogSourceEntitlement(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogSourceId")
    def catalog_source_id(self) -> pulumi.Output[str]:
        """
        The id of the catalog source to create the entitlement.
        """
        return pulumi.get(self, "catalog_source_id")

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Output[Sequence['outputs.CatalogSourceEntitlementDefinition']]:
        return pulumi.get(self, "definitions")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The id of the project this entity belongs to.
        """
        return pulumi.get(self, "project_id")

