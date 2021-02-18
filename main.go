package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	apps_v1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/apps/v1alpha"
	edgenetclientset "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	namespace := flag.String("namespace", "", "the namespace(s) to be observed")
	timeout := flag.Int("timeout", 1, "stop the script after X hours")
	flag.Parse()

	resources := flag.Args()
	clientset, edgenetclientset, err := createClientSets()
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	for _, resource := range resources {
		switch strings.ToLower(resource) {
		case "node", "nodes":
			watchNode, err := clientset.CoreV1().Nodes().Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchNode, "Node")
			} else {
				log.Println(err)
			}
		case "pod", "pods":
			watchPod, err := clientset.CoreV1().Pods(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchPod, "Pod")
			} else {
				log.Println(err)
			}
		case "deployment", "deployments":
			watchDeployment, err := clientset.AppsV1().Deployments(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchDeployment, "Deployment")
			} else {
				log.Println(err)
			}
		case "daemonset", "daemonsets":
			watchDaemonSet, err := clientset.AppsV1().DaemonSets(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchDaemonSet, "DaemonSet")
			} else {
				log.Println(err)
			}
		case "statefulset", "statefulsets":
			watchStatefulSet, err := clientset.AppsV1().StatefulSets(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchStatefulSet, "StatefulSet")
			} else {
				log.Println(err)
			}
		case "job", "jobs":
			watchJob, err := clientset.BatchV1().Jobs(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchJob, "Job")
			} else {
				log.Println(err)
			}
		case "cronjob", "cronjobs":
			watchCronJob, err := clientset.BatchV1beta1().CronJobs(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchCronJob, "CronJob")
			} else {
				log.Println(err)
			}
		case "selectivedeployment", "selectivedeployments":
			watchSelectiveDeployment, err := edgenetclientset.AppsV1alpha().SelectiveDeployments(string(*namespace)).Watch(context.TODO(), metav1.ListOptions{})
			if err == nil {
				go watchResource(watchSelectiveDeployment, "SelectiveDeployment")
			} else {
				log.Println(err)
			}
		}
	}
	//watchSlice, err := t.edgenetClientset.AppsV1alpha().Slices(sliceCopy.GetNamespace()).Watch(context.TODO(), metav1.ListOptions{FieldSelector: fmt.Sprintf("metadata.name==%s", sliceCopy.GetName())})
	timeAfter := time.After(time.Hour * time.Duration(int(*timeout)))
timeoutLoop:
	for {
		select {
		case <-timeAfter:
			break timeoutLoop
		}
	}
}

func watchResource(watch watch.Interface, resource string) {
	for event := range watch.ResultChan() {
		switch strings.ToLower(resource) {
		case "node":
			node, status := event.Object.(*corev1.Node)
			if status {
				log.Printf("Node Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %v\n\n", node.GetName(),
					event.Type, time.Now(), node.Status.Conditions)
			}
		case "pod":
			pod, status := event.Object.(*corev1.Pod)
			if status {
				log.Printf("Pod Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %s\n\n", pod.GetName(),
					event.Type, time.Now(), pod.Status.Phase)
			}
		case "deployment":
			deployment, status := event.Object.(*appsv1.Deployment)
			if status {
				log.Printf("Deployment Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %d/%d\n\n", deployment.GetName(),
					event.Type, time.Now(), deployment.Status.ReadyReplicas, deployment.Status.Replicas)
			}
		case "daemonset":
			daemonset, status := event.Object.(*appsv1.DaemonSet)
			if status {
				log.Printf("DaemonSet Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %d/%d\n\n", daemonset.GetName(),
					event.Type, time.Now(), daemonset.Status.CurrentNumberScheduled, daemonset.Status.DesiredNumberScheduled)
			}
		case "statefulset":
			statefulset, status := event.Object.(*appsv1.StatefulSet)
			if status {
				log.Printf("StatefulSet Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %d/%d\n\n", statefulset.GetName(),
					event.Type, time.Now(), statefulset.Status.ReadyReplicas, statefulset.Status.Replicas)
			}
		case "job":
			job, status := event.Object.(*batchv1.Job)
			if status {
				log.Printf("Job Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %d\n\n", job.GetName(),
					event.Type, time.Now(), job.Status.Active)
			}
		case "cronjob":
			cronjob, status := event.Object.(*batchv1beta.CronJob)
			if status {
				log.Printf("CronJob Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %v\n\n", cronjob.GetName(),
					event.Type, time.Now(), cronjob.Status.Active)
			}
		case "selectivedeployment":
			selectivedeployment, status := event.Object.(*apps_v1alpha.SelectiveDeployment)
			if status {
				log.Printf("SelectiveDeployment Name: %s\nEvent Type: %s\nTimestamp: %s\nStatus: %v\n\n", selectivedeployment.GetName(),
					event.Type, time.Now(), selectivedeployment.Status.Ready)
			}
		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func createClientSets() (*kubernetes.Clientset, *edgenetclientset.Clientset, error) {
	var path string
	if home := homeDir(); home != "" {
		path = filepath.Join(home, ".kube", "config")
	} else {
		path = "./"
	}
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	edgenetclientset, err := edgenetclientset.NewForConfig(config)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	return clientset, edgenetclientset, err
}
