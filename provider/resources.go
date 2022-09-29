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
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/schmidtw/pulumi-vra/provider/pkg/version"
	"github.com/vmware/terraform-provider-vra/vra"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vra"
	// modules:
	mainMod = "index" // the vra module
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(vra.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "vra",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "VMWare vRA",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "schmidtw",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/schmidtw/pulumi-vra/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing vmware vra cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "vra", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/schmidtw/pulumi-vra",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "vmware",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"vra_block_device":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BlockDevice")},
			"vra_block_device_snapshot":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BlockDeviceSnapshot")},
			"vra_blueprint":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Blueprint")},
			"vra_blueprint_version":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BlueprintVersion")},
			"vra_catalog_item_entitlement":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogItemEntitlement")},
			"vra_catalog_source_blueprint":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogSourceBlueprint")},
			"vra_catalog_source_entitlement": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CatalogSourceEntitlement")},
			"vra_cloud_account_aws":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountAws")},
			"vra_cloud_account_azure":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountAzure")},
			"vra_cloud_account_gcp":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountGcp")},
			"vra_cloud_account_nsxt":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountNsxt")},
			"vra_cloud_account_nsxv":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountNsxv")},
			"vra_cloud_account_vmc":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountVmc")},
			"vra_cloud_account_vsphere":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "CloudAccountVsphere")},
			"vra_content_source":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ContentSource")},
			"vra_deployment":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Deployment")},
			"vra_fabric_compute":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricCompute")},
			"vra_fabric_datastore_vsphere":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricDatastoreVsphere")},
			"vra_fabric_network_vsphere":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FabricNetworkVsphere")},
			"vra_flavor_profile":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FlavorProfile")},
			"vra_image_profile":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ImageProfile")},
			"vra_load_balancer":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LoadBalancer")},
			"vra_machine":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Machine")},
			"vra_network":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Network")},
			"vra_network_profile":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkProfile")},
			"vra_network_ip_range":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkIpRange")},
			"vra_project":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Project")},
			"vra_storage_profile":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageProfile")},
			"vra_storage_profile_aws":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageProfileAws")},
			"vra_storage_profile_azure":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageProfileAzure")},
			"vra_storage_profile_vsphere":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StorageProfileVsphere")},
			"vra_zone":                       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Zone")},
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"vra_block_device": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlockDevice")},
			"vra_block_device_snapshots": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlockDeviceSnapshots"),
				Docs: &tfbridge.DocInfo{
					Source: "vra_block_device_snapshot.html.markdown",
				},
			},
			"vra_blueprint":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlueprint")},
			"vra_blueprint_version":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlueprintVersion")},
			"vra_catalog_item":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogItem")},
			"vra_catalog_item_entitlement":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogItemEntitlement")},
			"vra_catalog_source_blueprint":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogSourceBlueprint")},
			"vra_catalog_source_entitlement":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCatalogSourceEntitlement")},
			"vra_cloud_account_aws":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountAws")},
			"vra_cloud_account_azure":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountAzure")},
			"vra_cloud_account_gcp":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountGcp")},
			"vra_cloud_account_nsxt":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountNsxt")},
			"vra_cloud_account_nsxv":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountNsxv")},
			"vra_cloud_account_vmc":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountVmc")},
			"vra_cloud_account_vsphere":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getCloudAccountVsphere")},
			"vra_data_collector":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDataCollector")},
			"vra_deployment":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDeployment")},
			"vra_fabric_compute":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFabricCompute")},
			"vra_fabric_datastore_vsphere":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFabricDatastoreVsphere")},
			"vra_fabric_network":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFabricNetwork")},
			"vra_fabric_storage_account_azure":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFabricStorageAccountAzure")},
			"vra_fabric_storage_policy_vsphere": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFabricStoragePolicyVsphere")},
			"vra_image":                         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getImage")},
			"vra_image_profile":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getImageProfile")},
			"vra_machine":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getMachine")},
			"vra_network":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetwork")},
			"vra_network_domain":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkDomain")},
			"vra_network_profile":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getNetworkProfile")},
			"vra_project":                       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getProject")},
			"vra_region":                        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegion")},
			"vra_region_enumeration":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumeration")},
			"vra_region_enumeration_aws":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumerationAws")},
			"vra_region_enumeration_azure":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumerationAzure")},
			"vra_region_enumeration_gcp":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumerationGcp")},
			"vra_region_enumeration_vmc":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumerationVmc")},
			"vra_region_enumeration_vsphere":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegionEnumerationVsphere")},
			"vra_security_group":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSecurityGroup")},
			"vra_storage_profile":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageProfile")},
			"vra_storage_profile_aws":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageProfileAws")},
			"vra_storage_profile_azure":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageProfileAzure")},
			"vra_storage_profile_vsphere":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStorageProfileVsphere")},
			"vra_zone":                          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getZone")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@schmidtw/vra",
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
			PackageName: "schmidtw_vra",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/schmidtw/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "schmidtw",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
