package test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	yaml "k8s.io/apimachinery/pkg/util/yaml"
)

type benchmarkConfig struct {
	Adminkubeconfig string `json:"adminKubeconfig"`
	Label           string `json:"label,omitempty"`
	TenantA         tenant `json:"tenantA,omitempty"`
	TenantB         tenant `json:"tenantB,omitempty"`
}

type tenant struct {
	Kubeconfig string `json:"kubeconfig"`
	Namespace  string `json:"namespace"`
}

func (c *benchmarkConfig) getValidTenant() tenant {
	if !reflect.DeepEqual(c.TenantA, tenant{}) {
		return c.TenantA
	}

	return c.TenantB
}

func readConfig(path string) (*benchmarkConfig, error) {
	var config *benchmarkConfig

	file, err := loadFile(path)
	if err != nil {
		return nil, err
	}

	data, err := yaml.ToJSON(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func loadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(path)
}
