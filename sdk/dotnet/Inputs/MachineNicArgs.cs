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

    public sealed class MachineNicArgs : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;

        /// <summary>
        /// List of IP addresses allocated or in use by this network interface.
        /// example:[ "10.1.2.190" ]
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// Additional properties that may be used to extend the base type.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// Human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The device index of this network interface.
        /// </summary>
        [Input("deviceIndex")]
        public Input<int>? DeviceIndex { get; set; }

        /// <summary>
        /// Human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the network instance that this network interface plugs into.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// List of security group ids which this network interface will be assigned to.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        public MachineNicArgs()
        {
        }
        public static new MachineNicArgs Empty => new MachineNicArgs();
    }
}
