// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./aws";
export * from "./azure";
export * from "./gcp";
export * from "./getAws";
export * from "./getAzure";
export * from "./getGcp";
export * from "./getNsxt";
export * from "./getNsxv";
export * from "./getVSphere";
export * from "./getVmc";
export * from "./nsxt";
export * from "./nsxv";
export * from "./vmc";
export * from "./vsphere";

// Import resources to register:
import { Aws } from "./aws";
import { Azure } from "./azure";
import { Gcp } from "./gcp";
import { Nsxt } from "./nsxt";
import { Nsxv } from "./nsxv";
import { VSphere } from "./vsphere";
import { Vmc } from "./vmc";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vra:cloudaccount/aws:Aws":
                return new Aws(name, <any>undefined, { urn })
            case "vra:cloudaccount/azure:Azure":
                return new Azure(name, <any>undefined, { urn })
            case "vra:cloudaccount/gcp:Gcp":
                return new Gcp(name, <any>undefined, { urn })
            case "vra:cloudaccount/nsxt:Nsxt":
                return new Nsxt(name, <any>undefined, { urn })
            case "vra:cloudaccount/nsxv:Nsxv":
                return new Nsxv(name, <any>undefined, { urn })
            case "vra:cloudaccount/vSphere:VSphere":
                return new VSphere(name, <any>undefined, { urn })
            case "vra:cloudaccount/vmc:Vmc":
                return new Vmc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vra", "cloudaccount/aws", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/azure", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/gcp", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/nsxt", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/nsxv", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/vSphere", _module)
pulumi.runtime.registerResourceModule("vra", "cloudaccount/vmc", _module)
