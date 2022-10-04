// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace pulumiverse.Vra.Inputs
{

    public sealed class MachineDisksListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the existing block device.
        /// </summary>
        [Input("blockDeviceId", required: true)]
        public Input<string> BlockDeviceId { get; set; } = null!;

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

        public MachineDisksListArgs()
        {
        }
        public static new MachineDisksListArgs Empty => new MachineDisksListArgs();
    }
}
