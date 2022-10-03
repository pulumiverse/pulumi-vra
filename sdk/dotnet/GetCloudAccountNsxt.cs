// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra
{
    public static class GetCloudAccountNsxt
    {
        /// <summary>
        /// Provides a VMware vRA vra.CloudAccountNsxt data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// **NSX-T cloud account data source by its id:**
        /// 
        /// This is an example of how to read the cloud account data source using its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountNsxt.Invoke(new()
        ///     {
        ///         Id = @var.Vra_cloud_account_nsxt_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **NSX-T cloud account data source by its name:**
        /// 
        /// This is an example of how to read the cloud account data source using its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountNsxt.Invoke(new()
        ///     {
        ///         Name = @var.Vra_cloud_account_nsxt_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudAccountNsxtResult> InvokeAsync(GetCloudAccountNsxtArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudAccountNsxtResult>("vra:index/getCloudAccountNsxt:getCloudAccountNsxt", args ?? new GetCloudAccountNsxtArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware vRA vra.CloudAccountNsxt data source.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// **NSX-T cloud account data source by its id:**
        /// 
        /// This is an example of how to read the cloud account data source using its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountNsxt.Invoke(new()
        ///     {
        ///         Id = @var.Vra_cloud_account_nsxt_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **NSX-T cloud account data source by its name:**
        /// 
        /// This is an example of how to read the cloud account data source using its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCloudAccountNsxt.Invoke(new()
        ///     {
        ///         Name = @var.Vra_cloud_account_nsxt_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudAccountNsxtResult> Invoke(GetCloudAccountNsxtInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudAccountNsxtResult>("vra:index/getCloudAccountNsxt:getCloudAccountNsxt", args ?? new GetCloudAccountNsxtInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudAccountNsxtArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of this NSX-T cloud account.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of this NSX-T cloud account.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private List<Inputs.GetCloudAccountNsxtTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public List<Inputs.GetCloudAccountNsxtTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetCloudAccountNsxtTagArgs>());
            set => _tags = value;
        }

        public GetCloudAccountNsxtArgs()
        {
        }
        public static new GetCloudAccountNsxtArgs Empty => new GetCloudAccountNsxtArgs();
    }

    public sealed class GetCloudAccountNsxtInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of this NSX-T cloud account.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of this NSX-T cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetCloudAccountNsxtTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.GetCloudAccountNsxtTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetCloudAccountNsxtTagInputArgs>());
            set => _tags = value;
        }

        public GetCloudAccountNsxtInvokeArgs()
        {
        }
        public static new GetCloudAccountNsxtInvokeArgs Empty => new GetCloudAccountNsxtInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudAccountNsxtResult
    {
        /// <summary>
        /// Cloud accounts associated with this cloud account.
        /// </summary>
        public readonly ImmutableArray<string> AssociatedCloudAccountIds;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Identifier of a data collector vm deployed in the on premise infrastructure.
        /// </summary>
        public readonly string DcId;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Host name for the NSX-T cloud account.
        /// </summary>
        public readonly string Hostname;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudAccountNsxtLinkResult> Links;
        public readonly bool ManagerMode;
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudAccountNsxtTagResult> Tags;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// Username to authenticate with the cloud account.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetCloudAccountNsxtResult(
            ImmutableArray<string> associatedCloudAccountIds,

            string createdAt,

            string dcId,

            string description,

            string hostname,

            string id,

            ImmutableArray<Outputs.GetCloudAccountNsxtLinkResult> links,

            bool managerMode,

            string name,

            string orgId,

            string owner,

            ImmutableArray<Outputs.GetCloudAccountNsxtTagResult> tags,

            string updatedAt,

            string username)
        {
            AssociatedCloudAccountIds = associatedCloudAccountIds;
            CreatedAt = createdAt;
            DcId = dcId;
            Description = description;
            Hostname = hostname;
            Id = id;
            Links = links;
            ManagerMode = managerMode;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            Tags = tags;
            UpdatedAt = updatedAt;
            Username = username;
        }
    }
}
