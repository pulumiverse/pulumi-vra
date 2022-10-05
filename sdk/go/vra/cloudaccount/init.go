// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudaccount

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vra/sdk/go/vra"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "vra:cloudaccount/aws:Aws":
		r = &Aws{}
	case "vra:cloudaccount/azure:Azure":
		r = &Azure{}
	case "vra:cloudaccount/gcp:Gcp":
		r = &Gcp{}
	case "vra:cloudaccount/nsxt:Nsxt":
		r = &Nsxt{}
	case "vra:cloudaccount/nsxv:Nsxv":
		r = &Nsxv{}
	case "vra:cloudaccount/vSphere:VSphere":
		r = &VSphere{}
	case "vra:cloudaccount/vmc:Vmc":
		r = &Vmc{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := vra.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/aws",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/azure",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/gcp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/nsxt",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/nsxv",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/vSphere",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vra",
		"cloudaccount/vmc",
		&module{version},
	)
}
