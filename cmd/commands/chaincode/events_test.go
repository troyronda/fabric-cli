/*
Copyright State Street Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	"github.com/hyperledger/fabric-cli/cmd/commands/chaincode"
	"github.com/hyperledger/fabric-cli/pkg/environment"
	"github.com/hyperledger/fabric-cli/pkg/fabric/mocks"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

var _ = Describe("ChaincodeEventsCommand", func() {
	var (
		cmd      *cobra.Command
		settings *environment.Settings
		out      *bytes.Buffer

		args []string
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)

		settings = &environment.Settings{
			Home: environment.Home(os.TempDir()),
			Streams: environment.Streams{
				Out: out,
			},
		}

		args = os.Args
	})

	JustBeforeEach(func() {
		cmd = chaincode.NewChaincodeEventsCommand(settings)
	})

	AfterEach(func() {
		os.Args = args
	})

	It("should create a chaincode events command", func() {
		Expect(cmd.Name()).To(Equal("events"))
		Expect(cmd.HasSubCommands()).To(BeFalse())
	})

	It("should provide a help prompt", func() {
		os.Args = append(os.Args, "--help")

		Expect(cmd.Execute()).Should(Succeed())
		Expect(fmt.Sprint(out)).To(ContainSubstring("events <chaincode-name>"))
	})
})

var _ = Describe("ChaincodeEventsImplementation", func() {
	var (
		impl     *chaincode.EventsCommand
		err      error
		out      *bytes.Buffer
		settings *environment.Settings
		factory  *mocks.Factory
		client   *mocks.Channel
	)

	BeforeEach(func() {
		out = new(bytes.Buffer)

		settings = &environment.Settings{
			Home: environment.Home(os.TempDir()),
			Streams: environment.Streams{
				Out: out,
			},
		}

		factory = &mocks.Factory{}
		client = &mocks.Channel{}

		impl = &chaincode.EventsCommand{}
		impl.Settings = settings
		impl.Factory = factory
	})

	It("should not be nil", func() {
		Expect(impl).ShouldNot(BeNil())
	})

	Describe("Validate", func() {
		JustBeforeEach(func() {
			err = impl.Validate()
		})

		It("should fail when name is not set", func() {
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("chaincode name not specified"))
		})

		Context("when chaincode name is set", func() {
			BeforeEach(func() {
				impl.ChaincodeName = "mycc"
			})

			It("should succeed with chaincode name is set", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Run", func() {
		var (
			eventch chan *fab.CCEvent
		)

		BeforeEach(func() {
			impl.ChaincodeName = "mycc"

			impl.Channel = client

			eventch = make(chan *fab.CCEvent, 1)
		})

		JustBeforeEach(func() {
			go func() {
				err = impl.Run()
			}()

			// give the test case a chance to read from the channel
			time.Sleep(1 * time.Millisecond)
			close(eventch)
		})

		Context("when channel client succeeds", func() {
			BeforeEach(func() {
				eventch <- &fab.CCEvent{
					TxID: "1",
				}

				client.RegisterChaincodeEventReturns(struct{}{}, eventch, nil)
			})

			It("should process chaincode events", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when channel client fails", func() {
			BeforeEach(func() {
				client.RegisterChaincodeEventReturns(nil, nil, errors.New("events error"))
			})

			It("should fail process chaincode events", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("events error"))
			})
		})
	})
})
