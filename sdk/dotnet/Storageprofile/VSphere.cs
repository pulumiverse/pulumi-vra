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
    /// This is an example of how to create a storage profile vsphere resource.
    /// 
    /// **Vra storage profile vsphere:**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // vSphere storage profile using generic vra_storage_profile resource.
    ///     var @this = new Vra.Storageprofile.VSphere("this", new()
    ///     {
    ///         Description = "vSphere Storage Profile with FCD disk.",
    ///         RegionId = data.Vra_region.This.Id,
    ///         DefaultItem = false,
    ///         DiskType = "firstClass",
    ///         ProvisioningType = "thin",
    ///         DatastoreId = data.Vra_fabric_datastore_vsphere.This.Id,
    ///         StoragePolicyId = data.Vra_fabric_storage_policy_vsphere.This.Id,
    ///         Tags = new[]
    ///         {
    ///             new Vra.Storageprofile.Inputs.VSphereTagArgs
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
    /// A storage profile vsphere resource supports the following arguments:
    /// </summary>
    [VraResourceType("vra:storageprofile/vSphere:VSphere")]
    public partial class VSphere : global::Pulumi.CustomResource
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
        /// Id of the vSphere Datastore for placing disk and VM.
        /// </summary>
        [Output("datastoreId")]
        public Output<string?> DatastoreId { get; private set; } = null!;

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
        /// Type of mode for the disk.
        /// </summary>
        [Output("diskMode")]
        public Output<string> DiskMode { get; private set; } = null!;

        /// <summary>
        /// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        /// </summary>
        [Output("diskType")]
        public Output<string?> DiskType { get; private set; } = null!;

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Output("externalRegionId")]
        public Output<string> ExternalRegionId { get; private set; } = null!;

        /// <summary>
        /// The upper bound for the I/O operations per second allocated for each virtual disk.
        /// </summary>
        [Output("limitIops")]
        public Output<string?> LimitIops { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.VSphereLink>> Links { get; private set; } = null!;

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
        /// Type of provisioning policy for the disk.
        /// </summary>
        [Output("provisioningType")]
        public Output<string?> ProvisioningType { get; private set; } = null!;

        /// <summary>
        /// The Id of the region that is associated with the storage profile.
        /// </summary>
        [Output("regionId")]
        public Output<string> RegionId { get; private set; } = null!;

        /// <summary>
        /// A specific number of shares assigned to each virtual machine.
        /// </summary>
        [Output("shares")]
        public Output<string?> Shares { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Output("sharesLevel")]
        public Output<string?> SharesLevel { get; private set; } = null!;

        /// <summary>
        /// Id of the vSphere Storage Policy to be applied.
        /// </summary>
        [Output("storagePolicyId")]
        public Output<string?> StoragePolicyId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this storage policy should support encryption or not.
        /// </summary>
        [Output("supportsEncryption")]
        public Output<bool?> SupportsEncryption { get; private set; } = null!;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.VSphereTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a VSphere resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VSphere(string name, VSphereArgs args, CustomResourceOptions? options = null)
            : base("vra:storageprofile/vSphere:VSphere", name, args ?? new VSphereArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VSphere(string name, Input<string> id, VSphereState? state = null, CustomResourceOptions? options = null)
            : base("vra:storageprofile/vSphere:VSphere", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VSphere resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VSphere Get(string name, Input<string> id, VSphereState? state = null, CustomResourceOptions? options = null)
        {
            return new VSphere(name, id, state, options);
        }
    }

    public sealed class VSphereArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the vSphere Datastore for placing disk and VM.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

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

        /// <summary>
        /// Type of mode for the disk.
        /// </summary>
        [Input("diskMode")]
        public Input<string>? DiskMode { get; set; }

        /// <summary>
        /// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The upper bound for the I/O operations per second allocated for each virtual disk.
        /// </summary>
        [Input("limitIops")]
        public Input<string>? LimitIops { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of provisioning policy for the disk.
        /// </summary>
        [Input("provisioningType")]
        public Input<string>? ProvisioningType { get; set; }

        /// <summary>
        /// The Id of the region that is associated with the storage profile.
        /// </summary>
        [Input("regionId", required: true)]
        public Input<string> RegionId { get; set; } = null!;

        /// <summary>
        /// A specific number of shares assigned to each virtual machine.
        /// </summary>
        [Input("shares")]
        public Input<string>? Shares { get; set; }

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("sharesLevel")]
        public Input<string>? SharesLevel { get; set; }

        /// <summary>
        /// Id of the vSphere Storage Policy to be applied.
        /// </summary>
        [Input("storagePolicyId")]
        public Input<string>? StoragePolicyId { get; set; }

        /// <summary>
        /// Indicates whether this storage policy should support encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.VSphereTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.VSphereTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VSphereTagArgs>());
            set => _tags = value;
        }

        public VSphereArgs()
        {
        }
        public static new VSphereArgs Empty => new VSphereArgs();
    }

    public sealed class VSphereState : global::Pulumi.ResourceArgs
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
        /// Id of the vSphere Datastore for placing disk and VM.
        /// </summary>
        [Input("datastoreId")]
        public Input<string>? DatastoreId { get; set; }

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

        /// <summary>
        /// Type of mode for the disk.
        /// </summary>
        [Input("diskMode")]
        public Input<string>? DiskMode { get; set; }

        /// <summary>
        /// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The id of the region as seen in the cloud provider for which this profile is defined.
        /// </summary>
        [Input("externalRegionId")]
        public Input<string>? ExternalRegionId { get; set; }

        /// <summary>
        /// The upper bound for the I/O operations per second allocated for each virtual disk.
        /// </summary>
        [Input("limitIops")]
        public Input<string>? LimitIops { get; set; }

        [Input("links")]
        private InputList<Inputs.VSphereLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.VSphereLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.VSphereLinkGetArgs>());
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
        /// Type of provisioning policy for the disk.
        /// </summary>
        [Input("provisioningType")]
        public Input<string>? ProvisioningType { get; set; }

        /// <summary>
        /// The Id of the region that is associated with the storage profile.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// A specific number of shares assigned to each virtual machine.
        /// </summary>
        [Input("shares")]
        public Input<string>? Shares { get; set; }

        /// <summary>
        /// Indicates whether this storage profile supports encryption or not.
        /// </summary>
        [Input("sharesLevel")]
        public Input<string>? SharesLevel { get; set; }

        /// <summary>
        /// Id of the vSphere Storage Policy to be applied.
        /// </summary>
        [Input("storagePolicyId")]
        public Input<string>? StoragePolicyId { get; set; }

        /// <summary>
        /// Indicates whether this storage policy should support encryption or not.
        /// </summary>
        [Input("supportsEncryption")]
        public Input<bool>? SupportsEncryption { get; set; }

        [Input("tags")]
        private InputList<Inputs.VSphereTagGetArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this Network Profile.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.VSphereTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VSphereTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Date when the entity was last updated. The date is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public VSphereState()
        {
        }
        public static new VSphereState Empty => new VSphereState();
    }
}
