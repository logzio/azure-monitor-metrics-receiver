package azuremonitormetricsreceiver

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectResourceTargetMetrics_AllDataWithValues(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1, testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	metrics, notCollectedMetrics, err := ammr.CollectResourceTargetMetrics(ammr.Targets.ResourceTargets[0])
	require.NoError(t, err)

	assert.Len(t, metrics, 2)
	assert.Len(t, notCollectedMetrics, 0)

	for _, metric := range metrics {
		assert.Contains(t, []string{"azure_monitor_microsoft_test_type1_metric1", "azure_monitor_microsoft_test_type1_metric2"}, metric.Name)
		assert.Len(t, metric.Fields, 3)

		for fieldKey := range metric.Fields {
			assert.Contains(t, []string{MetricFieldTotal, MetricFieldMaximum, MetricFieldTimeStamp}, fieldKey)
		}

		for tagKey := range metric.Tags {
			assert.Contains(t, []string{MetricTagSubscriptionID, MetricTagResourceGroup, MetricTagResourceName, MetricTagNamespace, MetricTagResourceRegion, MetricTagUnit}, tagKey)
		}

		if metric.Name == "azure_monitor_microsoft_test_type1_metric1" {
			assert.Equal(t, 5.0, metric.Fields[MetricFieldTotal])
			assert.Equal(t, 5.0, metric.Fields[MetricFieldMaximum])
			assert.Equal(t, "2022-02-22T22:59:00Z", metric.Fields[MetricFieldTimeStamp])

			assert.Equal(t, testSubscriptionID, metric.Tags[MetricTagSubscriptionID])
			assert.Equal(t, testResourceGroup1, metric.Tags[MetricTagResourceGroup])
			assert.Equal(t, testResource1Name, metric.Tags[MetricTagResourceName])
			assert.Equal(t, testResourceType1, metric.Tags[MetricTagNamespace])
			assert.Equal(t, testResourceRegion, metric.Tags[MetricTagResourceRegion])
			assert.Equal(t, string(armmonitor.MetricUnitCount), metric.Tags[MetricTagUnit])
		}

		if metric.Name == "azure_monitor_microsoft_test_type1_metric2" {
			assert.Equal(t, 2.5, metric.Fields[MetricFieldTotal])
			assert.Equal(t, 2.5, metric.Fields[MetricFieldMaximum])
			assert.Equal(t, "2022-02-22T22:59:00Z", metric.Fields[MetricFieldTimeStamp])

			assert.Equal(t, testSubscriptionID, metric.Tags[MetricTagSubscriptionID])
			assert.Equal(t, testResourceGroup1, metric.Tags[MetricTagResourceGroup])
			assert.Equal(t, testResource1Name, metric.Tags[MetricTagResourceName])
			assert.Equal(t, testResourceType1, metric.Tags[MetricTagNamespace])
			assert.Equal(t, testResourceRegion, metric.Tags[MetricTagResourceRegion])
			assert.Equal(t, string(armmonitor.MetricUnitCount), metric.Tags[MetricTagUnit])
		}
	}
}

func TestCollectResourceTargetMetrics_LastDataWithNoValue(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType1Resource3, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMinimum)}),
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

	metrics, notCollectedMetrics, err := ammr.CollectResourceTargetMetrics(ammr.Targets.ResourceTargets[0])
	require.NoError(t, err)

	assert.Len(t, metrics, 1)
	assert.Len(t, notCollectedMetrics, 0)

	assert.Equal(t, "azure_monitor_microsoft_test_type1_metric1", metrics[0].Name)
	assert.Len(t, metrics[0].Fields, 3)

	for fieldKey := range metrics[0].Fields {
		assert.Contains(t, []string{MetricFieldTotal, MetricFieldMinimum, MetricFieldTimeStamp}, fieldKey)
	}

	for tagKey := range metrics[0].Tags {
		assert.Contains(t, []string{MetricTagSubscriptionID, MetricTagResourceGroup, MetricTagResourceName, MetricTagNamespace, MetricTagResourceRegion, MetricTagUnit}, tagKey)
	}

	assert.Equal(t, 2.5, metrics[0].Fields[MetricFieldTotal])
	assert.Equal(t, 2.5, metrics[0].Fields[MetricFieldMinimum])
	assert.Equal(t, "2022-02-22T22:58:00Z", metrics[0].Fields[MetricFieldTimeStamp])

	assert.Equal(t, testSubscriptionID, metrics[0].Tags[MetricTagSubscriptionID])
	assert.Equal(t, testResourceGroup2, metrics[0].Tags[MetricTagResourceGroup])
	assert.Equal(t, testResource3Name, metrics[0].Tags[MetricTagResourceName])
	assert.Equal(t, testResourceType1, metrics[0].Tags[MetricTagNamespace])
	assert.Equal(t, testResourceRegion, metrics[0].Tags[MetricTagResourceRegion])
	assert.Equal(t, string(armmonitor.MetricUnitBytes), metrics[0].Tags[MetricTagUnit])
}

func TestCollectResourceTargetMetrics_AllDataWithNoValues(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType2Resource4, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	metrics, notCollectedMetrics, err := ammr.CollectResourceTargetMetrics(ammr.Targets.ResourceTargets[0])
	require.NoError(t, err)

	assert.Len(t, metrics, 0)
	assert.Len(t, notCollectedMetrics, 1)

	assert.Equal(t, testFullResourceGroup2ResourceType2Resource4+"/providers/Microsoft.Insights/metrics/metric1", notCollectedMetrics[0])
}

func TestCollectResourceTargetMetrics_EmptyData(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType2Resource5, []string{testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	metrics, notCollectedMetrics, err := ammr.CollectResourceTargetMetrics(ammr.Targets.ResourceTargets[0])
	require.NoError(t, err)

	assert.Len(t, metrics, 0)
	assert.Len(t, notCollectedMetrics, 1)

	assert.Equal(t, testFullResourceGroup2ResourceType2Resource5+"/providers/Microsoft.Insights/metrics/metric2", notCollectedMetrics[0])
}

func TestCollectResourceTargetMetrics_EmptyTimeseries(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType2Resource6, []string{testMetric2}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	metrics, notCollectedMetrics, err := ammr.CollectResourceTargetMetrics(ammr.Targets.ResourceTargets[0])
	require.NoError(t, err)

	assert.Len(t, metrics, 0)
	assert.Len(t, notCollectedMetrics, 1)

	assert.Equal(t, testFullResourceGroup2ResourceType2Resource6+"/providers/Microsoft.Insights/metrics/metric2", notCollectedMetrics[0])
}

func TestGetMetricName_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumAverage)}),
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricName, err := createMetricName(response.Value[0], &response)
	require.NoError(t, err)
	require.NotNil(t, metricName)

	assert.Equal(t, "azure_monitor_microsoft_test_type1_metric1", *metricName)
}

func TestGetMetricFields_AllTimeseriesWithData(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricFields := getMetricFields(response.Value[0].Timeseries[0].Data)
	require.NotNil(t, metricFields)

	assert.Len(t, metricFields, 3)

	assert.Equal(t, "2022-02-22T22:59:00Z", metricFields[MetricFieldTimeStamp])
	assert.Equal(t, 5.0, metricFields[MetricFieldTotal])
	assert.Equal(t, 5.0, metricFields[MetricFieldMaximum])
}

func TestGetMetricFields_LastTimeseriesWithoutData(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricFields := getMetricFields(response.Value[0].Timeseries[0].Data)
	require.NotNil(t, metricFields)

	assert.Len(t, metricFields, 3)

	assert.Equal(t, "2022-02-22T22:58:00Z", metricFields[MetricFieldTimeStamp])
	assert.Equal(t, 2.5, metricFields[MetricFieldTotal])
	assert.Equal(t, 2.5, metricFields[MetricFieldMinimum])
}

func TestGetMetricFields_AllTimeseriesWithoutData(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType2Resource4, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricFields := getMetricFields(response.Value[0].Timeseries[0].Data)
	require.Nil(t, metricFields)
}

func TestGetMetricFields_NoTimeseriesData(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup2ResourceType2Resource5, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricFields := getMetricFields(response.Value[0].Timeseries[0].Data)
	require.Nil(t, metricFields)
}

func TestGetMetricTags_Success(t *testing.T) {
	ammr := &AzureMonitorMetricsReceiver{
		Targets: NewTargets(
			[]*ResourceTarget{
				NewResourceTarget(testFullResourceGroup1ResourceType1Resource1, []string{testMetric1}, []string{string(armmonitor.AggregationTypeEnumTotal), string(armmonitor.AggregationTypeEnumMaximum)}),
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

	response, err := ammr.AzureClients.MetricsClient.List(ammr.AzureClients.Ctx, ammr.Targets.ResourceTargets[0].ResourceID, nil)
	assert.NoError(t, err)

	metricTags, err := getMetricTags(response.Value[0], &response)
	require.NoError(t, err)
	require.NotNil(t, metricTags)

	assert.Len(t, metricTags, 6)

	assert.Equal(t, testSubscriptionID, metricTags[MetricTagSubscriptionID])
	assert.Equal(t, testResourceGroup1, metricTags[MetricTagResourceGroup])
	assert.Equal(t, testResource1Name, metricTags[MetricTagResourceName])
	assert.Equal(t, testResourceType1, metricTags[MetricTagNamespace])
	assert.Equal(t, testResourceRegion, metricTags[MetricTagResourceRegion])
	assert.Equal(t, string(armmonitor.MetricUnitCount), metricTags[MetricTagUnit])
}
