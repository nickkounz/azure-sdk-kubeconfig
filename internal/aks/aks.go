package myaks

import (
	"azure-sdk-kubeconfig/internal/config"
	"azure-sdk-kubeconfig/internal/iam"
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
)

func GetAKSClient() (containerservice.ManagedClustersClient, error) {
	aksClient := containerservice.NewManagedClustersClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	aksClient.Authorizer = auth
	aksClient.AddToUserAgent(config.UserAgent())
	aksClient.PollingDuration = time.Hour * 1
	return aksClient, nil
}

func GetAKS(ctx context.Context, resourceGroupName, resourceName string) (c containerservice.ManagedCluster, err error) {
	aksClient, err := GetAKSClient()
	if err != nil {
		return c, fmt.Errorf("cannot get AKS client: %v", err)
	}

	c, err = aksClient.Get(ctx, resourceGroupName, resourceName)
	if err != nil {
		return c, fmt.Errorf("cannot get AKS managed cluster %v from resource group %v: %v", resourceName, resourceGroupName, err)
	}

	return c, nil
}

func GetAKSConfig(ctx context.Context, resourceGroupName, resourceName string, roleName string) (c containerservice.ManagedClusterAccessProfile, err error) {
	aksClient, err := GetAKSClient()
	if err != nil {
		return c, fmt.Errorf("cannot get AKS client: %v", err)
	}

	c, err = aksClient.GetAccessProfile(ctx, resourceGroupName, resourceName, roleName)
	if err != nil {
		return c, fmt.Errorf("cannot get AKS managed cluster %v from resource group %v: %v", resourceName, resourceGroupName, err)
	}

	return c, nil
}
