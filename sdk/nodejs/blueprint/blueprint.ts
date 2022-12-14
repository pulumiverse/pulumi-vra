// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a VMware vRealize Automation (vRA) cloud template resource, formerly known as a blueprint.
 *
 * ## Example Usage
 *
 * The following example shows how to create a blueprint resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * const _this = new vra.blueprint.Blueprint("this", {
 *     description: "Created by vRA terraform provider",
 *     projectId: vra_project["this"].id,
 *     content: `formatVersion: 1
 * inputs:
 *   image:
 *     type: string
 *     description: "Image"
 *   flavor:
 *     type: string
 *     description: "Flavor"
 * resources:
 *   Machine:
 *     type: Cloud.Machine
 *     properties:
 *       image: ${input.image}
 *       flavor: ${input.flavor}
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * To import the cloud template, use the ID as in the following example
 *
 * ```sh
 *  $ pulumi import vra:blueprint/blueprint:Blueprint this 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class Blueprint extends pulumi.CustomResource {
    /**
     * Get an existing Blueprint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlueprintState, opts?: pulumi.CustomResourceOptions): Blueprint {
        return new Blueprint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:blueprint/blueprint:Blueprint';

    /**
     * Returns true if the given object is an instance of Blueprint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Blueprint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Blueprint.__pulumiType;
    }

    /**
     * Blueprint YAML content.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * ID of content source.
     */
    public /*out*/ readonly contentSourceId!: pulumi.Output<string>;
    /**
     * Content source path.
     */
    public /*out*/ readonly contentSourcePath!: pulumi.Output<string>;
    /**
     * Date when content source was last synced. The date is in ISO 8601 and UTC.
     */
    public /*out*/ readonly contentSourceSyncAt!: pulumi.Output<string>;
    /**
     * Content source last sync messages.
     */
    public /*out*/ readonly contentSourceSyncMessages!: pulumi.Output<string[]>;
    /**
     * Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
     */
    public /*out*/ readonly contentSourceSyncStatus!: pulumi.Output<string>;
    /**
     * Content source type.
     */
    public /*out*/ readonly contentSourceType!: pulumi.Output<string>;
    /**
     * Date when entity was created. The date is in ISO 8601 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The user who created entity.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * Human-friendly description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of organization that entity belongs to.
     */
    public /*out*/ readonly orgId!: pulumi.Output<string>;
    /**
     * ID of project that entity belongs to.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Name of project that entity belongs to.
     */
    public /*out*/ readonly projectName!: pulumi.Output<string>;
    /**
     * Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
     */
    public readonly requestScopeOrg!: pulumi.Output<boolean>;
    /**
     * HATEOAS of entity.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Status of cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Total number of released versions.
     */
    public /*out*/ readonly totalReleasedVersions!: pulumi.Output<number>;
    /**
     * Total number of versions.
     */
    public /*out*/ readonly totalVersions!: pulumi.Output<number>;
    /**
     * Date when entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The user who last updated the entity.
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<string>;
    /**
     * Flag to indicate if the current content of the cloud template/blueprint is valid.
     */
    public /*out*/ readonly valid!: pulumi.Output<boolean>;
    /**
     * List of validations messages.
     * * message - Validation message.
     */
    public /*out*/ readonly validationMessages!: pulumi.Output<outputs.blueprint.BlueprintValidationMessage[]>;

    /**
     * Create a Blueprint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlueprintArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlueprintArgs | BlueprintState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlueprintState | undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentSourceId"] = state ? state.contentSourceId : undefined;
            resourceInputs["contentSourcePath"] = state ? state.contentSourcePath : undefined;
            resourceInputs["contentSourceSyncAt"] = state ? state.contentSourceSyncAt : undefined;
            resourceInputs["contentSourceSyncMessages"] = state ? state.contentSourceSyncMessages : undefined;
            resourceInputs["contentSourceSyncStatus"] = state ? state.contentSourceSyncStatus : undefined;
            resourceInputs["contentSourceType"] = state ? state.contentSourceType : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["createdBy"] = state ? state.createdBy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["requestScopeOrg"] = state ? state.requestScopeOrg : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["totalReleasedVersions"] = state ? state.totalReleasedVersions : undefined;
            resourceInputs["totalVersions"] = state ? state.totalVersions : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["updatedBy"] = state ? state.updatedBy : undefined;
            resourceInputs["valid"] = state ? state.valid : undefined;
            resourceInputs["validationMessages"] = state ? state.validationMessages : undefined;
        } else {
            const args = argsOrState as BlueprintArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["requestScopeOrg"] = args ? args.requestScopeOrg : undefined;
            resourceInputs["contentSourceId"] = undefined /*out*/;
            resourceInputs["contentSourcePath"] = undefined /*out*/;
            resourceInputs["contentSourceSyncAt"] = undefined /*out*/;
            resourceInputs["contentSourceSyncMessages"] = undefined /*out*/;
            resourceInputs["contentSourceSyncStatus"] = undefined /*out*/;
            resourceInputs["contentSourceType"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["totalReleasedVersions"] = undefined /*out*/;
            resourceInputs["totalVersions"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["updatedBy"] = undefined /*out*/;
            resourceInputs["valid"] = undefined /*out*/;
            resourceInputs["validationMessages"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Blueprint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Blueprint resources.
 */
export interface BlueprintState {
    /**
     * Blueprint YAML content.
     */
    content?: pulumi.Input<string>;
    /**
     * ID of content source.
     */
    contentSourceId?: pulumi.Input<string>;
    /**
     * Content source path.
     */
    contentSourcePath?: pulumi.Input<string>;
    /**
     * Date when content source was last synced. The date is in ISO 8601 and UTC.
     */
    contentSourceSyncAt?: pulumi.Input<string>;
    /**
     * Content source last sync messages.
     */
    contentSourceSyncMessages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Content source last sync status. Supported values: `SUCCESSFUL`, `FAILED`.
     */
    contentSourceSyncStatus?: pulumi.Input<string>;
    /**
     * Content source type.
     */
    contentSourceType?: pulumi.Input<string>;
    /**
     * Date when entity was created. The date is in ISO 8601 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The user who created entity.
     */
    createdBy?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of organization that entity belongs to.
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of project that entity belongs to.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Name of project that entity belongs to.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
     */
    requestScopeOrg?: pulumi.Input<boolean>;
    /**
     * HATEOAS of entity.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Status of cloud template. Supported values: `DRAFT`, `VERSIONED`, `RELEASED`.
     */
    status?: pulumi.Input<string>;
    /**
     * Total number of released versions.
     */
    totalReleasedVersions?: pulumi.Input<number>;
    /**
     * Total number of versions.
     */
    totalVersions?: pulumi.Input<number>;
    /**
     * Date when entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The user who last updated the entity.
     */
    updatedBy?: pulumi.Input<string>;
    /**
     * Flag to indicate if the current content of the cloud template/blueprint is valid.
     */
    valid?: pulumi.Input<boolean>;
    /**
     * List of validations messages.
     * * message - Validation message.
     */
    validationMessages?: pulumi.Input<pulumi.Input<inputs.blueprint.BlueprintValidationMessage>[]>;
}

/**
 * The set of arguments for constructing a Blueprint resource.
 */
export interface BlueprintArgs {
    /**
     * Blueprint YAML content.
     */
    content?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Human-friendly name used as an identifier in APIs that support this option.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of project that entity belongs to.
     */
    projectId: pulumi.Input<string>;
    /**
     * Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to.
     */
    requestScopeOrg?: pulumi.Input<boolean>;
}
