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
    public sealed class GetMachineTagResult
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private GetMachineTagResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
