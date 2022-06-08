package controllers

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cachev1alpha1 "github.com/subns/test-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"time"
)

var _ = Describe("Reconciliation", func() {
	var o cachev1alpha1.Memcached
	BeforeEach(func() {
		o = cachev1alpha1.Memcached{}
		o.Spec.Replicas = 4
		o.Name = "foo"
		o.Namespace = "default"
		k8sClient.Create(context.Background(), &o, &client.CreateOptions{})
	})

	Context("Create", func() {
		It("should update the status", func() {
			Eventually(func() int {
				err := k8sClient.Get(context.Background(),
					types.NamespacedName{Namespace: o.Namespace, Name: o.Name}, &o)
				if err != nil {
					logf.Log.Error(err, "no....")
					return -1
				}
				logf.Log.Info("status is", "object", o)
				return o.Status.FailedReplicas
			}).
				WithTimeout(time.Second * 5).
				WithPolling(time.Second * 2).
				Should(Equal(5))
		})

	})

})
