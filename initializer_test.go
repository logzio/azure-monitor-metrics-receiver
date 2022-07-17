package azure_monitor_metrics_receiver

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckConfigValidation_ResourceTargetsOnly(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.NoError(t, err)
}

func TestCheckConfigValidation_ResourceTargetWithNoResourceID(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget("", []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_ResourceTargetWithInvalidAggregation(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{testInvalidAggregation}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_ResourceGroupTargetsOnly(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{
						NewResource(testResourceType1, []string{}, []string{}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.NoError(t, err)
}

func TestCheckConfigValidation_ResourceGroupTargetWithoutResourceGroup(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					"",
					[]*Resource{
						NewResource(testResourceType1, []string{}, []string{}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_ResourceGroupTargetWithoutResources(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_ResourceGroupTargetWithResourceWithoutResourceType(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{
						NewResource("", []string{}, []string{}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_ResourceGroupTargetWithInvalidAggregation(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{
						NewResource(testResourceType1, []string{}, []string{testInvalidAggregation}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_SubscriptionTargetsOnly(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{
				NewResource(testResourceType1, []string{}, []string{}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.NoError(t, err)
}

func TestCheckConfigValidation_SubscriptionTargetWithoutResourceType(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{
				NewResource("", []string{}, []string{}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_SubscriptionTargetWithInvalidAggregation(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{
				NewResource(testResourceType1, []string{}, []string{testInvalidAggregation}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_AllTargetTypes(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{
						NewResource(testResourceType1, []string{}, []string{}),
					},
				),
			},
			[]*Resource{
				NewResource(testResourceType1, []string{}, []string{}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.NoError(t, err)
}

func TestCheckConfigValidation_NoTargets(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_NoSubscriptionID(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: "",
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_NoClientID(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       "",
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_NoClientSecret(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   "",
		tenantID:       testTenantID,
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestCheckConfigValidation_NoTenantID(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       "",
	}

	err := ammr.checkValidation()
	require.Error(t, err)
}

func TestAddPrefixToResourceTargetsResourceID_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testResourceGroup1ResourceType1Resource1, []string{}, []string{}),
				NewResourceTarget(testResourceGroup1ResourceType2Resource2, []string{}, []string{}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	ammr.addPrefixToResourceTargetsResourceID()

	assert.Equal(t, testFullResourceGroup1ResourceType1Resource1, ammr.Targets.ResourceTargets[0].ResourceID)
	assert.Equal(t, testFullResourceGroup1ResourceType2Resource2, ammr.Targets.ResourceTargets[1].ResourceID)
}

func TestCreateResourceTargetsFromResourceGroupTargets_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup1,
					[]*Resource{
						NewResource(testResourceType1, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
						NewResource(testResourceType2, []string{testMetric1}, []string{}),
					},
				),
				NewResourceGroupTarget(
					testResourceGroup2,
					[]*Resource{
						NewResource(testResourceType1, []string{testMetric3}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CreateResourceTargetsFromResourceGroupTargets()
	require.NoError(t, err)

	assert.Len(t, ammr.Targets.ResourceTargets, 3)

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Equal(t, ammr.Targets.resourceGroupTargets[0].resources[0].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.resourceGroupTargets[0].resources[0].aggregations, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, ammr.Targets.resourceGroupTargets[0].resources[1].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.resourceGroupTargets[0].resources[1].aggregations, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, ammr.Targets.resourceGroupTargets[1].resources[0].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.resourceGroupTargets[1].resources[0].aggregations, target.Aggregations)
		}
	}
}

func TestCreateResourceTargetsFromResourceGroupTargets_NoResourceFound(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{
				NewResourceGroupTarget(
					testResourceGroup2,
					[]*Resource{
						NewResource(testResourceType2, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
					},
				),
			},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CreateResourceTargetsFromResourceGroupTargets()
	require.Error(t, err)
}

func TestCreateResourceTargetsFromSubscriptionTargets_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{
				NewResource(testResourceType1, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
				NewResource(testResourceType2, []string{testMetric3}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeAverage)}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CreateResourceTargetsFromSubscriptionTargets()
	require.NoError(t, err)

	assert.Len(t, ammr.Targets.ResourceTargets, 3)

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Equal(t, ammr.Targets.subscriptionTargets[0].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.subscriptionTargets[0].aggregations, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, ammr.Targets.subscriptionTargets[1].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.subscriptionTargets[1].aggregations, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, ammr.Targets.subscriptionTargets[0].metrics, target.Metrics)
			assert.Equal(t, ammr.Targets.subscriptionTargets[0].aggregations, target.Aggregations)
		}
	}
}

func TestCreateResourceTargetsFromSubscriptionTargets_NoResourceFound(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{},
			[]*ResourceGroupTarget{},
			[]*Resource{
				NewResource(testResourceType3, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CreateResourceTargetsFromSubscriptionTargets()
	require.Error(t, err)
}

func TestCheckResourceTargetsMetricsValidation_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{}, []string{}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CheckResourceTargetsMetricsValidation()
	require.NoError(t, err)
}

func TestCheckResourceTargetsMetricsValidation_WithResourceTargetWithInvalidMetric(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{}, []string{}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testInvalidMetric}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.CheckResourceTargetsMetricsValidation()
	require.Error(t, err)
}

func TestSetResourceTargetsMetrics_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{}, []string{}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{}, []string{string(armmonitor.AggregationTypeAverage)}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.SetResourceTargetsMetrics()
	require.NoError(t, err)

	assert.Len(t, ammr.Targets.ResourceTargets, 3)

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Len(t, target.Metrics, 3)
			assert.Equal(t, []string{testMetric1, testMetric2, testMetric3}, target.Metrics)
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Len(t, target.Metrics, 2)
			assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Len(t, target.Metrics, 1)
			assert.Equal(t, []string{testMetric1}, target.Metrics)
		}
	}
}

func TestSplitResourceTargetsMetricsByMinTimeGrain_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1, testMetric2, testMetric3}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeAverage)}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	err := ammr.SplitResourceTargetsMetricsByMinTimeGrain()
	require.NoError(t, err)

	assert.Len(t, ammr.Targets.ResourceTargets, 4)

	for _, target := range ammr.Targets.ResourceTargets {
		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Contains(t, []int{1, 2}, len(target.Metrics))

			if len(target.Metrics) == 1 {
				assert.Equal(t, []string{testMetric3}, target.Metrics)
				assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
			}
			if len(target.Metrics) == 2 {
				assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
				assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
			}
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, []string{testMetric1}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal)}, target.Aggregations)
		}
	}
}

func TestSplitResourceTargetsWithMoreThanMaxMetrics_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeAverage)}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	for index := 1; index <= 25; index++ {
		ammr.Targets.ResourceTargets[0].Metrics = append(ammr.Targets.ResourceTargets[0].Metrics, testMetric1)
	}

	expectedResource1Metrics := make([]string, 0)
	for index := 1; index <= MaxMetricsPerRequest; index++ {
		expectedResource1Metrics = append(expectedResource1Metrics, testMetric1)
	}

	ammr.SplitResourceTargetsWithMoreThanMaxMetrics()

	assert.Len(t, ammr.Targets.ResourceTargets, 4)

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Contains(t, []int{5, MaxMetricsPerRequest}, len(target.Metrics))

			if len(target.Metrics) == MaxMetricsPerRequest {
				assert.Equal(t, expectedResource1Metrics, target.Metrics)
				assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
			}
			if len(target.Metrics) == 5 {
				assert.Equal(t, []string{testMetric1, testMetric1, testMetric1, testMetric1, testMetric1}, target.Metrics)
				assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
			}
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, []string{testMetric1}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal)}, target.Aggregations)
		}
	}
}

func TestChangeResourceTargetsMetricsWithComma(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1, testMetric2, testMetric3WithComma}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeAverage)}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	ammr.changeResourceTargetsMetricsWithComma()

	assert.Len(t, ammr.Targets.ResourceTargets, 3)

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Equal(t, []string{testMetric1, testMetric2, testMetric3ChangedComma}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, []string{testMetric1}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal)}, target.Aggregations)
		}
	}
}

func TestSetResourceTargetsAggregations_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1, testMetric2, testMetric3}, []string{}),
				NewResourceTarget(testFullResourceGroup1ResourceType2Resource2, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeAverage)}),
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
			},
			[]*ResourceGroupTarget{},
			[]*Resource{},
		),
		AzureClients:   setMockAzureClients(),
		subscriptionID: testSubscriptionID,
		clientID:       testClientID,
		clientSecret:   testClientSecret,
		tenantID:       testTenantID,
	}

	ammr.SetResourceTargetsAggregations()

	for _, target := range ammr.Targets.ResourceTargets {
		assert.Contains(t, []string{testFullResourceGroup1ResourceType1Resource1, testFullResourceGroup1ResourceType2Resource2, testFullResourceGroup2ResourceType1Resource3}, target.ResourceID)

		if target.ResourceID == testFullResourceGroup1ResourceType1Resource1 {
			assert.Equal(t, []string{testMetric1, testMetric2, testMetric3}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumAverage), string(armmonitor.AggregationTypeEnumCount), string(armmonitor.AggregationTypeEnumMaximum), string(armmonitor.AggregationTypeEnumMinimum), string(armmonitor.AggregationTypeEnumTotal)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup1ResourceType2Resource2 {
			assert.Equal(t, []string{testMetric1, testMetric2}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
		if target.ResourceID == testFullResourceGroup2ResourceType1Resource3 {
			assert.Equal(t, []string{testMetric1}, target.Metrics)
			assert.Equal(t, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}, target.Aggregations)
		}
	}
}
