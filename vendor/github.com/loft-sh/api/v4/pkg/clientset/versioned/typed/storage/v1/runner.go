// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// RunnersGetter has a method to return a RunnerInterface.
// A group's client should implement this interface.
type RunnersGetter interface {
	Runners() RunnerInterface
}

// RunnerInterface has methods to work with Runner resources.
type RunnerInterface interface {
	Create(ctx context.Context, runner *storagev1.Runner, opts metav1.CreateOptions) (*storagev1.Runner, error)
	Update(ctx context.Context, runner *storagev1.Runner, opts metav1.UpdateOptions) (*storagev1.Runner, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, runner *storagev1.Runner, opts metav1.UpdateOptions) (*storagev1.Runner, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.Runner, error)
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1.RunnerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *storagev1.Runner, err error)
	RunnerExpansion
}

// runners implements RunnerInterface
type runners struct {
	*gentype.ClientWithList[*storagev1.Runner, *storagev1.RunnerList]
}

// newRunners returns a Runners
func newRunners(c *StorageV1Client) *runners {
	return &runners{
		gentype.NewClientWithList[*storagev1.Runner, *storagev1.RunnerList](
			"runners",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *storagev1.Runner { return &storagev1.Runner{} },
			func() *storagev1.RunnerList { return &storagev1.RunnerList{} },
		),
	}
}
