// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace pulumiverse.Vra
{
    public static class GetDataCollector
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup a data collector.
        /// 
        /// **Data collector data source by its name:**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetDataCollector.Invoke(new()
        ///     {
        ///         Name = @var.Data_collector_name,
        ///     });
        /// 
        /// });
        /// ```
        /// The data collector data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataCollectorResult> InvokeAsync(GetDataCollectorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataCollectorResult>("vra:index/getDataCollector:getDataCollector", args ?? new GetDataCollectorArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to lookup a data collector.
        /// 
        /// **Data collector data source by its name:**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.GetDataCollector.Invoke(new()
        ///     {
        ///         Name = @var.Data_collector_name,
        ///     });
        /// 
        /// });
        /// ```
        /// The data collector data source supports the following arguments:
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDataCollectorResult> Invoke(GetDataCollectorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataCollectorResult>("vra:index/getDataCollector:getDataCollector", args ?? new GetDataCollectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataCollectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Data collector name. Example: Datacollector1
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDataCollectorArgs()
        {
        }
        public static new GetDataCollectorArgs Empty => new GetDataCollectorArgs();
    }

    public sealed class GetDataCollectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Data collector name. Example: Datacollector1
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDataCollectorInvokeArgs()
        {
        }
        public static new GetDataCollectorInvokeArgs Empty => new GetDataCollectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataCollectorResult
    {
        /// <summary>
        /// Data collector host name. Example: dc1-lnd.mycompany.com
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv4 Address of the data collector VM. Example: 10.0.0.1
        /// </summary>
        public readonly string IpAddress;
        public readonly string Name;
        /// <summary>
        /// Current status of the data collector. Example: ACTIVE, INACTIVE
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetDataCollectorResult(
            string hostname,

            string id,

            string ipAddress,

            string name,

            string status)
        {
            Hostname = hostname;
            Id = id;
            IpAddress = ipAddress;
            Name = name;
            Status = status;
        }
    }
}
