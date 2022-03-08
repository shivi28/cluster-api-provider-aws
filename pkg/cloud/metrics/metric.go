package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	DescribeSubnets = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "describe_subnets",
			Help: "DescribeSubnets",
		},
	)

	CreateNatGateway = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateNatGateway",
			Help: "CreateNatGateway",
		},
	)

	DeleteNatGateway = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteNatGateway",
			Help: "DeleteNatGateway",
		},
	)

	WaitUntilNatGatewayAvailable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "WaitUntilNatGatewayAvailable",
			Help: "WaitUntilNatGatewayAvailable",
		},
	)

	DescribeNatGateway = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeNatGateway",
			Help: "DescribeNatGateway",
		},
	)

	DescribeNatGatewaysPages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeNatGatewaysPages",
			Help: "DescribeNatGatewaysPages",
		},
	)

	DescribeAvailabilityZones = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeAvailabilityZones",
			Help: "DescribeAvailabilityZones",
		},
	)

	CreateTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateTags",
			Help: "CreateTags",
		},
	)

	CreateSubnet = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateSubnet",
			Help: "CreateSubnet",
		},
	)

	WaitUntilSubnetAvailable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "WaitUntilSubnetAvailable",
			Help: "WaitUntilSubnetAvailable",
		},
	)

	ModifySubnetAttribute = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ModifySubnetAttribute",
			Help: "ModifySubnetAttribute",
		},
	)

	AWSCall = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "aws_api_call_total",
			Help: "Number of calls made from AWSCluster controller to AWS",
		},
	)

	EnsureTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "aEnsureTags",
			Help: "EnsureTags",
		},
	)

	K8sAPIServer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "k8s_api_server_calls_total",
			Help: "Number of API server calls from AWSCluster controller",
		},
	)

	DescribeVpcsCalls = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "describe_vpc_calls",
			Help: "DescribeVpcs",
		})

	DescribeVpcAttributeCalls = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "describe_vpc_attribute_calls",
			Help: "DescribeVpcAttribute",
		})

	ModifyVpcAttributeCalls = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "modify_vpc_attribute_calls",
			Help: "ModifyVpcAttribute",
		})

	CreateVpcCalls = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "create_vpc_calls",
			Help: "CreateVpcCalls",
		})
	AssociateVpcCidrBlockCalls = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "associate_vpc_calls",
			Help: "AssociateVpcCidrBlock",
		})

	DescribeInternetGateways = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "describe_internet_gateways",
			Help: "DescribeInternetGateways",
		})

	CreateInternetGateway = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "create_internet_gateway",
			Help: "CreateInternetGateway",
		})

	AttachInternetGateway = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "attach_internet_gateway",
			Help: "AttachInternetGateway",
		})

	DescribeAddresses = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeAddresses",
			Help: "DescribeAddresses",
		})

	AllocateAddress = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "AllocateAddress",
			Help: "AllocateAddress",
		})

	CreateRouteTable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateRouteTable",
			Help: "CreateRouteTable",
		})

	CreateRoute = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateRoute",
			Help: "CreateRoute",
		})

	DescribeRouteTables = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeRouteTables",
			Help: "DescribeRouteTables",
		},
	)

	AssociateRouteTable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "AssociateRouteTable",
			Help: "AssociateRouteTable",
		},
	)
	DisassociateRouteTable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DisassociateRouteTable",
			Help: "DisassociateRouteTable",
		},
	)
	DeleteRouteTable = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteRouteTable",
			Help: "DeleteRouteTable",
		},
	)

	ReplaceRoute = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ReplaceRoute",
			Help: "ReplaceRoute",
		},
	)

	DescribeSecurityGroupsInput = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeSecurityGroupsInput",
			Help: "DescribeSecurityGroupsInput",
		})

	DeleteSecurityGroup = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteSecurityGroup",
			Help: "DeleteSecurityGroup",
		})

	DescribeSecurityGroupsPages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeSecurityGroupsPages",
			Help: "DescribeSecurityGroupsPages",
		})
	DescribeSecurityGroups = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeSecurityGroups",
			Help: "DescribeSecurityGroups",
		})

	CreateSecurityGroup = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateSecurityGroup",
			Help: "CreateSecurityGroup",
		})

	AuthorizeSecurityGroupIngress = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "AuthorizeSecurityGroupIngress",
			Help: "AuthorizeSecurityGroupIngress",
		})

	RevokeSecurityGroupIngress = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "RevokeSecurityGroupIngress",
			Help: "RevokeSecurityGroupIngress",
		})

	DescribeInstances = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeInstances",
			Help: "DescribeInstances",
		})
	TerminateInstances = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "TerminateInstances",
			Help: "TerminateInstances",
		})
	WaitUntilInstanceTerminated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "WaitUntilInstanceTerminated",
			Help: "WaitUntilInstanceTerminated",
		})
	RunInstances = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "RunInstances",
			Help: "RunInstances",
		})
	WaitUntilInstanceRunningWithContext = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "WaitUntilInstanceRunningWithContext",
			Help: "WaitUntilInstanceRunningWithContext",
		})
	DeleteTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteTags",
			Help: "DeleteTags",
		})
	DescribeNetworkInterfaces = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeNetworkInterfaces",
			Help: "DescribeNetworkInterfaces",
		})
	DescribeImages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeImages",
			Help: "DescribeImages",
		})
	DescribeNetworkInterfaceAttribute = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeNetworkInterfaceAttribute",
			Help: "DescribeNetworkInterfaceAttribute",
		})
	ModifyNetworkInterfaceAttribute = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ModifyNetworkInterfaceAttribute",
			Help: "ModifyNetworkInterfaceAttribute",
		})

	CreateQueue = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateQueue",
			Help: "CreateQueue",
		})

	DeleteQueue = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteQueue",
			Help: "DeleteQueue",
		})

	SetQueueAttributes = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "SetQueueAttributes",
			Help: "SetQueueAttributes",
		})

	DescribeRule = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeRule",
			Help: "DescribeRule",
		})

	GetQueueURL = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "GetQueueURL",
			Help: "GetQueueURL",
		})

	GetQueueAttributes = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "GetQueueAttributes",
			Help: "GetQueueAttributes",
		})

	ListTargetsByRule = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ListTargetsByRule",
			Help: "ListTargetsByRule",
		})

	PutTargets = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "PutTargets",
			Help: "PutTargets",
		})

	PutRule = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "PutRule",
			Help: "PutRule",
		})

	RemoveTargets = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "RemoveTargets",
			Help: "RemoveTargets",
		})

	DeleteRule = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteRule",
			Help: "DeleteRule",
		})

	AttachLoadBalancerToSubnets = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "AttachLoadBalancerToSubnets",
			Help: "AttachLoadBalancerToSubnets",
		},
	)

	DeleteLoadBalancer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeleteLoadBalancer",
			Help: "DeleteLoadBalancer",
		},
	)

	ApplySecurityGroupsToLoadBalancer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ApplySecurityGroupsToLoadBalancer",
			Help: "ApplySecurityGroupsToLoadBalancer",
		},
	)

	DescribeLoadBalancers = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeLoadBalancers",
			Help: "DescribeLoadBalancers",
		},
	)

	RegisterInstancesWithLoadBalancer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "RegisterInstancesWithLoadBalancer",
			Help: "RegisterInstancesWithLoadBalancer",
		},
	)

	DeregisterInstancesFromLoadBalancer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DeregisterInstancesFromLoadBalancer",
			Help: "DeregisterInstancesFromLoadBalancer",
		},
	)

	CreateLoadBalancer = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "CreateLoadBalancer",
			Help: "CreateLoadBalancer",
		},
	)

	ConfigureHealthCheck = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ConfigureHealthCheck",
			Help: "ConfigureHealthCheck",
		},
	)

	ModifyLoadBalancerAttributes = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ModifyLoadBalancerAttributes",
			Help: "ModifyLoadBalancerAttributes",
		},
	)

	DescribeLoadBalancersPages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeLoadBalancersPages",
			Help: "DescribeLoadBalancersPages",
		},
	)

	DescribeTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeTags",
			Help: "DescribeTags",
		},
	)

	DescribeLoadBalancerAttributes = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "DescribeLoadBalancerAttributes",
			Help: "DescribeLoadBalancerAttributes",
		},
	)

	AddTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "AddTags",
			Help: "AddTags",
		},
	)

	RemoveTags = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "RemoveTags",
			Help: "RemoveTags",
		},
	)
)
