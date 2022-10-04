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

    public sealed class BlueprintValidationMessageGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public BlueprintValidationMessageGetArgs()
        {
        }
        public static new BlueprintValidationMessageGetArgs Empty => new BlueprintValidationMessageGetArgs();
    }
}
