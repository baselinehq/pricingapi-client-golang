package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := openapiclient.NewGithubComBaselinehqGolangSharedTypesInstance() // *GithubComBaselinehqGolangSharedTypesInstance

	instance.SetInstanceType("s-2vcpu-2gb")
	instance.SetUsageType("ONDEMAND")
	instance.SetProvider("DigitalOcean")
	instance.SetOperatingSystem("linux")
	instance.SetService("Droplet")
	instance.SetRegion("nyc1")

	cpuCores := float32(2)
	ramGb := float32(2)
	instance.SetVm(openapiclient.GithubComBaselinehqGolangSharedTypesVM{CpuCores: &cpuCores, RamGb: &ramGb})

	token := os.Getenv("BASELINEHQ_CLOUD_API_KEY")
	if token == "" {
		fmt.Fprintf(os.Stderr, "Missing BASELINEHQ_CLOUD_API_KEY environment variable")
		os.Exit(1)
	}

	configuration := openapiclient.NewConfiguration()
	configuration.DefaultHeader["Authorization"] = fmt.Sprintf("Bearer %s", token)
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingComputePost(context.Background()).Instance(*instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingComputePost`: SchemaComputePricingsRow
	data, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingComputePost`: %v\n", string(data))
}
