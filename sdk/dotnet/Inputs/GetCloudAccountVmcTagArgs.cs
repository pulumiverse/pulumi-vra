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

    public sealed class GetCloudAccountVmcTagInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag’s key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Tag’s value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public GetCloudAccountVmcTagInputArgs()
        {
        }
        public static new GetCloudAccountVmcTagInputArgs Empty => new GetCloudAccountVmcTagInputArgs();
    }
}
