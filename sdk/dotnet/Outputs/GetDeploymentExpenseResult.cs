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
    public sealed class GetDeploymentExpenseResult
    {
        /// <summary>
        /// Additional expense incurred for the resource.
        /// </summary>
        public readonly double AdditionalExpense;
        /// <summary>
        /// Expense sync message code if any.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Compute expense of the entity.
        /// </summary>
        public readonly double ComputeExpense;
        /// <summary>
        /// Last expense sync time.
        /// </summary>
        public readonly string LastUpdateTime;
        /// <summary>
        /// Expense sync message if any.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Network expense of the entity.
        /// </summary>
        public readonly double NetworkExpense;
        /// <summary>
        /// Storage expense of the entity.
        /// </summary>
        public readonly double StorageExpense;
        /// <summary>
        /// Total expense of the entity.
        /// </summary>
        public readonly double TotalExpense;
        /// <summary>
        /// Monetary unit.
        /// </summary>
        public readonly string Unit;

        [OutputConstructor]
        private GetDeploymentExpenseResult(
            double additionalExpense,

            string code,

            double computeExpense,

            string lastUpdateTime,

            string message,

            double networkExpense,

            double storageExpense,

            double totalExpense,

            string unit)
        {
            AdditionalExpense = additionalExpense;
            Code = code;
            ComputeExpense = computeExpense;
            LastUpdateTime = lastUpdateTime;
            Message = message;
            NetworkExpense = networkExpense;
            StorageExpense = storageExpense;
            TotalExpense = totalExpense;
            Unit = unit;
        }
    }
}
