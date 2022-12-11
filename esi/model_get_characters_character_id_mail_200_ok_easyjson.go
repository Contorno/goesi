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

func easyjson69f56f1fDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdMail200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdMail200OkList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdMail200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdMail200Ok
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
func easyjson69f56f1fEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdMail200OkList) {
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
func (v GetCharactersCharacterIdMail200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson69f56f1fEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMail200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson69f56f1fEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMail200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson69f56f1fDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMail200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson69f56f1fDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson69f56f1fDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdMail200Ok) {
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
		case "from":
			out.From = int32(in.Int32())
		case "is_read":
			out.IsRead = bool(in.Bool())
		case "labels":
			if in.IsNull() {
				in.Skip()
				out.Labels = nil
			} else {
				in.Delim('[')
				if out.Labels == nil {
					if !in.IsDelim(']') {
						out.Labels = make([]int32, 0, 16)
					} else {
						out.Labels = []int32{}
					}
				} else {
					out.Labels = (out.Labels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Labels = append(out.Labels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "mail_id":
			out.MailId = int32(in.Int32())
		case "recipients":
			if in.IsNull() {
				in.Skip()
				out.Recipients = nil
			} else {
				in.Delim('[')
				if out.Recipients == nil {
					if !in.IsDelim(']') {
						out.Recipients = make([]GetCharactersCharacterIdMailRecipient, 0, 2)
					} else {
						out.Recipients = []GetCharactersCharacterIdMailRecipient{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetCharactersCharacterIdMailRecipient
					(v5).UnmarshalEasyJSON(in)
					out.Recipients = append(out.Recipients, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "subject":
			out.Subject = string(in.String())
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
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
func easyjson69f56f1fEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdMail200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.From != 0 {
		const prefix string = ",\"from\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.From))
	}
	if in.IsRead {
		const prefix string = ",\"is_read\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsRead))
	}
	if len(in.Labels) != 0 {
		const prefix string = ",\"labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Labels {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v7))
			}
			out.RawByte(']')
		}
	}
	if in.MailId != 0 {
		const prefix string = ",\"mail_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MailId))
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
			for v8, v9 := range in.Recipients {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
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
	if true {
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Timestamp).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdMail200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson69f56f1fEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdMail200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson69f56f1fEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdMail200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson69f56f1fDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdMail200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson69f56f1fDecodeGithubComContornoGoesiEsi1(l, v)
}
