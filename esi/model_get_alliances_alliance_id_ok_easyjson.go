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

func easyjson3ab323d2DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetAlliancesAllianceIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetAlliancesAllianceIdOkList, 0, 0)
			} else {
				*out = GetAlliancesAllianceIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetAlliancesAllianceIdOk
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
func easyjson3ab323d2EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetAlliancesAllianceIdOkList) {
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
func (v GetAlliancesAllianceIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3ab323d2EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3ab323d2EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3ab323d2DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3ab323d2DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson3ab323d2DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetAlliancesAllianceIdOk) {
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
		case "creator_corporation_id":
			out.CreatorCorporationId = int32(in.Int32())
		case "creator_id":
			out.CreatorId = int32(in.Int32())
		case "date_founded":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateFounded).UnmarshalJSON(data))
			}
		case "executor_corporation_id":
			out.ExecutorCorporationId = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "ticker":
			out.Ticker = string(in.String())
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
func easyjson3ab323d2EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetAlliancesAllianceIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CreatorCorporationId != 0 {
		const prefix string = ",\"creator_corporation_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CreatorCorporationId))
	}
	if in.CreatorId != 0 {
		const prefix string = ",\"creator_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CreatorId))
	}
	if true {
		const prefix string = ",\"date_founded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.DateFounded).MarshalJSON())
	}
	if in.ExecutorCorporationId != 0 {
		const prefix string = ",\"executor_corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ExecutorCorporationId))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Ticker != "" {
		const prefix string = ",\"ticker\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ticker))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAlliancesAllianceIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3ab323d2EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3ab323d2EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3ab323d2DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3ab323d2DecodeGithubComContornoGoesiEsi1(l, v)
}
