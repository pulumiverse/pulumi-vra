// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace pulumiverse.Vra
{
    public static class GetCatalogSourceEntitlement
    {
        /// <summary>
        /// This data source provides information about a catalog source entitlement in vRA.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to get a vRA catalog source entitlement by its id:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogSourceEntitlement.Invoke(new()
        ///     {
        ///         Id = @var.Catalog_source_entitlement_id,
        ///         ProjectId = @var.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This is an example of how to get a vRA catalog source entitlement by its catalog source id:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogSourceEntitlement.Invoke(new()
        ///     {
        ///         CatalogSourceId = @var.Catalog_source_id,
        ///         ProjectId = @var.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCatalogSourceEntitlementResult> InvokeAsync(GetCatalogSourceEntitlementArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCatalogSourceEntitlementResult>("vra:index/getCatalogSourceEntitlement:getCatalogSourceEntitlement", args ?? new GetCatalogSourceEntitlementArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides information about a catalog source entitlement in vRA.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to get a vRA catalog source entitlement by its id:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogSourceEntitlement.Invoke(new()
        ///     {
        ///         Id = @var.Catalog_source_entitlement_id,
        ///         ProjectId = @var.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// This is an example of how to get a vRA catalog source entitlement by its catalog source id:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetCatalogSourceEntitlement.Invoke(new()
        ///     {
        ///         CatalogSourceId = @var.Catalog_source_id,
        ///         ProjectId = @var.Project_id,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCatalogSourceEntitlementResult> Invoke(GetCatalogSourceEntitlementInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCatalogSourceEntitlementResult>("vra:index/getCatalogSourceEntitlement:getCatalogSourceEntitlement", args ?? new GetCatalogSourceEntitlementInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCatalogSourceEntitlementArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the catalog source to find the entitlement. One of `catalog_source_id` or `id` must be provided.
        /// </summary>
        [Input("catalogSourceId")]
        public string? CatalogSourceId { get; set; }

        /// <summary>
        /// The id of entitlement. One of `catalog_source_id` or `id` must be provided.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The id of the project that this entitlement belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetCatalogSourceEntitlementArgs()
        {
        }
        public static new GetCatalogSourceEntitlementArgs Empty => new GetCatalogSourceEntitlementArgs();
    }

    public sealed class GetCatalogSourceEntitlementInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the catalog source to find the entitlement. One of `catalog_source_id` or `id` must be provided.
        /// </summary>
        [Input("catalogSourceId")]
        public Input<string>? CatalogSourceId { get; set; }

        /// <summary>
        /// The id of entitlement. One of `catalog_source_id` or `id` must be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The id of the project that this entitlement belongs to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetCatalogSourceEntitlementInvokeArgs()
        {
        }
        public static new GetCatalogSourceEntitlementInvokeArgs Empty => new GetCatalogSourceEntitlementInvokeArgs();
    }


    [OutputType]
    public sealed class GetCatalogSourceEntitlementResult
    {
        public readonly string? CatalogSourceId;
        /// <summary>
        /// Represents a catalog source that is linked to a project via an entitlement.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCatalogSourceEntitlementDefinitionResult> Definitions;
        /// <summary>
        /// Id of the catalog source.
        /// </summary>
        public readonly string? Id;
        public readonly string ProjectId;

        [OutputConstructor]
        private GetCatalogSourceEntitlementResult(
            string? catalogSourceId,

            ImmutableArray<Outputs.GetCatalogSourceEntitlementDefinitionResult> definitions,

            string? id,

            string projectId)
        {
            CatalogSourceId = catalogSourceId;
            Definitions = definitions;
            Id = id;
            ProjectId = projectId;
        }
    }
}
