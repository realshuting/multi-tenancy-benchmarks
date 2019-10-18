package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/onsi/ginkgo"
	gomega "github.com/onsi/gomega"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/test/e2e/framework"
	ginkgowrapper "k8s.io/kubernetes/test/e2e/framework/ginkgowrapper"
)

const (
	expectedVal = "Error from server (Forbidden)"
	configPath  = "manifest/config.yaml"
)

var _ = framework.KubeDescribe("test tenant permission", func() {
	var config *benchmarkConfig
	var resourceList string
	var err error
	var dryrun = "--dry-run=true"
	var all = "--all=true"

	framework.KubeDescribe("test tenant operation on none namespaced resource", func() {
		ginkgo.BeforeEach(func() {
			ginkgo.By("get none namespaced api-resources")

			config, err = readConfig(configPath)
			framework.ExpectNoError(err)

			os.Setenv("KUBECONFIG", config.Adminkubeconfig)
			nsdFlag := fmt.Sprintf("--namespaced=false")
			outputFlag := fmt.Sprintf("-o=name")

			resourceList, err = framework.RunKubectl("api-resources", nsdFlag, outputFlag)
			framework.ExpectNoError(err)
		})

		framework.KubeDescribe("tenant admin", func() {
			// mkpath := func(file string) string {
			// 	return filepath.Join(manifestPath, file)
			// }

			ginkgo.BeforeEach(func() {
				tenantkubeconfig := config.getValidTenant()
				os.Setenv("KUBECONFIG", tenantkubeconfig.Kubeconfig)
			})

			ginkgo.It("cannot get none namespaced resources", func() {
				resources := strings.Fields(resourceList)
				for _, resource := range resources {
					_, errNew := framework.LookForString(expectedVal, time.Minute, func() string {
						_, err := framework.RunKubectl("get", resource)
						return err.Error()
					})

					framework.ExpectNoError(errNew)
				}
			})

			ginkgo.It("cannot edit none namespaced resources", func() {
				resources := strings.Fields(resourceList)
				annotation := "test=multi-tenancy"
				for _, resource := range resources {
					_, errNew := framework.LookForString(expectedVal, time.Minute, func() string {
						_, err := framework.RunKubectl("annotate", resource, annotation, dryrun, all)
						return err.Error()
					})

					framework.ExpectNoError(errNew)
				}
			})
		})
	})
})

func RunE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	ginkgo.RunSpecs(t, "Multi-Tenancy Benchmarks")
}
