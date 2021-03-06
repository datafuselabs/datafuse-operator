// Copyright 2020-2021 The Datafuse Authors.
//
// SPDX-License-Identifier: Apache-2.0.
package framework

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/pkg/errors"
)

func PathToOSFile(relativPath string) (*os.File, error) {
	path, err := filepath.Abs(relativPath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed generate absolute file path of %s", relativPath))
	}

	manifest, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to open file %s", path))
	}

	return manifest, nil
}

// PodRunningAndReady returns whether a pod is running and each container has
// passed it's ready state.
func PodRunningAndReady(pod v1.Pod) (bool, error) {
	switch pod.Status.Phase {
	case v1.PodFailed, v1.PodSucceeded:
		return false, fmt.Errorf("pod completed")
	case v1.PodRunning:
		for _, cond := range pod.Status.Conditions {
			if cond.Type != v1.PodReady {
				continue
			}
			return cond.Status == v1.ConditionTrue, nil
		}
		return false, fmt.Errorf("pod ready condition not found")
	}
	return false, nil
}

// WaitForPodsReady waits for a selection of Pods to be running and each
// container to pass its readiness check.
func WaitForPodsReady(kubeClient kubernetes.Interface, namespace string, timeout time.Duration, expectedReplicas int, opts metav1.ListOptions) error {
	return wait.Poll(time.Second, timeout, func() (bool, error) {
		pl, err := kubeClient.CoreV1().Pods(namespace).List(context.TODO(), opts)
		if err != nil {
			return false, err
		}

		runningAndReady := 0
		for _, p := range pl.Items {
			isRunningAndReady, err := PodRunningAndReady(p)
			if err != nil {
				return false, err
			}

			if isRunningAndReady {
				runningAndReady++
			}
		}

		if runningAndReady == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}

func WaitForPodsRunImage(kubeClient kubernetes.Interface, namespace string, expectedReplicas int, image string, opts metav1.ListOptions) error {
	return wait.Poll(time.Second, time.Minute*5, func() (bool, error) {
		pl, err := kubeClient.CoreV1().Pods(namespace).List(context.TODO(), opts)
		if err != nil {
			return false, err
		}

		runningImage := 0
		for _, p := range pl.Items {
			if podRunsImage(p, image) {
				runningImage++
			}
		}

		if runningImage == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}

func WaitForHTTPSuccessStatusCode(timeout time.Duration, url string) error {
	var resp *http.Response
	err := wait.Poll(time.Second, timeout, func() (bool, error) {
		var err error
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return true, nil
		}
		return false, nil
	})

	return errors.Wrap(err, fmt.Sprintf(
		"waiting for %v to return a successful status code timed out. Last response from server was: %v",
		url,
		resp,
	))
}

func podRunsImage(p v1.Pod, image string) bool {
	for _, c := range p.Spec.Containers {
		if image == c.Image {
			return true
		}
	}

	return false
}

func GetLogs(kubeClient kubernetes.Interface, namespace string, podName, containerName string) (string, error) {
	logs, err := kubeClient.CoreV1().RESTClient().Get().
		Resource("pods").
		Namespace(namespace).
		Name(podName).SubResource("log").
		Param("container", containerName).
		Do(context.TODO()).
		Raw()
	if err != nil {
		return "", err
	}
	return string(logs), err
}

func ProxyGetPod(kubeClient kubernetes.Interface, namespace string, podName string, port string, path string) *rest.Request {
	return kubeClient.CoreV1().RESTClient().Get().Prefix("proxy").Namespace(namespace).Resource("pods").Name(podName + ":" + port).Suffix(path)
}
