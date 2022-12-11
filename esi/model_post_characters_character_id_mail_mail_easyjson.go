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

func easyjson28a4c0c5DecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PostCharactersCharacterIdMailMailList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostCharactersCharacterIdMailMailList, 0, 1)
			} else {
				*out = PostCharactersCharacterIdMailMailList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostCharactersCharacterIdMailMail
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
func easyjson28a4c0c5EncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PostCharactersCharacterIdMailMailList) {
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
func (v PostCharactersCharacterIdMailMailList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28a4c0c5EncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailMailList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28a4c0c5EncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailMailList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28a4c0c5DecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailMailList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28a4c0c5DecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson28a4c0c5DecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PostCharactersCharacterIdMailMail) {
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
		case "approved_cost":
			out.ApprovedCost = int64(in.Int64())
		case "body":
			out.Body = string(in.String())
		case "recipients":
			if in.IsNull() {
				in.Skip()
				out.Recipients = nil
			} else {
				in.Delim('[')
				if out.Recipients == nil {
					if !in.IsDelim(']') {
						out.Recipients = make([]PostCharactersCharacterIdMailRecipient, 0, 2)
					} else {
						out.Recipients = []PostCharactersCharacterIdMailRecipient{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v4 PostCharactersCharacterIdMailRecipient
					(v4).UnmarshalEasyJSON(in)
					out.Recipients = append(out.Recipients, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "subject":
			out.Subject = string(in.String())
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
func easyjson28a4c0c5EncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PostCharactersCharacterIdMailMail) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ApprovedCost != 0 {
		const prefix string = ",\"approved_cost\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ApprovedCost))
	}
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if len(in.Recipients) != 0 {
		const prefix string = ",\"recipients\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Recipients {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Subject != "" {
		const prefix string = ",\"subject\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Subject))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostCharactersCharacterIdMailMail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson28a4c0c5EncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostCharactersCharacterIdMailMail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson28a4c0c5EncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostCharactersCharacterIdMailMail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson28a4c0c5DecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostCharactersCharacterIdMailMail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson28a4c0c5DecodeGithubComContornoGoesiEsi1(l, v)
}
