package entity

import (
	"github.com/howood/moggiecollector/domain/model"
)

type MfaType model.MfaType

const (
	MfaTypeTOTP     MfaType = MfaType(model.MfaTypeTOTP)
	MfaTypeWebAuthn MfaType = MfaType(model.MfaTypeWebAuthn)
)
