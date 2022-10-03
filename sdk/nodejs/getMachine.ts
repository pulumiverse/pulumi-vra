// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## ---layout: "vra"
 *
 * page_title: "VMware vRealize Automation: vra.Machine"
 * description: |-
 *   Provides a data lookup for vra_machine.
 * ---
 *
 * # Data Source: vra.Machine
 * ## Example Usage
 * ### S
 *
 * This is an example of how to read a machine data source.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getMachine({
 *     id: _var.my_machine_id,
 * });
 * ```
 *
 * **Machine data source filter by name:**
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisMachine = pulumi.output(vra.getMachine({
 *     filter: `name eq '${var_machine_name}'`,
 * }));
 * ```
 */
export function getMachine(args?: GetMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getMachine:getMachine", {
        "description": args.description,
        "filter": args.filter,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getMachine.
 */
export interface GetMachineArgs {
    /**
     * A human-friendly description.
     */
    description?: string;
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: string;
    /**
     * The id of the image profile instance.
     */
    id?: string;
    /**
     * A set of tag keys and optional values that were set on this resource.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: inputs.GetMachineTag[];
}

/**
 * A collection of values returned by getMachine.
 */
export interface GetMachineResult {
    /**
     * Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
     */
    readonly address: string;
    /**
     * Set of ids of the cloud accounts this resource belongs to.
     */
    readonly cloudAccountIds: string[];
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * Additional properties that may be used to extend the base resource.
     */
    readonly customProperties: {[key: string]: any};
    /**
     * Deployment id that is associated with this resource.
     */
    readonly deploymentId: string;
    readonly description?: string;
    /**
     * External entity Id on the provider side.
     */
    readonly externalId: string;
    /**
     * The external regionId of the resource.
     */
    readonly externalRegionId: string;
    /**
     * The external zoneId of the resource.
     */
    readonly externalZoneId: string;
    readonly filter?: string;
    readonly id: string;
    /**
     * HATEOAS of the entity
     */
    readonly links: outputs.GetMachineLink[];
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    readonly name: string;
    /**
     * The id of the organization this entity belongs to.
     */
    readonly orgId: string;
    /**
     * Email of the user that owns the entity.
     */
    readonly owner: string;
    /**
     * Power state of machine.
     */
    readonly powerState: string;
    /**
     * The id of the project this resource belongs to.
     */
    readonly projectId: string;
    /**
     * A set of tag keys and optional values that were set on this resource.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    readonly tags: outputs.GetMachineTag[];
    readonly updatedAt: string;
}

export function getMachineOutput(args?: GetMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMachineResult> {
    return pulumi.output(args).apply(a => getMachine(a, opts))
}

/**
 * A collection of arguments for invoking getMachine.
 */
export interface GetMachineOutputArgs {
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: pulumi.Input<string>;
    /**
     * The id of the image profile instance.
     */
    id?: pulumi.Input<string>;
    /**
     * A set of tag keys and optional values that were set on this resource.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.GetMachineTagArgs>[]>;
}
