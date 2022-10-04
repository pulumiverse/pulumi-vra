// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace pulumiverse.Vra
{
    public static class GetNetwork
    {
        /// <summary>
        /// ## ---layout: "vra"
        /// 
        /// page_title: "VMware vRealize Automation: vra.Network"
        /// description: |-
        ///   Provides a data lookup for vra_network.
        /// ---
        /// 
        /// # Data Source: vra.Network
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a network data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test_network = Vra.GetNetwork.Invoke(new()
        ///     {
        ///         Name = @var.Network_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("vra:index/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// ## ---layout: "vra"
        /// 
        /// page_title: "VMware vRealize Automation: vra.Network"
        /// description: |-
        ///   Provides a data lookup for vra_network.
        /// ---
        /// 
        /// # Data Source: vra.Network
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a network data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test_network = Vra.GetNetwork.Invoke(new()
        ///     {
        ///         Name = @var.Network_name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("vra:index/getNetwork:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        [Input("constraints")]
        private List<Inputs.GetNetworkConstraintArgs>? _constraints;

        /// <summary>
        /// List of storage, network and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        public List<Inputs.GetNetworkConstraintArgs> Constraints
        {
            get => _constraints ?? (_constraints = new List<Inputs.GetNetworkConstraintArgs>());
            set => _constraints = value;
        }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private List<Inputs.GetNetworkTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public List<Inputs.GetNetworkTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetNetworkTagArgs>());
            set => _tags = value;
        }

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("constraints")]
        private InputList<Inputs.GetNetworkConstraintInputArgs>? _constraints;

        /// <summary>
        /// List of storage, network and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        public InputList<Inputs.GetNetworkConstraintInputArgs> Constraints
        {
            get => _constraints ?? (_constraints = new InputList<Inputs.GetNetworkConstraintInputArgs>());
            set => _constraints = value;
        }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetNetworkTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.GetNetworkTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetNetworkTagInputArgs>());
            set => _tags = value;
        }

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// IPv4 address range of the network in CIDR format.
        /// </summary>
        public readonly string Cidr;
        /// <summary>
        /// List of storage, network and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkConstraintResult> Constraints;
        /// <summary>
        /// Additional properties that may be used to extend the base resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> CustomProperties;
        /// <summary>
        /// Deployment id that is associated with this resource.
        /// </summary>
        public readonly string DeploymentId;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// External entity Id on the provider side.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The external zoneId of the resource.
        /// </summary>
        public readonly string ExternalZoneId;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkLinkResult> Links;
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty
        /// </summary>
        public readonly bool OutboundAccess;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// The id of the project this resource belongs to.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Self link of this request.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkTagResult> Tags;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetNetworkResult(
            string cidr,

            ImmutableArray<Outputs.GetNetworkConstraintResult> constraints,

            ImmutableDictionary<string, object> customProperties,

            string deploymentId,

            string description,

            string externalId,

            string externalZoneId,

            string id,

            ImmutableArray<Outputs.GetNetworkLinkResult> links,

            string name,

            string organizationId,

            bool outboundAccess,

            string owner,

            string projectId,

            string selfLink,

            ImmutableArray<Outputs.GetNetworkTagResult> tags,

            string updatedAt)
        {
            Cidr = cidr;
            Constraints = constraints;
            CustomProperties = customProperties;
            DeploymentId = deploymentId;
            Description = description;
            ExternalId = externalId;
            ExternalZoneId = externalZoneId;
            Id = id;
            Links = links;
            Name = name;
            OrganizationId = organizationId;
            OutboundAccess = outboundAccess;
            Owner = owner;
            ProjectId = projectId;
            SelfLink = selfLink;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
