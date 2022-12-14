// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Zone
{
    /// <summary>
    /// ## Example Usage
    /// ### S
    /// 
    /// This is an example of how to create a zone resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Zone.Zone("this", new()
    ///     {
    ///         Description = "my terraform test cloud zone",
    ///         RegionId = data.Vra_region.This.Id,
    ///         Tags = new[]
    ///         {
    ///             new Vra.Zone.Inputs.ZoneTagArgs
    ///             {
    ///                 Key = "my-tf-key",
    ///                 Value = "my-tf-value",
    ///             },
    ///             new Vra.Zone.Inputs.ZoneTagArgs
    ///             {
    ///                 Key = "tf-foo",
    ///                 Value = "tf-bar",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// A zone resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:zone/zone:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the cloud account this zone belongs to.
        /// </summary>
        [Output("cloudAccountId")]
        public Output<string> CloudAccountId { get; private set; } = null!;

        /// <summary>
        /// The ids of the compute resources that will be explicitly assigned to this zone.
        /// </summary>
        [Output("computeIds")]
        public Output<ImmutableArray<string>> ComputeIds { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A list of key value pair of properties that will be used.
        /// </summary>
        [Output("customProperties")]
        public Output<ImmutableDictionary<string, object>> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The id of the region for which this zone is defined.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
        /// </summary>
        [Output("folder")]
        public Output<string?> Folder { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.ZoneLink>> Links { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier for the zone resource instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`. Default is `DEFAULT`.
        /// </summary>
        [Output("placementPolicy")]
        public Output<string?> PlacementPolicy { get; private set; } = null!;

        /// <summary>
        /// The id of the region for which this zone is created.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ZoneTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        [Output("tagsToMatches")]
        public Output<ImmutableArray<Outputs.ZoneTagsToMatch>> TagsToMatches { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("vra:zone/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("vra:zone/zone:Zone", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        [Input("computeIds")]
        private InputList<string>? _computeIds;

        /// <summary>
        /// The ids of the compute resources that will be explicitly assigned to this zone.
        /// </summary>
        public InputList<string> ComputeIds
        {
            get => _computeIds ?? (_computeIds = new InputList<string>());
            set => _computeIds = value;
        }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// A list of key value pair of properties that will be used.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier for the zone resource instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`. Default is `DEFAULT`.
        /// </summary>
        [Input("placementPolicy")]
        public Input<string>? PlacementPolicy { get; set; }

        /// <summary>
        /// The id of the region for which this zone is created.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.ZoneTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.ZoneTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ZoneTagArgs>());
            set => _tags = value;
        }

        [Input("tagsToMatches")]
        private InputList<Inputs.ZoneTagsToMatchArgs>? _tagsToMatches;

        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        public InputList<Inputs.ZoneTagsToMatchArgs> TagsToMatches
        {
            get => _tagsToMatches ?? (_tagsToMatches = new InputList<Inputs.ZoneTagsToMatchArgs>());
            set => _tagsToMatches = value;
        }

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }

    public sealed class ZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cloud account this zone belongs to.
        /// </summary>
        [Input("cloudAccountId")]
        public Input<string>? CloudAccountId { get; set; }

        [Input("computeIds")]
        private InputList<string>? _computeIds;

        /// <summary>
        /// The ids of the compute resources that will be explicitly assigned to this zone.
        /// </summary>
        public InputList<string> ComputeIds
        {
            get => _computeIds ?? (_computeIds = new InputList<string>());
            set => _computeIds = value;
        }

        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// A list of key value pair of properties that will be used.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The id of the region for which this zone is defined.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        [Input("links")]
        private InputList<Inputs.ZoneLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        public InputList<Inputs.ZoneLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.ZoneLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier for the zone resource instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`. Default is `DEFAULT`.
        /// </summary>
        [Input("placementPolicy")]
        public Input<string>? PlacementPolicy { get; set; }

        /// <summary>
        /// The id of the region for which this zone is created.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        [Input("tags")]
        private InputList<Inputs.ZoneTagGetArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.ZoneTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ZoneTagGetArgs>());
            set => _tags = value;
        }

        [Input("tagsToMatches")]
        private InputList<Inputs.ZoneTagsToMatchGetArgs>? _tagsToMatches;

        /// <summary>
        /// A set of tag keys and optional values for compute resource filtering:
        /// </summary>
        public InputList<Inputs.ZoneTagsToMatchGetArgs> TagsToMatches
        {
            get => _tagsToMatches ?? (_tagsToMatches = new InputList<Inputs.ZoneTagsToMatchGetArgs>());
            set => _tagsToMatches = value;
        }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ZoneState()
        {
        }
        public static new ZoneState Empty => new ZoneState();
    }
}
