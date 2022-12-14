// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### S
 * This is an example of how to lookup fabric Azure storage account.
 *
 * **Fabric Azure storage account by Id:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.fabric.getStorageAccountAzure({
 *     id: _var.fabric_storage_account_azure_id,
 * });
 * ```
 *
 * **Fabric Azure storage by filter query:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * // Lookup fabric Azure storage account using its name
 * const thisStorageAccountAzure = pulumi.output(vra.fabric.getStorageAccountAzure({
 *     filter: `name eq '${var_name}'`,
 * }));
 * ```
 *
 * A fabric Azure storage account supports the following arguments:
 */
export function getStorageAccountAzure(args?: GetStorageAccountAzureArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageAccountAzureResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:fabric/getStorageAccountAzure:getStorageAccountAzure", {
        "filter": args.filter,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getStorageAccountAzure.
 */
export interface GetStorageAccountAzureArgs {
    /**
     * Search criteria to narrow down the fabric Azure storage accounts. Only one of 'filter' or 'id' must be specified.
     */
    filter?: string;
    /**
     * The id of the fabric Azure storage account. Only one of 'filter' or 'id' must be specified.
     */
    id?: string;
}

/**
 * A collection of values returned by getStorageAccountAzure.
 */
export interface GetStorageAccountAzureResult {
    /**
     * Set of ids of the cloud accounts this entity belongs to.
     */
    readonly cloudAccountIds: string[];
    /**
     * Date when the entity was created. The date is in ISO 6801 and UTC.
     */
    readonly createdAt: string;
    /**
     * A human-friendly description of the fabric Azure storage account.
     */
    readonly description: string;
    /**
     * External entity Id on the provider side.
     */
    readonly externalId: string;
    /**
     * The id of the region for which this entity is defined.
     */
    readonly externalRegionId: string;
    readonly filter?: string;
    readonly id: string;
    /**
     * HATEOAS of the entity
     */
    readonly links: outputs.fabric.GetStorageAccountAzureLink[];
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
     * Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed. example: Standard_LRS / Premium_LRS
     */
    readonly type: string;
    /**
     * Date when the entity was last updated. The date is ISO 8601 and UTC.
     */
    readonly updatedAt: string;
}

export function getStorageAccountAzureOutput(args?: GetStorageAccountAzureOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStorageAccountAzureResult> {
    return pulumi.output(args).apply(a => getStorageAccountAzure(a, opts))
}

/**
 * A collection of arguments for invoking getStorageAccountAzure.
 */
export interface GetStorageAccountAzureOutputArgs {
    /**
     * Search criteria to narrow down the fabric Azure storage accounts. Only one of 'filter' or 'id' must be specified.
     */
    filter?: pulumi.Input<string>;
    /**
     * The id of the fabric Azure storage account. Only one of 'filter' or 'id' must be specified.
     */
    id?: pulumi.Input<string>;
}
