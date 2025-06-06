// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	managementv1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// EventsGetter has a method to return a EventInterface.
// A group's client should implement this interface.
type EventsGetter interface {
	Events() EventInterface
}

// EventInterface has methods to work with Event resources.
type EventInterface interface {
	Create(ctx context.Context, event *managementv1.Event, opts metav1.CreateOptions) (*managementv1.Event, error)
	Update(ctx context.Context, event *managementv1.Event, opts metav1.UpdateOptions) (*managementv1.Event, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, event *managementv1.Event, opts metav1.UpdateOptions) (*managementv1.Event, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementv1.Event, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementv1.EventList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementv1.Event, err error)
	EventExpansion
}

// events implements EventInterface
type events struct {
	*gentype.ClientWithList[*managementv1.Event, *managementv1.EventList]
}

// newEvents returns a Events
func newEvents(c *ManagementV1Client) *events {
	return &events{
		gentype.NewClientWithList[*managementv1.Event, *managementv1.EventList](
			"events",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementv1.Event { return &managementv1.Event{} },
			func() *managementv1.EventList { return &managementv1.EventList{} },
		),
	}
}
