// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toppings

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

// Run with: PACKER_ACC=1 go test -count 1 -v ./provisioner/toppings/provisioner_acc_test.go  -timeout=120m
func TestToppingsProvisioner(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "toppings_provisioner_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testProvisionerHCL2Basic,
		Type:     "hashicups-toppings",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}

			logs, err := os.Open(logfile)
			if err != nil {
				return fmt.Errorf("Unable find %s", logfile)
			}
			defer logs.Close()

			logsBytes, err := ioutil.ReadAll(logs)
			if err != nil {
				return fmt.Errorf("Unable to read %s", logfile)
			}
			logsString := string(logsBytes)

			provisionerOutputLog := "Pouring cinnamon..."
			if matched, _ := regexp.MatchString(provisionerOutputLog+".*", logsString); !matched {
				t.Fatalf("logs doesn't contain expected output %q", logsString)
			}
			provisionerOutputLog = "Pouring marshmellow..."
			if matched, _ := regexp.MatchString(provisionerOutputLog+".*", logsString); !matched {
				t.Fatalf("logs doesn't contain expected output %q", logsString)
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

const testProvisionerHCL2Basic = `
source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = 5
      name = "my custom vagrante"
      ingredient {
        id = 1
        quantity = 50
      }
    }
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]

  provisioner "hashicups-toppings" {
    toppings = ["cinnamon", "marshmellow"]
  }
}
`
