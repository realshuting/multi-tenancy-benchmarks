package test

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/onsi/ginkgo"
	"k8s.io/kubernetes/test/e2e/framework"
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

	ginkgo.BeforeEach(func() {
		ginkgo.By("get cluster wide api-resources")

		config, err = readConfig(configPath)
		framework.ExpectNoError(err)

		os.Setenv("KUBECONFIG", config.Adminkubeconfig)
		nsdFlag := fmt.Sprintf("--namespaced=false")
		outputFlag := fmt.Sprintf("-o=name")

		resourceList, err = framework.RunKubectl("api-resources", nsdFlag, outputFlag)
		framework.ExpectNoError(err)
	})

	framework.KubeDescribe("tenant cannot accesss cluster wide resources", func() {
		var user string

		ginkgo.BeforeEach(func() {
			tenantkubeconfig := config.getValidTenant()
			os.Setenv("KUBECONFIG", tenantkubeconfig.Kubeconfig)
			user = getContextFromKubeconfig(tenantkubeconfig.Kubeconfig)
		})

		ginkgo.It("get cluster wide resources", func() {
			ginkgo.By(fmt.Sprintf("tenant %s cannot get cluster wide resources", user))
			resources := strings.Fields(resourceList)
			for _, resource := range resources {
				_, errNew := framework.LookForString(expectedVal, time.Minute, func() string {
					_, err := framework.RunKubectl("get", resource)
					return err.Error()
				})

				framework.ExpectNoError(errNew)
			}
		})

		ginkgo.It("edit cluster wide resources", func() {
			ginkgo.By(fmt.Sprintf("tenant %s cannot edit cluster wide resources", user))
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
