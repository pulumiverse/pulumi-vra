// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a VMware vRealize Automation machine resource.
 *
 * ## Example Usage
 * ### S
 *
 * The following example shows how to create a machine resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * const _this = new vra.Machine("this", {
 *     description: "terrafrom test machine",
 *     projectId: data.vra_project["this"].id,
 *     image: "ubuntu2",
 *     flavor: "medium",
 *     bootConfig: {
 *         content: `#cloud-config
 *   users:
 *   - default
 *   - name: myuser
 *     sudo: ['ALL=(ALL) NOPASSWD:ALL']
 *     groups: [wheel, sudo, admin]
 *     shell: '/bin/bash'
 *     ssh-authorized-keys: |
 *       ssh-rsa your-ssh-rsa:
 *     - sudo sed -e 's/.*PasswordAuthentication yes.*&#47;PasswordAuthentication no/' -i /etc/ssh/sshd_config
 *     - sudo service sshd restart
 * `,
 *     },
 *     nics: [{
 *         networkId: data.vra_network["this"].id,
 *     }],
 *     constraints: [{
 *         mandatory: true,
 *         expression: "AWS",
 *     }],
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 *     disks: [
 *         {
 *             blockDeviceId: vra_block_device.disk1.id,
 *         },
 *         {
 *             blockDeviceId: vra_block_device.disk2.id,
 *         },
 *     ],
 * });
 * ```
 * A machine resource supports the following resource:
 */
export class Machine extends pulumi.CustomResource {
    /**
     * Get an existing Machine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineState, opts?: pulumi.CustomResourceOptions): Machine {
        return new Machine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:index/machine:Machine';

    /**
     * Returns true if the given object is an instance of Machine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Machine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Machine.__pulumiType;
    }

    /**
     * Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
     */
    public /*out*/ readonly address!: pulumi.Output<string>;
    /**
     * Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
     */
    public readonly bootConfig!: pulumi.Output<outputs.MachineBootConfig | undefined>;
    /**
     * Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
     * Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
     */
    public readonly constraints!: pulumi.Output<outputs.MachineConstraint[] | undefined>;
    /**
     * Date when the entity was created. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Additional properties that may be used to extend the base type.
     */
    public readonly customProperties!: pulumi.Output<{[key: string]: any}>;
    /**
     * Describes machine within the scope of your organization and is not propagated to the cloud.
     */
    public readonly deploymentId!: pulumi.Output<string>;
    /**
     * Human-friendly description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specification for attaching/detaching disks to a machine.
     */
    public readonly disks!: pulumi.Output<outputs.MachineDisk[] | undefined>;
    /**
     * List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
     */
    public /*out*/ readonly disksLists!: pulumi.Output<outputs.MachineDisksList[]>;
    /**
     * External entity ID on the provider side.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * External regionId of the resource.
     */
    public /*out*/ readonly externalRegionId!: pulumi.Output<string>;
    /**
     * External zoneId of the resource.
     */
    public /*out*/ readonly externalZoneId!: pulumi.Output<string>;
    /**
     * Flavor of machine instance.
     */
    public readonly flavor!: pulumi.Output<string>;
    /**
     * Type of image used for this machine.
     */
    public readonly image!: pulumi.Output<string | undefined>;
    /**
     * Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
     */
    public readonly imageDiskConstraints!: pulumi.Output<outputs.MachineImageDiskConstraint[] | undefined>;
    /**
     * Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
     */
    public readonly imageRef!: pulumi.Output<string | undefined>;
    /**
     * HATEOAS of the entity
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.MachineLink[]>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
     */
    public readonly nics!: pulumi.Output<outputs.MachineNic[] | undefined>;
    /**
     * ID of the organization this entity belongs to.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * Email of entity owner.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Power state of machine.
     */
    public /*out*/ readonly powerState!: pulumi.Output<string>;
    /**
     * ID of project that resource belongs to.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
     */
    public readonly tags!: pulumi.Output<outputs.MachineTag[]>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Machine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachineArgs | MachineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["bootConfig"] = state ? state.bootConfig : undefined;
            resourceInputs["constraints"] = state ? state.constraints : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["customProperties"] = state ? state.customProperties : undefined;
            resourceInputs["deploymentId"] = state ? state.deploymentId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disks"] = state ? state.disks : undefined;
            resourceInputs["disksLists"] = state ? state.disksLists : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["externalRegionId"] = state ? state.externalRegionId : undefined;
            resourceInputs["externalZoneId"] = state ? state.externalZoneId : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["image"] = state ? state.image : undefined;
            resourceInputs["imageDiskConstraints"] = state ? state.imageDiskConstraints : undefined;
            resourceInputs["imageRef"] = state ? state.imageRef : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nics"] = state ? state.nics : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["powerState"] = state ? state.powerState : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as MachineArgs | undefined;
            if ((!args || args.flavor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavor'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["bootConfig"] = args ? args.bootConfig : undefined;
            resourceInputs["constraints"] = args ? args.constraints : undefined;
            resourceInputs["customProperties"] = args ? args.customProperties : undefined;
            resourceInputs["deploymentId"] = args ? args.deploymentId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disks"] = args ? args.disks : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["image"] = args ? args.image : undefined;
            resourceInputs["imageDiskConstraints"] = args ? args.imageDiskConstraints : undefined;
            resourceInputs["imageRef"] = args ? args.imageRef : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nics"] = args ? args.nics : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["address"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["disksLists"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["externalRegionId"] = undefined /*out*/;
            resourceInputs["externalZoneId"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["powerState"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Machine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Machine resources.
 */
export interface MachineState {
    /**
     * Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
     */
    address?: pulumi.Input<string>;
    /**
     * Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
     */
    bootConfig?: pulumi.Input<inputs.MachineBootConfig>;
    /**
     * Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
     * Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
     */
    constraints?: pulumi.Input<pulumi.Input<inputs.MachineConstraint>[]>;
    /**
     * Date when the entity was created. Date and time format is ISO 8601 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Additional properties that may be used to extend the base type.
     */
    customProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Describes machine within the scope of your organization and is not propagated to the cloud.
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specification for attaching/detaching disks to a machine.
     */
    disks?: pulumi.Input<pulumi.Input<inputs.MachineDisk>[]>;
    /**
     * List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.
     */
    disksLists?: pulumi.Input<pulumi.Input<inputs.MachineDisksList>[]>;
    /**
     * External entity ID on the provider side.
     */
    externalId?: pulumi.Input<string>;
    /**
     * External regionId of the resource.
     */
    externalRegionId?: pulumi.Input<string>;
    /**
     * External zoneId of the resource.
     */
    externalZoneId?: pulumi.Input<string>;
    /**
     * Flavor of machine instance.
     */
    flavor?: pulumi.Input<string>;
    /**
     * Type of image used for this machine.
     */
    image?: pulumi.Input<string>;
    /**
     * Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
     */
    imageDiskConstraints?: pulumi.Input<pulumi.Input<inputs.MachineImageDiskConstraint>[]>;
    /**
     * Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
     */
    imageRef?: pulumi.Input<string>;
    /**
     * HATEOAS of the entity
     */
    links?: pulumi.Input<pulumi.Input<inputs.MachineLink>[]>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
     */
    nics?: pulumi.Input<pulumi.Input<inputs.MachineNic>[]>;
    /**
     * ID of the organization this entity belongs to.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Email of entity owner.
     */
    owner?: pulumi.Input<string>;
    /**
     * Power state of machine.
     */
    powerState?: pulumi.Input<string>;
    /**
     * ID of project that resource belongs to.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.MachineTag>[]>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Machine resource.
 */
export interface MachineArgs {
    /**
     * Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.
     */
    bootConfig?: pulumi.Input<inputs.MachineBootConfig>;
    /**
     * Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets.  
     * Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]
     */
    constraints?: pulumi.Input<pulumi.Input<inputs.MachineConstraint>[]>;
    /**
     * Additional properties that may be used to extend the base type.
     */
    customProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Describes machine within the scope of your organization and is not propagated to the cloud.
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Specification for attaching/detaching disks to a machine.
     */
    disks?: pulumi.Input<pulumi.Input<inputs.MachineDisk>[]>;
    /**
     * Flavor of machine instance.
     */
    flavor: pulumi.Input<string>;
    /**
     * Type of image used for this machine.
     */
    image?: pulumi.Input<string>;
    /**
     * Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.
     */
    imageDiskConstraints?: pulumi.Input<pulumi.Input<inputs.MachineImageDiskConstraint>[]>;
    /**
     * Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
     */
    imageRef?: pulumi.Input<string>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
     */
    nics?: pulumi.Input<pulumi.Input<inputs.MachineNic>[]>;
    /**
     * ID of project that resource belongs to.
     */
    projectId: pulumi.Input<string>;
    /**
     * Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.MachineTag>[]>;
}
