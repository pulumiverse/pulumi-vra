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
    public sealed class GetBlockDeviceSnapshotLinkResult
    {
        public readonly string? Href;
        public readonly ImmutableArray<string> Hrefs;
        public readonly string Rel;

        [OutputConstructor]
        private GetBlockDeviceSnapshotLinkResult(
            string? href,

            ImmutableArray<string> hrefs,

            string rel)
        {
            Href = href;
            Hrefs = hrefs;
            Rel = rel;
        }
    }
}
