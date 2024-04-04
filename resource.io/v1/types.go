package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:storageversions
// +kubebuilder:resource:scope=Cluster,shortName=pr
// +kubebuilder:subresource:status
type Priority struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	//
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Sub SubPriority `json:"sub,omitempty"`
}

type SubPriority struct {
	// +optional
	Num int `json:"num,omitempty"`
}

func (*Priority) DeepCopy() *Priority {
	return nil
}

func (*Priority) DeepCopyObject() runtime.Object {
	return nil
}
