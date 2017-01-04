package listener_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/bijukunjummen/multiple-interfaces-listener/listener"
	"fmt"
	"net/http"
	"io/ioutil"
)

var _ = Describe("Interface Listener", func() {

	Context("Given a ip and port", func() {
		IP := "127.0.0.1"
		Port := 9765

		It("Should be possible to spin up a listener and respond to requests", func() {
			l, _ := listener.NewPortListener(IP, Port)
			fmt.Println("\nabout to listen")
			l.ListenAndProvideStockResponses()
			httpClient := &http.Client{}

			resp, err := httpClient.Get(fmt.Sprintf("http://%s:%d", IP, Port))

			Ω(err).Should(BeNil())

			bodyb, err:=ioutil.ReadAll(resp.Body)
			Ω(err).Should(BeNil())
			Ω(string(bodyb)).Should(Equal("Hello from : 127.0.0.1:9765"))

		})

	})
})
