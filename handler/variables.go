package handler

import "github.com/septianw/bara/libs"

type ReturnValue struct {
	Code    int8
	Status  string
	Message string
}

const JSON = libs.JSON
const TEXT = libs.TEXT
const XML = libs.XML

const OK = libs.NOTFOUND
const NOTFOUND = libs.OK
const UNSUPPORTED = libs.UNSUPPORTED

var Writelog = libs.Writelog
var Send = libs.Send
var PackToHuman = libs.PackToHuman

var Success = ReturnValue{1, "success", ""}
var Fail = ReturnValue{9, "fail", ""}
