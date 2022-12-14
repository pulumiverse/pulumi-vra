// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Updates a VMware vRealize Automation fabricDatastoreVsphere resource.
 *
 * ## Example Usage
 * ### S
 *
 * You cannot create a fabric datastore vSphere resource, however you can import it using the command specified in the import section below.
 *
 * Once a resource is imported, you can update it as shown below:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisDatastoreVSphere = new vra.fabric.DatastoreVSphere("this", {
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * To import the fabric datastore vSphere resource, use the ID as in the following example
 *
 * ```sh
 *  $ pulumi import vra:fabric/datastoreVSphere:DatastoreVSphere this 8e0c9a4c-3ab8-48e8-b9d5-0751c871e282
 * ```
 */
export class DatastoreVSphere extends pulumi.CustomResource {
    /**
     * Get an existing DatastoreVSphere resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatastoreVSphereState, opts?: pulumi.CustomResourceOptions): DatastoreVSphere {
        return new DatastoreVSphere(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:fabric/datastoreVSphere:DatastoreVSphere';

    /**
     * Returns true if the given object is an instance of DatastoreVSphere.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatastoreVSphere {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatastoreVSphere.__pulumiType;
    }

    /**
     * Set of ids of the cloud accounts this entity belongs to.
     */
    public /*out*/ readonly cloudAccountIds!: pulumi.Output<string[]>;
    /**
     * Date when the entity was created. The date is in ISO 8601 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A human-friendly description.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * External entity Id on the provider side.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * Id of datacenter in which the datastore is present.
     */
    public /*out*/ readonly externalRegionId!: pulumi.Output<string>;
    /**
     * Indicates free size available in datastore.
     */
    public /*out*/ readonly freeSizeGb!: pulumi.Output<string>;
    /**
     * HATEOAS of the entity
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.fabric.DatastoreVSphereLink[]>;
    /**
     * A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    public /*out*/ readonly orgId!: pulumi.Output<string>;
    /**
     * Email of the user that owns the entity.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    public readonly tags!: pulumi.Output<outputs.fabric.DatastoreVSphereTag[]>;
    /**
     * Type of datastore.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a DatastoreVSphere resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatastoreVSphereArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatastoreVSphereArgs | DatastoreVSphereState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatastoreVSphereState | undefined;
            resourceInputs["cloudAccountIds"] = state ? state.cloudAccountIds : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["externalRegionId"] = state ? state.externalRegionId : undefined;
            resourceInputs["freeSizeGb"] = state ? state.freeSizeGb : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as DatastoreVSphereArgs | undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["cloudAccountIds"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["externalRegionId"] = undefined /*out*/;
            resourceInputs["freeSizeGb"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatastoreVSphere.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatastoreVSphere resources.
 */
export interface DatastoreVSphereState {
    /**
     * Set of ids of the cloud accounts this entity belongs to.
     */
    cloudAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date when the entity was created. The date is in ISO 8601 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * External entity Id on the provider side.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Id of datacenter in which the datastore is present.
     */
    externalRegionId?: pulumi.Input<string>;
    /**
     * Indicates free size available in datastore.
     */
    freeSizeGb?: pulumi.Input<string>;
    /**
     * HATEOAS of the entity
     */
    links?: pulumi.Input<pulumi.Input<inputs.fabric.DatastoreVSphereLink>[]>;
    /**
     * A human-friendly name used as an identifier for the vSphere fabric datastore resource instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the organization this entity belongs to.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Email of the user that owns the entity.
     */
    owner?: pulumi.Input<string>;
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    tags?: pulumi.Input<pulumi.Input<inputs.fabric.DatastoreVSphereTag>[]>;
    /**
     * Type of datastore.
     */
    type?: pulumi.Input<string>;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatastoreVSphere resource.
 */
export interface DatastoreVSphereArgs {
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    tags?: pulumi.Input<pulumi.Input<inputs.fabric.DatastoreVSphereTag>[]>;
}
