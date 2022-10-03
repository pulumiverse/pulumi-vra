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
    /// Creates a VMware vRealize Automation Azure cloud account resource.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// The following example shows how to create an Azure cloud account resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.CloudAccountAzure("this", new()
    ///     {
    ///         ApplicationId = "sample-application-id",
    ///         ApplicationKey = "sample-application=key",
    ///         Description = "test cloud account",
    ///         Regions = new[]
    ///         {
    ///             "centralus",
    ///         },
    ///         SubscriptionId = "sample-subscription-id",
    ///         TenantId = "sample-tenant-id",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import the Azure cloud account, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:index/cloudAccountAzure:CloudAccountAzure new_azure 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:index/cloudAccountAzure:CloudAccountAzure")]
    public partial class CloudAccountAzure : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Azure Client Application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Azure Client Application Secret Key.
        /// </summary>
        [Output("applicationKey")]
        public Output<string> ApplicationKey { get; private set; } = null!;

        /// <summary>
        /// Date when entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.CloudAccountAzureLink>> Links { get; private set; } = null!;

        /// <summary>
        /// Name of Azure cloud account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// Azure Subscription ID.
        /// </summary>
        [Output("subscriptionId")]
        public Output<string> SubscriptionId { get; private set; } = null!;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CloudAccountAzureTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure Tenant ID.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a CloudAccountAzure resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudAccountAzure(string name, CloudAccountAzureArgs args, CustomResourceOptions? options = null)
            : base("vra:index/cloudAccountAzure:CloudAccountAzure", name, args ?? new CloudAccountAzureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudAccountAzure(string name, Input<string> id, CloudAccountAzureState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/cloudAccountAzure:CloudAccountAzure", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CloudAccountAzure resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudAccountAzure Get(string name, Input<string> id, CloudAccountAzureState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudAccountAzure(name, id, state, options);
        }
    }

    public sealed class CloudAccountAzureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Client Application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Azure Client Application Secret Key.
        /// </summary>
        [Input("applicationKey", required: true)]
        public Input<string> ApplicationKey { get; set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of Azure cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions", required: true)]
        private InputList<string>? _regions;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Azure Subscription ID.
        /// </summary>
        [Input("subscriptionId", required: true)]
        public Input<string> SubscriptionId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.CloudAccountAzureTagArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.CloudAccountAzureTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CloudAccountAzureTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Azure Tenant ID.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public CloudAccountAzureArgs()
        {
        }
        public static new CloudAccountAzureArgs Empty => new CloudAccountAzureArgs();
    }

    public sealed class CloudAccountAzureState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Azure Client Application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Azure Client Application Secret Key.
        /// </summary>
        [Input("applicationKey")]
        public Input<string>? ApplicationKey { get; set; }

        /// <summary>
        /// Date when entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("links")]
        private InputList<Inputs.CloudAccountAzureLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        public InputList<Inputs.CloudAccountAzureLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.CloudAccountAzureLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Name of Azure cloud account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Email of entity owner.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// Set of region names enabled for the cloud account.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Azure Subscription ID.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        [Input("tags")]
        private InputList<Inputs.CloudAccountAzureTagGetArgs>? _tags;

        /// <summary>
        /// Set of tag keys and values to apply to the cloud account.  
        /// Example:[ { "key" : "vmware", "value": "provider" } ]
        /// </summary>
        public InputList<Inputs.CloudAccountAzureTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CloudAccountAzureTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Azure Tenant ID.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public CloudAccountAzureState()
        {
        }
        public static new CloudAccountAzureState Empty => new CloudAccountAzureState();
    }
}
