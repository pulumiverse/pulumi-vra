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
    public sealed class LoadBalancerRouteHealthCheckConfiguration
    {
        /// <summary>
        /// Number of consecutive successful checks before considering a particular back-end instance as healthy.
        /// </summary>
        public readonly int? HealthyThreshold;
        /// <summary>
        /// Interval (in seconds) at which the health checks will be performed.
        /// </summary>
        public readonly int? IntervalSeconds;
        /// <summary>
        /// Port which the load balancer is listening to.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// The protocol of the incoming load balancer requests.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Timeout (in seconds) to wait for a response from the back-end instance.
        /// </summary>
        public readonly int? TimeoutSeconds;
        public readonly int? UnhealthyThreshold;
        public readonly string? UrlPath;

        [OutputConstructor]
        private LoadBalancerRouteHealthCheckConfiguration(
            int? healthyThreshold,

            int? intervalSeconds,

            string port,

            string protocol,

            int? timeoutSeconds,

            int? unhealthyThreshold,

            string? urlPath)
        {
            HealthyThreshold = healthyThreshold;
            IntervalSeconds = intervalSeconds;
            Port = port;
            Protocol = protocol;
            TimeoutSeconds = timeoutSeconds;
            UnhealthyThreshold = unhealthyThreshold;
            UrlPath = urlPath;
        }
    }
}
