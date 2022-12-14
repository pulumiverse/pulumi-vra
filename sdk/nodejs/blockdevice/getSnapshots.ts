// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## ---layout: "vra"
 *
 * page_title: "VMware vRealize Automation: vra.blockdevice.getSnapshots"
 * description: |-
 *   Provides a data lookup for vra_block_device_snapshots.
 * ---
 *
 * # Data Source: vra.blockdevice.getSnapshots
 * ## Example Usage
 */
export function getSnapshots(args: GetSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vra:blockdevice/getSnapshots:getSnapshots", {
        "blockDeviceId": args.blockDeviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshots.
 */
export interface GetSnapshotsArgs {
    blockDeviceId: string;
}

/**
 * A collection of values returned by getSnapshots.
 */
export interface GetSnapshotsResult {
    readonly blockDeviceId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly snapshots: outputs.blockdevice.GetSnapshotsSnapshot[];
}

export function getSnapshotsOutput(args: GetSnapshotsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotsResult> {
    return pulumi.output(args).apply(a => getSnapshots(a, opts))
}

/**
 * A collection of arguments for invoking getSnapshots.
 */
export interface GetSnapshotsOutputArgs {
    blockDeviceId: pulumi.Input<string>;
}
