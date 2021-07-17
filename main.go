package main

import (
	"azure-sdk-kubeconfig/interal/myaks"
	"azure-sdk-kubeconfig/internal/config"
	"azure-sdk-kubeconfig/internal/utils"
	"context"
	"fmt"
	"io/ioutil"
	"time"
)

func init() {
	config.ParseEnvironment()
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*1))
	defer cancel()
	r, err := myaks.GetAKSConfig(ctx, config.AKSResourceGroup(), config.AKSClusterName(), config.AKSRoleName())

	if err != nil {
		fmt.Println("error:", err)
	}

	kubeconfig := *(r.AccessProfile.KubeConfig)
	homeDir := utils.HomeDir()

	ioutil.WriteFile(homeDir+"/.kube/config", kubeconfig, 0644)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(*(r.AccessProfile.KubeConfig)))
}
