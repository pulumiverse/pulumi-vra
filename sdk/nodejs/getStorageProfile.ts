// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to create a storage profile data source.
 *
 * **Storage profile data source by its id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getStorageProfile({
 *     id: vra_storage_profile["this"].id,
 * });
 * ```
 *
 * **Vra storage profile data source filter by external region id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisStorageProfile = pulumi.output(vra.getStorageProfile({
 *     filter: "externalRegionId eq 'foobar'",
 * }));
 * ```
 *
 * A storage profile data source supports the following arguments:
 */
export function getStorageProfile(args?: GetStorageProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageProfileResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getStorageProfile:getStorageProfile", {
        "description": args.description,
        "diskProperties": args.diskProperties,
        "filter": args.filter,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageProfile.
 */
export interface GetStorageProfileArgs {
    /**
     * A human-friendly description.
     */
    description?: string;
    /**
     * Map of storage properties that are to be applied on disk while provisioning.
     */
    diskProperties?: {[key: string]: any};
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: string;
    /**
     * The id of the image profile instance.
     */
    id?: string;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: inputs.GetStorageProfileTag[];
}

/**
 * A collection of values returned by getStorageProfile.
 */
export interface GetStorageProfileResult {
    /**
     * Id of the cloud account this storage profile belongs to.
     */
    readonly cloudAccountId: string;
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * Indicates if this storage profile is a default profile.
     */
    readonly defaultItem: boolean;
    readonly description?: string;
    readonly diskProperties: {[key: string]: any};
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    readonly externalRegionId: string;
    readonly filter?: string;
    readonly id: string;
    /**
     * HATEOAS of the entity
     */
    readonly links: outputs.GetStorageProfileLink[];
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
     * Indicates whether this storage profile supports encryption or not.
     */
    readonly supportsEncryption: boolean;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    readonly tags: outputs.GetStorageProfileTag[];
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
}

export function getStorageProfileOutput(args?: GetStorageProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStorageProfileResult> {
    return pulumi.output(args).apply(a => getStorageProfile(a, opts))
}

/**
 * A collection of arguments for invoking getStorageProfile.
 */
export interface GetStorageProfileOutputArgs {
    /**
     * A human-friendly description.
     */
    description?: pulumi.Input<string>;
    /**
     * Map of storage properties that are to be applied on disk while provisioning.
     */
    diskProperties?: pulumi.Input<{[key: string]: any}>;
    /**
     * Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.
     */
    filter?: pulumi.Input<string>;
    /**
     * The id of the image profile instance.
     */
    id?: pulumi.Input<string>;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    tags?: pulumi.Input<pulumi.Input<inputs.GetStorageProfileTagArgs>[]>;
}
