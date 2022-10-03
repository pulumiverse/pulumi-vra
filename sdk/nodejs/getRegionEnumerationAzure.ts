// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 * ### S
 *
 * This is an example of how to lookup a region enumeration data source for Azure cloud account.
 *
 * **Region enumeration data source for Azure**
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vra from "@pulumi/vra";
 *
 * const this = vra.getRegionEnumerationAzure({
 *     applicationId: _var.application_id,
 *     applicationKey: _var.application_key,
 *     subscriptionId: _var.subscription_id,
 *     tenantId: _var.tenant_id,
 * });
 * ```
 *
 * The region enumeration data source for Azure cloud account supports the following arguments:
 */
export function getRegionEnumerationAzure(args: GetRegionEnumerationAzureArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionEnumerationAzureResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:index/getRegionEnumerationAzure:getRegionEnumerationAzure", {
        "applicationId": args.applicationId,
        "applicationKey": args.applicationKey,
        "subscriptionId": args.subscriptionId,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegionEnumerationAzure.
 */
export interface GetRegionEnumerationAzureArgs {
    /**
     * Azure Client Application ID
     */
    applicationId: string;
    /**
     * Azure Client Application Secret Key
     */
    applicationKey: string;
    /**
     * Azure Subscription ID
     */
    subscriptionId: string;
    /**
     * Azure Tenant ID
     */
    tenantId: string;
}

/**
 * A collection of values returned by getRegionEnumerationAzure.
 */
export interface GetRegionEnumerationAzureResult {
    readonly applicationId: string;
    readonly applicationKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A set of Region names to enable provisioning on. Example: northamerica-northeast1
     */
    readonly regions: string[];
    readonly subscriptionId: string;
    readonly tenantId: string;
}

export function getRegionEnumerationAzureOutput(args: GetRegionEnumerationAzureOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionEnumerationAzureResult> {
    return pulumi.output(args).apply(a => getRegionEnumerationAzure(a, opts))
}

/**
 * A collection of arguments for invoking getRegionEnumerationAzure.
 */
export interface GetRegionEnumerationAzureOutputArgs {
    /**
     * Azure Client Application ID
     */
    applicationId: pulumi.Input<string>;
    /**
     * Azure Client Application Secret Key
     */
    applicationKey: pulumi.Input<string>;
    /**
     * Azure Subscription ID
     */
    subscriptionId: pulumi.Input<string>;
    /**
     * Azure Tenant ID
     */
    tenantId: pulumi.Input<string>;
}
