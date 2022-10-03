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
    public static class GetCatalogItem
    {
        /// <summary>
        /// This data source provides information about a catalog item in vRA.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to get a vRA catalog item by its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogItem.Invoke(new()
        ///     {
        ///         Name = @var.Catalog_item_name,
        ///         ExpandVersions = true,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This is an example of how to get a vRA catalog item by its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogItem.Invoke(new()
        ///     {
        ///         Id = @var.Catalog_item_id,
        ///         ExpandVersions = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCatalogItemResult> InvokeAsync(GetCatalogItemArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCatalogItemResult>("vra:index/getCatalogItem:getCatalogItem", args ?? new GetCatalogItemArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides information about a catalog item in vRA.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to get a vRA catalog item by its name.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogItem.Invoke(new()
        ///     {
        ///         Name = @var.Catalog_item_name,
        ///         ExpandVersions = true,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This is an example of how to get a vRA catalog item by its id.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogItem.Invoke(new()
        ///     {
        ///         Id = @var.Catalog_item_id,
        ///         ExpandVersions = true,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCatalogItemResult> Invoke(GetCatalogItemInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCatalogItemResult>("vra:index/getCatalogItem:getCatalogItem", args ?? new GetCatalogItemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCatalogItemArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Flag to indicate whether to expand detailed project data for the catalog item.
        /// </summary>
        [Input("expandProjects")]
        public bool? ExpandProjects { get; set; }

        /// <summary>
        /// Flag to indicate whether to expand detailed versions of the catalog item.
        /// </summary>
        [Input("expandVersions")]
        public bool? ExpandVersions { get; set; }

        /// <summary>
        /// The id of catalog item. One of `id`, or `name` must be provided.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Name of the catalog item. One of `id`, or `name` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCatalogItemArgs()
        {
        }
        public static new GetCatalogItemArgs Empty => new GetCatalogItemArgs();
    }

    public sealed class GetCatalogItemInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Flag to indicate whether to expand detailed project data for the catalog item.
        /// </summary>
        [Input("expandProjects")]
        public Input<bool>? ExpandProjects { get; set; }

        /// <summary>
        /// Flag to indicate whether to expand detailed versions of the catalog item.
        /// </summary>
        [Input("expandVersions")]
        public Input<bool>? ExpandVersions { get; set; }

        /// <summary>
        /// The id of catalog item. One of `id`, or `name` must be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the catalog item. One of `id`, or `name` must be provided.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCatalogItemInvokeArgs()
        {
        }
        public static new GetCatalogItemInvokeArgs Empty => new GetCatalogItemInvokeArgs();
    }


    [OutputType]
    public sealed class GetCatalogItemResult
    {
        /// <summary>
        /// Date-time when catalog item version was created at.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The user the entity was created by.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string Description;
        public readonly bool? ExpandProjects;
        public readonly bool? ExpandVersions;
        /// <summary>
        /// Id of the catalog item version.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Date-time when the entity was last updated.
        /// </summary>
        public readonly string LastUpdatedAt;
        /// <summary>
        /// The user the entity was last updated by.
        /// </summary>
        public readonly string LastUpdatedBy;
        /// <summary>
        /// Name of the entity.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of associated project IDs that can be used for requesting this catalog item.
        /// </summary>
        public readonly ImmutableArray<string> ProjectIds;
        /// <summary>
        /// List of associated projects that can be used for requesting this catalog item.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCatalogItemProjectResult> Projects;
        /// <summary>
        /// Json schema describing request parameters, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5
        /// </summary>
        public readonly string Schema;
        /// <summary>
        /// LibraryItem source ID.
        /// </summary>
        public readonly string SourceId;
        /// <summary>
        /// LibraryItem source name.
        /// </summary>
        public readonly string SourceName;
        public readonly ImmutableArray<Outputs.GetCatalogItemTypeResult> Types;
        /// <summary>
        /// Catalog item versions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCatalogItemVersionResult> Versions;

        [OutputConstructor]
        private GetCatalogItemResult(
            string createdAt,

            string createdBy,

            string description,

            bool? expandProjects,

            bool? expandVersions,

            string id,

            string lastUpdatedAt,

            string lastUpdatedBy,

            string name,

            ImmutableArray<string> projectIds,

            ImmutableArray<Outputs.GetCatalogItemProjectResult> projects,

            string schema,

            string sourceId,

            string sourceName,

            ImmutableArray<Outputs.GetCatalogItemTypeResult> types,

            ImmutableArray<Outputs.GetCatalogItemVersionResult> versions)
        {
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Description = description;
            ExpandProjects = expandProjects;
            ExpandVersions = expandVersions;
            Id = id;
            LastUpdatedAt = lastUpdatedAt;
            LastUpdatedBy = lastUpdatedBy;
            Name = name;
            ProjectIds = projectIds;
            Projects = projects;
            Schema = schema;
            SourceId = sourceId;
            SourceName = sourceName;
            Types = types;
            Versions = versions;
        }
    }
}
