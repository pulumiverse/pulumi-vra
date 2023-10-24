// Copyright 2016-2018, Pulumi Corporation.
//
// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vra

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumiverse/pulumi-vra/provider/pkg/version"

	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/vmware/terraform-provider-vra/vra"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vra"
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	vraP := vra.Provider()
	cf := vraP.ConfigureFunc
	vraP.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		// override vra configuration cache, cache includes
		// - access_token (access token for API operations) with env var VRA_ACCESS_TOKEN
		// - refresh_token (refresh token for API operations) with env var VRA_REFRESH_TOKEN
		// - url (base url for API operations) with env var VRA_URL
		// - reauthorizeTimeout (timeout for how often to reauthorize the access token) with env var VRA7_REAUTHORIZE_TIMEOUT
		envAccessToken := os.Getenv("VRA_ACCESS_TOKEN")
		if len(envAccessToken) != 0 {
			d.Set("access_token", envAccessToken)
		}

		envRefreshToken := os.Getenv("VRA_REFRESH_TOKEN")
		if len(envRefreshToken) != 0 {
			d.Set("refresh_token", envRefreshToken)
		}

		envURL := os.Getenv("VRA_URL")
		if len(envURL) != 0 {
			d.Set("url", envURL)
		}

		envReauthorizeTimeout := os.Getenv("VRA7_REAUTHORIZE_TIMEOUT")
		if len(envReauthorizeTimeout) != 0 {
			d.Set("reauthorize_timeout", envReauthorizeTimeout)
		}

		return cf(d)
	}
	p := shimv2.NewProvider(vraP)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "vra",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "VMWare vra",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing VMware VRA cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "vra", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumiverse/pulumi-vra",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "vmware",
		Config: map[string]*tfbridge.SchemaInfo{
			"access_token": {
				Secret: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_ACCESS_TOKEN"},
				},
			},
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_URL"},
				},
			},
			"refresh_token": {
				Secret: tfbridge.True(),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_REFRESH_TOKEN"},
				},
			},
			"insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_INSECURE", "VRA7_INSECURE"},
				},
			},
			"reauthorizeTimeout": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_REAUTHORIZE_TIMEOUT", "VRA7_REAUTHORIZE_TIMEOUT"},
				},
			},
			"api_timeout": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VRA_API_TIMEOUT"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"vra_block_device":               {Tok: tfbridge.MakeResource(mainPkg, "blockdevice", "BlockDevice")},
			"vra_block_device_snapshot":      {Tok: tfbridge.MakeResource(mainPkg, "blockdevice", "Snapshot")},
			"vra_blueprint":                  {Tok: tfbridge.MakeResource(mainPkg, "blueprint", "Blueprint")},
			"vra_blueprint_version":          {Tok: tfbridge.MakeResource(mainPkg, "blueprint", "BlueprintVersion")},
			"vra_catalog_item_entitlement":   {Tok: tfbridge.MakeResource(mainPkg, "catalog", "ItemEntitlement")},
			"vra_catalog_source_blueprint":   {Tok: tfbridge.MakeResource(mainPkg, "catalog", "SourceBlueprint")},
			"vra_catalog_source_entitlement": {Tok: tfbridge.MakeResource(mainPkg, "catalog", "SourceEntitlement")},
			"vra_cloud_account_aws":          {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Aws")},
			"vra_cloud_account_azure":        {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Azure")},
			"vra_cloud_account_gcp":          {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Gcp")},
			"vra_cloud_account_nsxt":         {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Nsxt")},
			"vra_cloud_account_nsxv":         {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Nsxv")},
			"vra_cloud_account_vmc":          {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "Vmc")},
			"vra_cloud_account_vsphere":      {Tok: tfbridge.MakeResource(mainPkg, "cloudaccount", "VSphere")},
			"vra_content_source":             {Tok: tfbridge.MakeResource(mainPkg, "contentsource", "ContentSource")},
			"vra_deployment":                 {Tok: tfbridge.MakeResource(mainPkg, "deployment", "Deployment")},
			"vra_fabric_compute":             {Tok: tfbridge.MakeResource(mainPkg, "fabric", "Compute")},
			"vra_fabric_datastore_vsphere":   {Tok: tfbridge.MakeResource(mainPkg, "fabric", "DatastoreVSphere")},
			"vra_fabric_network_vsphere":     {Tok: tfbridge.MakeResource(mainPkg, "fabric", "NetworkVSphere")},
			"vra_flavor_profile":             {Tok: tfbridge.MakeResource(mainPkg, "flavor", "Profile")},
			"vra_image_profile":              {Tok: tfbridge.MakeResource(mainPkg, "image", "Profile")},
			"vra_load_balancer":              {Tok: tfbridge.MakeResource(mainPkg, "loadbalancer", "LoadBalancer")},
			"vra_machine":                    {Tok: tfbridge.MakeResource(mainPkg, "machine", "Machine")},
			"vra_network":                    {Tok: tfbridge.MakeResource(mainPkg, "network", "Network")},
			"vra_network_profile":            {Tok: tfbridge.MakeResource(mainPkg, "network", "Profile")},
			"vra_network_ip_range":           {Tok: tfbridge.MakeResource(mainPkg, "network", "IpRange")},
			"vra_project":                    {Tok: tfbridge.MakeResource(mainPkg, "project", "Project")},
			"vra_storage_profile":            {Tok: tfbridge.MakeResource(mainPkg, "storageprofile", "StorageProfile")},
			"vra_storage_profile_aws":        {Tok: tfbridge.MakeResource(mainPkg, "storageprofile", "Aws")},
			"vra_storage_profile_azure":      {Tok: tfbridge.MakeResource(mainPkg, "storageprofile", "Azure")},
			"vra_storage_profile_vsphere":    {Tok: tfbridge.MakeResource(mainPkg, "storageprofile", "VSphere")},
			"vra_zone":                       {Tok: tfbridge.MakeResource(mainPkg, "zone", "Zone")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"vra_block_device": {Tok: tfbridge.MakeDataSource(mainPkg, "blockdevice", "getBlockDevice")},
			"vra_block_device_snapshots": {
				Tok: tfbridge.MakeDataSource(mainPkg, "blockdevice", "getSnapshots"),
				Docs: &tfbridge.DocInfo{
					Source: "vra_block_device_snapshot.html.markdown",
				},
			},
			"vra_blueprint":                  {Tok: tfbridge.MakeDataSource(mainPkg, "blueprint", "getBlueprint")},
			"vra_blueprint_version":          {Tok: tfbridge.MakeDataSource(mainPkg, "blueprint", "getVersion")},
			"vra_catalog_item":               {Tok: tfbridge.MakeDataSource(mainPkg, "catalog", "getItem")},
			"vra_catalog_item_entitlement":   {Tok: tfbridge.MakeDataSource(mainPkg, "catalog", "getItemEntitlement")},
			"vra_catalog_source_blueprint":   {Tok: tfbridge.MakeDataSource(mainPkg, "catalog", "getSourceBlueprint")},
			"vra_catalog_source_entitlement": {Tok: tfbridge.MakeDataSource(mainPkg, "catalog", "getSourceEntitlement")},
			"vra_cloud_account_aws":          {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getAws")},
			"vra_cloud_account_azure":        {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getAzure")},
			"vra_cloud_account_gcp":          {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getGcp")},
			"vra_cloud_account_nsxt":         {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getNsxt")},
			"vra_cloud_account_nsxv":         {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getNsxv")},
			"vra_cloud_account_vmc":          {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getVmc")},
			"vra_cloud_account_vsphere":      {Tok: tfbridge.MakeDataSource(mainPkg, "cloudaccount", "getVSphere")},
			"vra_data_collector":             {Tok: tfbridge.MakeDataSource(mainPkg, "datacollector", "getDataCollector")},
			"vra_deployment":                 {Tok: tfbridge.MakeDataSource(mainPkg, "deployment", "getDeployment")},
			"vra_fabric_compute":             {Tok: tfbridge.MakeDataSource(mainPkg, "fabric", "getCompute")},
			"vra_fabric_datastore_vsphere":   {Tok: tfbridge.MakeDataSource(mainPkg, "fabric", "getDatastoreVSphere")},
			"vra_fabric_network":             {Tok: tfbridge.MakeDataSource(mainPkg, "fabric", "getNetwork")},
			"vra_fabric_storage_account_azure": {
				Tok: tfbridge.MakeDataSource(mainPkg, "fabric", "getStorageAccountAzure"),
			},
			"vra_fabric_storage_policy_vsphere": {
				Tok: tfbridge.MakeDataSource(mainPkg, "fabric", "getStoragePolicyVSphere"),
			},
			"vra_image":                      {Tok: tfbridge.MakeDataSource(mainPkg, "image", "getImage")},
			"vra_image_profile":              {Tok: tfbridge.MakeDataSource(mainPkg, "image", "getProfile")},
			"vra_machine":                    {Tok: tfbridge.MakeDataSource(mainPkg, "machine", "getMachine")},
			"vra_network":                    {Tok: tfbridge.MakeDataSource(mainPkg, "network", "getNetwork")},
			"vra_network_domain":             {Tok: tfbridge.MakeDataSource(mainPkg, "network", "getDomain")},
			"vra_network_profile":            {Tok: tfbridge.MakeDataSource(mainPkg, "network", "getProfile")},
			"vra_project":                    {Tok: tfbridge.MakeDataSource(mainPkg, "project", "getProject")},
			"vra_region":                     {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getRegion")},
			"vra_region_enumeration":         {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumeration")},
			"vra_region_enumeration_aws":     {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumerationAws")},
			"vra_region_enumeration_azure":   {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumerationAzure")},
			"vra_region_enumeration_gcp":     {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumerationGcp")},
			"vra_region_enumeration_vmc":     {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumerationVmc")},
			"vra_region_enumeration_vsphere": {Tok: tfbridge.MakeDataSource(mainPkg, "region", "getEnumerationVSphere")},
			"vra_security_group":             {Tok: tfbridge.MakeDataSource(mainPkg, "securitygroup", "getSecurityGroup")},
			"vra_storage_profile":            {Tok: tfbridge.MakeDataSource(mainPkg, "storageprofile", "getStorageProfile")},
			"vra_storage_profile_aws":        {Tok: tfbridge.MakeDataSource(mainPkg, "storageprofile", "getAws")},
			"vra_storage_profile_azure":      {Tok: tfbridge.MakeDataSource(mainPkg, "storageprofile", "getAzure")},
			"vra_storage_profile_vsphere":    {Tok: tfbridge.MakeDataSource(mainPkg, "storageprofile", "getVSphere")},
			"vra_zone":                       {Tok: tfbridge.MakeDataSource(mainPkg, "zone", "getZone")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/vra",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_vra",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
