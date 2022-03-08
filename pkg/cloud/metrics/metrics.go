/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/awserrors"
)

const (
	metricAWSSubsystem       = "aws"
	metricRequestCountKey    = "api_requests_total"
	metricRequestDurationKey = "api_request_duration_seconds"
	metricAPICallRetries     = "api_call_retries"
	metricServiceLabel       = "service"
	metricRegionLabel        = "region"
	metricOperationLabel     = "operation"
	metricControllerLabel    = "controller"
	metricStatusCodeLabel    = "status_code"
	metricErrorCodeLabel     = "error_code"
)

var (
	awsRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricAWSSubsystem,
		Name:      metricRequestCountKey,
		Help:      "Total number of AWS requests",
	}, []string{metricControllerLabel, metricServiceLabel, metricRegionLabel, metricOperationLabel, metricStatusCodeLabel, metricErrorCodeLabel})

	awsRequestDurationSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricAWSSubsystem,
		Name:      metricRequestDurationKey,
		Help:      "Latency of HTTP requests to AWS",
	}, []string{metricControllerLabel, metricServiceLabel, metricRegionLabel, metricOperationLabel})

	awsCallRetries = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricAWSSubsystem,
		Name:      metricAPICallRetries,
		Help:      "Number of retries made against an AWS API",
		Buckets:   []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}, []string{metricControllerLabel, metricServiceLabel, metricRegionLabel, metricOperationLabel})
)

func init() {
	metrics.Registry.MustRegister(awsRequestCount)
	metrics.Registry.MustRegister(awsRequestDurationSeconds)
	metrics.Registry.MustRegister(awsCallRetries)
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(DescribeSubnets)
	metrics.Registry.MustRegister(CreateNatGateway)
	metrics.Registry.MustRegister(DeleteNatGateway)
	metrics.Registry.MustRegister(WaitUntilNatGatewayAvailable)
	metrics.Registry.MustRegister(DescribeNatGateway)
	metrics.Registry.MustRegister(DescribeNatGatewaysPages)
	metrics.Registry.MustRegister(DescribeAvailabilityZones)
	metrics.Registry.MustRegister(CreateTags)
	metrics.Registry.MustRegister(CreateSubnet)
	metrics.Registry.MustRegister(WaitUntilSubnetAvailable)
	metrics.Registry.MustRegister(ModifySubnetAttribute)

	metrics.Registry.MustRegister(AWSCall)
	metrics.Registry.MustRegister(EnsureTags)
	metrics.Registry.MustRegister(K8sAPIServer)
	metrics.Registry.MustRegister(DescribeVpcsCalls)
	metrics.Registry.MustRegister(DescribeVpcAttributeCalls)
	metrics.Registry.MustRegister(ModifyVpcAttributeCalls)
	metrics.Registry.MustRegister(CreateVpcCalls)
	metrics.Registry.MustRegister(AssociateVpcCidrBlockCalls)
	metrics.Registry.MustRegister(DescribeInternetGateways)
	metrics.Registry.MustRegister(CreateInternetGateway)
	metrics.Registry.MustRegister(AttachInternetGateway)
	metrics.Registry.MustRegister(DescribeAddresses)
	metrics.Registry.MustRegister(AllocateAddress)
	metrics.Registry.MustRegister(CreateRouteTable)
	metrics.Registry.MustRegister(CreateRoute)
	metrics.Registry.MustRegister(DescribeRouteTables)
	metrics.Registry.MustRegister(AssociateRouteTable)
	metrics.Registry.MustRegister(DisassociateRouteTable)
	metrics.Registry.MustRegister(DeleteRouteTable)
	metrics.Registry.MustRegister(ReplaceRoute)
	metrics.Registry.MustRegister(DescribeSecurityGroupsInput)
	metrics.Registry.MustRegister(DeleteSecurityGroup)
	metrics.Registry.MustRegister(DescribeSecurityGroupsPages)
	metrics.Registry.MustRegister(DescribeSecurityGroups)
	metrics.Registry.MustRegister(CreateSecurityGroup)
	metrics.Registry.MustRegister(AuthorizeSecurityGroupIngress)
	metrics.Registry.MustRegister(RevokeSecurityGroupIngress)
	metrics.Registry.MustRegister(DescribeInstances)
	metrics.Registry.MustRegister(TerminateInstances)
	metrics.Registry.MustRegister(WaitUntilInstanceTerminated)
	metrics.Registry.MustRegister(RunInstances)
	metrics.Registry.MustRegister(WaitUntilInstanceRunningWithContext)
	metrics.Registry.MustRegister(DeleteTags)
	metrics.Registry.MustRegister(DescribeNetworkInterfaces)
	metrics.Registry.MustRegister(DescribeImages)
	metrics.Registry.MustRegister(DescribeNetworkInterfaceAttribute)
	metrics.Registry.MustRegister(ModifyNetworkInterfaceAttribute)
	metrics.Registry.MustRegister(CreateQueue)
	metrics.Registry.MustRegister(DeleteQueue)
	metrics.Registry.MustRegister(SetQueueAttributes)
	metrics.Registry.MustRegister(DescribeRule)
	metrics.Registry.MustRegister(GetQueueURL)
	metrics.Registry.MustRegister(GetQueueAttributes)
	metrics.Registry.MustRegister(ListTargetsByRule)
	metrics.Registry.MustRegister(PutTargets)
	metrics.Registry.MustRegister(PutRule)
	metrics.Registry.MustRegister(RemoveTargets)
	metrics.Registry.MustRegister(DeleteRule)
	metrics.Registry.MustRegister(AttachLoadBalancerToSubnets)
	metrics.Registry.MustRegister(DeleteLoadBalancer)
	metrics.Registry.MustRegister(ApplySecurityGroupsToLoadBalancer)
	metrics.Registry.MustRegister(DescribeLoadBalancers)
	metrics.Registry.MustRegister(RegisterInstancesWithLoadBalancer)
	metrics.Registry.MustRegister(DeregisterInstancesFromLoadBalancer)
	metrics.Registry.MustRegister(CreateLoadBalancer)
	metrics.Registry.MustRegister(ConfigureHealthCheck)
	metrics.Registry.MustRegister(ModifyLoadBalancerAttributes)
	metrics.Registry.MustRegister(DescribeLoadBalancersPages)
	metrics.Registry.MustRegister(DescribeTags)
	metrics.Registry.MustRegister(DescribeLoadBalancerAttributes)
	metrics.Registry.MustRegister(AddTags)
	metrics.Registry.MustRegister(RemoveTags)
}

// CaptureRequestMetrics will monitor and capture request metrics.
func CaptureRequestMetrics(controller string) func(r *request.Request) {
	return func(r *request.Request) {
		duration := time.Since(r.AttemptTime)
		operation := r.Operation.Name
		region := aws.StringValue(r.Config.Region)
		service := endpointToService(r.ClientInfo.Endpoint)
		statusCode := "0"
		errorCode := ""
		if r.HTTPResponse != nil {
			statusCode = strconv.Itoa(r.HTTPResponse.StatusCode)
		}
		if r.Error != nil {
			var ok bool
			if errorCode, ok = awserrors.Code(r.Error); !ok {
				errorCode = "internal"
			}
		}
		awsRequestCount.WithLabelValues(controller, service, region, operation, statusCode, errorCode).Inc()
		awsRequestDurationSeconds.WithLabelValues(controller, service, region, operation).Observe(duration.Seconds())
		awsCallRetries.WithLabelValues(controller, service, region, operation).Observe(float64(r.RetryCount))
	}
}

func endpointToService(endpoint string) string {
	endpointURL, err := url.Parse(endpoint)
	// If possible extract the service name, else return entire endpoint address
	if err == nil {
		host := endpointURL.Host
		components := strings.Split(host, ".")
		if len(components) > 0 {
			return components[0]
		}
	}
	return endpoint
}
