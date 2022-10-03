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
    public sealed class GetSecurityGroupRuleResult
    {
        public readonly string Access;
        public readonly string Direction;
        public readonly int IpRangeCidr;
        /// <summary>
        /// Name of the security group.
        /// </summary>
        public readonly string? Name;
        public readonly string Ports;
        public readonly string Protocol;
        public readonly string? Service;

        [OutputConstructor]
        private GetSecurityGroupRuleResult(
            string access,

            string direction,

            int ipRangeCidr,

            string? name,

            string ports,

            string protocol,

            string? service)
        {
            Access = access;
            Direction = direction;
            IpRangeCidr = ipRangeCidr;
            Name = name;
            Ports = ports;
            Protocol = protocol;
            Service = service;
        }
    }
}
