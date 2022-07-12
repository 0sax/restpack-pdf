package restpack_pdf

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestHtmlToPDF_Convert(t *testing.T) {
	type fields struct {
		BodyHtml   *string
		FooterHtml *string
		BodyCss    *string
	}
	tests := []struct {
		name        string
		setDefaults bool
		returnUrl   bool
		wantErr     bool
		fields      fields
	}{
		// Test Cases
		{
			name:        "Test 1",
			setDefaults: true,
			returnUrl:   false,
			wantErr:     false,
			fields: fields{
				BodyHtml:   &boody,
				FooterHtml: &footer,
				BodyCss:    &css,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHtmlToPDFJob(apiKey, tt.setDefaults)
			if tt.fields.BodyHtml != nil {
				h.SetBodyHtml(*tt.fields.BodyHtml)
			}
			if tt.fields.FooterHtml != nil {
				h.SetFooterHtml(*tt.fields.FooterHtml)
			}
			if tt.fields.BodyCss != nil {
				h.SetBodyCss(*tt.fields.BodyCss)
			}
			h.SetMargins("30px 30px 30px 30px")
			gotResultJson, gotFile, err := h.Convert(tt.returnUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.returnUrl && gotResultJson == nil && !tt.wantErr {
				t.Errorf("Convert() Json result was nil, not quite what we want")
			}
			if !tt.returnUrl && gotFile == nil && !tt.wantErr {
				t.Errorf("Convert() File was nil, not quite what we want")
			}
			fmt.Printf("JsonResult: %+v\n", gotResultJson)

			if gotFile != nil {
				fmt.Printf("FileResuly: %+v\n", gotFile)
				writeTofile(gotFile)
			}
		})
	}
}

func writeTofile(bytes []byte) {
	// Open a new file for writing only
	file, err := os.Create("testresult.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	bytesWritten, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
