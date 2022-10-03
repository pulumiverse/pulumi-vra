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
    public static class GetZone
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a zone data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test_zone = Vra.GetZone.Invoke(new()
        ///     {
        ///         Name = @var.Zone_name,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A zone data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZoneResult> InvokeAsync(GetZoneArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZoneResult>("vra:index/getZone:getZone", args ?? new GetZoneArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a zone data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test_zone = Vra.GetZone.Invoke(new()
        ///     {
        ///         Name = @var.Zone_name,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// A zone data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZoneResult> Invoke(GetZoneInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetZoneResult>("vra:index/getZone:getZone", args ?? new GetZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the zone resource instance.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier for the zone resource instance.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private List<Inputs.GetZoneTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public List<Inputs.GetZoneTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetZoneTagArgs>());
            set => _tags = value;
        }

        [Input("tagsToMatches")]
        private List<Inputs.GetZoneTagsToMatchArgs>? _tagsToMatches;

        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        public List<Inputs.GetZoneTagsToMatchArgs> TagsToMatches
        {
            get => _tagsToMatches ?? (_tagsToMatches = new List<Inputs.GetZoneTagsToMatchArgs>());
            set => _tagsToMatches = value;
        }

        public GetZoneArgs()
        {
        }
        public static new GetZoneArgs Empty => new GetZoneArgs();
    }

    public sealed class GetZoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the zone resource instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier for the zone resource instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetZoneTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.GetZoneTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetZoneTagInputArgs>());
            set => _tags = value;
        }

        [Input("tagsToMatches")]
        private InputList<Inputs.GetZoneTagsToMatchInputArgs>? _tagsToMatches;

        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        public InputList<Inputs.GetZoneTagsToMatchInputArgs> TagsToMatches
        {
            get => _tagsToMatches ?? (_tagsToMatches = new InputList<Inputs.GetZoneTagsToMatchInputArgs>());
            set => _tagsToMatches = value;
        }

        public GetZoneInvokeArgs()
        {
        }
        public static new GetZoneInvokeArgs Empty => new GetZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetZoneResult
    {
        /// <summary>
        /// The ID of the cloud account this zone belongs to.
        /// </summary>
        public readonly string CloudAccountId;
        /// <summary>
        /// The ids of the compute resources that has been explicitly assigned to this zone.
        /// </summary>
        public readonly ImmutableArray<string> ComputeIds;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// A list of key value pair of properties that will be used.
        /// </summary>
        public readonly ImmutableDictionary<string, object> CustomProperties;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the region for which this zone is defined.
        /// </summary>
        public readonly string ExternalRegionId;
        /// <summary>
        /// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
        /// </summary>
        public readonly string Folder;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneLinkResult> Links;
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
        /// The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`.
        /// </summary>
        public readonly string PlacementPolicy;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneTagResult> Tags;
        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneTagsToMatchResult> TagsToMatches;
        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetZoneResult(
            string cloudAccountId,

            ImmutableArray<string> computeIds,

            string createdAt,

            ImmutableDictionary<string, object> customProperties,

            string description,

            string externalRegionId,

            string folder,

            string id,

            ImmutableArray<Outputs.GetZoneLinkResult> links,

            string name,

            string orgId,

            string owner,

            string placementPolicy,

            ImmutableArray<Outputs.GetZoneTagResult> tags,

            ImmutableArray<Outputs.GetZoneTagsToMatchResult> tagsToMatches,

            string updatedAt)
        {
            CloudAccountId = cloudAccountId;
            ComputeIds = computeIds;
            CreatedAt = createdAt;
            CustomProperties = customProperties;
            Description = description;
            ExternalRegionId = externalRegionId;
            Folder = folder;
            Id = id;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            PlacementPolicy = placementPolicy;
            Tags = tags;
            TagsToMatches = tagsToMatches;
            UpdatedAt = updatedAt;
        }
    }
}
