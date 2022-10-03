// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra.Outputs
{

    [OutputType]
    public sealed class CatalogSourceEntitlementDefinition
    {
        public readonly string? Description;
        public readonly string? IconId;
        public readonly string? Id;
        public readonly string? Name;
        public readonly int? NumberOfItems;
        public readonly string? SourceName;
        public readonly string? SourceType;
        public readonly string? Type;

        [OutputConstructor]
        private CatalogSourceEntitlementDefinition(
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
