// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a VMware vRealize Automation NSX-T cloud account resource.
 *
 * ## Example Usage
 * ### S
 *
 * The following example shows how to create an NSX-T cloud account resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumiverse/vra";
 *
 * const _this = new vra.CloudAccountNsxt("this", {
 *     description: "foobar",
 *     username: _var.username,
 *     password: _var.password,
 *     hostname: _var.hostname,
 *     dcId: _var.vra_data_collector_id,
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
 * To import the NSX-T cloud account, use the ID as in the following example
 *
 * ```sh
 *  $ pulumi import vra:index/cloudAccountNsxt:CloudAccountNsxt new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15`
 * ```
 */
export class CloudAccountNsxt extends pulumi.CustomResource {
    /**
     * Get an existing CloudAccountNsxt resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudAccountNsxtState, opts?: pulumi.CustomResourceOptions): CloudAccountNsxt {
        return new CloudAccountNsxt(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vra:index/cloudAccountNsxt:CloudAccountNsxt';

    /**
     * Returns true if the given object is an instance of CloudAccountNsxt.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudAccountNsxt {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudAccountNsxt.__pulumiType;
    }

    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    public readonly acceptSelfSignedCert!: pulumi.Output<boolean | undefined>;
    /**
     * Cloud accounts associated with the cloud account.
     */
    public /*out*/ readonly associatedCloudAccountIds!: pulumi.Output<string[]>;
    /**
     * Date when entity was created. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Identifier of a data collector VM deployed in the on premise infrastructure.
     */
    public readonly dcId!: pulumi.Output<string | undefined>;
    /**
     * Human-friendly description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Host name for NSX-T cloud account.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * HATEOAS of entity.
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.CloudAccountNsxtLink[]>;
    /**
     * Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
     * Mode cannot be changed after cloud account is created. Default value is false.
     */
    public readonly managerMode!: pulumi.Output<boolean | undefined>;
    /**
     * Name of NSX-T cloud account.
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
     * Password used to authenticate to the cloud Account.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Set of tag keys and values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    public readonly tags!: pulumi.Output<outputs.CloudAccountNsxtTag[]>;
    /**
     * Date when entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Username used to authenticate to the cloud account.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a CloudAccountNsxt resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudAccountNsxtArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudAccountNsxtArgs | CloudAccountNsxtState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudAccountNsxtState | undefined;
            resourceInputs["acceptSelfSignedCert"] = state ? state.acceptSelfSignedCert : undefined;
            resourceInputs["associatedCloudAccountIds"] = state ? state.associatedCloudAccountIds : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dcId"] = state ? state.dcId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["managerMode"] = state ? state.managerMode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as CloudAccountNsxtArgs | undefined;
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["acceptSelfSignedCert"] = args ? args.acceptSelfSignedCert : undefined;
            resourceInputs["dcId"] = args ? args.dcId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["managerMode"] = args ? args.managerMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["associatedCloudAccountIds"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["orgId"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudAccountNsxt.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudAccountNsxt resources.
 */
export interface CloudAccountNsxtState {
    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    acceptSelfSignedCert?: pulumi.Input<boolean>;
    /**
     * Cloud accounts associated with the cloud account.
     */
    associatedCloudAccountIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Date when entity was created. Date and time format is ISO 8601 and UTC.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Identifier of a data collector VM deployed in the on premise infrastructure.
     */
    dcId?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Host name for NSX-T cloud account.
     */
    hostname?: pulumi.Input<string>;
    /**
     * HATEOAS of entity.
     */
    links?: pulumi.Input<pulumi.Input<inputs.CloudAccountNsxtLink>[]>;
    /**
     * Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
     * Mode cannot be changed after cloud account is created. Default value is false.
     */
    managerMode?: pulumi.Input<boolean>;
    /**
     * Name of NSX-T cloud account.
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
     * Password used to authenticate to the cloud Account.
     */
    password?: pulumi.Input<string>;
    /**
     * Set of tag keys and values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.CloudAccountNsxtTag>[]>;
    /**
     * Date when entity was last updated. Date and time format is ISO 8601 and UTC.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Username used to authenticate to the cloud account.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudAccountNsxt resource.
 */
export interface CloudAccountNsxtArgs {
    /**
     * Accept self-signed certificate when connecting to the cloud account.
     */
    acceptSelfSignedCert?: pulumi.Input<boolean>;
    /**
     * Identifier of a data collector VM deployed in the on premise infrastructure.
     */
    dcId?: pulumi.Input<string>;
    /**
     * Human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Host name for NSX-T cloud account.
     */
    hostname: pulumi.Input<string>;
    /**
     * Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account is created in Manager mode.
     * Mode cannot be changed after cloud account is created. Default value is false.
     */
    managerMode?: pulumi.Input<boolean>;
    /**
     * Name of NSX-T cloud account.
     */
    name?: pulumi.Input<string>;
    /**
     * Password used to authenticate to the cloud Account.
     */
    password: pulumi.Input<string>;
    /**
     * Set of tag keys and values to apply to the cloud account.  
     * Example:[ { "key" : "vmware", "value": "provider" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.CloudAccountNsxtTag>[]>;
    /**
     * Username used to authenticate to the cloud account.
     */
    username: pulumi.Input<string>;
}
