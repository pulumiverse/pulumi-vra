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

    public sealed class MachineConstraintGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// Indicates whether this constraint should be strictly enforced or not.
        /// </summary>
        [Input("mandatory", required: true)]
        public Input<bool> Mandatory { get; set; } = null!;

        public MachineConstraintGetArgs()
        {
        }
        public static new MachineConstraintGetArgs Empty => new MachineConstraintGetArgs();
    }
}
