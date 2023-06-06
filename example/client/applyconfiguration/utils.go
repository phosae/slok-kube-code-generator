// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/phosae/kube-code-generator/example/apis/comic/v1"
	comicv1 "github.com/phosae/kube-code-generator/example/client/applyconfiguration/comic/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=comic.kube-code-generator.slok.dev, Version=v1
	case v1.SchemeGroupVersion.WithKind("Hero"):
		return &comicv1.HeroApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HeroSpec"):
		return &comicv1.HeroSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HeroStatus"):
		return &comicv1.HeroStatusApplyConfiguration{}

	}
	return nil
}
