package integration_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/baggageclaim"
)

var _ = Describe("Copy On Write Strategy", func() {
	var (
		runner *BaggageClaimRunner
		client baggageclaim.Client
	)

	BeforeEach(func() {
		runner = NewRunner(baggageClaimPath)
		runner.Start()
		client = runner.Client()
	})

	AfterEach(func() {
		runner.Stop()
		runner.Cleanup()
	})

	Describe("API", func() {
		writeData := func(volumePath string) string {
			filename := randSeq(10)
			newFilePath := filepath.Join(volumePath, filename)

			err := ioutil.WriteFile(newFilePath, []byte(filename), 0755)
			Expect(err).NotTo(HaveOccurred())

			return filename
		}

		dataExistsInVolume := func(filename, volumePath string) bool {
			_, err := os.Stat(filepath.Join(volumePath, filename))
			return err == nil
		}

		Describe("POST /volumes with strategy: cow", func() {
			It("creates a copy of the volume", func() {
				parentVolume, err := client.CreateVolume(logger, "some-handle", baggageclaim.VolumeSpec{})
				Expect(err).NotTo(HaveOccurred())

				dataInParent := writeData(parentVolume.Path())
				Expect(dataExistsInVolume(dataInParent, parentVolume.Path())).To(BeTrue())

				childVolume, err := client.CreateVolume(logger, "another-handle", baggageclaim.VolumeSpec{
					Strategy: baggageclaim.COWStrategy{
						Parent: parentVolume,
					},
				})
				Expect(err).NotTo(HaveOccurred())

				Expect(dataExistsInVolume(dataInParent, childVolume.Path())).To(BeTrue())

				newDataInParent := writeData(parentVolume.Path())
				Expect(dataExistsInVolume(newDataInParent, parentVolume.Path())).To(BeTrue())
				Expect(dataExistsInVolume(newDataInParent, childVolume.Path())).To(BeFalse())

				dataInChild := writeData(childVolume.Path())
				Expect(dataExistsInVolume(dataInChild, childVolume.Path())).To(BeTrue())
				Expect(dataExistsInVolume(dataInChild, parentVolume.Path())).To(BeFalse())
			})
		})
	})
})

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
