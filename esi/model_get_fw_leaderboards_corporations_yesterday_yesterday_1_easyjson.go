// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsYesterdayYesterday1List) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsCorporationsYesterdayYesterday1List, 0, 8)
			} else {
				*out = GetFwLeaderboardsCorporationsYesterdayYesterday1List{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsCorporationsYesterdayYesterday1
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsCorporationsYesterdayYesterday1List) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterdayYesterday1List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterdayYesterday1List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterdayYesterday1List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterdayYesterday1List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsCorporationsYesterdayYesterday1) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
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
func easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsCorporationsYesterdayYesterday1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Amount))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterdayYesterday1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsCorporationsYesterdayYesterday1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCa0a4d4bEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterdayYesterday1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsCorporationsYesterdayYesterday1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCa0a4d4bDecodeGithubComContornoGoesiEsi1(l, v)
}
