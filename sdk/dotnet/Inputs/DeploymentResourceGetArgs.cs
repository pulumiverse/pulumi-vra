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

    public sealed class DeploymentResourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("dependsOns")]
        private InputList<string>? _dependsOns;

        /// <summary>
        /// A list of other resources this resource depends on.
        /// </summary>
        public InputList<string> DependsOns
        {
            get => _dependsOns ?? (_dependsOns = new InputList<string>());
            set => _dependsOns = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expenses")]
        private InputList<Inputs.DeploymentResourceExpenseGetArgs>? _expenses;

        /// <summary>
        /// Expense incurred for the deployment.
        /// </summary>
        public InputList<Inputs.DeploymentResourceExpenseGetArgs> Expenses
        {
            get => _expenses ?? (_expenses = new InputList<Inputs.DeploymentResourceExpenseGetArgs>());
            set => _expenses = value;
        }

        /// <summary>
        /// Unique identifier of the resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// List of properties in the encoded JSON string format.
        /// </summary>
        [Input("propertiesJson")]
        public Input<string>? PropertiesJson { get; set; }

        /// <summary>
        /// The current state of the resource. Supported values are `PARTIAL`, `TAINTED`, `OK.`
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The current sync status. Supported values are `SUCCESS`, `MISSING`, `STALE`.
        /// </summary>
        [Input("syncStatus")]
        public Input<string>? SyncStatus { get; set; }

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DeploymentResourceGetArgs()
        {
        }
        public static new DeploymentResourceGetArgs Empty => new DeploymentResourceGetArgs();
    }
}
