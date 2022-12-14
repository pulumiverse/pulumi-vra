// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./aws";
export * from "./azure";
export * from "./getAws";
export * from "./getAzure";
export * from "./getStorageProfile";
export * from "./getVSphere";
export * from "./storageProfile";
export * from "./vsphere";

// Import resources to register:
import { Aws } from "./aws";
import { Azure } from "./azure";
import { StorageProfile } from "./storageProfile";
import { VSphere } from "./vsphere";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vra:storageprofile/aws:Aws":
                return new Aws(name, <any>undefined, { urn })
            case "vra:storageprofile/azure:Azure":
                return new Azure(name, <any>undefined, { urn })
            case "vra:storageprofile/storageProfile:StorageProfile":
                return new StorageProfile(name, <any>undefined, { urn })
            case "vra:storageprofile/vSphere:VSphere":
                return new VSphere(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vra", "storageprofile/aws", _module)
pulumi.runtime.registerResourceModule("vra", "storageprofile/azure", _module)
pulumi.runtime.registerResourceModule("vra", "storageprofile/storageProfile", _module)
pulumi.runtime.registerResourceModule("vra", "storageprofile/vSphere", _module)
