// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Deployment.Inputs
{

    public sealed class DeploymentResourceExpenseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Additional expense incurred for the resource.
        /// </summary>
        [Input("additionalExpense")]
        public Input<double>? AdditionalExpense { get; set; }

        /// <summary>
        /// Expense sync message code if any.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// Compute expense of the entity.
        /// </summary>
        [Input("computeExpense")]
        public Input<double>? ComputeExpense { get; set; }

        /// <summary>
        /// Last expense sync time.
        /// </summary>
        [Input("lastUpdateTime")]
        public Input<string>? LastUpdateTime { get; set; }

        /// <summary>
        /// Expense sync message if any.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Network expense of the entity.
        /// </summary>
        [Input("networkExpense")]
        public Input<double>? NetworkExpense { get; set; }

        /// <summary>
        /// Storage expense of the entity.
        /// </summary>
        [Input("storageExpense")]
        public Input<double>? StorageExpense { get; set; }

        /// <summary>
        /// Total expense of the entity.
        /// </summary>
        [Input("totalExpense")]
        public Input<double>? TotalExpense { get; set; }

        /// <summary>
        /// Monetary unit.
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public DeploymentResourceExpenseGetArgs()
        {
        }
        public static new DeploymentResourceExpenseGetArgs Empty => new DeploymentResourceExpenseGetArgs();
    }
}
