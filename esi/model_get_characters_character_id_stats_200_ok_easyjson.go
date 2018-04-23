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

func easyjson19558e8dDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStats200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStats200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdStats200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStats200Ok
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
func easyjson19558e8dEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStats200OkList) {
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
func (v GetCharactersCharacterIdStats200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19558e8dEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStats200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19558e8dEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStats200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19558e8dDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStats200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19558e8dDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson19558e8dDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStats200Ok) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "year":
			out.Year = int32(in.Int32())
		case "character":
			(out.Character).UnmarshalEasyJSON(in)
		case "combat":
			(out.Combat).UnmarshalEasyJSON(in)
		case "industry":
			(out.Industry).UnmarshalEasyJSON(in)
		case "inventory":
			(out.Inventory).UnmarshalEasyJSON(in)
		case "isk":
			(out.Isk).UnmarshalEasyJSON(in)
		case "market":
			(out.Market).UnmarshalEasyJSON(in)
		case "mining":
			easyjson19558e8dDecodeGithubComAntihaxGoesiEsi2(in, &out.Mining)
		case "module":
			(out.Module).UnmarshalEasyJSON(in)
		case "orbital":
			(out.Orbital).UnmarshalEasyJSON(in)
		case "pve":
			(out.Pve).UnmarshalEasyJSON(in)
		case "social":
			easyjson19558e8dDecodeGithubComAntihaxGoesiEsi3(in, &out.Social)
		case "travel":
			easyjson19558e8dDecodeGithubComAntihaxGoesiEsi4(in, &out.Travel)
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
func easyjson19558e8dEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStats200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Year != 0 {
		const prefix string = ",\"year\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Year))
	}
	if true {
		const prefix string = ",\"character\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Character).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"combat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Combat).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"industry\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Industry).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"inventory\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Inventory).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"isk\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Isk).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"market\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Market).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"mining\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19558e8dEncodeGithubComAntihaxGoesiEsi2(out, in.Mining)
	}
	if true {
		const prefix string = ",\"module\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Module).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"orbital\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Orbital).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"pve\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Pve).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"social\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19558e8dEncodeGithubComAntihaxGoesiEsi3(out, in.Social)
	}
	if true {
		const prefix string = ",\"travel\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19558e8dEncodeGithubComAntihaxGoesiEsi4(out, in.Travel)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStats200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19558e8dEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStats200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19558e8dEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStats200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19558e8dDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStats200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19558e8dDecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson19558e8dDecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsTravel) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "acceleration_gate_activations":
			out.AccelerationGateActivations = int64(in.Int64())
		case "align_to":
			out.AlignTo = int64(in.Int64())
		case "distance_warped_high_sec":
			out.DistanceWarpedHighSec = int64(in.Int64())
		case "distance_warped_low_sec":
			out.DistanceWarpedLowSec = int64(in.Int64())
		case "distance_warped_null_sec":
			out.DistanceWarpedNullSec = int64(in.Int64())
		case "distance_warped_wormhole":
			out.DistanceWarpedWormhole = int64(in.Int64())
		case "docks_high_sec":
			out.DocksHighSec = int64(in.Int64())
		case "docks_low_sec":
			out.DocksLowSec = int64(in.Int64())
		case "docks_null_sec":
			out.DocksNullSec = int64(in.Int64())
		case "jumps_stargate_high_sec":
			out.JumpsStargateHighSec = int64(in.Int64())
		case "jumps_stargate_low_sec":
			out.JumpsStargateLowSec = int64(in.Int64())
		case "jumps_stargate_null_sec":
			out.JumpsStargateNullSec = int64(in.Int64())
		case "jumps_wormhole":
			out.JumpsWormhole = int64(in.Int64())
		case "warps_high_sec":
			out.WarpsHighSec = int64(in.Int64())
		case "warps_low_sec":
			out.WarpsLowSec = int64(in.Int64())
		case "warps_null_sec":
			out.WarpsNullSec = int64(in.Int64())
		case "warps_to_bookmark":
			out.WarpsToBookmark = int64(in.Int64())
		case "warps_to_celestial":
			out.WarpsToCelestial = int64(in.Int64())
		case "warps_to_fleet_member":
			out.WarpsToFleetMember = int64(in.Int64())
		case "warps_to_scan_result":
			out.WarpsToScanResult = int64(in.Int64())
		case "warps_wormhole":
			out.WarpsWormhole = int64(in.Int64())
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
func easyjson19558e8dEncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetCharactersCharacterIdStatsTravel) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AccelerationGateActivations != 0 {
		const prefix string = ",\"acceleration_gate_activations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AccelerationGateActivations))
	}
	if in.AlignTo != 0 {
		const prefix string = ",\"align_to\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AlignTo))
	}
	if in.DistanceWarpedHighSec != 0 {
		const prefix string = ",\"distance_warped_high_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DistanceWarpedHighSec))
	}
	if in.DistanceWarpedLowSec != 0 {
		const prefix string = ",\"distance_warped_low_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DistanceWarpedLowSec))
	}
	if in.DistanceWarpedNullSec != 0 {
		const prefix string = ",\"distance_warped_null_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DistanceWarpedNullSec))
	}
	if in.DistanceWarpedWormhole != 0 {
		const prefix string = ",\"distance_warped_wormhole\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DistanceWarpedWormhole))
	}
	if in.DocksHighSec != 0 {
		const prefix string = ",\"docks_high_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DocksHighSec))
	}
	if in.DocksLowSec != 0 {
		const prefix string = ",\"docks_low_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DocksLowSec))
	}
	if in.DocksNullSec != 0 {
		const prefix string = ",\"docks_null_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DocksNullSec))
	}
	if in.JumpsStargateHighSec != 0 {
		const prefix string = ",\"jumps_stargate_high_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JumpsStargateHighSec))
	}
	if in.JumpsStargateLowSec != 0 {
		const prefix string = ",\"jumps_stargate_low_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JumpsStargateLowSec))
	}
	if in.JumpsStargateNullSec != 0 {
		const prefix string = ",\"jumps_stargate_null_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JumpsStargateNullSec))
	}
	if in.JumpsWormhole != 0 {
		const prefix string = ",\"jumps_wormhole\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.JumpsWormhole))
	}
	if in.WarpsHighSec != 0 {
		const prefix string = ",\"warps_high_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsHighSec))
	}
	if in.WarpsLowSec != 0 {
		const prefix string = ",\"warps_low_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsLowSec))
	}
	if in.WarpsNullSec != 0 {
		const prefix string = ",\"warps_null_sec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsNullSec))
	}
	if in.WarpsToBookmark != 0 {
		const prefix string = ",\"warps_to_bookmark\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsToBookmark))
	}
	if in.WarpsToCelestial != 0 {
		const prefix string = ",\"warps_to_celestial\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsToCelestial))
	}
	if in.WarpsToFleetMember != 0 {
		const prefix string = ",\"warps_to_fleet_member\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsToFleetMember))
	}
	if in.WarpsToScanResult != 0 {
		const prefix string = ",\"warps_to_scan_result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsToScanResult))
	}
	if in.WarpsWormhole != 0 {
		const prefix string = ",\"warps_wormhole\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WarpsWormhole))
	}
	out.RawByte('}')
}
func easyjson19558e8dDecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsSocial) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "add_contact_bad":
			out.AddContactBad = int64(in.Int64())
		case "add_contact_good":
			out.AddContactGood = int64(in.Int64())
		case "add_contact_high":
			out.AddContactHigh = int64(in.Int64())
		case "add_contact_horrible":
			out.AddContactHorrible = int64(in.Int64())
		case "add_contact_neutral":
			out.AddContactNeutral = int64(in.Int64())
		case "add_note":
			out.AddNote = int64(in.Int64())
		case "added_as_contact_bad":
			out.AddedAsContactBad = int64(in.Int64())
		case "added_as_contact_good":
			out.AddedAsContactGood = int64(in.Int64())
		case "added_as_contact_high":
			out.AddedAsContactHigh = int64(in.Int64())
		case "added_as_contact_horrible":
			out.AddedAsContactHorrible = int64(in.Int64())
		case "added_as_contact_neutral":
			out.AddedAsContactNeutral = int64(in.Int64())
		case "calendar_event_created":
			out.CalendarEventCreated = int64(in.Int64())
		case "chat_messages_alliance":
			out.ChatMessagesAlliance = int64(in.Int64())
		case "chat_messages_constellation":
			out.ChatMessagesConstellation = int64(in.Int64())
		case "chat_messages_corporation":
			out.ChatMessagesCorporation = int64(in.Int64())
		case "chat_messages_fleet":
			out.ChatMessagesFleet = int64(in.Int64())
		case "chat_messages_region":
			out.ChatMessagesRegion = int64(in.Int64())
		case "chat_messages_solarsystem":
			out.ChatMessagesSolarsystem = int64(in.Int64())
		case "chat_messages_warfaction":
			out.ChatMessagesWarfaction = int64(in.Int64())
		case "chat_total_message_length":
			out.ChatTotalMessageLength = int64(in.Int64())
		case "direct_trades":
			out.DirectTrades = int64(in.Int64())
		case "fleet_broadcasts":
			out.FleetBroadcasts = int64(in.Int64())
		case "fleet_joins":
			out.FleetJoins = int64(in.Int64())
		case "mails_received":
			out.MailsReceived = int64(in.Int64())
		case "mails_sent":
			out.MailsSent = int64(in.Int64())
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
func easyjson19558e8dEncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetCharactersCharacterIdStatsSocial) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AddContactBad != 0 {
		const prefix string = ",\"add_contact_bad\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactBad))
	}
	if in.AddContactGood != 0 {
		const prefix string = ",\"add_contact_good\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactGood))
	}
	if in.AddContactHigh != 0 {
		const prefix string = ",\"add_contact_high\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactHigh))
	}
	if in.AddContactHorrible != 0 {
		const prefix string = ",\"add_contact_horrible\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactHorrible))
	}
	if in.AddContactNeutral != 0 {
		const prefix string = ",\"add_contact_neutral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddContactNeutral))
	}
	if in.AddNote != 0 {
		const prefix string = ",\"add_note\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddNote))
	}
	if in.AddedAsContactBad != 0 {
		const prefix string = ",\"added_as_contact_bad\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactBad))
	}
	if in.AddedAsContactGood != 0 {
		const prefix string = ",\"added_as_contact_good\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactGood))
	}
	if in.AddedAsContactHigh != 0 {
		const prefix string = ",\"added_as_contact_high\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactHigh))
	}
	if in.AddedAsContactHorrible != 0 {
		const prefix string = ",\"added_as_contact_horrible\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactHorrible))
	}
	if in.AddedAsContactNeutral != 0 {
		const prefix string = ",\"added_as_contact_neutral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AddedAsContactNeutral))
	}
	if in.CalendarEventCreated != 0 {
		const prefix string = ",\"calendar_event_created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.CalendarEventCreated))
	}
	if in.ChatMessagesAlliance != 0 {
		const prefix string = ",\"chat_messages_alliance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesAlliance))
	}
	if in.ChatMessagesConstellation != 0 {
		const prefix string = ",\"chat_messages_constellation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesConstellation))
	}
	if in.ChatMessagesCorporation != 0 {
		const prefix string = ",\"chat_messages_corporation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesCorporation))
	}
	if in.ChatMessagesFleet != 0 {
		const prefix string = ",\"chat_messages_fleet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesFleet))
	}
	if in.ChatMessagesRegion != 0 {
		const prefix string = ",\"chat_messages_region\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesRegion))
	}
	if in.ChatMessagesSolarsystem != 0 {
		const prefix string = ",\"chat_messages_solarsystem\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesSolarsystem))
	}
	if in.ChatMessagesWarfaction != 0 {
		const prefix string = ",\"chat_messages_warfaction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatMessagesWarfaction))
	}
	if in.ChatTotalMessageLength != 0 {
		const prefix string = ",\"chat_total_message_length\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ChatTotalMessageLength))
	}
	if in.DirectTrades != 0 {
		const prefix string = ",\"direct_trades\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DirectTrades))
	}
	if in.FleetBroadcasts != 0 {
		const prefix string = ",\"fleet_broadcasts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FleetBroadcasts))
	}
	if in.FleetJoins != 0 {
		const prefix string = ",\"fleet_joins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.FleetJoins))
	}
	if in.MailsReceived != 0 {
		const prefix string = ",\"mails_received\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MailsReceived))
	}
	if in.MailsSent != 0 {
		const prefix string = ",\"mails_sent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MailsSent))
	}
	out.RawByte('}')
}
func easyjson19558e8dDecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsMining) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "drone_mine":
			out.DroneMine = int64(in.Int64())
		case "ore_arkonor":
			out.OreArkonor = int64(in.Int64())
		case "ore_bistot":
			out.OreBistot = int64(in.Int64())
		case "ore_crokite":
			out.OreCrokite = int64(in.Int64())
		case "ore_dark_ochre":
			out.OreDarkOchre = int64(in.Int64())
		case "ore_gneiss":
			out.OreGneiss = int64(in.Int64())
		case "ore_harvestable_cloud":
			out.OreHarvestableCloud = int64(in.Int64())
		case "ore_hedbergite":
			out.OreHedbergite = int64(in.Int64())
		case "ore_hemorphite":
			out.OreHemorphite = int64(in.Int64())
		case "ore_ice":
			out.OreIce = int64(in.Int64())
		case "ore_jaspet":
			out.OreJaspet = int64(in.Int64())
		case "ore_kernite":
			out.OreKernite = int64(in.Int64())
		case "ore_mercoxit":
			out.OreMercoxit = int64(in.Int64())
		case "ore_omber":
			out.OreOmber = int64(in.Int64())
		case "ore_plagioclase":
			out.OrePlagioclase = int64(in.Int64())
		case "ore_pyroxeres":
			out.OrePyroxeres = int64(in.Int64())
		case "ore_scordite":
			out.OreScordite = int64(in.Int64())
		case "ore_spodumain":
			out.OreSpodumain = int64(in.Int64())
		case "ore_veldspar":
			out.OreVeldspar = int64(in.Int64())
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
func easyjson19558e8dEncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdStatsMining) {
	out.RawByte('{')
	first := true
	_ = first
	if in.DroneMine != 0 {
		const prefix string = ",\"drone_mine\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.DroneMine))
	}
	if in.OreArkonor != 0 {
		const prefix string = ",\"ore_arkonor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreArkonor))
	}
	if in.OreBistot != 0 {
		const prefix string = ",\"ore_bistot\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreBistot))
	}
	if in.OreCrokite != 0 {
		const prefix string = ",\"ore_crokite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreCrokite))
	}
	if in.OreDarkOchre != 0 {
		const prefix string = ",\"ore_dark_ochre\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreDarkOchre))
	}
	if in.OreGneiss != 0 {
		const prefix string = ",\"ore_gneiss\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreGneiss))
	}
	if in.OreHarvestableCloud != 0 {
		const prefix string = ",\"ore_harvestable_cloud\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreHarvestableCloud))
	}
	if in.OreHedbergite != 0 {
		const prefix string = ",\"ore_hedbergite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreHedbergite))
	}
	if in.OreHemorphite != 0 {
		const prefix string = ",\"ore_hemorphite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreHemorphite))
	}
	if in.OreIce != 0 {
		const prefix string = ",\"ore_ice\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreIce))
	}
	if in.OreJaspet != 0 {
		const prefix string = ",\"ore_jaspet\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreJaspet))
	}
	if in.OreKernite != 0 {
		const prefix string = ",\"ore_kernite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreKernite))
	}
	if in.OreMercoxit != 0 {
		const prefix string = ",\"ore_mercoxit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreMercoxit))
	}
	if in.OreOmber != 0 {
		const prefix string = ",\"ore_omber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreOmber))
	}
	if in.OrePlagioclase != 0 {
		const prefix string = ",\"ore_plagioclase\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrePlagioclase))
	}
	if in.OrePyroxeres != 0 {
		const prefix string = ",\"ore_pyroxeres\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OrePyroxeres))
	}
	if in.OreScordite != 0 {
		const prefix string = ",\"ore_scordite\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreScordite))
	}
	if in.OreSpodumain != 0 {
		const prefix string = ",\"ore_spodumain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreSpodumain))
	}
	if in.OreVeldspar != 0 {
		const prefix string = ",\"ore_veldspar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.OreVeldspar))
	}
	out.RawByte('}')
}
