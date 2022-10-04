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
    public sealed class ContentSourceConfig
    {
        /// <summary>
        /// Content source branch name.
        /// </summary>
        public readonly string? Branch;
        /// <summary>
        /// Content source type. Supported values are `BLUEPRINT`, `IMAGE`, `ABX_SCRIPTS`, `TERRAFORM_CONFIGURATION`.
        /// </summary>
        public readonly string? ContentType;
        /// <summary>
        /// Content source integration id as seen in vRA integrations.
        /// </summary>
        public readonly string IntegrationId;
        /// <summary>
        /// Path to refer to in the content source repository and branch.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Name of the project.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// Content source repository.
        /// </summary>
        public readonly string? Repository;

        [OutputConstructor]
        private ContentSourceConfig(
            string? branch,

            string? contentType,

            string integrationId,

            string path,

            string projectName,

            string? repository)
        {
            Branch = branch;
            ContentType = contentType;
            IntegrationId = integrationId;
            Path = path;
            ProjectName = projectName;
            Repository = repository;
        }
    }
}
