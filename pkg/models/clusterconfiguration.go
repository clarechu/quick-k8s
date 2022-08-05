package models

type ClusterConfiguration struct {
	ClusterName       string `yaml:"cluster"`
	KubernetesVersion string `yaml:"kubernetesVersion"`
	// The ControlPlaneEndpoint, that is the address of the external loadbalancer
	// if defined or the bootstrap node
	ControlPlaneEndpoint string `yaml:"controlPlaneEndpoint"`
	// 这个地方展示 k8s所需核心镜像 key 为镜像的名称 value为镜像地址
	KubernetesImages []Images `yaml:"kubernetesImages"`
	// 第三方镜像
	AddonImages []Images `yaml:"addonImages"`

	// 二进制包
	BinaryURI []Binary `yaml:"binaryUri"`

	RedHatPackageManagerURI []RedHatPackageManager `yaml:"redHatPackageManagerUri"`
	DebianPackageManagerURI []RedHatPackageManager `yaml:"debianPackageManagerUri"`
}

type Images struct {
	Name       string `json:"name" yaml:"name"`
	Repository string `json:"repository" yaml:"repository"`
	Host       string `json:"host" yaml:"host"`
}

type Binary struct {
	Name string `json:"name" yaml:"name"`
	URI  string `json:"uri" yaml:"uri"`
}

type RedHatPackageManager struct {
	Name string `json:"name" yaml:"name"`
	URI  string `json:"uri" yaml:"uri"`
}
