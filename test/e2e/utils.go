package test

import (
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/test/e2e/framework"
)

func getContextFromKubeconfig(kubeconfigpath string) string {
	apiConfig := clientcmd.GetConfigFromFileOrDie(kubeconfigpath)

	if apiConfig.CurrentContext == "" {
		framework.Failf("current-context is not set in %s", kubeconfigpath)
	}

	return apiConfig.CurrentContext
}

func newKubeClientWithKubeconfig(kubeconfigpath string) *kubernetes.Clientset {
	clientConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	framework.ExpectNoError(err)

	kclient, err := kubernetes.NewForConfig(clientConfig)
	framework.ExpectNoError(err)

	return kclient
}
