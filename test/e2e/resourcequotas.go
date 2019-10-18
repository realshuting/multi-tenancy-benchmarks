package test

import (
	"strings"

	"github.com/onsi/ginkgo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/test/e2e/framework"
)

var _ = framework.KubeDescribe("A tenant setup must comply with cluster wide resource configurations", func() {
	var config *benchmarkConfig
	var err error

	ginkgo.BeforeEach(func() {
		config, err = readConfig(configPath)
		framework.ExpectNoError(err)
	})

	ginkgo.It("tenant must have resourcequotas configured same with cluster wide resource", func() {
		resourceNameList := getResourceNameList(config.Adminkubeconfig)
		tenanta := config.getValidTenant()
		tenantResourcequotas := getTenantResoureQuotas(tenanta)
		expectedVal := strings.Join(tenantResourcequotas, " ")
		for _, r := range resourceNameList {
			if !strings.Contains(expectedVal, r) {
				framework.Failf("%s must be configured in tenant resourcequotas", r)
			}
		}
	})
})

func getTenantResoureQuotas(t tenant) []string {
	var tmpList string
	var tenantResourceQuotas []string

	kclient := newKubeClientWithKubeconfig(t.Kubeconfig)
	resourcequotaList, err := kclient.CoreV1().ResourceQuotas(t.Namespace).List(metav1.ListOptions{})
	framework.ExpectNoError(err)

	for _, resourcequota := range resourcequotaList.Items {
		for name, _ := range resourcequota.Spec.Hard {
			if strings.Contains(tmpList, name.String()) {
				continue
			}

			tenantResourceQuotas = append(tenantResourceQuotas, name.String())
			tmpList = tmpList + name.String()
		}
	}

	return tenantResourceQuotas
}

func getResourceNameList(kubeconfigpath string) []string {
	kclient := newKubeClientWithKubeconfig(kubeconfigpath)
	nodes, err := kclient.CoreV1().Nodes().List(metav1.ListOptions{})
	framework.ExpectNoError(err)

	return getResourcequotaFromNodes(*nodes)
}

func getResourcequotaFromNodes(nodeList corev1.NodeList) []string {
	var resourceNameList []string
	var tmpList string
	for _, node := range nodeList.Items {
		for resourceName, _ := range node.Status.Capacity {
			if strings.Contains(tmpList, resourceName.String()) {
				continue
			}

			resourceNameList = append(resourceNameList, resourceName.String())
			tmpList = tmpList + resourceName.String()
		}
	}
	return resourceNameList
}

func newKubeClientWithKubeconfig(kubeconfigpath string) *kubernetes.Clientset {
	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	framework.ExpectNoError(err)

	kclient, err := kubernetes.NewForConfig(clientConfig)
	framework.ExpectNoError(err)

	return kclient
}
