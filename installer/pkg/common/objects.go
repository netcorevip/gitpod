// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package common

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

func DefaultServiceAccount(component string) RenderFunc {
	return func(cfg *RenderContext) ([]runtime.Object, error) {
		return []runtime.Object{
			&corev1.ServiceAccount{
				TypeMeta: TypeMetaServiceAccount,
				ObjectMeta: metav1.ObjectMeta{
					Name:      component,
					Namespace: cfg.Namespace,
					Labels:    DefaultLabels(component),
				},
				AutomountServiceAccountToken: pointer.Bool(true),
			},
		}, nil
	}
}

type ServicePort struct {
	ContainerPort int32
	ServicePort   int32
}

func GenerateService(component string, ports map[string]ServicePort, assignClusterIp bool) RenderFunc {
	return func(cfg *RenderContext) ([]runtime.Object, error) {
		labels := DefaultLabels(component)

		var servicePorts []corev1.ServicePort
		for name, port := range ports {
			servicePorts = append(servicePorts, corev1.ServicePort{
				Protocol:   *TCPProtocol,
				Name:       name,
				Port:       port.ServicePort,
				TargetPort: intstr.IntOrString{IntVal: port.ContainerPort},
			})
		}

		spec := corev1.ServiceSpec{
			Ports:    servicePorts,
			Selector: labels,
			Type:     corev1.ServiceTypeClusterIP,
		}

		if assignClusterIp == false {
			spec.ClusterIP = "None"
		}

		return []runtime.Object{&corev1.Service{
			TypeMeta: TypeMetaService,
			ObjectMeta: metav1.ObjectMeta{
				Name:      component,
				Namespace: cfg.Namespace,
				Labels:    labels,
			},
			Spec: spec,
		}}, nil
	}
}

// GlobalObjects is any objects which are outside the scope of components, but
// required for the application to function. Typically, these will be ClusterRole,
// ClusterRoleBindings and similar cluster-level objects
func GlobalObjects(ctx *RenderContext) ([]runtime.Object, error) {
	return []runtime.Object{
		&rbacv1.ClusterRole{
			TypeMeta: TypeMetaClusterRole,
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("%s-kube-rbac-proxy", ctx.Namespace),
			},
			Rules: []rbacv1.PolicyRule{{
				APIGroups: []string{"authentication.k8s.io"},
				Resources: []string{"tokenreviews"},
				Verbs:     []string{"create"},
			}, {
				APIGroups: []string{"authorization.k8s.io"},
				Resources: []string{"subjectaccessreviews"},
				Verbs:     []string{"create"},
			}},
		},
		&corev1.ResourceQuota{
			TypeMeta: TypeMetaResourceQuota,
			ObjectMeta: metav1.ObjectMeta{
				Name:      "gitpod-resource-quota",
				Namespace: ctx.Namespace,
			},
			Spec: corev1.ResourceQuotaSpec{
				Hard: map[corev1.ResourceName]resource.Quantity{
					"pods": resource.MustParse("10k"),
				},
				ScopeSelector: &corev1.ScopeSelector{
					MatchExpressions: []corev1.ScopedResourceSelectorRequirement{{
						Operator:  corev1.ScopeSelectorOpIn,
						ScopeName: corev1.ResourceQuotaScopePriorityClass,
						Values:    []string{SystemNodeCritical},
					}},
				},
			},
		},
	}, nil
}
