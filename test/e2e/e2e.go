package test

import (
	"testing"

	"github.com/onsi/ginkgo"
	gomega "github.com/onsi/gomega"

	"k8s.io/component-base/logs"
	ginkgowrapper "k8s.io/kubernetes/test/e2e/framework/ginkgowrapper"

)

func RunE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	gomega.RegisterFailHandler(ginkgowrapper.Fail)
	ginkgo.RunSpecs(t, "Multi-Tenancy Benchmarks")
}