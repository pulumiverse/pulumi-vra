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
    public sealed class GetProjectConstraintsExtensibilityResult
    {
        public readonly string Expression;
        public readonly bool Mandatory;

        [OutputConstructor]
        private GetProjectConstraintsExtensibilityResult(
            string expression,

            bool mandatory)
        {
            Expression = expression;
            Mandatory = mandatory;
        }
    }
}
