package clb

type DescribeLoadBalancersArgs struct {
	LoadBalancerIds  LoadBalancerIds  `url:"loadBalancerIds,omitempty"`
	LoadBalancerType int              `url:"loadBalancerType,omitempty"`
	LoadBalancerName string           `url:"loadBalancerName,omitempty"`
	Domain           string           `url:"domain,omitempty"`
	LoadBalancerVips LoadBalancerVips `url:"loadBalancerVips,omitempty"`
	BackendWanIps    BackendWanIps    `url:"backendWanIps,omitempty"`
	Offset           int              `url:"offset,omitempty"`
	Limit            int              `url:"limit,omitempty"`
	OrderBy          string           `url:"orderBy,omitempty"`
	OrderType        int              `url:"orderType,omitempty"`
	SearchKey        string           `url:"searchKey,omitempty"`
	ProjectId        int              `url:"projectId,omitempty"`
	Forward          int              `url:"forward,omitempty"`
	WithRs           int              `url:"withRs,omitempty"`
}

type LoadBalancerIds []string

type LoadBalancerVips []string

type BackendWanIps []string

type DescribeLoadBalancersResponse struct {
	Code            int            `json:"code"`
	Message         string         `json:"message"`
	TotalCount      int            `json:"totalCount"`
	LoadBalancerSet []LoadBalancer `json:"loadBalancerSet"`
	CodeDesc        string         `json:"codeDesc"`
}

type LoadBalancer struct {
	LoadBalancerId   string           `json:"loadBalancerId"`
	UnLoadBalancerId string           `json:"unLoadBalancerId"`
	LoadBalancerName string           `json:"loadBalancerName"`
	LoadBalancerType int              `json:"loadBalancerType"`
	Domain           string           `json:"domain"`
	LoadBalancerVips LoadBalancerVips `json:"loadBalancerVips"`
	Status           int              `json:"status"`
	CreateTime       string           `json:"createTime"`
	StatusTime       string           `json:"statusTime"`
	ProjectId        int              `json:"projectId"`
	VpcId            int              `json:"vpcId"`
	SubnetId         int              `json:"subnetId"`
}

func (client *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResponse, error) {
	response := &DescribeLoadBalancersResponse{}
	err := client.Invoke("DescribeLoadBalancers", args, response)
	if err != nil {
		return &DescribeLoadBalancersResponse{}, err
	}
	return response, nil
}

type InquiryLBPriceArgs struct {
	LoadBalancerType int `url:"loadBalancerType"`
}

type InquiryLBPriceResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Price   int    `json:"price"`
}

func (client *Client) InquiryLBPrice(args *InquiryLBPriceArgs) (*InquiryLBPriceResponse, error) {
	response := &InquiryLBPriceResponse{}
	err := client.Invoke("InquiryLBPrice", args, response)
	if err != nil {
		return &InquiryLBPriceResponse{}, err
	}
	return response, nil
}

type CreateLoadBalancerArgs struct {
	LoadBalancerType int    `url:"loadBalancerType"`
	Forward          int    `url:"forward"`
	LoadBalancerName string `url:"loadBalancerName"`
	DomainPrefix     string `url:"domainPrefix"`
	VpcId            string `url:"vpcId"`
	SubnetId         string `url:"subnetId"`
	ProjectId        int    `url:"projectId"`
	Number           int    `url:"number"`
}

type CreateLoadBalancerResponse struct {
	Code              int                 `json:"code"`
	Message           string              `json:"message"`
	CodeDesc          string              `json:"codeDesc"`
	UnLoadBalancerIds map[string][]string `json:"unLoadBalancerIds"`
	DealIds           []string            `json:"dealIds"`
}

func (client *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (*CreateLoadBalancerResponse, error) {
	response := &CreateLoadBalancerResponse{}
	err := client.Invoke("CreateLoadBalancer", args, response)
	if err != nil {
		return &CreateLoadBalancerResponse{}, err
	}
	return response, nil
}

type ModifyLoadBalancerAttributesArgs struct {
	LoadBalancerId   string `url:"loadBalancerId"`
	LoadBalancerName string `url:"loadBalancerName"`
	DomainPrefix     string `url:"domainPrefix"`
}

type ModifyLoadBalancerAttributesResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId int    `json:"requestId"`
}

func (client *Client) ModifyLoadBalancerAttributes(args *ModifyLoadBalancerAttributesArgs) (*ModifyLoadBalancerAttributesResponse, error) {
	response := &ModifyLoadBalancerAttributesResponse{}
	err := client.Invoke("ModifyLoadBalancerAttributes", args, response)
	if err != nil {
		return &ModifyLoadBalancerAttributesResponse{}, err
	}
	return response, nil
}

type DeleteLoadBalancersArgs struct {
	LoadBalancerIds []string `url:"loadBalancerIds"`
}

type DeleteLoadBalancersResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId int    `json:"requestId"`
}

func (client *Client) DeleteLoadBalancers(args *DeleteLoadBalancersArgs) (*DeleteLoadBalancersResponse, error) {
	response := &DeleteLoadBalancersResponse{}
	err := client.Invoke("DeleteLoadBalancers", args, response)
	if err != nil {
		return &DeleteLoadBalancersResponse{}, err
	}
	return response, nil
}

type DescribeLoadBalancersTaskResultArgs struct {
	RequestId int `url:"requestId"`
}

type DescribeLoadBalancersTaskResultResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Status int `json:"status"`
	} `json:"data"`
}

func (client *Client) DescribeLoadBalancersTaskResult(args *DescribeLoadBalancersTaskResultArgs) (*DescribeLoadBalancersTaskResultResponse, error) {
	response := &DescribeLoadBalancersTaskResultResponse{}
	err := client.Invoke("DescribeLoadBalancersTaskResult", args, response)
	if err != nil {
		return &DescribeLoadBalancersTaskResultResponse{}, err
	}
	return response, nil
}
