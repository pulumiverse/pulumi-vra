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
    public sealed class GetCatalogItemVersionResult
    {
        /// <summary>
        /// Date-time when catalog item version was created at.
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The id of catalog item. One of `id`, or `name` must be provided.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private GetCatalogItemVersionResult(
            string? createdAt,

            string? description,

            string? id)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
        }
    }
}
