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

    public sealed class ProjectConstraintsExtensibilityArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("mandatory", required: true)]
        public Input<bool> Mandatory { get; set; } = null!;

        public ProjectConstraintsExtensibilityArgs()
        {
        }
        public static new ProjectConstraintsExtensibilityArgs Empty => new ProjectConstraintsExtensibilityArgs();
    }
}
