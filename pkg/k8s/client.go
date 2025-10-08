package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	clientset *kubernetes.Clientset
}

func NewClient(kubeconfigPath string) (*Client, error) {
	var (
		config *rest.Config
		err    error
	)

	config, err = rest.InClusterConfig()
	if err != nil {
		if kubeconfigPath != "" {
			config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)

			if err != nil {
				return nil, fmt.Errorf("failed to load kubeconfig from %s: %w", kubeconfigPath, err)
			}
		} else {
			return nil, fmt.Errorf("kubeconfig was not found in the cluster: %w", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s clientset: %w", err)
	}

	return &Client{clientset: clientset}, nil
}
