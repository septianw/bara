package libs

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	//	"github.com/septianw/golok"
	//	"github.com/septianw/log15"
)

type Payload struct {
	Id     int8
	Name   string
	Active bool
}

type PayloadSequence struct {
	Text      string
	Type      uint8
	PlainType string
}

func CaptureOutput(f func()) string {
	var old = os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}

// convert from map to struct found on http://play.golang.org/p/tN8mxT_V9h http://stackoverflow.com/questions/26744873/converting-map-to-struct

func TestSend(t *testing.T) {
	var paid [3]PayloadSequence

	record := Payload{
		Id:     12,
		Name:   "Robert Murdock",
		Active: true,
	}

	paid[0].Text = "Hi. I'm Robert. Robert Murdock"
	paid[0].Type = TEXT
	paid[0].PlainType = "text/plain"

	tmpjson, jsonr := json.Marshal(&record)
	paid[1].Text = string(tmpjson)
	paid[1].Type = JSON
	paid[1].PlainType = "application/json"

	tmpxml, xmlr := xml.Marshal(record)
	paid[2].Text = string(tmpxml)
	paid[2].Type = XML
	paid[2].PlainType = "application/xml"

	t.Logf("jsonr: %+v; xmlr: %+v", jsonr, xmlr)

	for _, pay := range paid {
		hts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Send(w, OK, pay.Type, pay.Text)
		}))

		res, err := http.Get(hts.URL)
		if err != nil {
			t.Errorf("Unknown error: %+v", err)
			t.FailNow()
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Errorf("ioutil fail to read body: %+v", err)
			t.Fail()
			return
		}

		if strings.Compare(string(body), pay.Text) != 0 {
			t.Errorf("Expected %s, have %s", pay.Text, string(body))
		}

		t.Logf("headers: %+v", res.Header)
		t.Logf("Body: %+v", string(body))
		if strings.Compare(res.Header.Get("Content-Type"), pay.PlainType) != 0 {
			t.Errorf("Expected %s, have %s", pay.PlainType, (res.Header.Get("Content-Type")))
		}

		hts.Close()
	}
}

func TestPackToHuman(t *testing.T) {
	person := Payload{15, "Sonoya Mizuno", true}
	//	var result Payload
	personToHuman := PackToHuman(interface{}(person))
	var humanToPerson Payload

	fail := PackToHuman(person)
	t.Log(fail)
	if strings.Compare(fail, "{}") == 0 {
		t.Errorf("PackToHuman input is %+v instead of %+v", reflect.TypeOf(person).Kind(), interface{}(person))
	}

	t.Logf("PersonToHuman: %+v", personToHuman)
	t.Logf("personToHuman type: %+v", reflect.TypeOf(personToHuman).Kind())
	t.Logf("Person type: %+v", reflect.TypeOf(interface{}(person)).Kind())

	err := json.Unmarshal([]byte(personToHuman), &humanToPerson)

	//	_ = FillStruct(map[string]interface{}(humanToPerson), result)

	if err != nil {
		t.Errorf("Unmarshal person fail. Reason: %+v", err)
		t.Fail()
		return
	} else {
		if !reflect.DeepEqual(person, humanToPerson) {
			t.Errorf("Person unequal %+v", humanToPerson)
			t.Fail()
			return
		}
	}

}
