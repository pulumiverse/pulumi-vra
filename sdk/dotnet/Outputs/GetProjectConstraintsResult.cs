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
    public sealed class GetProjectConstraintsResult
    {
        public readonly ImmutableArray<Outputs.GetProjectConstraintsExtensibilityResult> Extensibilities;
        public readonly ImmutableArray<Outputs.GetProjectConstraintsNetworkResult> Networks;
        public readonly ImmutableArray<Outputs.GetProjectConstraintsStorageResult> Storages;

        [OutputConstructor]
        private GetProjectConstraintsResult(
            ImmutableArray<Outputs.GetProjectConstraintsExtensibilityResult> extensibilities,

            ImmutableArray<Outputs.GetProjectConstraintsNetworkResult> networks,

            ImmutableArray<Outputs.GetProjectConstraintsStorageResult> storages)
        {
            Extensibilities = extensibilities;
            Networks = networks;
            Storages = storages;
        }
    }
}
