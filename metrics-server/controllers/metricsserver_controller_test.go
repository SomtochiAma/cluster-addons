/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"testing"

	api "sigs.k8s.io/cluster-addons/metrics-server/api/v1alpha1"

	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/test/golden"
)

func TestMetricsServer(t *testing.T) {
	v := golden.NewValidator(t, api.SchemeBuilder)
	dr := &MetricsServerReconciler{
		Client: v.Manager().GetClient(),
	}
	err := dr.setupReconciler(v.Manager())
	if err != nil {
		t.Fatalf("creating reconciler: %v", err)
	}

	v.Validate(dr.Reconciler)
}