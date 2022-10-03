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
    /// Creates a VMware vRealize Automation (vRA) cloud template resource, formerly known as a blueprint.
    /// 
    /// ## Example Usage
    /// 
    /// The following example shows how to create a blueprint resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = schmidtw.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Blueprint("this", new()
    ///     {
    ///         Description = "Created by vRA terraform provider",
    ///         ProjectId = vra_project.This.Id,
    ///         Content = @$"formatVersion: 1
    /// inputs:
    ///   image:
    ///     type: string
    ///     description: ""Image""
    ///   flavor:
    ///     type: string
    ///     description: ""Flavor""
    /// resources:
    ///   Machine:
    ///     type: Cloud.Machine
    ///     properties:
    ///       image: {input.Image}
    ///       flavor: {input.Flavor}
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// To import the cloud template, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:index/blueprint:Blueprint this 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:index/blueprint:Blueprint")]
    public partial class Blueprint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Blueprint YAML content.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// ID of content source.
        /// </summary>
        [Output("contentSourceId")]
        public Output<string> ContentSourceId { get; private set; } = null!;

        /// <summary>
        /// Content source path.
        /// </summary>
        [Output("contentSourcePath")]
        public Output<string> ContentSourcePath { get; private set; } = null!;

        /// <summary>
        /// Date when content source was last synced. The date is in ISO 8601 and UTC.
        /// </summary>
        [Output("contentSourceSyncAt")]
        public Output<string> ContentSourceSyncAt { get; private set; } = null!;

        /// <summary>
        /// Content source last sync messages.
        /// </summary>
        [Output("contentSourceSyncMessages")]
        public Output<ImmutableArray<string>> ContentSourceSyncMessages { get; private set; } = null!;

        /// <summary>
        /// Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
        /// </summary>
        [Output("contentSourceSyncStatus")]
        public Output<string> ContentSourceSyncStatus { get; private set; } = null!;

        /// <summary>
        /// Content source type.
        /// </summary>
        [Output("contentSourceType")]
        public Output<string> ContentSourceType { get; private set; } = null!;

        /// <summary>
        /// Date when entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The user who created entity.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of project that entity belongs to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Name of project that entity belongs to.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
        /// </summary>
        [Output("requestScopeOrg")]
        public Output<bool> RequestScopeOrg { get; private set; } = null!;

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Status of cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Total number of released versions.
        /// </summary>
        [Output("totalReleasedVersions")]
        public Output<int> TotalReleasedVersions { get; private set; } = null!;

        /// <summary>
        /// Total number of versions.
        /// </summary>
        [Output("totalVersions")]
        public Output<int> TotalVersions { get; private set; } = null!;

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The user who last updated the entity.
        /// </summary>
        [Output("updatedBy")]
        public Output<string> UpdatedBy { get; private set; } = null!;

        /// <summary>
        /// Flag to indicate if the current content of the cloud template/blueprint is valid.
        /// </summary>
        [Output("valid")]
        public Output<bool> Valid { get; private set; } = null!;

        /// <summary>
        /// List of validations messages.
        /// * message - Validation message.
        /// </summary>
        [Output("validationMessages")]
        public Output<ImmutableArray<Outputs.BlueprintValidationMessage>> ValidationMessages { get; private set; } = null!;


        /// <summary>
        /// Create a Blueprint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Blueprint(string name, BlueprintArgs args, CustomResourceOptions? options = null)
            : base("vra:index/blueprint:Blueprint", name, args ?? new BlueprintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Blueprint(string name, Input<string> id, BlueprintState? state = null, CustomResourceOptions? options = null)
            : base("vra:index/blueprint:Blueprint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Blueprint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Blueprint Get(string name, Input<string> id, BlueprintState? state = null, CustomResourceOptions? options = null)
        {
            return new Blueprint(name, id, state, options);
        }
    }

    public sealed class BlueprintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Blueprint YAML content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of project that entity belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
        /// </summary>
        [Input("requestScopeOrg")]
        public Input<bool>? RequestScopeOrg { get; set; }

        public BlueprintArgs()
        {
        }
        public static new BlueprintArgs Empty => new BlueprintArgs();
    }

    public sealed class BlueprintState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Blueprint YAML content.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// ID of content source.
        /// </summary>
        [Input("contentSourceId")]
        public Input<string>? ContentSourceId { get; set; }

        /// <summary>
        /// Content source path.
        /// </summary>
        [Input("contentSourcePath")]
        public Input<string>? ContentSourcePath { get; set; }

        /// <summary>
        /// Date when content source was last synced. The date is in ISO 8601 and UTC.
        /// </summary>
        [Input("contentSourceSyncAt")]
        public Input<string>? ContentSourceSyncAt { get; set; }

        [Input("contentSourceSyncMessages")]
        private InputList<string>? _contentSourceSyncMessages;

        /// <summary>
        /// Content source last sync messages.
        /// </summary>
        public InputList<string> ContentSourceSyncMessages
        {
            get => _contentSourceSyncMessages ?? (_contentSourceSyncMessages = new InputList<string>());
            set => _contentSourceSyncMessages = value;
        }

        /// <summary>
        /// Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
        /// </summary>
        [Input("contentSourceSyncStatus")]
        public Input<string>? ContentSourceSyncStatus { get; set; }

        /// <summary>
        /// Content source type.
        /// </summary>
        [Input("contentSourceType")]
        public Input<string>? ContentSourceType { get; set; }

        /// <summary>
        /// Date when entity was created. The date is in ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The user who created entity.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of organization that entity belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of project that entity belongs to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Name of project that entity belongs to.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
        /// </summary>
        [Input("requestScopeOrg")]
        public Input<bool>? RequestScopeOrg { get; set; }

        /// <summary>
        /// HATEOAS of entity.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Status of cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Total number of released versions.
        /// </summary>
        [Input("totalReleasedVersions")]
        public Input<int>? TotalReleasedVersions { get; set; }

        /// <summary>
        /// Total number of versions.
        /// </summary>
        [Input("totalVersions")]
        public Input<int>? TotalVersions { get; set; }

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The user who last updated the entity.
        /// </summary>
        [Input("updatedBy")]
        public Input<string>? UpdatedBy { get; set; }

        /// <summary>
        /// Flag to indicate if the current content of the cloud template/blueprint is valid.
        /// </summary>
        [Input("valid")]
        public Input<bool>? Valid { get; set; }

        [Input("validationMessages")]
        private InputList<Inputs.BlueprintValidationMessageGetArgs>? _validationMessages;

        /// <summary>
        /// List of validations messages.
        /// * message - Validation message.
        /// </summary>
        public InputList<Inputs.BlueprintValidationMessageGetArgs> ValidationMessages
        {
            get => _validationMessages ?? (_validationMessages = new InputList<Inputs.BlueprintValidationMessageGetArgs>());
            set => _validationMessages = value;
        }

        public BlueprintState()
        {
        }
        public static new BlueprintState Empty => new BlueprintState();
    }
}
