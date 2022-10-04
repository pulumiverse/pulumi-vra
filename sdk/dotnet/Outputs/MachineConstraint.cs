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
    public sealed class MachineConstraint
    {
        /// <summary>
        /// Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// Indicates whether this constraint should be strictly enforced or not.
        /// </summary>
        public readonly bool Mandatory;

        [OutputConstructor]
        private MachineConstraint(
            string expression,

            bool mandatory)
        {
            Expression = expression;
            Mandatory = mandatory;
        }
    }
}
