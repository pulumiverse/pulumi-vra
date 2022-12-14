// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Region
{
    public static class GetEnumeration
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup a region enumeration data source.
        /// 
        /// DeprecationMessage: 'region_enumeration' is deprecated. Use 'region_enumeration_vsphere' instead.
        /// 
        /// **Region enumeration data source for vSphere**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Region.GetEnumerationVSphere.Invoke(new()
        ///     {
        ///         AcceptSelfSignedCert = false,
        ///         Dcid = @var.Vra_data_collector_id,
        ///         Hostname = @var.Hostname,
        ///         Password = @var.Password,
        ///         Username = @var.Username,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The region enumeration data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEnumerationResult> InvokeAsync(GetEnumerationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEnumerationResult>("vra:region/getEnumeration:getEnumeration", args ?? new GetEnumerationArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup a region enumeration data source.
        /// 
        /// DeprecationMessage: 'region_enumeration' is deprecated. Use 'region_enumeration_vsphere' instead.
        /// 
        /// **Region enumeration data source for vSphere**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Region.GetEnumerationVSphere.Invoke(new()
        ///     {
        ///         AcceptSelfSignedCert = false,
        ///         Dcid = @var.Vra_data_collector_id,
        ///         Hostname = @var.Hostname,
        ///         Password = @var.Password,
        ///         Username = @var.Username,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// The region enumeration data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEnumerationResult> Invoke(GetEnumerationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEnumerationResult>("vra:region/getEnumeration:getEnumeration", args ?? new GetEnumerationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnumerationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Accept self signed certificate when connecting to vSphere. Example: false
        /// </summary>
        [Input("acceptSelfSignedCert")]
        public bool? AcceptSelfSignedCert { get; set; }

        /// <summary>
        /// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
        /// </summary>
        [Input("dcid")]
        public string? Dcid { get; set; }

        /// <summary>
        /// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
        /// </summary>
        [Input("hostname", required: true)]
        public string Hostname { get; set; } = null!;

        /// <summary>
        /// Password for the user used to authenticate with the cloud Account
        /// </summary>
        [Input("password", required: true)]
        public string Password { get; set; } = null!;

        /// <summary>
        /// Username to authenticate with the cloud account
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetEnumerationArgs()
        {
        }
        public static new GetEnumerationArgs Empty => new GetEnumerationArgs();
    }

    public sealed class GetEnumerationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Accept self signed certificate when connecting to vSphere. Example: false
        /// </summary>
        [Input("acceptSelfSignedCert")]
        public Input<bool>? AcceptSelfSignedCert { get; set; }

        /// <summary>
        /// ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
        /// </summary>
        [Input("dcid")]
        public Input<string>? Dcid { get; set; }

        /// <summary>
        /// Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// Password for the user used to authenticate with the cloud Account
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Username to authenticate with the cloud account
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetEnumerationInvokeArgs()
        {
        }
        public static new GetEnumerationInvokeArgs Empty => new GetEnumerationInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnumerationResult
    {
        public readonly bool? AcceptSelfSignedCert;
        public readonly string? Dcid;
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Password;
        /// <summary>
        /// A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        public readonly string Username;

        [OutputConstructor]
        private GetEnumerationResult(
            bool? acceptSelfSignedCert,

            string? dcid,

            string hostname,

            string id,

            string password,

            ImmutableArray<string> regions,

            string username)
        {
            AcceptSelfSignedCert = acceptSelfSignedCert;
            Dcid = dcid;
            Hostname = hostname;
            Id = id;
            Password = password;
            Regions = regions;
            Username = username;
        }
    }
}
