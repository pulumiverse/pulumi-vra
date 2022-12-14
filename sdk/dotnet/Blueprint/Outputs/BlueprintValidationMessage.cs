// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Blueprint.Outputs
{

    [OutputType]
    public sealed class BlueprintValidationMessage
    {
        public readonly string? Message;
        public readonly ImmutableDictionary<string, string>? Metadata;
        public readonly string? Path;
        public readonly string? ResourceName;
        public readonly string? Type;

        [OutputConstructor]
        private BlueprintValidationMessage(
            string? message,

            ImmutableDictionary<string, string>? metadata,

            string? path,

            string? resourceName,

            string? type)
        {
            Message = message;
            Metadata = metadata;
            Path = path;
            ResourceName = resourceName;
            Type = type;
        }
    }
}
