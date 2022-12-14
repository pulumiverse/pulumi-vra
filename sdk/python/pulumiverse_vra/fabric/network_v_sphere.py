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

__all__ = ['NetworkVSphereArgs', 'NetworkVSphere']

@pulumi.input_type
class NetworkVSphereArgs:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 default_ipv6_gateway: Optional[pulumi.Input[str]] = None,
                 dns_search_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_server_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]] = None):
        """
        The set of arguments for constructing a NetworkVSphere resource.
        :param pulumi.Input[str] cidr: Network CIDR to be used.
        :param pulumi.Input[str] default_gateway: IPv4 default gateway to be used.
        :param pulumi.Input[str] default_ipv6_gateway: IPv6 default gateway to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_search_domains: List of dns search domains for the vSphere network.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_server_addresses: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] domain: Domain for the vSphere network.
        :param pulumi.Input[str] ipv6_cidr: Network IPv6 CIDR to be used.
        :param pulumi.Input[bool] is_default: Indicates whether this is the default subnet for the zone.
        :param pulumi.Input[bool] is_public: Indicates whether the sub-network supports public IP assignment.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]] tags: Set of tag keys and values to apply to the resource.
               Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if default_gateway is not None:
            pulumi.set(__self__, "default_gateway", default_gateway)
        if default_ipv6_gateway is not None:
            pulumi.set(__self__, "default_ipv6_gateway", default_ipv6_gateway)
        if dns_search_domains is not None:
            pulumi.set(__self__, "dns_search_domains", dns_search_domains)
        if dns_server_addresses is not None:
            pulumi.set(__self__, "dns_server_addresses", dns_server_addresses)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if ipv6_cidr is not None:
            pulumi.set(__self__, "ipv6_cidr", ipv6_cidr)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Network CIDR to be used.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 default gateway to be used.
        """
        return pulumi.get(self, "default_gateway")

    @default_gateway.setter
    def default_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_gateway", value)

    @property
    @pulumi.getter(name="defaultIpv6Gateway")
    def default_ipv6_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 default gateway to be used.
        """
        return pulumi.get(self, "default_ipv6_gateway")

    @default_ipv6_gateway.setter
    def default_ipv6_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_ipv6_gateway", value)

    @property
    @pulumi.getter(name="dnsSearchDomains")
    def dns_search_domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of dns search domains for the vSphere network.
        """
        return pulumi.get(self, "dns_search_domains")

    @dns_search_domains.setter
    def dns_search_domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_search_domains", value)

    @property
    @pulumi.getter(name="dnsServerAddresses")
    def dns_server_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "dns_server_addresses")

    @dns_server_addresses.setter
    def dns_server_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_server_addresses", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain for the vSphere network.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="ipv6Cidr")
    def ipv6_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Network IPv6 CIDR to be used.
        """
        return pulumi.get(self, "ipv6_cidr")

    @ipv6_cidr.setter
    def ipv6_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this is the default subnet for the zone.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the sub-network supports public IP assignment.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]]:
        """
        Set of tag keys and values to apply to the resource.
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _NetworkVSphereState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 cloud_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 custom_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 default_ipv6_gateway: Optional[pulumi.Input[str]] = None,
                 dns_search_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_server_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 external_id: Optional[pulumi.Input[str]] = None,
                 external_region_id: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereLinkArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkVSphere resources.
        :param pulumi.Input[str] cidr: Network CIDR to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cloud_account_ids: Set of ids of the cloud accounts this entity belongs to.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[str] default_gateway: IPv4 default gateway to be used.
        :param pulumi.Input[str] default_ipv6_gateway: IPv6 default gateway to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_search_domains: List of dns search domains for the vSphere network.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_server_addresses: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] domain: Domain for the vSphere network.
        :param pulumi.Input[str] external_id: External entity Id on the provider side.
        :param pulumi.Input[str] external_region_id: The id of the region for which this network is defined.
        :param pulumi.Input[str] ipv6_cidr: Network IPv6 CIDR to be used.
        :param pulumi.Input[bool] is_default: Indicates whether this is the default subnet for the zone.
        :param pulumi.Input[bool] is_public: Indicates whether the sub-network supports public IP assignment.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkVSphereLinkArgs']]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] organization_id: ID of organization that entity belongs to.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]] tags: Set of tag keys and values to apply to the resource.
               Example:[ { "key" : "vmware", "value": "provider" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if cloud_account_ids is not None:
            pulumi.set(__self__, "cloud_account_ids", cloud_account_ids)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if custom_properties is not None:
            pulumi.set(__self__, "custom_properties", custom_properties)
        if default_gateway is not None:
            pulumi.set(__self__, "default_gateway", default_gateway)
        if default_ipv6_gateway is not None:
            pulumi.set(__self__, "default_ipv6_gateway", default_ipv6_gateway)
        if dns_search_domains is not None:
            pulumi.set(__self__, "dns_search_domains", dns_search_domains)
        if dns_server_addresses is not None:
            pulumi.set(__self__, "dns_server_addresses", dns_server_addresses)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if external_region_id is not None:
            pulumi.set(__self__, "external_region_id", external_region_id)
        if ipv6_cidr is not None:
            pulumi.set(__self__, "ipv6_cidr", ipv6_cidr)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Network CIDR to be used.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="cloudAccountIds")
    def cloud_account_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of ids of the cloud accounts this entity belongs to.
        """
        return pulumi.get(self, "cloud_account_ids")

    @cloud_account_ids.setter
    def cloud_account_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cloud_account_ids", value)

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
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "custom_properties")

    @custom_properties.setter
    def custom_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "custom_properties", value)

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 default gateway to be used.
        """
        return pulumi.get(self, "default_gateway")

    @default_gateway.setter
    def default_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_gateway", value)

    @property
    @pulumi.getter(name="defaultIpv6Gateway")
    def default_ipv6_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 default gateway to be used.
        """
        return pulumi.get(self, "default_ipv6_gateway")

    @default_ipv6_gateway.setter
    def default_ipv6_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_ipv6_gateway", value)

    @property
    @pulumi.getter(name="dnsSearchDomains")
    def dns_search_domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of dns search domains for the vSphere network.
        """
        return pulumi.get(self, "dns_search_domains")

    @dns_search_domains.setter
    def dns_search_domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_search_domains", value)

    @property
    @pulumi.getter(name="dnsServerAddresses")
    def dns_server_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "dns_server_addresses")

    @dns_server_addresses.setter
    def dns_server_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_server_addresses", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Domain for the vSphere network.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[str]]:
        """
        External entity Id on the provider side.
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_id", value)

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the region for which this network is defined.
        """
        return pulumi.get(self, "external_region_id")

    @external_region_id.setter
    def external_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_region_id", value)

    @property
    @pulumi.getter(name="ipv6Cidr")
    def ipv6_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Network IPv6 CIDR to be used.
        """
        return pulumi.get(self, "ipv6_cidr")

    @ipv6_cidr.setter
    def ipv6_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this is the default subnet for the zone.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the sub-network supports public IP assignment.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereLinkArgs']]]]:
        """
        HATEOAS of the entity
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereLinkArgs']]]]):
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
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of organization that entity belongs to.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]]:
        """
        Set of tag keys and values to apply to the resource.
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkVSphereTagArgs']]]]):
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


class NetworkVSphere(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 default_ipv6_gateway: Optional[pulumi.Input[str]] = None,
                 dns_search_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_server_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereTagArgs']]]]] = None,
                 __props__=None):
        """
        Updates a VMware vRealize Automation fabric_network_vsphere resource.

        ## Example Usage
        ### S

        You cannot create a vSphere fabric network resource, however you can import using the command specified in the import section below.
        Once a resource is imported, you can update it as shown below:

        ```python
        import pulumi
        import pulumiverse_vra as vra

        simple = vra.fabric.NetworkVSphere("simple",
            cidr=var["cidr"],
            default_gateway=var["gateway"],
            domain=var["domain"],
            tags=[vra.fabric.NetworkVSphereTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        ## Import

        To import the vSphere fabric network resource, use the ID as in the following example

        ```sh
         $ pulumi import vra:fabric/networkVSphere:NetworkVSphere new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: Network CIDR to be used.
        :param pulumi.Input[str] default_gateway: IPv4 default gateway to be used.
        :param pulumi.Input[str] default_ipv6_gateway: IPv6 default gateway to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_search_domains: List of dns search domains for the vSphere network.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_server_addresses: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] domain: Domain for the vSphere network.
        :param pulumi.Input[str] ipv6_cidr: Network IPv6 CIDR to be used.
        :param pulumi.Input[bool] is_default: Indicates whether this is the default subnet for the zone.
        :param pulumi.Input[bool] is_public: Indicates whether the sub-network supports public IP assignment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereTagArgs']]]] tags: Set of tag keys and values to apply to the resource.
               Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkVSphereArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Updates a VMware vRealize Automation fabric_network_vsphere resource.

        ## Example Usage
        ### S

        You cannot create a vSphere fabric network resource, however you can import using the command specified in the import section below.
        Once a resource is imported, you can update it as shown below:

        ```python
        import pulumi
        import pulumiverse_vra as vra

        simple = vra.fabric.NetworkVSphere("simple",
            cidr=var["cidr"],
            default_gateway=var["gateway"],
            domain=var["domain"],
            tags=[vra.fabric.NetworkVSphereTagArgs(
                key="foo",
                value="bar",
            )])
        ```

        ## Import

        To import the vSphere fabric network resource, use the ID as in the following example

        ```sh
         $ pulumi import vra:fabric/networkVSphere:NetworkVSphere new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
        ```

        :param str resource_name: The name of the resource.
        :param NetworkVSphereArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkVSphereArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 default_gateway: Optional[pulumi.Input[str]] = None,
                 default_ipv6_gateway: Optional[pulumi.Input[str]] = None,
                 dns_search_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_server_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkVSphereArgs.__new__(NetworkVSphereArgs)

            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["default_gateway"] = default_gateway
            __props__.__dict__["default_ipv6_gateway"] = default_ipv6_gateway
            __props__.__dict__["dns_search_domains"] = dns_search_domains
            __props__.__dict__["dns_server_addresses"] = dns_server_addresses
            __props__.__dict__["domain"] = domain
            __props__.__dict__["ipv6_cidr"] = ipv6_cidr
            __props__.__dict__["is_default"] = is_default
            __props__.__dict__["is_public"] = is_public
            __props__.__dict__["tags"] = tags
            __props__.__dict__["cloud_account_ids"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["custom_properties"] = None
            __props__.__dict__["external_id"] = None
            __props__.__dict__["external_region_id"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["org_id"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["updated_at"] = None
        super(NetworkVSphere, __self__).__init__(
            'vra:fabric/networkVSphere:NetworkVSphere',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            cloud_account_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            custom_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            default_gateway: Optional[pulumi.Input[str]] = None,
            default_ipv6_gateway: Optional[pulumi.Input[str]] = None,
            dns_search_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            dns_server_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            external_id: Optional[pulumi.Input[str]] = None,
            external_region_id: Optional[pulumi.Input[str]] = None,
            ipv6_cidr: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereLinkArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereTagArgs']]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'NetworkVSphere':
        """
        Get an existing NetworkVSphere resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: Network CIDR to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cloud_account_ids: Set of ids of the cloud accounts this entity belongs to.
        :param pulumi.Input[str] created_at: Date when the entity was created. The date is in ISO 6801 and UTC.
        :param pulumi.Input[str] default_gateway: IPv4 default gateway to be used.
        :param pulumi.Input[str] default_ipv6_gateway: IPv6 default gateway to be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_search_domains: List of dns search domains for the vSphere network.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_server_addresses: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] domain: Domain for the vSphere network.
        :param pulumi.Input[str] external_id: External entity Id on the provider side.
        :param pulumi.Input[str] external_region_id: The id of the region for which this network is defined.
        :param pulumi.Input[str] ipv6_cidr: Network IPv6 CIDR to be used.
        :param pulumi.Input[bool] is_default: Indicates whether this is the default subnet for the zone.
        :param pulumi.Input[bool] is_public: Indicates whether the sub-network supports public IP assignment.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereLinkArgs']]]] links: HATEOAS of the entity
        :param pulumi.Input[str] name: A human-friendly name used as an identifier in APIs that support this option.
        :param pulumi.Input[str] organization_id: ID of organization that entity belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkVSphereTagArgs']]]] tags: Set of tag keys and values to apply to the resource.
               Example:[ { "key" : "vmware", "value": "provider" } ]
        :param pulumi.Input[str] updated_at: Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkVSphereState.__new__(_NetworkVSphereState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["cloud_account_ids"] = cloud_account_ids
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["custom_properties"] = custom_properties
        __props__.__dict__["default_gateway"] = default_gateway
        __props__.__dict__["default_ipv6_gateway"] = default_ipv6_gateway
        __props__.__dict__["dns_search_domains"] = dns_search_domains
        __props__.__dict__["dns_server_addresses"] = dns_server_addresses
        __props__.__dict__["domain"] = domain
        __props__.__dict__["external_id"] = external_id
        __props__.__dict__["external_region_id"] = external_region_id
        __props__.__dict__["ipv6_cidr"] = ipv6_cidr
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["is_public"] = is_public
        __props__.__dict__["links"] = links
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        return NetworkVSphere(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        Network CIDR to be used.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="cloudAccountIds")
    def cloud_account_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Set of ids of the cloud accounts this entity belongs to.
        """
        return pulumi.get(self, "cloud_account_ids")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was created. The date is in ISO 6801 and UTC.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter(name="defaultGateway")
    def default_gateway(self) -> pulumi.Output[str]:
        """
        IPv4 default gateway to be used.
        """
        return pulumi.get(self, "default_gateway")

    @property
    @pulumi.getter(name="defaultIpv6Gateway")
    def default_ipv6_gateway(self) -> pulumi.Output[str]:
        """
        IPv6 default gateway to be used.
        """
        return pulumi.get(self, "default_ipv6_gateway")

    @property
    @pulumi.getter(name="dnsSearchDomains")
    def dns_search_domains(self) -> pulumi.Output[Sequence[str]]:
        """
        List of dns search domains for the vSphere network.
        """
        return pulumi.get(self, "dns_search_domains")

    @property
    @pulumi.getter(name="dnsServerAddresses")
    def dns_server_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A human-friendly name used as an identifier in APIs that support this option.
        """
        return pulumi.get(self, "dns_server_addresses")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[Optional[str]]:
        """
        Domain for the vSphere network.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[str]:
        """
        External entity Id on the provider side.
        """
        return pulumi.get(self, "external_id")

    @property
    @pulumi.getter(name="externalRegionId")
    def external_region_id(self) -> pulumi.Output[str]:
        """
        The id of the region for which this network is defined.
        """
        return pulumi.get(self, "external_region_id")

    @property
    @pulumi.getter(name="ipv6Cidr")
    def ipv6_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        Network IPv6 CIDR to be used.
        """
        return pulumi.get(self, "ipv6_cidr")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this is the default subnet for the zone.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the sub-network supports public IP assignment.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.NetworkVSphereLink']]:
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
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        ID of organization that entity belongs to.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.NetworkVSphereTag']]:
        """
        Set of tag keys and values to apply to the resource.
        Example:[ { "key" : "vmware", "value": "provider" } ]
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date when the entity was last updated. The date is ISO 8601 and UTC.
        """
        return pulumi.get(self, "updated_at")

