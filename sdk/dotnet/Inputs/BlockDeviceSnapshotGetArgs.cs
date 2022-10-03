// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra.Inputs
{

    public sealed class BlockDeviceSnapshotGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date when entity was created. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Describes machine within the scope of your organization and is not propagated to the cloud.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of the block device snapshot.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Indicates whether snapshot on block device is current.
        /// </summary>
        [Input("isCurrent")]
        public Input<bool>? IsCurrent { get; set; }

        [Input("links")]
        private InputList<Inputs.BlockDeviceSnapshotLinkGetArgs>? _links;

        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public InputList<Inputs.BlockDeviceSnapshotLinkGetArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.BlockDeviceSnapshotLinkGetArgs>());
            set => _links = value;
        }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of organization that block device snapshot belongs to.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Email of block device snapshot owner.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// Date when entity was last updated. Date and time format is ISO 8601 and UTC.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public BlockDeviceSnapshotGetArgs()
        {
        }
        public static new BlockDeviceSnapshotGetArgs Empty => new BlockDeviceSnapshotGetArgs();
    }
}
