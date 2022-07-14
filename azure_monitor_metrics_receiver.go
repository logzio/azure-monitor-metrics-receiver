package azure_monitor_metrics_receiver

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type AzureMonitorMetricsReceiver struct {
	Targets 	         *Targets
	AzureClients         *AzureClients

	subscriptionID       string
	clientID             string
	clientSecret         string
	tenantID             string
}

type Targets struct {
	ResourceTargets      []*ResourceTarget

	resourceGroupTargets []*ResourceGroupTarget
	subscriptionTargets  []*Resource
}

type ResourceTarget struct {
	ResourceID   string
	Metrics      []string
	Aggregations []string
}

type ResourceGroupTarget struct {
	resourceGroup string
	resources     []*Resource
}

type Resource struct {
	resourceType string
	metrics      []string
	aggregations []string
}

type AzureClients struct {
	Ctx						context.Context
	ResourcesClient         ResourcesClient
	MetricDefinitionsClient MetricDefinitionsClient
	MetricsClient           MetricsClient
}

type Metric struct {
	Name string
	Fields map[string]interface{}
	Tags map[string]string
}

type azureResourcesClient struct {
	client *armresources.Client
}

type ResourcesClient interface {
	List(context.Context, *armresources.ClientListOptions) ([]*armresources.ClientListResponse, error)
	ListByResourceGroup(context.Context, string, *armresources.ClientListByResourceGroupOptions) ([]*armresources.ClientListByResourceGroupResponse, error)
}

type MetricDefinitionsClient interface {
	List(context.Context, string, *armmonitor.MetricDefinitionsClientListOptions) (armmonitor.MetricDefinitionsClientListResponse, error)
}

type MetricsClient interface {
	List(context.Context, string, *armmonitor.MetricsClientListOptions) (armmonitor.MetricsClientListResponse, error)
}

func NewAzureMonitorMetricsReceiver(subscriptionID string, clientID string, clientSecret string, tenantID string, targets *Targets, azureClients *AzureClients) (*AzureMonitorMetricsReceiver, error) {
	azureMonitorMetricsReceiver := &AzureMonitorMetricsReceiver{
		Targets: targets,
		AzureClients: azureClients,
		subscriptionID: subscriptionID,
		clientID: clientID,
		clientSecret: clientSecret,
		tenantID: tenantID,
	}

	if err := azureMonitorMetricsReceiver.checkValidation(); err != nil {
		return nil, fmt.Errorf("got validation error: %v", err)
	}

	azureMonitorMetricsReceiver.addPrefixToResourceTargetsResourceID()
	return azureMonitorMetricsReceiver, nil
}

func NewTargets(resourceTargets []*ResourceTarget, resourceGroupTargets []*ResourceGroupTarget, subscriptionTargets []*Resource) *Targets {
	return &Targets{
		ResourceTargets: resourceTargets,
		resourceGroupTargets: resourceGroupTargets,
		subscriptionTargets: subscriptionTargets,
	}
}

func NewResourceTarget(resourceID string, metrics []string, aggregations []string) *ResourceTarget {
	return &ResourceTarget{
		ResourceID:   resourceID,
		Metrics:      metrics,
		Aggregations: aggregations,
	}
}

func NewResourceGroupTarget(resourceGroup string, resources []*Resource) *ResourceGroupTarget {
	return &ResourceGroupTarget{
		resourceGroup: resourceGroup,
		resources: resources,
	}
}

func NewResource(resourceType string, metrics []string, aggregations []string) *Resource {
	return &Resource{
		resourceType: resourceType,
		metrics: metrics,
		aggregations: aggregations,
	}
}

func newAzureResourcesClient(subscriptionID string, credential *azidentity.ClientSecretCredential) *azureResourcesClient {
	return &azureResourcesClient{
		client: armresources.NewClient(subscriptionID, credential, nil),
	}
}
