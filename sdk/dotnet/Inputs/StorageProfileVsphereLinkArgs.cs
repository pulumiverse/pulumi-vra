// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace schmidtw.Vra.Inputs
{

    public sealed class StorageProfileVsphereLinkArgs : global::Pulumi.ResourceArgs
    {
        [Input("href")]
        public Input<string>? Href { get; set; }

        [Input("hrefs")]
        private InputList<string>? _hrefs;
        public InputList<string> Hrefs
        {
            get => _hrefs ?? (_hrefs = new InputList<string>());
            set => _hrefs = value;
        }

        [Input("rel", required: true)]
        public Input<string> Rel { get; set; } = null!;

        public StorageProfileVsphereLinkArgs()
        {
        }
        public static new StorageProfileVsphereLinkArgs Empty => new StorageProfileVsphereLinkArgs();
    }
}
