package k8sdemo

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"
)

func GetClients(c *rest.Config) (*kubernetes.Clientset, error) {
	configCopy := metadata.ConfigFor(c)
	clientset, err := kubernetes.NewForConfig(configCopy)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func listPods(ctx context.Context, client rest.Interface, namespace string, options metav1.ListOptions) (*v1.PodList, error) {
	podList := &v1.PodList{}
	return podList, nil
}
