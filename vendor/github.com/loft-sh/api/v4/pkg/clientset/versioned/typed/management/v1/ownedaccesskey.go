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

// OwnedAccessKeysGetter has a method to return a OwnedAccessKeyInterface.
// A group's client should implement this interface.
type OwnedAccessKeysGetter interface {
	OwnedAccessKeys() OwnedAccessKeyInterface
}

// OwnedAccessKeyInterface has methods to work with OwnedAccessKey resources.
type OwnedAccessKeyInterface interface {
	Create(ctx context.Context, ownedAccessKey *managementv1.OwnedAccessKey, opts metav1.CreateOptions) (*managementv1.OwnedAccessKey, error)
	Update(ctx context.Context, ownedAccessKey *managementv1.OwnedAccessKey, opts metav1.UpdateOptions) (*managementv1.OwnedAccessKey, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ownedAccessKey *managementv1.OwnedAccessKey, opts metav1.UpdateOptions) (*managementv1.OwnedAccessKey, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementv1.OwnedAccessKey, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementv1.OwnedAccessKeyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementv1.OwnedAccessKey, err error)
	OwnedAccessKeyExpansion
}

// ownedAccessKeys implements OwnedAccessKeyInterface
type ownedAccessKeys struct {
	*gentype.ClientWithList[*managementv1.OwnedAccessKey, *managementv1.OwnedAccessKeyList]
}

// newOwnedAccessKeys returns a OwnedAccessKeys
func newOwnedAccessKeys(c *ManagementV1Client) *ownedAccessKeys {
	return &ownedAccessKeys{
		gentype.NewClientWithList[*managementv1.OwnedAccessKey, *managementv1.OwnedAccessKeyList](
			"ownedaccesskeys",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementv1.OwnedAccessKey { return &managementv1.OwnedAccessKey{} },
			func() *managementv1.OwnedAccessKeyList { return &managementv1.OwnedAccessKeyList{} },
		),
	}
}
