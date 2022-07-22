// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE34310f8DecodeOzonTestAppModels(in *jlexer.Lexer, out *CustomError) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "error":
			out.CustomErr = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE34310f8EncodeOzonTestAppModels(out *jwriter.Writer, in CustomError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix[1:])
		out.String(string(in.CustomErr))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CustomError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE34310f8EncodeOzonTestAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CustomError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE34310f8EncodeOzonTestAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CustomError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE34310f8DecodeOzonTestAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CustomError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE34310f8DecodeOzonTestAppModels(l, v)
}
func easyjsonE34310f8DecodeOzonTestAppModels1(in *jlexer.Lexer, out *Created) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "created":
			out.CreatedInfo = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE34310f8EncodeOzonTestAppModels1(out *jwriter.Writer, in Created) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.CreatedInfo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Created) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE34310f8EncodeOzonTestAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Created) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE34310f8EncodeOzonTestAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Created) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE34310f8DecodeOzonTestAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Created) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE34310f8DecodeOzonTestAppModels1(l, v)
}
