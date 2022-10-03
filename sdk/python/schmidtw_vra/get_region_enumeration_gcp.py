# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRegionEnumerationGcpResult',
    'AwaitableGetRegionEnumerationGcpResult',
    'get_region_enumeration_gcp',
    'get_region_enumeration_gcp_output',
]

@pulumi.output_type
class GetRegionEnumerationGcpResult:
    """
    A collection of values returned by getRegionEnumerationGcp.
    """
    def __init__(__self__, client_email=None, id=None, private_key=None, private_key_id=None, project_id=None, regions=None):
        if client_email and not isinstance(client_email, str):
            raise TypeError("Expected argument 'client_email' to be a str")
        pulumi.set(__self__, "client_email", client_email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if private_key_id and not isinstance(private_key_id, str):
            raise TypeError("Expected argument 'private_key_id' to be a str")
        pulumi.set(__self__, "private_key_id", private_key_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)

    @property
    @pulumi.getter(name="clientEmail")
    def client_email(self) -> str:
        return pulumi.get(self, "client_email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyId")
    def private_key_id(self) -> str:
        return pulumi.get(self, "private_key_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[str]:
        """
        A set of Region names to enable provisioning on. Example: northamerica-northeast1
        """
        return pulumi.get(self, "regions")


class AwaitableGetRegionEnumerationGcpResult(GetRegionEnumerationGcpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionEnumerationGcpResult(
            client_email=self.client_email,
            id=self.id,
            private_key=self.private_key,
            private_key_id=self.private_key_id,
            project_id=self.project_id,
            regions=self.regions)


def get_region_enumeration_gcp(client_email: Optional[str] = None,
                               private_key: Optional[str] = None,
                               private_key_id: Optional[str] = None,
                               project_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionEnumerationGcpResult:
    """
    ## Example Usage
    ### S

    This is an example of how to lookup a region enumeration data source for GCP cloud account.

    **Region enumeration data source for GCP**
    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.get_region_enumeration_gcp(client_email=var["client_email"],
        private_key_id=var["private_key_id"],
        private_key=var["private_key"],
        project_id=var["project_id"])
    ```

    The region enumeration data source for GCP cloud account supports the following arguments:


    :param str client_email: Client E-mail ID.
    :param str private_key: GCP Private key.
    :param str private_key_id: GCP Private key ID.
    :param str project_id: GCP Project ID.
    """
    __args__ = dict()
    __args__['clientEmail'] = client_email
    __args__['privateKey'] = private_key
    __args__['privateKeyId'] = private_key_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vra:index/getRegionEnumerationGcp:getRegionEnumerationGcp', __args__, opts=opts, typ=GetRegionEnumerationGcpResult).value

    return AwaitableGetRegionEnumerationGcpResult(
        client_email=__ret__.client_email,
        id=__ret__.id,
        private_key=__ret__.private_key,
        private_key_id=__ret__.private_key_id,
        project_id=__ret__.project_id,
        regions=__ret__.regions)


@_utilities.lift_output_func(get_region_enumeration_gcp)
def get_region_enumeration_gcp_output(client_email: Optional[pulumi.Input[str]] = None,
                                      private_key: Optional[pulumi.Input[str]] = None,
                                      private_key_id: Optional[pulumi.Input[str]] = None,
                                      project_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionEnumerationGcpResult]:
    """
    ## Example Usage
    ### S

    This is an example of how to lookup a region enumeration data source for GCP cloud account.

    **Region enumeration data source for GCP**
    ```python
    import pulumi
    import pulumi_vra as vra

    this = vra.get_region_enumeration_gcp(client_email=var["client_email"],
        private_key_id=var["private_key_id"],
        private_key=var["private_key"],
        project_id=var["project_id"])
    ```

    The region enumeration data source for GCP cloud account supports the following arguments:


    :param str client_email: Client E-mail ID.
    :param str private_key: GCP Private key.
    :param str private_key_id: GCP Private key ID.
    :param str project_id: GCP Project ID.
    """
    ...
