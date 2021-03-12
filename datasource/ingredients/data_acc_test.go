package ingredients

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/ingredients/data_acc_test.go  -timeout=120m
func TestIngredientsDatasource(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "ingredients_datasource_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "hashicups-ingredients",
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
data "hashicups-ingredients" "vagrante-ingredients" {
  username = "education"
  password = "test123"
  coffee = "Vagrante espresso"
}

locals {
  espresso = data.hashicups-ingredients.vagrante-ingredients.map["Espresso"]
}

source "hashicups-order" "my-custom-order" {
  username = "education"
  password = "test123"

  item {
    coffee {
      id = 5
      name = "my custom vagrante"
      ingredient {
        id = local.espresso
        quantity = 50
      }
    }
  }
}

build {
  sources = ["sources.hashicups-order.my-custom-order"]
}
`
