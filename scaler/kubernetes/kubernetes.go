package kubernetes

import (
	"fmt"
	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/rest"
)

// Represents a connection to a kubernetes cluster.
type KubernetesConnection struct {
	InCluster bool         `yaml:"in_cluster"`
	Config    *rest.Config `yaml:"connection_info"`

	client *kubernetes.Clientset
}

// Gets the client from this connection,
// only opens a new one if the existing doesnt exist
func (k *KubernetesConnection) GetClient() (*kubernetes.Clientset, error) {
	if k.client != nil {
		return k.client, nil
	}

	fmt.Println(k.InCluster)

	// We're in the cluster, so we'll just use the inbuilt tokens to connect.
	if k.InCluster {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		k.Config = config
	}

	clientset, err := kubernetes.NewForConfig(k.Config)
	if err != nil {
		return nil, err
	}

	k.client = clientset

	return clientset, nil
}
