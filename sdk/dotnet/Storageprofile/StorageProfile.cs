// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Storageprofile
{
    /// <summary>
    /// ## Example Usage
    /// ### S
    /// This is an example of how to create a storage profile resource.
    /// 
    /// **Vra storage profile:**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // vSphere storage profile using generic vra_storage_profile resource.
    ///     var thisStorageProfile = new Vra.Storageprofile.StorageProfile("thisStorageProfile", new()
    ///     {
    ///         Description = "vSphere Storage Profile with standard independent non-persistent disk.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         DiskProperties = 
    ///         {
    ///             { "independent", "true" },
    ///             { "persistent", "false" },
    ///             { "limitIops", "2000" },
    ///             { "provisioningType", "eagerZeroedThick" },
    ///             { "sharesLevel", "custom" },
    ///             { "shares", "1500" },
    ///         },
    ///         DiskTargetProperties = 
    ///         {
    ///             { "datastoreId", data.Vra_fabric_datastore_vsphere.This.Id },
    ///             { "storagePolicyId", data.Vra_fabric_storage_policy_vsphere.This.Id },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Vra.Storageprofile.Inputs.StorageProfileTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    ///     // AWS storage profile using generic vra_storage_profile resource.
    ///     var thisStorageprofile_storageProfileStorageProfile = new Vra.Storageprofile.StorageProfile("thisStorageprofile/storageProfileStorageProfile", new()
    ///     {
    ///         Description = "AWS Storage Profile with instance store device type.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         DiskProperties = 
    ///         {
    ///             { "deviceType", "instance-store" },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Vra.Storageprofile.Inputs.StorageProfileTagArgs
    ///             {
    ///                 Key = "foo",
    ///                 Value = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Azure storage profile using generic vra_storage_profile resource.
    ///     var thisVraStorageprofile_storageProfileStorageProfile = new Vra.Storageprofile.StorageProfile("thisVraStorageprofile/storageProfileStorageProfile", new()
    ///     {
    ///         Description = "Azure Storage Profile with managed disks.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         SupportsEncryption = false,
    ///         DiskProperties = 
    ///         {
    ///             { "azureDataDiskCaching", "None" },
    ///             { "azureManagedDiskType", "Standard_LRS" },
    ///             { "azureOsDiskCaching", "None" },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new Vra.Storageprofile.Inputs.StorageProfileTagArgs
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
    /// A storage profile resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:storageprofile/storageProfile:StorageProfile")]
    public partial class StorageProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the cloud account this storage profile belongs to.
        /// </summary>
        [Output("cloudAccountId")]
        public Output<string> CloudAccountId { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Output("defaultItem")]
        public Output<bool> DefaultItem { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Map of storage properties that are to be applied on disk while provisioning.
        /// </summary>
        [Output("diskProperties")]
        public Output<ImmutableDictionary<string, object>?> DiskProperties { get; private set; } = null!;

        /// <summary>
        /// Map of storage placements to know where the disk is provisioned.
        /// </summary>
        [Output("diskTargetProperties")]
        public Output<ImmutableDictionary<string, object>?> DiskTargetProperties { get; private set; } = null!;

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.StorageProfileLink>> Links { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
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
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Output("supportsEncryption")]
        public Output<bool> SupportsEncryption { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.StorageProfileTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a StorageProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageProfile(string name, StorageProfileArgs args, CustomResourceOptions? options = null)
            : base("vra:storageprofile/storageProfile:StorageProfile", name, args ?? new StorageProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageProfile(string name, Input<string> id, StorageProfileState? state = null, CustomResourceOptions? options = null)
            : base("vra:storageprofile/storageProfile:StorageProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StorageProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageProfile Get(string name, Input<string> id, StorageProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new StorageProfile(name, id, state, options);
        }
    }

    public sealed class StorageProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Input("defaultItem", required: true)]
        public Input<bool> DefaultItem { get; set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskProperties")]
        private InputMap<object>? _diskProperties;

        /// <summary>
        /// Map of storage properties that are to be applied on disk while provisioning.
        /// </summary>
        public InputMap<object> DiskProperties
        {
            get => _diskProperties ?? (_diskProperties = new InputMap<object>());
            set => _diskProperties = value;
        }

        [Input("diskTargetProperties")]
        private InputMap<object>? _diskTargetProperties;

        /// <summary>
        /// Map of storage placements to know where the disk is provisioned.
        /// </summary>
        public InputMap<object> DiskTargetProperties
        {
            get => _diskTargetProperties ?? (_diskTargetProperties = new InputMap<object>());
            set => _diskTargetProperties = value;
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

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.StorageProfileTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.StorageProfileTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StorageProfileTagArgs>());
            set => _tags = value;
        }

        public StorageProfileArgs()
        {
        }
        public static new StorageProfileArgs Empty => new StorageProfileArgs();
    }

    public sealed class StorageProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the cloud account this storage profile belongs to.
        /// </summary>
        [Input("cloudAccountId")]
        public Input<string>? CloudAccountId { get; set; }

        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Indicates if this storage profile is a default profile.
        /// </summary>
        [Input("defaultItem")]
        public Input<bool>? DefaultItem { get; set; }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("diskProperties")]
        private InputMap<object>? _diskProperties;

        /// <summary>
        /// Map of storage properties that are to be applied on disk while provisioning.
        /// </summary>
        public InputMap<object> DiskProperties
        {
            get => _diskProperties ?? (_diskProperties = new InputMap<object>());
            set => _diskProperties = value;
        }

        [Input("diskTargetProperties")]
        private InputMap<object>? _diskTargetProperties;

        /// <summary>
        /// Map of storage placements to know where the disk is provisioned.
        /// </summary>
        public InputMap<object> DiskTargetProperties
        {
            get => _diskTargetProperties ?? (_diskTargetProperties = new InputMap<object>());
            set => _diskTargetProperties = value;
        }

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        [Input("links")]
        private InputList<Inputs.StorageProfileLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.StorageProfileLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.StorageProfileLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
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
        /// The id of the region for which this profile is defined as in vRealize Automation(vRA).
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.StorageProfileTagGetArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.StorageProfileTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StorageProfileTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public StorageProfileState()
        {
        }
        public static new StorageProfileState Empty => new StorageProfileState();
    }
}
