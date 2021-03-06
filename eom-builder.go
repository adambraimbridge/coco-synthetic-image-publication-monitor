package main

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"text/template"
	"time"
)

//Eom struct matching Methode eom files
type Eom struct {
	UUID             string        `json:"uuid"`
	EomType          string        `json:"type"`
	Value            string        `json:"value"`
	Attributes       string        `json:"attributes"`
	WorkflowStatus   string        `json:"workflowStatus"`
	SystemAttributes string        `json:"systemAttributes"`
	UsageTickets     string        `json:"usageTickets"`
	LinkedObjects    []interface{} `json:"linkedObjects"`
}

// BuildRandomEOMImage builds and populates an Eom struct.
// The EOM value is a randomly generated 1000 length byte array encoded with base64 in string representation
func BuildRandomEOMImage(uuid string) (*Eom, time.Time) {
	t := time.Now()
	return &Eom{
		UUID:             uuid,
		EomType:          "Image",
		Value:            base64.StdEncoding.EncodeToString(randomBytes(1000)),
		Attributes:       populateTemplate("attributes.template", uuid),
		WorkflowStatus:   "",
		SystemAttributes: populateTemplate("systemAttributes.template", t.Format("20060102")),
		UsageTickets: populateTemplate("usageTickets.template", struct {
			UUID, Date, FormattedDate string
		}{
			uuid,
			t.Format("20060102150405"),
			t.Format(time.UnixDate),
		}),
		LinkedObjects: make([]interface{}, 0),
	}, t
}

func randomBytes(length int) []byte {
	b := make([]byte, length)
	for i := 0; i < len(b); i++ {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func populateTemplate(fileTempl string, data interface{}) string {
	tmpl, err := template.ParseFiles(fileTempl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
