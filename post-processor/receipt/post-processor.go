//go:generate mapstructure-to-hcl2 -type Config

package receipt

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/jung-kurt/gofpdf"
	"github.com/mitchellh/mapstructure"
	"github.com/sylviamoss/hashicups-client-go"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	// Must contain extension. Defaults to `receipt.pdf`.
	Filename string `mapstructure:"filename"`
	// Should be `pdf` or `txt`. Defaults to pdf.
	Format string `mapstructure:"format"`
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec { return p.config.FlatMapstructure().HCL2Spec() }

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, nil, raws...)
	if err != nil {
		return err
	}

	if p.config.Format == "" {
		p.config.Format = "pdf"
	}

	if p.config.Filename == "" {
		p.config.Filename = "receipt.pdf"
	}

	return nil
}

func (p *PostProcessor) PostProcess(_ context.Context, ui packersdk.Ui, source packersdk.Artifact) (packersdk.Artifact, bool, bool, error) {
	order := hashicups.Order{}
	if err := mapstructure.Decode(source.State("order"), &order); err != nil {
		err := fmt.Errorf("Failed to decoder order")
		ui.Error(err.Error())
		return source, false, false, err
	}

	client := new(hashicups.Client)
	if err := mapstructure.Decode(source.State("client"), client); err != nil {
		err := fmt.Errorf("Failed to decoder client")
		ui.Error(err.Error())
		return source, false, false, err
	}

	// Get ingredients information
	for i, item := range order.Items {
		ingredients, err := client.GetCoffeeIngredients(strconv.Itoa(item.Coffee.ID))
		if err != nil {
			continue
		}
		order.Items[i].Coffee.Ingredient = ingredients
	}

	if err := writePDF(p.config.Filename, prettifyReceipt(order)); err != nil {
		ui.Error(fmt.Sprintf("Failed to write: %q", err.Error()))
		return source, false, false, err
	}
	return source, true, true, nil
}

func prettifyReceipt(order hashicups.Order) string {
	builder := strings.Builder{}

	builder.WriteString("========== ORDER RECEIPT ===========\n")
	builder.WriteString(fmt.Sprintf("ID: %d \n", order.ID))
	builder.WriteString("\n")
	total := float64(0)
	for i, item := range order.Items {
		builder.WriteString(fmt.Sprintf("-> Item %d \n", i))
		builder.WriteString(fmt.Sprintf("    x%d %s (id: %d)     \n", item.Quantity, item.Coffee.Name, item.Coffee.ID))
		for _, ingredient := range item.Coffee.Ingredient {
			builder.WriteString(fmt.Sprintf("          %d: %s   %d%s\n", ingredient.ID, ingredient.Name, ingredient.Quantity, ingredient.Unit))
		}
		builder.WriteString(fmt.Sprintf("    Price: $%.2f", item.Coffee.Price))
		total += float64(item.Quantity) * item.Coffee.Price
	}

	builder.WriteString("\n\n")
	builder.WriteString(fmt.Sprintf("Total: $%.2f", total))

	return builder.String()
}

func writePDF(filename, content string) error {
	if info, err := os.Stat(filename); err == nil {
		filename = fmt.Sprintf("%s-%d.%s", info.Name(), time.Now().Unix(), filepath.Ext(filename))
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.MultiCell(190, 7, content, "5", "L", false)

	if err := pdf.OutputFileAndClose(filename); err != nil {
		return err
	}
	return nil
}
