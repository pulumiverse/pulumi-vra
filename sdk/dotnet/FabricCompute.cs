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
    /// <summary>
    /// Updates a VMware vRealize Automation fabric_compute resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// You cannot create a fabric compute resource, however you can import it using the command specified in the import section below.
    /// 
    /// Once a resource is imported, you can update it as shown below:
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.FabricCompute("this", new()
    ///     {
    ///         Tags = new[]
    ///         {
    ///             new Vra.Inputs.FabricComputeTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import the fabric compute resource, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:index/fabricCompute:FabricCompute this 88fdea8b-92ed-4aa9-b6ee-4670412961b0
    /// ```
    /// </summary>
    [VraResourceType("vra:index/fabricCompute:FabricCompute")]
    public partial class FabricCompute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A list of key value pair of custom properties for the fabric compute resource.
        /// </summary>
        [Output("customProperties")]
        public Output<ImmutableDictionary<string, object>> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The id of the external entity on the provider side.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The external region id of the fabric compute.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// The external zone id of the fabric compute.
        /// </summary>
        [Output("externalZoneId")]
        public Output<string> ExternalZoneId { get; private set; } = null!;

        /// <summary>
        /// Lifecycle status of the compute instance.
        /// </summary>
        [Output("lifecycleState")]
        public Output<string> LifecycleState { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.FabricComputeLink>> Links { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier for the fabric compute resource instance.
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
        /// Power state of fabric compute instance.
        /// </summary>
        [Output("powerState")]
        public Output<string> PowerState { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FabricComputeTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the fabric compute instance.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a FabricCompute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FabricCompute(string name, FabricComputeArgs? args = null, CustomResourceOptions? options = null)
            : base("vra:index/fabricCompute:FabricCompute", name, args ?? new FabricComputeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FabricCompute(string name, Input<string> id, FabricComputeState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/fabricCompute:FabricCompute", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FabricCompute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FabricCompute Get(string name, Input<string> id, FabricComputeState? state = null, CustomResourceOptions? options = null)
        {
            return new FabricCompute(name, id, state, options);
        }
    }

    public sealed class FabricComputeArgs : global::Pulumi.ResourceArgs
    {
        [Input("tags")]
        private InputList<Inputs.FabricComputeTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.FabricComputeTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FabricComputeTagArgs>());
            set => _tags = value;
        }

        public FabricComputeArgs()
        {
        }
        public static new FabricComputeArgs Empty => new FabricComputeArgs();
    }

    public sealed class FabricComputeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// A list of key value pair of custom properties for the fabric compute resource.
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
        /// The id of the external entity on the provider side.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The external region id of the fabric compute.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// The external zone id of the fabric compute.
        /// </summary>
        [Input("externalZoneId")]
        public Input<string>? ExternalZoneId { get; set; }

        /// <summary>
        /// Lifecycle status of the compute instance.
        /// </summary>
        [Input("lifecycleState")]
        public Input<string>? LifecycleState { get; set; }

        [Input("links")]
        private InputList<Inputs.FabricComputeLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity.
        /// </summary>
        public InputList<Inputs.FabricComputeLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.FabricComputeLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier for the fabric compute resource instance.
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
        /// Power state of fabric compute instance.
        /// </summary>
        [Input("powerState")]
        public Input<string>? PowerState { get; set; }

        [Input("tags")]
        private InputList<Inputs.FabricComputeTagGetArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource:
        /// </summary>
        public InputList<Inputs.FabricComputeTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FabricComputeTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of the fabric compute instance.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public FabricComputeState()
        {
        }
        public static new FabricComputeState Empty => new FabricComputeState();
    }
}
