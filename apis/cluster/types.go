package cluster

type BasicInfo struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type ProviderOptions struct {
	Name          string `json:"name"`
	Credential    string `json:"credential,omitempty"`
	ClusterID     string `json:"clusterID,omitempty"`
	Project       string `json:"project,omitempty"`
	Region        string `json:"region,omitempty"`
	ResourceGroup string `json:"resourceGroup,omitempty"`
	KubeConfig    string `json:"kubeConfig,omitempty"`
}

type ComponentOptions struct {
	FluxCD        bool `json:"flux-cd,omitempty"`
	LicenseServer bool `json:"license-server,omitempty"`
}

type ListOptions struct {
	Provider string `json:"provider,omitempty"`
}

type GetOptions struct {
	Name string `json:"name,omitempty"`
}

type ImportOptions struct {
	BasicInfo  BasicInfo        `json:"basicInfo,omitempty"`
	Provider   ProviderOptions  `json:"provider,omitempty"`
	Components ComponentOptions `json:"components,omitempty"`
}

type ConnectOptions struct {
	Name       string `json:"name"`
	Credential string `json:"credential,omitempty"`
	KubeConfig string `json:"kubeConfig,omitempty"`
}

type ReconfigureOptions struct {
	Name       string           `json:"name"`
	Components ComponentOptions `json:"components,omitempty"`
}

type RemovalOptions struct {
	Name       string           `json:"name"`
	Components ComponentOptions `json:"components,omitempty"`
}
