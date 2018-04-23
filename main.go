package main

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/storage/mgmt/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/golang/glog"
)

func main() {
	//fmt.Printf("Hello, world.\n")
	fmt.Printf(os.Getenv("APIVERSION_ARM_NETWORK"))
	fmt.Printf(os.Getenv("GOPATH"))
	os.Setenv("AZURE_ENVIRONMENT_FILEPATH", "C:\\gopath\\src\\github.com\\honcao\\cloudprovider\\config\\azurestackcloud.json")
	os.Setenv("APIVERSION_ARM_COMPUTE", "2016-03-30")
	os.Setenv("APIVERSION_ARM_COMPUTE_CONTAINERSERVICE", "2016-03-30")
	os.Setenv("APIVERSION_ARM_CONTAINERREGISTRY", "2017-10-01")
	os.Setenv("APIVERSION_ARM_DISK", "2016-04-30-preview")
	os.Setenv("APIVERSION_ARM_NETWORK", "2015-06-15")
	os.Setenv("APIVERSION_ARM_NETWORK_SCALESET", "2015-06-15")
	os.Setenv("APIVERSION_ARM_STORAGE", "2016-01-01")
	os.Setenv("APIVERSION_STORAGE", "2016-05-31")
	resourceManagerEndpoint := "https://management.local.azurestack.external"
	subscriptionID := "110054c2-21bc-442c-b214-c31c2578a371"
	serviceManagementEndpoint := "https://management.azurestackci11.onmicrosoft.com/8d887891-6596-46c4-bdb6-1fbde1edbc7e"
	tenantID := "d9b782d5-d098-4374-8f2c-3907cc34611c"
	activeDirectoryEndpoint := "https://login.windows.net/"
	aADClientID := "a7a77abf-ad26-4bb8-9abd-329d03d14804"
	aADClientSecret := "7iiAzT+66U3zazrlnZjNwAVqf7tFscThEVOx1TrGunc="
	servicePrincipalToken, _ := GetServicePrincipalToken(activeDirectoryEndpoint, tenantID, serviceManagementEndpoint, aADClientID, aADClientSecret)

	storageAccountClient := storage.NewAccountsClientWithBaseURI(resourceManagerEndpoint, subscriptionID)
	storageAccountClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	configureUserAgent(&storageAccountClient.Client)

	SAName := "pvc2068417454001"
	resourceGroup := "kl214"
	listKeysResult, err := storageAccountClient.ListKeys(resourceGroup, SAName)

	if err != nil {
		fmt.Println(err)
	}
	if listKeysResult.Keys == nil {
		fmt.Printf("azureDisk - empty listKeysResult in storage account:%s keys", SAName)
		return
	}
	for _, v := range *listKeysResult.Keys {
		fmt.Printf(" Key Name: %s  key vaule: %s", *v.KeyName, *v.Value)
	}
}

// GetServicePrincipalToken creates a new service principal token based on the configuration
func GetServicePrincipalToken(activeDirectoryEndpoint string, tenantID string, serviceManagementEndpoint string, aADClientID string, aADClientSecret string) (*adal.ServicePrincipalToken, error) {
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, fmt.Errorf("creating the OAuth config: %v", err)
	}

	glog.V(2).Infoln("azure: using client_id+client_secret to retrieve access token")
	return adal.NewServicePrincipalToken(
		*oauthConfig,
		aADClientID,
		aADClientSecret,
		serviceManagementEndpoint)
}

// configureUserAgent configures the autorest client with a user agent that
// includes "kubernetes" and the full kubernetes git version string
// example:
// Azure-SDK-for-Go/7.0.1-beta arm-network/2016-09-01; kubernetes-cloudprovider/v1.7.0-alpha.2.711+a2fadef8170bb0-dirty;
func configureUserAgent(client *autorest.Client) {
	k8sVersion := "1.9"
	client.UserAgent = fmt.Sprintf("%s; kubernetes-cloudprovider/%s", client.UserAgent, k8sVersion)
}
