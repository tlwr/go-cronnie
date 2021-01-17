package cronnie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestCronnie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cronnie Suite")
}

var _ = Describe("Cronnie", func() {
	var (
		path string
	)

	BeforeSuite(func() {
		var err error
		path, err = Build("github.com/tlwr/go-cronnie/integration")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	Context("when unsignalled", func() {
		It("generates work after the duration", func() {
			start := time.Now()

			command := exec.Command(path)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			Eventually(session).Should(Exit(0))
			end := time.Now()

			Expect(session).To(Say("done"))
			Expect(end).To(BeTemporally("~", start.Add(250*time.Millisecond), 250*time.Millisecond))
		})
	})

	Context("when using SIGUSR1", func() {
		It("generates work when signalled", func() {
			start := time.Now()

			command := exec.Command(path)
			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(25 * time.Millisecond)
			session.Signal(syscall.SIGUSR1)

			Eventually(session).Should(Exit(0))
			end := time.Now()

			Expect(session).To(Say("done"))
			Expect(end).To(BeTemporally("~", start.Add(25*time.Millisecond), 25*time.Millisecond))
		})
	})
})
