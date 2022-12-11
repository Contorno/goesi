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

func easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi(in *jlexer.Lexer, out *PostUiOpenwindowNewmailNewMailList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostUiOpenwindowNewmailNewMailList, 0, 1)
			} else {
				*out = PostUiOpenwindowNewmailNewMailList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 PostUiOpenwindowNewmailNewMail
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
func easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi(out *jwriter.Writer, in PostUiOpenwindowNewmailNewMailList) {
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
func (v PostUiOpenwindowNewmailNewMailList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUiOpenwindowNewmailNewMailList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMailList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMailList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi(l, v)
}
func easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi1(in *jlexer.Lexer, out *PostUiOpenwindowNewmailNewMail) {
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
						out.Recipients = make([]int32, 0, 16)
					} else {
						out.Recipients = []int32{}
					}
				} else {
					out.Recipients = (out.Recipients)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Recipients = append(out.Recipients, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "subject":
			out.Subject = string(in.String())
		case "to_corp_or_alliance_id":
			out.ToCorpOrAllianceId = int32(in.Int32())
		case "to_mailing_list_id":
			out.ToMailingListId = int32(in.Int32())
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
func easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi1(out *jwriter.Writer, in PostUiOpenwindowNewmailNewMail) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Body != "" {
		const prefix string = ",\"body\":"
		first = false
		out.RawString(prefix[1:])
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
				out.Int32(int32(v6))
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
	if in.ToCorpOrAllianceId != 0 {
		const prefix string = ",\"to_corp_or_alliance_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ToCorpOrAllianceId))
	}
	if in.ToMailingListId != 0 {
		const prefix string = ",\"to_mailing_list_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ToMailingListId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUiOpenwindowNewmailNewMail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUiOpenwindowNewmailNewMail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6b8ca5bfEncodeGithubComContornoGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUiOpenwindowNewmailNewMail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6b8ca5bfDecodeGithubComContornoGoesiEsi1(l, v)
}
