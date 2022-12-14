// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Image
{
    /// <summary>
    /// ## Example Usage
    /// ### S
    /// This is an example of how to create an image profile resource.
    /// 
    /// **Image profile:**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Image.Profile("this", new()
    ///     {
    ///         Description = "test image profile",
    ///         RegionId = data.Vra_region.This.Id,
    ///         ImageMappings = new[]
    ///         {
    ///             new Vra.Image.Inputs.ProfileImageMappingArgs
    ///             {
    ///                 Name = "centos",
    ///                 ImageId = data.Vra_image.Centos.Id,
    ///                 Constraints = new[]
    ///                 {
    ///                     new Vra.Image.Inputs.ProfileImageMappingConstraintArgs
    ///                     {
    ///                         Mandatory = true,
    ///                         Expression = "!env:Test",
    ///                     },
    ///                     new Vra.Image.Inputs.ProfileImageMappingConstraintArgs
    ///                     {
    ///                         Mandatory = false,
    ///                         Expression = "foo:bar",
    ///                     },
    ///                 },
    ///             },
    ///             new Vra.Image.Inputs.ProfileImageMappingArgs
    ///             {
    ///                 Name = "photon",
    ///                 ImageId = data.Vra_image.Photon.Id,
    ///                 CloudConfig = "runcmd echo 'Hello'",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// An image profile resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:image/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        [Output("imageMappings")]
        public Output<ImmutableArray<Outputs.ProfileImageMapping>> ImageMappings { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs args, CustomResourceOptions? options = null)
            : base("vra:image/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("vra:image/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("imageMappings")]
        private InputList<Inputs.ProfileImageMappingArgs>? _imageMappings;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        public InputList<Inputs.ProfileImageMappingArgs> ImageMappings
        {
            get => _imageMappings ?? (_imageMappings = new InputList<Inputs.ProfileImageMappingArgs>());
            set => _imageMappings = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        [Input("imageMappings")]
        private InputList<Inputs.ProfileImageMappingGetArgs>? _imageMappings;

        /// <summary>
        /// Image mapping defined for the corresponding region.
        /// </summary>
        public InputList<Inputs.ProfileImageMappingGetArgs> ImageMappings
        {
            get => _imageMappings ?? (_imageMappings = new InputList<Inputs.ProfileImageMappingGetArgs>());
            set => _imageMappings = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}
