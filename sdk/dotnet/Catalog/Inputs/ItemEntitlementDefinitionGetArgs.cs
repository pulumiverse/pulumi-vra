// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Catalog.Inputs
{

    public sealed class ItemEntitlementDefinitionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the catalog item.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Icon id of associated catalog item.
        /// </summary>
        [Input("iconId")]
        public Input<string>? IconId { get; set; }

        /// <summary>
        /// Id of the catalog item.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the catalog item.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of items in the associated catalog source.
        /// </summary>
        [Input("numberOfItems")]
        public Input<int>? NumberOfItems { get; set; }

        /// <summary>
        /// Catalog source name.
        /// </summary>
        [Input("sourceName")]
        public Input<string>? SourceName { get; set; }

        /// <summary>
        /// Catalog source type.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// Content definition type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ItemEntitlementDefinitionGetArgs()
        {
        }
        public static new ItemEntitlementDefinitionGetArgs Empty => new ItemEntitlementDefinitionGetArgs();
    }
}
