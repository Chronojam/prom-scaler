package kubernetes

import (
	"github.com/chronojam/prometheus-scaler/config"
	"github.com/chronojam/prometheus-scaler/scaler"

	"gopkg.in/yaml.v2"
	//extensionsv1 "k8s.io/client-go/1.5/pkg/apis/extensions/v1beta1"
)

func init() {
	scaler.Register("kubernetes-deployment", driver)
}

func driver(cfg config.RegistrableScalarConfig) (scaler.ScalableResource, error) {
	var ds DeploymentScaler

	bytes, err := yaml.Marshal(cfg.Options)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &ds)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

type DeploymentScaler struct {
	Namespace      string
	DeploymentName string
	ScaleNames     []string
	KubernetesConnection
}

func (ds DeploymentScaler) Names() []string {
	return ds.ScaleNames
}

func (ds DeploymentScaler) Scale(count int) error {
	client, err := ds.GetClient()
	if err != nil {
		return err
	}

	d, err := client.Extensions().Deployments(ds.Namespace).Get(ds.DeploymentName)
	if err != nil {
		return err
	}

	newReplicas := *d.Spec.Replicas + int32(count)

	d.Spec.Replicas = &newReplicas
	d, err = client.Extensions().Deployments(ds.Namespace).Update(d)
	if err != nil {
		return err
	}

	return nil
}
