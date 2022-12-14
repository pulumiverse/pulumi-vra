// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Catalog
{
    /// <summary>
    /// Creates a VMware vRealize Automation catalog source resource of type cloud template, formerly known as a blueprint.
    /// 
    /// ## Example Usage
    /// ### S
    /// 
    /// The following example shows how to create a catalog source resource.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vra = Pulumiverse.Vra;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Vra.Catalog.SourceBlueprint("this", new()
    ///     {
    ///         ProjectId = @var.Vra_project_id,
    ///     });
    /// 
    /// });
    /// ```
    /// ## Attribute Reference
    /// 
    /// * `created_at` - Date when entity was created. Date and time format is ISO 8601 and UTC.
    /// 
    /// * `created_by` - User who created the entity.
    /// 
    /// * `global` - Flag indicating that all items can be requested across all projects.
    /// 
    /// * `id` - ID of catalog source.
    /// 
    /// * `items_found` - Number of items found in the catalog source.
    /// 
    /// * `items_imported` - Number of items imported from the catalog source.
    /// 
    /// * `last_import_completed_at` - Time at which the last import completed.
    /// 
    /// * `last_import_errors` - List of errors seen when the catalog source was last imported.
    /// 
    /// * `last_import_started_at` - Time at which the last import started.
    /// 
    /// * `last_updated_by` - User who last updated the catalog source.
    /// 
    /// * `type_id` - Type of catalog source. Example: `blueprint`, `CFT`, etc.
    /// 
    /// ## Import
    /// 
    /// To import the cloud template catalog source, use the ID as in the following example
    /// 
    /// ```sh
    ///  $ pulumi import vra:catalog/sourceBlueprint:SourceBlueprint this 05956583-6488-4e7d-84c9-92a7b7219a15`
    /// ```
    /// </summary>
    [VraResourceType("vra:catalog/sourceBlueprint:SourceBlueprint")]
    public partial class SourceBlueprint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Custom configuration of the catalog source as a map of key values.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, string>> Config { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("global")]
        public Output<bool> Global { get; private set; } = null!;

        [Output("itemsFound")]
        public Output<string> ItemsFound { get; private set; } = null!;

        [Output("itemsImported")]
        public Output<string> ItemsImported { get; private set; } = null!;

        [Output("lastImportCompletedAt")]
        public Output<string> LastImportCompletedAt { get; private set; } = null!;

        [Output("lastImportErrors")]
        public Output<ImmutableArray<string>> LastImportErrors { get; private set; } = null!;

        [Output("lastImportStartedAt")]
        public Output<string> LastImportStartedAt { get; private set; } = null!;

        [Output("lastUpdatedBy")]
        public Output<string> LastUpdatedBy { get; private set; } = null!;

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the project this entity belongs to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("typeId")]
        public Output<string> TypeId { get; private set; } = null!;


        /// <summary>
        /// Create a SourceBlueprint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceBlueprint(string name, SourceBlueprintArgs args, CustomResourceOptions? options = null)
            : base("vra:catalog/sourceBlueprint:SourceBlueprint", name, args ?? new SourceBlueprintArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceBlueprint(string name, Input<string> id, SourceBlueprintState? state = null, CustomResourceOptions? options = null)
            : base("vra:catalog/sourceBlueprint:SourceBlueprint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SourceBlueprint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceBlueprint Get(string name, Input<string> id, SourceBlueprintState? state = null, CustomResourceOptions? options = null)
        {
            return new SourceBlueprint(name, id, state, options);
        }
    }

    public sealed class SourceBlueprintArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// Custom configuration of the catalog source as a map of key values.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

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
        /// ID of the project this entity belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public SourceBlueprintArgs()
        {
        }
        public static new SourceBlueprintArgs Empty => new SourceBlueprintArgs();
    }

    public sealed class SourceBlueprintState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// Custom configuration of the catalog source as a map of key values.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("global")]
        public Input<bool>? Global { get; set; }

        [Input("itemsFound")]
        public Input<string>? ItemsFound { get; set; }

        [Input("itemsImported")]
        public Input<string>? ItemsImported { get; set; }

        [Input("lastImportCompletedAt")]
        public Input<string>? LastImportCompletedAt { get; set; }

        [Input("lastImportErrors")]
        private InputList<string>? _lastImportErrors;
        public InputList<string> LastImportErrors
        {
            get => _lastImportErrors ?? (_lastImportErrors = new InputList<string>());
            set => _lastImportErrors = value;
        }

        [Input("lastImportStartedAt")]
        public Input<string>? LastImportStartedAt { get; set; }

        [Input("lastUpdatedBy")]
        public Input<string>? LastUpdatedBy { get; set; }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the project this entity belongs to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("typeId")]
        public Input<string>? TypeId { get; set; }

        public SourceBlueprintState()
        {
        }
        public static new SourceBlueprintState Empty => new SourceBlueprintState();
    }
}
