package metrics

//CollectionRequest - JSON payload posted by the scheduler
type CollectionRequest struct {
	JobRunID         string  `json:"jobRunId"`
	JobType          string  `json:"jobType"`
	CollectorPayload Payload `json:"collectorPayload"`
}

// Payload struct received by the collector
type Payload struct {
	JobID                 string       `json:"jobId"`
	JobName               string       `json:"jobName"`
	TenantID              string       `json:"tenantId"`
	Anchors               []Anchor     `json:"anchors"`
	Target                TargetConfig `json:"target"`
	StartTimeOffsetInSecs int64        `json:"startTimeOffsetInSeconds"`
	PeriodSecs            int64        `json:"periodSecs"`
	FilterConf            FilterConfig `json:"filterConfig"`
	ScopeConf             ScopeConfig  `json:"scopeConfig"`
	MetricConf            MetricConfig `json:"metricConfig"`
}

//Anchor struct contains the anchor info
type Anchor struct {
	AnchorID   string `json:"anchorId"`
	AnchorName string `json:"anchorName"`
	AnchorType string `json:"anchorType"`
}

//TargetConfig struct contains the target configuration
type TargetConfig struct {
	ID                      string                 `json:"id"`
	Name                    string                 `json:"name"`
	Description             string                 `json:"description"`
	Type                    string                 `json:"type"`
	Host                    string                 `json:"host"`
	Port                    int                    `json:"port"`
	Protocol                string                 `json:"protocol"`
	Endpoint                string                 `json:"endpoint"`
	RequestTimeoutInSeconds int                    `json:"requestTimeoutInSeconds"`
	ProxyURL                string                 `json:"proxyUrl"`
	ProxyCredentials        ProxyCredentialsConfig `json:"proxyCredentials"`
	Credentials             CredentialsConfig      `json:"credentials"`
}

//ProxyCredentialsConfig struct contains the proxy configuration
type ProxyCredentialsConfig struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Principal  string `json:"principal"`
	Credential string `json:"credential"`
}

//CredentialsConfig struct contains the credential configuration
type CredentialsConfig struct {
	ID      string                   `json:"id"`
	Name    string                   `json:"name"`
	Type    string                   `json:"type"`
	Details []CredentialDetailConfig `json:"details"`
}

//CredentialDetailConfig contains the key values pairs of the different credentails
type CredentialDetailConfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//FilterConfig contains the filter related Info such as tags , service list and resource list
type FilterConfig struct {
	TargetTags     []Tag      `json:"targetTags"`
	ServiceList    []string   `json:"serviceList"`
	ResourceList   []Resource `json:"resourceList"`
	CustomTagsList []string   `json:"customTagsList"`
}

//Resource Resource List
type Resource struct {
	CIID       string                 `json:"ciId"`
	CIName     string                 `json:"ciname"`
	CIType     string                 `json:"ciType"`
	TenantID   string                 `json:"tenantId"`
	Tags       map[string][]string    `json:"tags"`
	Properties map[string]interface{} `json:"properties"`
}

//Tag struct contains the name and values of tag
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//ScopeConfig struct
type ScopeConfig struct {
	EntityQuery string   `json:"entityQuery"`
	EntityList  []string `json:"entityList"`
}

//MetricConfig contains the metric configuration information
type MetricConfig struct {
	Expressions []string `json:"expressions"`
	MetricList  []string `json:"metricList"`
}

//TimeInfo struct
type TimeInfo struct {
	OffsetSecs int64
	PeriodSecs int64
}
