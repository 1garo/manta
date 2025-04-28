package manta_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestManta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manta Suite")
}
