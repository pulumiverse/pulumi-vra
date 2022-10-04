// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace pulumiverse.Vra.Outputs
{

    [OutputType]
    public sealed class CatalogItemEntitlementDefinition
    {
        /// <summary>
        /// Description of the catalog item.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Icon id of associated catalog item.
        /// </summary>
        public readonly string? IconId;
        /// <summary>
        /// Id of the catalog item.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the catalog item.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Number of items in the associated catalog source.
        /// </summary>
        public readonly int? NumberOfItems;
        /// <summary>
        /// Catalog source name.
        /// </summary>
        public readonly string? SourceName;
        /// <summary>
        /// Catalog source type.
        /// </summary>
        public readonly string? SourceType;
        /// <summary>
        /// Content definition type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private CatalogItemEntitlementDefinition(
            string? description,

            string? iconId,

            string? id,

            string? name,

            int? numberOfItems,

            string? sourceName,

            string? sourceType,

            string? type)
        {
            Description = description;
            IconId = iconId;
            Id = id;
            Name = name;
            NumberOfItems = numberOfItems;
            SourceName = sourceName;
            SourceType = sourceType;
            Type = type;
        }
    }
}
