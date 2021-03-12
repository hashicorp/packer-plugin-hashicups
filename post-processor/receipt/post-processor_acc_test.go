package receipt

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
	"github.com/hashicorp/packer-plugin-sdk/acctest/testutils"
)

// Run with: PACKER_ACC=1 go test -count 1 -v ./post-processor/receipt/post-processor_acc_test.go  -timeout=120m
func TestReceiptPostProcessor(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "scaffolding_post-processor_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			testutils.CleanupFiles("receipt.txt", "receipt.pdf")
			return nil
		},
		Template: testPostProcessorHCL2Basic,
		Type:     "scaffolding-my-post-processor",
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

			postProcessorOutputLog := "Printing order receipt..."
			if matched, _ := regexp.MatchString(postProcessorOutputLog+".*", logsString); !matched {
				t.Fatalf("logs doesn't contain expected output %q", logsString)
			}

			if !testutils.FileExists("receipt.txt") {
				t.Fatal("couldn't find expected receipt file: receipt.txt")
			}
			if !testutils.FileExists("receipt.pdf") {
				t.Fatal("couldn't find expected receipt file: receipt.pdf")
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

const testPostProcessorHCL2Basic = `
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

  post-processor "hashicups-receipt" {
    format = "pdf"
  }
  post-processor "hashicups-receipt" {
    format = "txt"
  }
}
`
