package restpack_pdf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	convertEndpoint = "/html2pdf/v7/convert"

	//EmulateMediaOptions
	MediaScreen = "screen"
	MediaPrint  = "print"

	//PageSizes
	PageSizeA0      = "A0"
	PageSizeA1      = "A1"
	PageSizeA2      = "A2"
	PageSizeA3      = "A3"
	PageSizeA4      = "A4"
	PageSizeA5      = "A5"
	PageSizeA6      = "A6"
	PageSizeLegal   = "Legal"
	PageSizeLetter  = "Letter"
	PageSizeTabloid = "Tabloid"
	PageSizeLedger  = "Ledger"
	PageSizeFull    = "Full"
)

type HtmlToPDF struct {
	client       *client
	Url          *string `json:"url,omitempty"`
	BodyHtml     *string `json:"html,omitempty"`
	FooterHtml   *string `json:"pdf_footer,omitempty"`
	BodyCss      *string `json:"css,omitempty"`
	PdfPage      *string `json:"pdf_page,omitempty"`
	NoModify     *bool   `json:"no_modify,omitempty"`
	EmulateMedia *string `json:"emulate_media,omitempty"`
	Margins      *string `json:"pdf_margins,omitempty"`
	ReturnJSON   *bool   `json:"json,omitempty"`
}

func NewHtmlToPDFJob(apiKey string, setDefaults bool) *HtmlToPDF {
	c := &HtmlToPDF{client: initClient(apiKey)}
	if setDefaults {
		c.setDefaults()
	}
	return c
}

func (h *HtmlToPDF) setDefaults() {
	margins := "30px"
	h.Margins = &margins
	defaultEmulateMedia := MediaScreen
	h.EmulateMedia = &defaultEmulateMedia
	defaultNoModify := false
	h.NoModify = &defaultNoModify
	h.SetPageSize(PageSizeA4)
}
func (h *HtmlToPDF) SetMargins(margins string) {
	h.Margins = &margins
}
func (h *HtmlToPDF) SetPageSize(size string) {
	h.PdfPage = &size
}
func (h *HtmlToPDF) SetBodyHtml(bodyHtml string) {
	h.BodyHtml = &bodyHtml
}
func (h *HtmlToPDF) SetFooterHtml(footer string) {
	h.FooterHtml = &footer
}
func (h *HtmlToPDF) SetBodyCss(css string) {
	h.BodyCss = &css
}

type HTMLToPDFCaptureResult struct {
	Image        string `json:"image,omitempty"`
	Width        string `json:"width,omitempty"`
	Height       string `json:"height,omitempty"`
	RemoteStatus string `json:"remote_status,omitempty"`
	Cached       bool   `json:"cached,string,omitempty"`
	URL          string `json:"url,omitempty"`
}

func (h *HtmlToPDF) Convert(returnUrl bool) (resultJson *HTMLToPDFCaptureResult, file []byte, err error) {
	h.ReturnJSON = &returnUrl

	var resp *http.Response
	resp, err = h.client.request(h, convertEndpoint)
	if err != nil {
		return
	}

	if returnUrl {
		err = json.NewDecoder(resp.Body).Decode(&resultJson)
		return
	}

	file, err = ioutil.ReadAll(resp.Body)
	return
}

//func (h *HtmlToPDF) BuildPayload() {
//
//	payload := url.Values{}
//	if h.BodyCss != nil {
//	payload.Set("css", *h.BodyCss)
//	}
//	payload.Set("pdf_page", "A4")
//	payload.Set("html", doc.String())
//	if restpack {
//		payload.Set("json", "true")
//	}
//	payload.Set("no_modify", "true")
//	payload.Set("emulate_media", "screen")
//	payload.Set("pdf_margins", "30px 30px 30px 30px")
//
//
//}
