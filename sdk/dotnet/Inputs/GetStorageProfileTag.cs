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

    public sealed class GetStorageProfileTagArgs : global::Pulumi.InvokeArgs
    {
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetStorageProfileTagArgs()
        {
        }
        public static new GetStorageProfileTagArgs Empty => new GetStorageProfileTagArgs();
    }
}
