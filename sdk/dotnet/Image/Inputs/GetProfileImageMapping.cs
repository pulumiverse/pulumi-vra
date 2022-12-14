// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Image.Inputs
{

    public sealed class GetProfileImageMappingArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudConfig")]
        public string? CloudConfig { get; set; }

        [Input("constraints")]
        private List<Inputs.GetProfileImageMappingConstraintArgs>? _constraints;
        public List<Inputs.GetProfileImageMappingConstraintArgs> Constraints
        {
            get => _constraints ?? (_constraints = new List<Inputs.GetProfileImageMappingConstraintArgs>());
            set => _constraints = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description", required: true)]
        public string Description { get; set; } = null!;

        [Input("externalId", required: true)]
        public string ExternalId { get; set; } = null!;

        /// <summary>
        /// The external regionId of the resource.
        /// </summary>
        [Input("externalRegionId", required: true)]
        public string ExternalRegionId { get; set; } = null!;

        [Input("imageId")]
        public string? ImageId { get; set; }

        [Input("imageName")]
        public string? ImageName { get; set; }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("organization", required: true)]
        public string Organization { get; set; } = null!;

        [Input("osFamily", required: true)]
        public string OsFamily { get; set; } = null!;

        /// <summary>
        /// Email of the user that owns the entity.
        /// </summary>
        [Input("owner", required: true)]
        public string Owner { get; set; } = null!;

        [Input("private", required: true)]
        public bool Private { get; set; }

        public GetProfileImageMappingArgs()
        {
        }
        public static new GetProfileImageMappingArgs Empty => new GetProfileImageMappingArgs();
    }
}
