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

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/controlplane/v1alpha1"
	scheme "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// NodeStatsSummariesGetter has a method to return a NodeStatsSummaryInterface.
// A group's client should implement this interface.
type NodeStatsSummariesGetter interface {
	NodeStatsSummaries() NodeStatsSummaryInterface
}

// NodeStatsSummaryInterface has methods to work with NodeStatsSummary resources.
type NodeStatsSummaryInterface interface {
	Create(ctx context.Context, nodeStatsSummary *v1alpha1.NodeStatsSummary, opts v1.CreateOptions) (*v1alpha1.NodeStatsSummary, error)
	NodeStatsSummaryExpansion
}

// nodeStatsSummaries implements NodeStatsSummaryInterface
type nodeStatsSummaries struct {
	client rest.Interface
}

// newNodeStatsSummaries returns a NodeStatsSummaries
func newNodeStatsSummaries(c *ControlplaneV1alpha1Client) *nodeStatsSummaries {
	return &nodeStatsSummaries{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a nodeStatsSummary and creates it.  Returns the server's representation of the nodeStatsSummary, and an error, if there is any.
func (c *nodeStatsSummaries) Create(ctx context.Context, nodeStatsSummary *v1alpha1.NodeStatsSummary, opts v1.CreateOptions) (result *v1alpha1.NodeStatsSummary, err error) {
	result = &v1alpha1.NodeStatsSummary{}
	err = c.client.Post().
		Resource("nodestatssummaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeStatsSummary).
		Do(ctx).
		Into(result)
	return
}