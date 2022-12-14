// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to create a storage profile azure resource.
 *
 * **Storage profile azure data source by its id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.storageprofile.getAzure({
 *     id: vra_storage_profile_azure["this"].id,
 * });
 * ```
 *
 * **Vra storage profile data source filter by external region id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const thisAzure = pulumi.output(vra.storageprofile.getAzure({
 *     filter: "externalRegionId eq 'foobar'",
 * }));
 * ```
 *
 * A storage profile azure data source supports the following arguments:
 */
export function getAzure(args?: GetAzureArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:storageprofile/getAzure:getAzure", {
        "filter": args.filter,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzure.
 */
export interface GetAzureArgs {
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
    tags?: inputs.storageprofile.GetAzureTag[];
}

/**
 * A collection of values returned by getAzure.
 */
export interface GetAzureResult {
    readonly cloudAccountId: string;
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * Indicates the caching mechanism for additional disk.
     */
    readonly dataDiskCaching: string;
    /**
     * Indicates if this storage profile is a default profile.
     */
    readonly defaultItem: boolean;
    /**
     * A human-friendly description.
     */
    readonly description: string;
    /**
     * Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
     */
    readonly diskType: string;
    /**
     * The id of the region as seen in the cloud provider for which this profile is defined.
     */
    readonly externalRegionId: string;
    readonly filter?: string;
    readonly id: string;
    /**
     * HATEOAS of the entity
     */
    readonly links: outputs.storageprofile.GetAzureLink[];
    /**
     * A human-friendly name used as an identifier in APIs that support this option.
     */
    readonly name: string;
    readonly orgId: string;
    /**
     * Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
     */
    readonly osDiskCaching: string;
    /**
     * Email of the user that owns the entity.
     */
    readonly owner: string;
    /**
     * Indicates whether this storage policy should support encryption or not.
     */
    readonly supportsEncryption: boolean;
    /**
     * A set of tag keys and optional values that were set on this Network Profile.
     * example:[ { "key" : "ownedBy", "value": "Rainpole" } ]
     */
    readonly tags: outputs.storageprofile.GetAzureTag[];
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
}

export function getAzureOutput(args?: GetAzureOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAzureResult> {
    return pulumi.output(args).apply(a => getAzure(a, opts))
}

/**
 * A collection of arguments for invoking getAzure.
 */
export interface GetAzureOutputArgs {
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
    tags?: pulumi.Input<pulumi.Input<inputs.storageprofile.GetAzureTagArgs>[]>;
}
