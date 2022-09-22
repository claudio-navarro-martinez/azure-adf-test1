package main

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

const (
	subid = ""
	rg = "CNMResourceGroup"
)
func main() {

	//ExampleFactoriesClient_NewListPager()
	vName := ExampleFactoriesClient_NewListByResourceGroupPager()
	vPipe := ExamplePipelinesClient_NewListByFactoryPager(vName)
	vRunId := ExamplePipelinesClient_CreateRun(vName, vPipe)
	ExamplePipelineRunsClient_Get(vName, vRunId)	
}

func ExamplePipelineRunsClient_Get(adfname string, runid string) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelineRunsClient(subid, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	
	for {
		res, err := client.Get(ctx, rg, adfname, runid , nil)
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}
		log.Println(*res.Status)
		time.Sleep(60 * time.Second)
		if *res.Status == "kk" {
			continue
		}	
	}	
}

func ExamplePipelinesClient_CreateRun(vName string, pipe string) string{
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient(subid, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateRun(ctx, rg, vName, pipe, &armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
		IsRecovery:        nil,
		StartActivityName: nil,
		StartFromFailure:  nil,
		Parameters: map[string]interface{}{
			"OutputBlobNameList": []interface{}{
				"exampleoutput.csv",
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
	log.Println(*res.RunID)
	return *res.RunID
}
func ExamplePipelinesClient_NewListByFactoryPager(vName string) string{
	var vPipe string
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient(subid, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager(rg, vName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
			log.Println(*v.Name)
			vPipe = *v.Name
		}
	}
	return vPipe
}

func ExampleFactoriesClient_NewListByResourceGroupPager() string {
	var vName string
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient(subid, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager(rg, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
			log.Println(*v.Identity.PrincipalID,*v.ID, *v.Name)
			vName = *v.Name
		}
	}
	return vName
}

func ExampleFactoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient(subid, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

func ExampleFactoriesClient_ConfigureFactoryRepo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ConfigureFactoryRepo(ctx, "East US", armdatafactory.FactoryRepoUpdate{
		FactoryResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
			Type:                to.Ptr("FactoryVSTSConfiguration"),
			AccountName:         to.Ptr("ADF"),
			CollaborationBranch: to.Ptr("master"),
			LastCommitID:        to.Ptr(""),
			RepositoryName:      to.Ptr("repo"),
			RootFolder:          to.Ptr("/"),
			ProjectName:         to.Ptr("project"),
			TenantID:            to.Ptr(""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

func ExampleFactoriesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.Factory{
		Location: to.Ptr("East US"),
	}, &armdatafactory.FactoriesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

func ExampleFactoriesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.FactoryUpdateParameters{
		Tags: map[string]*string{
			"exampleTag": to.Ptr("exampleValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

func ExampleFactoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "exampleResourceGroup", "exampleFactoryName", &armdatafactory.FactoriesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}


func ExampleFactoriesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "exampleResourceGroup", "exampleFactoryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}


func ExampleFactoriesClient_GetGitHubAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetGitHubAccessToken(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.GitHubAccessTokenRequest{
		GitHubAccessCode:         to.Ptr("some"),
		GitHubAccessTokenBaseURL: to.Ptr("some"),
		GitHubClientID:           to.Ptr("some"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}


func ExampleFactoriesClient_GetDataPlaneAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewFactoriesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetDataPlaneAccess(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.UserAccessPolicy{
		AccessResourcePath: to.Ptr(""),
		ExpireTime:         to.Ptr("2018-11-10T09:46:20.2659347Z"),
		Permissions:        to.Ptr("r"),
		ProfileName:        to.Ptr("DefaultProfile"),
		StartTime:          to.Ptr("2018-11-10T02:46:20.2659347Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
