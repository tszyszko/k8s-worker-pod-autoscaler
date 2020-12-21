/*
Copyright 2019 Practo Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	workerpodautoscalerv1 "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscaler/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkerPodAutoScalers implements WorkerPodAutoScalerInterface
type FakeWorkerPodAutoScalers struct {
	Fake *FakeK8sV1
	ns   string
}

var workerpodautoscalersResource = schema.GroupVersionResource{Group: "k8s.practo.dev", Version: "v1", Resource: "workerpodautoscalers"}

var workerpodautoscalersKind = schema.GroupVersionKind{Group: "k8s.practo.dev", Version: "v1", Kind: "WorkerPodAutoScaler"}

// Get takes name of the workerPodAutoScaler, and returns the corresponding workerPodAutoScaler object, and an error if there is any.
func (c *FakeWorkerPodAutoScalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *workerpodautoscalerv1.WorkerPodAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workerpodautoscalersResource, c.ns, name), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalerv1.WorkerPodAutoScaler), err
}

// List takes label and field selectors, and returns the list of WorkerPodAutoScalers that match those selectors.
func (c *FakeWorkerPodAutoScalers) List(ctx context.Context, opts v1.ListOptions) (result *workerpodautoscalerv1.WorkerPodAutoScalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workerpodautoscalersResource, workerpodautoscalersKind, c.ns, opts), &workerpodautoscalerv1.WorkerPodAutoScalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &workerpodautoscalerv1.WorkerPodAutoScalerList{ListMeta: obj.(*workerpodautoscalerv1.WorkerPodAutoScalerList).ListMeta}
	for _, item := range obj.(*workerpodautoscalerv1.WorkerPodAutoScalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workerPodAutoScalers.
func (c *FakeWorkerPodAutoScalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workerpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a workerPodAutoScaler and creates it.  Returns the server's representation of the workerPodAutoScaler, and an error, if there is any.
func (c *FakeWorkerPodAutoScalers) Create(ctx context.Context, workerPodAutoScaler *workerpodautoscalerv1.WorkerPodAutoScaler, opts v1.CreateOptions) (result *workerpodautoscalerv1.WorkerPodAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workerpodautoscalersResource, c.ns, workerPodAutoScaler), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalerv1.WorkerPodAutoScaler), err
}

// Update takes the representation of a workerPodAutoScaler and updates it. Returns the server's representation of the workerPodAutoScaler, and an error, if there is any.
func (c *FakeWorkerPodAutoScalers) Update(ctx context.Context, workerPodAutoScaler *workerpodautoscalerv1.WorkerPodAutoScaler, opts v1.UpdateOptions) (result *workerpodautoscalerv1.WorkerPodAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workerpodautoscalersResource, c.ns, workerPodAutoScaler), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalerv1.WorkerPodAutoScaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkerPodAutoScalers) UpdateStatus(ctx context.Context, workerPodAutoScaler *workerpodautoscalerv1.WorkerPodAutoScaler, opts v1.UpdateOptions) (*workerpodautoscalerv1.WorkerPodAutoScaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workerpodautoscalersResource, "status", c.ns, workerPodAutoScaler), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalerv1.WorkerPodAutoScaler), err
}

// Delete takes name of the workerPodAutoScaler and deletes it. Returns an error if one occurs.
func (c *FakeWorkerPodAutoScalers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workerpodautoscalersResource, c.ns, name), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkerPodAutoScalers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workerpodautoscalersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &workerpodautoscalerv1.WorkerPodAutoScalerList{})
	return err
}

// Patch applies the patch and returns the patched workerPodAutoScaler.
func (c *FakeWorkerPodAutoScalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *workerpodautoscalerv1.WorkerPodAutoScaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workerpodautoscalersResource, c.ns, name, pt, data, subresources...), &workerpodautoscalerv1.WorkerPodAutoScaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalerv1.WorkerPodAutoScaler), err
}
