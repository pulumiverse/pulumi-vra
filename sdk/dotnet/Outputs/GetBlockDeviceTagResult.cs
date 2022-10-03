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
    public sealed class GetBlockDeviceTagResult
    {
        /// <summary>
        /// Tag’s key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Tag’s value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetBlockDeviceTagResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
