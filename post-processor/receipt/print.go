// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package receipt

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Printer interface {
	Print(content string) error
}

type TextPrinter struct {
	Filename string
}

func (p *TextPrinter) Print(content string) error {
	if info, err := os.Stat(p.Filename); err == nil {
		p.Filename = fmt.Sprintf("%s-%d.%s", info.Name(), time.Now().Unix(), filepath.Ext(p.Filename))
	}
	f, err := os.Create(p.Filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

type PDFPrinter struct {
	Filename string
}

func (p *PDFPrinter) Print(content string) error {
	if info, err := os.Stat(p.Filename); err == nil {
		p.Filename = fmt.Sprintf("%s-%d.%s", info.Name(), time.Now().Unix(), filepath.Ext(p.Filename))
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.MultiCell(190, 7, content, "5", "L", false)

	if err := pdf.OutputFileAndClose(p.Filename); err != nil {
		return err
	}
	return nil
}
