// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a VMware vRealize Automation vSphere cloud account resource.
 *
 * ## Example Usage
 * ### S
 *
 * The following example shows how to create a vSphere cloud account resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * const _this = new vra.cloudaccount.VSphere("this", {
 *     description: "foobar",
 *     username: _var.username,
 *     password: _var.password,
 *     hostname: _var.hostname,
 *     dcid: _var.vra_data_collector_id,
 *     regions: _var.regions,
 *     associatedCloudAccountIds: [_var.vra_cloud_account_nsxt_id],
 *     acceptSelfSignedCert: true,
 *     tags: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * To import the vSphere cloud account, use the ID as in the following example
 *
 * ```sh
 *  $ pulumi import vra:cloudaccount/vSphere:VSphere new_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class VSphere extends pulumi.CustomResource {
    /**
     * Get an existing VSphere resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VSphereState, opts?: pulumi.CustomResourceOptions): VSphere {
        return new VSphere(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:cloudaccount/vSphere:VSphere';

    /**
     * Returns true if the given object is an instance of VSphere.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VSphere {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VSphere.__pulumiType;
    }

    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    public readonly acceptSelfSignedCert!: pulumi.Output<boolean | undefined>;
    /**
     * Cloud accounts associated with the cloud account.
     */
    public readonly associatedCloudAccountIds!: pulumi.Output<string[] | undefined>;
    /**
     * Date when  entity was created. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Identifier of a data collector vm deployed in the on premise infrastructure.
     */
    public readonly dcid!: pulumi.Output<string | undefined>;
    /**
     * Human-friendly description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * HATEOAS of entity.
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.cloudaccount.VSphereLink[]>;
    /**
     * Name of the vSphere cloud account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of organization that entity belongs to.
     */
    public /*out*/ readonly orgId!: pulumi.Output<string>;
    /**
     * Email of entity owner.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * Password used to authenticate to the cloud account.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * A set of region names that are enabled for the cloud account.
     */
    public readonly regions!: pulumi.Output<string[]>;
    /**
     * A set of tag keys and optional values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    public readonly tags!: pulumi.Output<outputs.cloudaccount.VSphereTag[]>;
    /**
     * Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * vSphere username used to authenticate to the cloud account.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a VSphere resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VSphereArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VSphereArgs | VSphereState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VSphereState | undefined;
            resourceInputs["acceptSelfSignedCert"] = state ? state.acceptSelfSignedCert : undefined;
            resourceInputs["associatedCloudAccountIds"] = state ? state.associatedCloudAccountIds : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dcid"] = state ? state.dcid : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as VSphereArgs | undefined;
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.regions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regions'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["acceptSelfSignedCert"] = args ? args.acceptSelfSignedCert : undefined;
            resourceInputs["associatedCloudAccountIds"] = args ? args.associatedCloudAccountIds : undefined;
            resourceInputs["dcid"] = args ? args.dcid : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VSphere.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VSphere resources.
 */
export interface VSphereState {
    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    acceptSelfSignedCert?: pulumi.Input<boolean>;
    /**
     * Cloud accounts associated with the cloud account.
     */
    associatedCloudAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date when  entity was created. Date and time format is ISO 8601 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Identifier of a data collector vm deployed in the on premise infrastructure.
     */
    dcid?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
     */
    hostname?: pulumi.Input<string>;
    /**
     * HATEOAS of entity.
     */
    links?: pulumi.Input<pulumi.Input<inputs.cloudaccount.VSphereLink>[]>;
    /**
     * Name of the vSphere cloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of organization that entity belongs to.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Email of entity owner.
     */
    owner?: pulumi.Input<string>;
    /**
     * Password used to authenticate to the cloud account.
     */
    password?: pulumi.Input<string>;
    /**
     * A set of region names that are enabled for the cloud account.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of tag keys and optional values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cloudaccount.VSphereTag>[]>;
    /**
     * Date when the entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * vSphere username used to authenticate to the cloud account.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VSphere resource.
 */
export interface VSphereArgs {
    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    acceptSelfSignedCert?: pulumi.Input<boolean>;
    /**
     * Cloud accounts associated with the cloud account.
     */
    associatedCloudAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier of a data collector vm deployed in the on premise infrastructure.
     */
    dcid?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.
     */
    hostname: pulumi.Input<string>;
    /**
     * Name of the vSphere cloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * Password used to authenticate to the cloud account.
     */
    password: pulumi.Input<string>;
    /**
     * A set of region names that are enabled for the cloud account.
     */
    regions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of tag keys and optional values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cloudaccount.VSphereTag>[]>;
    /**
     * vSphere username used to authenticate to the cloud account.
     */
    username: pulumi.Input<string>;
}
