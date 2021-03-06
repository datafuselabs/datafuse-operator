// Copyright 2020-2021 The Datafuse Authors.
//
// SPDX-License-Identifier: Apache-2.0.
package framework

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

func CreateNamespace(kubeClient kubernetes.Interface, name string) (*v1.Namespace, error) {
	namespace, err := kubeClient.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to create namespace with name %v", name))
	}
	return namespace, nil
}

func (ctx *TestCtx) CreateNamespace(t *testing.T, kubeClient kubernetes.Interface) string {
	name := ctx.GetObjID()
	if _, err := CreateNamespace(kubeClient, name); err != nil {
		t.Fatal(err)
	}

	namespaceFinalizerFn := func() error {
		return DeleteNamespace(kubeClient, name)
	}

	ctx.AddFinalizerFn(namespaceFinalizerFn)

	return name
}

func DeleteNamespace(kubeClient kubernetes.Interface, name string) error {
	return kubeClient.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func AddLabelsToNamespace(kubeClient kubernetes.Interface, name string, additionalLabels map[string]string) error {
	ns, err := kubeClient.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if ns.Labels == nil {
		ns.Labels = map[string]string{}
	}

	for k, v := range additionalLabels {
		ns.Labels[k] = v
	}

	_, err = kubeClient.CoreV1().Namespaces().Update(context.TODO(), ns, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func RemoveLabelsFromNamespace(kubeClient kubernetes.Interface, name string, labels ...string) error {
	ns, err := kubeClient.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if len(ns.Labels) == 0 {
		return nil
	}

	type patch struct {
		Op   string `json:"op"`
		Path string `json:"path"`
	}

	var patches []patch
	for _, l := range labels {
		patches = append(patches, patch{Op: "remove", Path: "/metadata/labels/" + l})
	}
	b, err := json.Marshal(patches)
	if err != nil {
		return err
	}

	_, err = kubeClient.CoreV1().Namespaces().Patch(context.TODO(), name, types.JSONPatchType, b, metav1.PatchOptions{})
	if err != nil {
		return err
	}

	return nil
}
