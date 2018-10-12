package kube

import (
	"fmt"
	"log"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Client struct
type Client struct {
	k8s *kubernetes.Clientset
}

// New create new kube client
func New() *Client {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	// creates the clientset
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		k8s: client,
	}
}

// DeleteEvictPods delete evicted pods from all namespace
func (client *Client) DeleteEvictPods() error {
	podList, err := client.k8s.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, pod := range podList.Items {
		if strings.Contains(pod.Status.Reason, "Evicted") {
			fmt.Println(pod.Namespace + " - " + pod.Name + " - " + pod.Status.Message)
			err := client.k8s.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
