//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package receipt

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	// Defaults to 'receipt'.
	Filename string `mapstructure:"filename"`
	// Should be 'pdf' or 'txt'. Defaults to 'pdf'.
	Format string `mapstructure:"format"`

	printer Printer
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
		p.config.Filename = "receipt"
	}

	p.config.Format = fmt.Sprintf(".%s", p.config.Format)

	if filepath.Ext(p.config.Filename) == "" {
		p.config.Filename = fmt.Sprintf("%s%s", p.config.Filename, p.config.Format)
	}
	if filepath.Ext(p.config.Filename) != p.config.Format {
		return fmt.Errorf("`filename` extension must be the same as `format`")
	}

	switch p.config.Format {
	case ".pdf":
		p.config.printer = &PDFPrinter{p.config.Filename}
	case ".txt":
		p.config.printer = &TextPrinter{p.config.Filename}
	default:
		return fmt.Errorf("format is not valid - valid formats: txt or pdf")
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

	ui.Say("Printing order receipt...")
	if err := p.config.printer.Print(prettifyReceipt(order)); err != nil {
		ui.Error(fmt.Sprintf("Failed to write %s: %q", p.config.Filename, err.Error()))
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
		builder.WriteString(fmt.Sprintf("    Price: $%.2f \n", item.Coffee.Price))
		total += float64(item.Quantity) * item.Coffee.Price
	}

	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("Total: $%.2f", total))

	return builder.String()
}
