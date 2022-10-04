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

    public sealed class GetProjectConstraintsArgs : global::Pulumi.InvokeArgs
    {
        [Input("extensibilities")]
        private List<Inputs.GetProjectConstraintsExtensibilityArgs>? _extensibilities;
        public List<Inputs.GetProjectConstraintsExtensibilityArgs> Extensibilities
        {
            get => _extensibilities ?? (_extensibilities = new List<Inputs.GetProjectConstraintsExtensibilityArgs>());
            set => _extensibilities = value;
        }

        [Input("networks")]
        private List<Inputs.GetProjectConstraintsNetworkArgs>? _networks;
        public List<Inputs.GetProjectConstraintsNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new List<Inputs.GetProjectConstraintsNetworkArgs>());
            set => _networks = value;
        }

        [Input("storages")]
        private List<Inputs.GetProjectConstraintsStorageArgs>? _storages;
        public List<Inputs.GetProjectConstraintsStorageArgs> Storages
        {
            get => _storages ?? (_storages = new List<Inputs.GetProjectConstraintsStorageArgs>());
            set => _storages = value;
        }

        public GetProjectConstraintsArgs()
        {
        }
        public static new GetProjectConstraintsArgs Empty => new GetProjectConstraintsArgs();
    }
}
