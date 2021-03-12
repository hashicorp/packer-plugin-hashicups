package coffees

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/coffees/data_acc_test.go  -timeout=120m
func TestCoffeesDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "coffees_datasource_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "hashicups-coffees",
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

			buildGeneratedDataLog := "hashicups-order.my-custom-order: Order [0-9]+ created!"
			if matched, _ := regexp.MatchString(buildGeneratedDataLog+".*", logsString); !matched {
				t.Fatalf("logs doesn't contain expected output %q", logsString)
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

const testDatasourceHCL2Basic = `
data "hashicups-coffees" "coffees" {
  username = "education"
  password = "test123"
}

locals {
  vagrante_espresso = data.hashicups-coffees.coffees.map["Vagrante espresso"]
}

source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = local.vagrante_espresso
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
}
`
