// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 *
 * This is an example of how to read a zone data source.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const test-zone = vra.getZone({
 *     name: _var.zone_name,
 * });
 * ```
 *
 * A zone data source supports the following arguments:
 */
export function getZone(args?: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getZone:getZone", {
        "id": args.id,
        "name": args.name,
        "tags": args.tags,
        "tagsToMatches": args.tagsToMatches,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    /**
     * The id of the zone resource instance.
     */
    id?: string;
    /**
     * A human-friendly name used as an identifier for the zone resource instance.
     */
    name?: string;
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    tags?: inputs.GetZoneTag[];
    /**
     * A set of tag keys and optional values for compute resource filtering:
     */
    tagsToMatches?: inputs.GetZoneTagsToMatch[];
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * The ID of the cloud account this zone belongs to.
     */
    readonly cloudAccountId: string;
    /**
     * The ids of the compute resources that has been explicitly assigned to this zone.
     */
    readonly computeIds: string[];
    /**
     * Date when the entity was created. The date is in ISO 8601 and UTC.
     */
    readonly createdAt: string;
    /**
     * A list of key value pair of properties that will be used.
     */
    readonly customProperties: {[key: string]: any};
    /**
     * A human-friendly description.
     */
    readonly description: string;
    /**
     * The id of the region for which this zone is defined.
     */
    readonly externalRegionId: string;
    /**
     * The folder relative path to the datacenter where resources are deployed to (only applicable for vSphere cloud zones).
     */
    readonly folder: string;
    readonly id: string;
    /**
     * HATEOAS of the entity.
     */
    readonly links: outputs.GetZoneLink[];
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
     * The placement policy for the zone. One of `DEFAULT`, `SPREAD` or `BINPACK`.
     */
    readonly placementPolicy: string;
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    readonly tags: outputs.GetZoneTag[];
    /**
     * A set of tag keys and optional values for compute resource filtering:
     */
    readonly tagsToMatches: outputs.GetZoneTagsToMatch[];
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
}

export function getZoneOutput(args?: GetZoneOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZoneResult> {
    return pulumi.output(args).apply(a => getZone(a, opts))
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    /**
     * The id of the zone resource instance.
     */
    id?: pulumi.Input<string>;
    /**
     * A human-friendly name used as an identifier for the zone resource instance.
     */
    name?: pulumi.Input<string>;
    /**
     * A set of tag keys and optional values that were set on this resource:
     */
    tags?: pulumi.Input<pulumi.Input<inputs.GetZoneTagArgs>[]>;
    /**
     * A set of tag keys and optional values for compute resource filtering:
     */
    tagsToMatches?: pulumi.Input<pulumi.Input<inputs.GetZoneTagsToMatchArgs>[]>;
}
