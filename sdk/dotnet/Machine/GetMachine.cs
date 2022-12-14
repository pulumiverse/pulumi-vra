// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Machine
{
    public static class GetMachine
    {
        /// <summary>
        /// ## ---layout: "vra"
        /// 
        /// page_title: "VMware vRealize Automation: vra.machine.Machine"
        /// description: |-
        ///   Provides a data lookup for vra_machine.
        /// ---
        /// 
        /// # Data Source: vra.machine.Machine
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a machine data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Machine.GetMachine.Invoke(new()
        ///     {
        ///         Id = @var.My_machine_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Machine data source filter by name:**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Machine.GetMachine.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Machine_name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMachineResult> InvokeAsync(GetMachineArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMachineResult>("vra:machine/getMachine:getMachine", args ?? new GetMachineArgs(), options.WithDefaults());

        /// <summary>
        /// ## ---layout: "vra"
        /// 
        /// page_title: "VMware vRealize Automation: vra.machine.Machine"
        /// description: |-
        ///   Provides a data lookup for vra_machine.
        /// ---
        /// 
        /// # Data Source: vra.machine.Machine
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### S
        /// 
        /// This is an example of how to read a machine data source.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Machine.GetMachine.Invoke(new()
        ///     {
        ///         Id = @var.My_machine_id,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// **Machine data source filter by name:**
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Vra = Pulumi.Vra;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Vra.Machine.GetMachine.Invoke(new()
        ///     {
        ///         Filter = $"name eq '{@var.Machine_name}'",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMachineResult> Invoke(GetMachineInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMachineResult>("vra:machine/getMachine:getMachine", args ?? new GetMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '&lt;regionId&gt;' and cloudAccountId eq '&lt;cloudAccountId&gt;'.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("tags")]
        private List<Inputs.GetMachineTagArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public List<Inputs.GetMachineTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.GetMachineTagArgs>());
            set => _tags = value;
        }

        public GetMachineArgs()
        {
        }
        public static new GetMachineArgs Empty => new GetMachineArgs();
    }

    public sealed class GetMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '&lt;regionId&gt;' and cloudAccountId eq '&lt;cloudAccountId&gt;'.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The id of the image profile instance.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tags")]
        private InputList<Inputs.GetMachineTagInputArgs>? _tags;

        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public InputList<Inputs.GetMachineTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GetMachineTagInputArgs>());
            set => _tags = value;
        }

        public GetMachineInvokeArgs()
        {
        }
        public static new GetMachineInvokeArgs Empty => new GetMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetMachineResult
    {
        /// <summary>
        /// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Set of ids of the cloud accounts this resource belongs to.
        /// </summary>
        public readonly ImmutableArray<string> CloudAccountIds;
        /// <summary>
        /// Date when the entity was created. The date is in ISO 6801 and UTC.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Additional properties that may be used to extend the base resource.
        /// </summary>
        public readonly ImmutableDictionary<string, object> CustomProperties;
        /// <summary>
        /// Deployment id that is associated with this resource.
        /// </summary>
        public readonly string DeploymentId;
        public readonly string? Description;
        /// <summary>
        /// External entity Id on the provider side.
        /// </summary>
        public readonly string ExternalId;
        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        public readonly string ExternalRegionId;
        /// <summary>
        /// The external zoneId of the resource.
        /// </summary>
        public readonly string ExternalZoneId;
        public readonly string? Filter;
        public readonly string Id;
        /// <summary>
        /// HATEOAS of the entity
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMachineLinkResult> Links;
        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the organization this entity belongs to.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Power state of machine.
        /// </summary>
        public readonly string PowerState;
        /// <summary>
        /// The id of the project this resource belongs to.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// A set of tag keys and optional values that were set on this resource.
        /// example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMachineTagResult> Tags;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetMachineResult(
            string address,

            ImmutableArray<string> cloudAccountIds,

            string createdAt,

            ImmutableDictionary<string, object> customProperties,

            string deploymentId,

            string? description,

            string externalId,

            string externalRegionId,

            string externalZoneId,

            string? filter,

            string id,

            ImmutableArray<Outputs.GetMachineLinkResult> links,

            string name,

            string orgId,

            string owner,

            string powerState,

            string projectId,

            ImmutableArray<Outputs.GetMachineTagResult> tags,

            string updatedAt)
        {
            Address = address;
            CloudAccountIds = cloudAccountIds;
            CreatedAt = createdAt;
            CustomProperties = customProperties;
            DeploymentId = deploymentId;
            Description = description;
            ExternalId = externalId;
            ExternalRegionId = externalRegionId;
            ExternalZoneId = externalZoneId;
            Filter = filter;
            Id = id;
            Links = links;
            Name = name;
            OrgId = orgId;
            Owner = owner;
            PowerState = powerState;
            ProjectId = projectId;
            Tags = tags;
            UpdatedAt = updatedAt;
        }
    }
}
