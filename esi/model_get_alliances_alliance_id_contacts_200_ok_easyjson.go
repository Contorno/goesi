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

func easyjson8f4a097DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetAlliancesAllianceIdContacts200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetAlliancesAllianceIdContacts200OkList, 0, 1)
			} else {
				*out = GetAlliancesAllianceIdContacts200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetAlliancesAllianceIdContacts200Ok
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
func easyjson8f4a097EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetAlliancesAllianceIdContacts200OkList) {
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
func (v GetAlliancesAllianceIdContacts200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8f4a097EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdContacts200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8f4a097EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdContacts200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8f4a097DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdContacts200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8f4a097DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson8f4a097DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetAlliancesAllianceIdContacts200Ok) {
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
		case "contact_id":
			out.ContactId = int32(in.Int32())
		case "contact_type":
			out.ContactType = string(in.String())
		case "label_ids":
			if in.IsNull() {
				in.Skip()
				out.LabelIds = nil
			} else {
				in.Delim('[')
				if out.LabelIds == nil {
					if !in.IsDelim(']') {
						out.LabelIds = make([]int64, 0, 8)
					} else {
						out.LabelIds = []int64{}
					}
				} else {
					out.LabelIds = (out.LabelIds)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int64
					v4 = int64(in.Int64())
					out.LabelIds = append(out.LabelIds, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "standing":
			out.Standing = float32(in.Float32())
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
func easyjson8f4a097EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetAlliancesAllianceIdContacts200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ContactId != 0 {
		const prefix string = ",\"contact_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.ContactId))
	}
	if in.ContactType != "" {
		const prefix string = ",\"contact_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContactType))
	}
	if len(in.LabelIds) != 0 {
		const prefix string = ",\"label_ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.LabelIds {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v6))
			}
			out.RawByte(']')
		}
	}
	if in.Standing != 0 {
		const prefix string = ",\"standing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Standing))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetAlliancesAllianceIdContacts200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8f4a097EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetAlliancesAllianceIdContacts200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8f4a097EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetAlliancesAllianceIdContacts200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8f4a097DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetAlliancesAllianceIdContacts200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8f4a097DecodeGithubComContornoGoesiEsi1(l, v)
}
