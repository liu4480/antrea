// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/metrics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAntreaNetworkPolicyMetricses implements AntreaNetworkPolicyMetricsInterface
type FakeAntreaNetworkPolicyMetricses struct {
	Fake *FakeMetricsV1alpha1
	ns   string
}

var antreanetworkpolicymetricsesResource = schema.GroupVersionResource{Group: "metrics.antrea.tanzu.vmware.com", Version: "v1alpha1", Resource: "antreanetworkpolicymetrics"}

var antreanetworkpolicymetricsesKind = schema.GroupVersionKind{Group: "metrics.antrea.tanzu.vmware.com", Version: "v1alpha1", Kind: "AntreaNetworkPolicyMetrics"}

// Get takes name of the antreaNetworkPolicyMetrics, and returns the corresponding antreaNetworkPolicyMetrics object, and an error if there is any.
func (c *FakeAntreaNetworkPolicyMetricses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AntreaNetworkPolicyMetrics, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(antreanetworkpolicymetricsesResource, c.ns, name), &v1alpha1.AntreaNetworkPolicyMetrics{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AntreaNetworkPolicyMetrics), err
}

// List takes label and field selectors, and returns the list of AntreaNetworkPolicyMetricses that match those selectors.
func (c *FakeAntreaNetworkPolicyMetricses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AntreaNetworkPolicyMetricsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(antreanetworkpolicymetricsesResource, antreanetworkpolicymetricsesKind, c.ns, opts), &v1alpha1.AntreaNetworkPolicyMetricsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AntreaNetworkPolicyMetricsList{ListMeta: obj.(*v1alpha1.AntreaNetworkPolicyMetricsList).ListMeta}
	for _, item := range obj.(*v1alpha1.AntreaNetworkPolicyMetricsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested antreaNetworkPolicyMetricses.
func (c *FakeAntreaNetworkPolicyMetricses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(antreanetworkpolicymetricsesResource, c.ns, opts))

}