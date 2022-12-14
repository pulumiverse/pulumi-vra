// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vra.Project
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [VraResourceType("vra:project/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrator users or groups associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        [Output("administratorRoles")]
        public Output<ImmutableArray<Outputs.ProjectAdministratorRole>> AdministratorRoles { get; private set; } = null!;

        /// <summary>
        /// A list of administrator users associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        [Output("administrators")]
        public Output<ImmutableArray<string>> Administrators { get; private set; } = null!;

        /// <summary>
        /// A list of storage, network, and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        [Output("constraints")]
        public Output<Outputs.ProjectConstraints?> Constraints { get; private set; } = null!;

        /// <summary>
        /// The project custom properties which are added to all requests in this project.
        /// </summary>
        [Output("customProperties")]
        public Output<ImmutableDictionary<string, object>?> CustomProperties { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The naming template to be used for resources provisioned in this project.
        /// </summary>
        [Output("machineNamingTemplate")]
        public Output<string?> MachineNamingTemplate { get; private set; } = null!;

        /// <summary>
        /// Member users or groups associated with the project.
        /// </summary>
        [Output("memberRoles")]
        public Output<ImmutableArray<Outputs.ProjectMemberRole>> MemberRoles { get; private set; } = null!;

        /// <summary>
        /// A list of member users associated with the project.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The timeout that should be used for cloud template operations and provisioning tasks. The timeout is measured in seconds.
        /// </summary>
        [Output("operationTimeout")]
        public Output<int?> OperationTimeout { get; private set; } = null!;

        /// <summary>
        /// The placement policy that will be applied when selecting a cloud zone for provisioning. Must be one of `DEFAULT` or `SPREAD`.
        /// </summary>
        [Output("placementPolicy")]
        public Output<string?> PlacementPolicy { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the resources in this projects are shared or not. If not set default will be used.
        /// </summary>
        [Output("sharedResources")]
        public Output<bool?> SharedResources { get; private set; } = null!;

        /// <summary>
        /// Viewer users or groups associated with the project.
        /// </summary>
        [Output("viewerRoles")]
        public Output<ImmutableArray<Outputs.ProjectViewerRole>> ViewerRoles { get; private set; } = null!;

        /// <summary>
        /// A list of viewer users associated with the project.
        /// </summary>
        [Output("viewers")]
        public Output<ImmutableArray<string>> Viewers { get; private set; } = null!;

        /// <summary>
        /// A list of configurations for zone assignment to a project.
        /// </summary>
        [Output("zoneAssignments")]
        public Output<ImmutableArray<Outputs.ProjectZoneAssignment>> ZoneAssignments { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("vra:project/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("vra:project/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("administratorRoles")]
        private InputList<Inputs.ProjectAdministratorRoleArgs>? _administratorRoles;

        /// <summary>
        /// Administrator users or groups associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        public InputList<Inputs.ProjectAdministratorRoleArgs> AdministratorRoles
        {
            get => _administratorRoles ?? (_administratorRoles = new InputList<Inputs.ProjectAdministratorRoleArgs>());
            set => _administratorRoles = value;
        }

        [Input("administrators")]
        private InputList<string>? _administrators;

        /// <summary>
        /// A list of administrator users associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer administrator_roles.")]
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// A list of storage, network, and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        [Input("constraints")]
        public Input<Inputs.ProjectConstraintsArgs>? Constraints { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// The project custom properties which are added to all requests in this project.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The naming template to be used for resources provisioned in this project.
        /// </summary>
        [Input("machineNamingTemplate")]
        public Input<string>? MachineNamingTemplate { get; set; }

        [Input("memberRoles")]
        private InputList<Inputs.ProjectMemberRoleArgs>? _memberRoles;

        /// <summary>
        /// Member users or groups associated with the project.
        /// </summary>
        public InputList<Inputs.ProjectMemberRoleArgs> MemberRoles
        {
            get => _memberRoles ?? (_memberRoles = new InputList<Inputs.ProjectMemberRoleArgs>());
            set => _memberRoles = value;
        }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A list of member users associated with the project.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer member_roles.")]
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The timeout that should be used for cloud template operations and provisioning tasks. The timeout is measured in seconds.
        /// </summary>
        [Input("operationTimeout")]
        public Input<int>? OperationTimeout { get; set; }

        /// <summary>
        /// The placement policy that will be applied when selecting a cloud zone for provisioning. Must be one of `DEFAULT` or `SPREAD`.
        /// </summary>
        [Input("placementPolicy")]
        public Input<string>? PlacementPolicy { get; set; }

        /// <summary>
        /// Specifies whether the resources in this projects are shared or not. If not set default will be used.
        /// </summary>
        [Input("sharedResources")]
        public Input<bool>? SharedResources { get; set; }

        [Input("viewerRoles")]
        private InputList<Inputs.ProjectViewerRoleArgs>? _viewerRoles;

        /// <summary>
        /// Viewer users or groups associated with the project.
        /// </summary>
        public InputList<Inputs.ProjectViewerRoleArgs> ViewerRoles
        {
            get => _viewerRoles ?? (_viewerRoles = new InputList<Inputs.ProjectViewerRoleArgs>());
            set => _viewerRoles = value;
        }

        [Input("viewers")]
        private InputList<string>? _viewers;

        /// <summary>
        /// A list of viewer users associated with the project.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer viewer_roles.")]
        public InputList<string> Viewers
        {
            get => _viewers ?? (_viewers = new InputList<string>());
            set => _viewers = value;
        }

        [Input("zoneAssignments")]
        private InputList<Inputs.ProjectZoneAssignmentArgs>? _zoneAssignments;

        /// <summary>
        /// A list of configurations for zone assignment to a project.
        /// </summary>
        public InputList<Inputs.ProjectZoneAssignmentArgs> ZoneAssignments
        {
            get => _zoneAssignments ?? (_zoneAssignments = new InputList<Inputs.ProjectZoneAssignmentArgs>());
            set => _zoneAssignments = value;
        }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        [Input("administratorRoles")]
        private InputList<Inputs.ProjectAdministratorRoleGetArgs>? _administratorRoles;

        /// <summary>
        /// Administrator users or groups associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        public InputList<Inputs.ProjectAdministratorRoleGetArgs> AdministratorRoles
        {
            get => _administratorRoles ?? (_administratorRoles = new InputList<Inputs.ProjectAdministratorRoleGetArgs>());
            set => _administratorRoles = value;
        }

        [Input("administrators")]
        private InputList<string>? _administrators;

        /// <summary>
        /// A list of administrator users associated with the project. Only administrators can manage project's configuration.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer administrator_roles.")]
        public InputList<string> Administrators
        {
            get => _administrators ?? (_administrators = new InputList<string>());
            set => _administrators = value;
        }

        /// <summary>
        /// A list of storage, network, and extensibility constraints to be applied when provisioning through this project.
        /// </summary>
        [Input("constraints")]
        public Input<Inputs.ProjectConstraintsGetArgs>? Constraints { get; set; }

        [Input("customProperties")]
        private InputMap<object>? _customProperties;

        /// <summary>
        /// The project custom properties which are added to all requests in this project.
        /// </summary>
        public InputMap<object> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<object>());
            set => _customProperties = value;
        }

        /// <summary>
        /// A human-friendly description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The naming template to be used for resources provisioned in this project.
        /// </summary>
        [Input("machineNamingTemplate")]
        public Input<string>? MachineNamingTemplate { get; set; }

        [Input("memberRoles")]
        private InputList<Inputs.ProjectMemberRoleGetArgs>? _memberRoles;

        /// <summary>
        /// Member users or groups associated with the project.
        /// </summary>
        public InputList<Inputs.ProjectMemberRoleGetArgs> MemberRoles
        {
            get => _memberRoles ?? (_memberRoles = new InputList<Inputs.ProjectMemberRoleGetArgs>());
            set => _memberRoles = value;
        }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A list of member users associated with the project.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer member_roles.")]
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// A human-friendly name used as an identifier in APIs that support this option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The timeout that should be used for cloud template operations and provisioning tasks. The timeout is measured in seconds.
        /// </summary>
        [Input("operationTimeout")]
        public Input<int>? OperationTimeout { get; set; }

        /// <summary>
        /// The placement policy that will be applied when selecting a cloud zone for provisioning. Must be one of `DEFAULT` or `SPREAD`.
        /// </summary>
        [Input("placementPolicy")]
        public Input<string>? PlacementPolicy { get; set; }

        /// <summary>
        /// Specifies whether the resources in this projects are shared or not. If not set default will be used.
        /// </summary>
        [Input("sharedResources")]
        public Input<bool>? SharedResources { get; set; }

        [Input("viewerRoles")]
        private InputList<Inputs.ProjectViewerRoleGetArgs>? _viewerRoles;

        /// <summary>
        /// Viewer users or groups associated with the project.
        /// </summary>
        public InputList<Inputs.ProjectViewerRoleGetArgs> ViewerRoles
        {
            get => _viewerRoles ?? (_viewerRoles = new InputList<Inputs.ProjectViewerRoleGetArgs>());
            set => _viewerRoles = value;
        }

        [Input("viewers")]
        private InputList<string>? _viewers;

        /// <summary>
        /// A list of viewer users associated with the project.
        /// </summary>
        [Obsolete(@"To specify the type of principal, please refer viewer_roles.")]
        public InputList<string> Viewers
        {
            get => _viewers ?? (_viewers = new InputList<string>());
            set => _viewers = value;
        }

        [Input("zoneAssignments")]
        private InputList<Inputs.ProjectZoneAssignmentGetArgs>? _zoneAssignments;

        /// <summary>
        /// A list of configurations for zone assignment to a project.
        /// </summary>
        public InputList<Inputs.ProjectZoneAssignmentGetArgs> ZoneAssignments
        {
            get => _zoneAssignments ?? (_zoneAssignments = new InputList<Inputs.ProjectZoneAssignmentGetArgs>());
            set => _zoneAssignments = value;
        }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
