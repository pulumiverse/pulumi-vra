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
    public sealed class GetProjectAdministratorRoleResult
    {
        public readonly string Email;
        public readonly string? Type;

        [OutputConstructor]
        private GetProjectAdministratorRoleResult(
            string email,

            string? type)
        {
            Email = email;
            Type = type;
        }
    }
}
