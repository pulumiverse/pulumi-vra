// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Project.Inputs
{

    public sealed class GetProjectConstraintsInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("extensibilities")]
        private InputList<Inputs.GetProjectConstraintsExtensibilityInputArgs>? _extensibilities;
        public InputList<Inputs.GetProjectConstraintsExtensibilityInputArgs> Extensibilities
        {
            get => _extensibilities ?? (_extensibilities = new InputList<Inputs.GetProjectConstraintsExtensibilityInputArgs>());
            set => _extensibilities = value;
        }

        [Input("networks")]
        private InputList<Inputs.GetProjectConstraintsNetworkInputArgs>? _networks;
        public InputList<Inputs.GetProjectConstraintsNetworkInputArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.GetProjectConstraintsNetworkInputArgs>());
            set => _networks = value;
        }

        [Input("storages")]
        private InputList<Inputs.GetProjectConstraintsStorageInputArgs>? _storages;
        public InputList<Inputs.GetProjectConstraintsStorageInputArgs> Storages
        {
            get => _storages ?? (_storages = new InputList<Inputs.GetProjectConstraintsStorageInputArgs>());
            set => _storages = value;
        }

        public GetProjectConstraintsInputArgs()
        {
        }
        public static new GetProjectConstraintsInputArgs Empty => new GetProjectConstraintsInputArgs();
    }
}
