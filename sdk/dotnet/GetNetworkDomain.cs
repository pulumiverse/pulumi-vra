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
    public static class GetNetworkDomain
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// This is an example of how to lookup Network domain objects.
        /// 
        /// **Network Domain by filter query:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetNetworkDomain.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A network domain object supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkDomainResult> InvokeAsync(GetNetworkDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkDomainResult>("vra:index/getNetworkDomain:getNetworkDomain", args ?? new GetNetworkDomainArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// This is an example of how to lookup Network domain objects.
        /// 
        /// **Network Domain by filter query:**
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetNetworkDomain.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A network domain object supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkDomainResult> Invoke(GetNetworkDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkDomainResult>("vra:index/getNetworkDomain:getNetworkDomain", args ?? new GetNetworkDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkDomainArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the network domain objects.
        /// </summary>
        [Input("filter", required: true)]
        public string Filter { get; set; } = null!;

        [Input("tags")]
        private List<Inputs.GetNetworkDomainTagArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public List<Inputs.GetNetworkDomainTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetNetworkDomainTagArgs>());
            set => _tags = value;
        }

        public GetNetworkDomainArgs()
        {
        }
        public static new GetNetworkDomainArgs Empty => new GetNetworkDomainArgs();
    }

    public sealed class GetNetworkDomainInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search criteria to narrow down the network domain objects.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.GetNetworkDomainTagInputArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.GetNetworkDomainTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetNetworkDomainTagInputArgs>());
            set => _tags = value;
        }

        public GetNetworkDomainInvokeArgs()
        {
        }
        public static new GetNetworkDomainInvokeArgs Empty => new GetNetworkDomainInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkDomainResult
    {
        public readonly string Cidr;
        /// <summary>
        /// Set of ids of the cloud accounts this entity belongs to.
        /// </summary>
        public readonly ImmutableArray<string> CloudAccountIds;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        public readonly ImmutableDictionary<string, object> CustomProperties;
        /// <summary>
        /// A human-friendly description of the fabric vSphere storage account.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// External entity Id on the provider side.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The id of the region for which this entity is defined.
        /// </summary>
        public readonly string ExternalRegionId;
        public readonly string Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkDomainLinkResult> Links;
        /// <summary>
        /// Name of the network domain.
        /// </summary>
        public readonly string Name;
        public readonly string OrganizationId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Set of tag keys and values to apply to the resource.
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkDomainTagResult> Tags;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetNetworkDomainResult(
            string cidr,

            ImmutableArray<string> cloudAccountIds,

            string createdAt,

            ImmutableDictionary<string, object> customProperties,

            string description,

            string externalId,

            string externalRegionId,

            string filter,

            string id,

            ImmutableArray<Outputs.GetNetworkDomainLinkResult> links,

            string name,

            string organizationId,

            string owner,

            ImmutableArray<Outputs.GetNetworkDomainTagResult> tags,

            string updatedAt)
        {
            Cidr = cidr;
            CloudAccountIds = cloudAccountIds;
            CreatedAt = createdAt;
            CustomProperties = customProperties;
            Description = description;
            ExternalId = externalId;
            ExternalRegionId = externalRegionId;
            Filter = filter;
            Id = id;
            Links = links;
            Name = name;
            OrganizationId = organizationId;
            Owner = owner;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
