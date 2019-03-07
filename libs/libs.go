package libs

import (
	"encoding/json"
	"net/http"

	"github.com/septianw/golok"
)

const TEXT uint8 = 1
const JSON uint8 = 2
const XML uint8 = 3

const NOTFOUND = 404
const OK = 200
const UNSUPPORTED = 415

var Loglevel = &golok.Loglevel
var Logfile = &golok.Logfile

/*
  Get human config and convert it to machine config
*/
var Getconfig = golok.Getconfig

/*
 Just write to log, and don't care about config.
 get config and write log.
*/
var Writelog = golok.Writelog

/*
 Send payload response to user via http response writer.
*/
func Send(w http.ResponseWriter, h int, payloadType uint8, payload string) {
	var contentType string

	//	if payload == nil {
	//		payload = ""
	//	}

	switch payloadType {
	case TEXT:
		contentType = "text/plain"
	case JSON:
		contentType = "application/json"
	case XML:
		contentType = "application/xml"
	default:
		contentType = "text/plain"
	}
	Writelog("debug", "type", "ContentType", contentType)

	w.Header().Set("content-type", contentType)
	w.WriteHeader(h)
	w.Write([]byte(payload))
}

/*
   Pack a goods to deliver to human
*/

func PackToHuman(goods interface{}) string {
	Writelog("debug", "PackToHuman input", "goods", goods)

	Package, err := json.Marshal(goods)
	Writelog("debug", "PackToHuman output", "Package", string(Package), "err", err)
	if err != nil {
		Writelog("debug", "json marshaller fail on PackToHuman", "err", err)
		Writelog("error", "Unable to pack to human.", "Reason:", err)
	}

	return string(Package)
}
