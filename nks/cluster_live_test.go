package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAwsCluster = Cluster{
	Name:                  "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:              "aws",
	MasterCount:           1,
	MasterSize:            "t2.medium",
	WorkerCount:           2,
	WorkerSize:            "t2.medium",
	Region:                "us-east-1",
	Zone:                  "us-east-1a",
	ProviderNetworkID:     "__new__",
	ProviderNetworkCdr:    "172.23.0.0/16",
	ProviderSubnetID:      "__new__",
	ProviderSubnetCidr:    "172.23.1.0/24",
	KubernetesVersion:     "v1.15.5",
	KubernetesPodCidr:     "10.2.0.0",
	KubernetesServiceCidr: "10.3.0.0",
	RbacEnabled:           true,
	DashboardEnabled:      true,
	EtcdType:              "classic",
	Platform:              "coreos",
	Channel:               "stable",
	NetworkComponents:     []NetworkComponent{},
	Solutions:             []Solution{Solution{Solution: "helm_tiller"}},
}

var testAzureCluster = Cluster{
	Name:                  "Test Azure Cluster Go SDK " + GetTicks(),
	Provider:              "azure",
	MasterCount:           1,
	MasterSize:            "standard_d2s_v3",
	WorkerCount:           2,
	WorkerSize:            "standard_d2s_v3",
	Region:                "eastus",
	ProviderResourceGp:    "__new__",
	ProviderNetworkID:     "__new__",
	ProviderNetworkCdr:    "10.0.0.0/16",
	ProviderSubnetID:      "__new__",
	ProviderSubnetCidr:    "10.0.0.0/24",
	KubernetesVersion:     "v1.15.5",
	KubernetesPodCidr:     "10.2.0.0",
	KubernetesServiceCidr: "10.3.0.0",
	RbacEnabled:           true,
	DashboardEnabled:      true,
	EtcdType:              "classic",
	Platform:              "coreos",
	Channel:               "stable",
	NetworkComponents:     []NetworkComponent{},
	Solutions:             []Solution{Solution{Solution: "helm_tiller"}},
}

var testGCECluster = Cluster{
	Name:                  "Test GCE Cluster Go SDK " + GetTicks(),
	Provider:              "gce",
	MasterCount:           1,
	MasterSize:            "n1-standard-1",
	WorkerCount:           2,
	WorkerSize:            "n1-standard-1",
	Region:                "us-east1-c",
	ProviderNetworkID:     "__new__",
	ProviderNetworkCdr:    "172.23.0.0/16",
	ProviderSubnetID:      "__new__",
	ProviderSubnetCidr:    "172.23.1.0/24",
	KubernetesVersion:     "v1.15.5",
	KubernetesPodCidr:     "10.2.0.0",
	KubernetesServiceCidr: "10.3.0.0",
	RbacEnabled:           true,
	DashboardEnabled:      true,
	EtcdType:              "classic",
	Platform:              "coreos",
	Channel:               "stable",
	NetworkComponents:     []NetworkComponent{},
	Solutions:             []Solution{Solution{Solution: "helm_tiller"}},
}

var clusterIds = make([]int, 0)

var timeout = 3600

func TestLiveBasicCluster(t *testing.T) {
	t.Run("create clusters", func(t *testing.T) {
		t.Run("aws", testClusterCreateAWS)
		t.Run("azure", testClusterCreateAzure)
		t.Run("gce", testClusterCreateGCE)
	})

	t.Run("get clusters", func(t *testing.T) {
		t.Run("list", testClusterList)
		t.Run("get", testClusterGet)
	})

	t.Run("delete clusters", func(t *testing.T) {
		t.Run("delete", testClusterDelete)
	})
}

func testClusterCreateAWS(t *testing.T) {
	t.Parallel()

	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testAwsCluster.ProviderKey = awsKeysetID
	testAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testAwsCluster)
	if err != nil {
		t.Fatal(err)
	}

	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)
}

func testClusterCreateAzure(t *testing.T) {
	t.Parallel()

	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	azureKeysetID, err := GetIDFromEnv("NKS_AZR_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	testAzureCluster.ProviderKey = azureKeysetID
	testAzureCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testAzureCluster)
	fmt.Println("No orgs found, but no error", err)
	fmt.Println("AZR", cluster.ID, err)
	if err != nil {
		t.Fatalf("failed to create azure cluster with error %d", err)
	}

	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)
}

func testClusterCreateGCE(t *testing.T) {
	t.Parallel()

	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	gceKeysetID, err := GetIDFromEnv("NKS_GCE_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	testGCECluster.ProviderKey = gceKeysetID
	testGCECluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testGCECluster)
	fmt.Println("GKE", cluster.ID)
	if err != nil {
		t.Fatal(err)
	}

	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Fatal(err)
	}

	clusterIds = append(clusterIds, cluster.ID)
}

func testClusterList(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	clusters, err := c.GetClusters(orgID)
	fmt.Println(clusters)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, len(clusters) > 0, "There should be at lease one cluster")
}

func testClusterGet(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	if len(clusterIds) == 0 {
		t.Fatal("no clusters where created to get")
	}

	cluster, err := c.GetCluster(orgID, clusterIds[0])
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, cluster, "Cluster does not exists")
}

func testClusterDelete(t *testing.T) {
	for _, clusterID := range clusterIds {
		t.Run(string(clusterID), func(t *testing.T) {
			clusterDelete(t, clusterID)
		})
	}
}

func clusterDelete(t *testing.T, clusterID int) {
	t.Parallel()

	c, err := NewClientFromEnv()

	if err != nil {
		t.Fatal(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	err = c.DeleteCluster(orgID, clusterID)
	if err != nil {
		t.Fatal(err)
	}

	err = c.WaitClusterDeleted(orgID, clusterID, timeout)
	if err != nil {
		t.Fatal(err)
	}
}
