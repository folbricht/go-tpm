package tpm2

import (
	"fmt"

	"github.com/google/go-tpm/tpmutil"
)

type (
	RCFmt0  uint8 // Format 0 error codes
	RCFmt1  uint8 // Format 1 error codes
	RCWarn  uint8 // Warning codes
	RcIndex uint8 // Indexes for arguments, handles and sessions in errors
)

// Format 0 error codes.
const (
	RCInitialize      RCFmt0 = 0x00
	RCFailure                = 0x01
	RCSequence               = 0x03
	RCPrivate                = 0x0B
	RCHMAC                   = 0x19
	RCDisabled               = 0x20
	RCExclusive              = 0x21
	RCAuthType               = 0x24
	RCAuthMissing            = 0x25
	RCPolicy                 = 0x26
	RCPCR                    = 0x27
	RCPCRChanged             = 0x28
	RCUpgrade                = 0x2D
	RCTooManyContexts        = 0x2E
	RCAuthUnavailable        = 0x2F
	RCReboot                 = 0x30
	RCUnbalanced             = 0x31
	RCCommandSize            = 0x42
	RCCommandCode            = 0x43
	RCAuthSize               = 0x44
	RCAuthContext            = 0x45
	RCNVRange                = 0x46
	RCNVSize                 = 0x47
	RCNVLocked               = 0x48
	RCNVAuthorization        = 0x49
	RCNVUninitialized        = 0x4A
	RCNVSpace                = 0x4B
	RCNVDefined              = 0x4C
	RCBadContext             = 0x50
	RCCPHash                 = 0x51
	RCParent                 = 0x52
	RCNeedsTest              = 0x53
	RCNoResult               = 0x54
	RCSensitive              = 0x55
)

var fmt0Msg = map[RCFmt0]string{
	RCInitialize:      "TPM not initialized by TPM2_Startup or already initialized",
	RCFailure:         "commands not being accepted because of a TPM failure",
	RCSequence:        "improper use of a sequence handle",
	RCPrivate:         "not currently used",
	RCHMAC:            "not currently used",
	RCDisabled:        "the command is disabled",
	RCExclusive:       "command failed because audit sequence required exclusivity",
	RCAuthType:        "authorization handle is not correct for command",
	RCAuthMissing:     "5 command requires an authorization session for handle and it is not present",
	RCPolicy:          "policy failure in math operation or an invalid authPolicy value",
	RCPCR:             "PCR check fail",
	RCPCRChanged:      "PCR have changed since checked",
	RCUpgrade:         "TPM is in field upgrade mode unless called via TPM2_FieldUpgradeData(), then it is not in field upgrade mode",
	RCTooManyContexts: "context ID counter is at maximum",
	RCAuthUnavailable: "authValue or authPolicy is not available for selected entity",
	RCReboot:          "a _TPM_Init and Startup(CLEAR) is required before the TPM can resume operation",
	RCUnbalanced:      "the protection algorithms (hash and symmetric) are not reasonably balanced; the digest size of the hash must be larger than the key size of the symmetric algorithm",
	RCCommandSize:     "command commandSize value is inconsistent with contents of the command buffer; either the size is not the same as the octets loaded by the hardware interface layer or the value is not large enough to hold a command header",
	RCCommandCode:     "command code not supported",
	RCAuthSize:        "the value of authorizationSize is out of range or the number of octets in the Authorization Area is greater than required",
	RCAuthContext:     "use of an authorization session with a context command or another command that cannot have an authorization session",
	RCNVRange:         "NV offset+size is out of range",
	RCNVSize:          "Requested allocation size is larger than allowed",
	RCNVLocked:        "NV access locked",
	RCNVAuthorization: "NV access authorization fails in command actions",
	RCNVUninitialized: "an NV Index is used before being initialized or the state saved by TPM2_Shutdown(STATE) could not be restored",
	RCNVSpace:         "insufficient space for NV allocation",
	RCNVDefined:       "NV Index or persistent object already defined",
	RCBadContext:      "context in TPM2_ContextLoad() is not valid",
	RCCPHash:          "cpHash value already set or not correct for use",
	RCParent:          "handle for parent is not a valid parent",
	RCNeedsTest:       "some function needs testing",
	RCNoResult:        "returned when an internal function cannot process a request due to an unspecified problem; this code is usually related to invalid parameters that are not properly filtered by the input unmarshaling code",
	RCSensitive:       "the sensitive area did not unmarshal correctly after decryption",
}

// Format 1 error codes.
const (
	RCAsymmetric   = 0x01
	RCAttributes   = 0x02
	RCHash         = 0x03
	RCValue        = 0x04
	RCHierarchy    = 0x05
	RCKeySize      = 0x07
	RCMGF          = 0x08
	RCMode         = 0x09
	RCType         = 0x0A
	RCHandle       = 0x0B
	RCKDF          = 0x0C
	RCRange        = 0x0D
	RCAuthFail     = 0x0E
	RCNonce        = 0x0F
	RCPP           = 0x10
	RCScheme       = 0x12
	RCSize         = 0x15
	RCSymmetric    = 0x16
	RCTag          = 0x17
	RCSelector     = 0x18
	RCInsufficient = 0x1A
	RCSignature    = 0x1B
	RCKey          = 0x1C
	RCPolicyFail   = 0x1D
	RCIntegrity    = 0x1F
	RCTicket       = 0x20
	RCReservedBits = 0x21
	RCBadAuth      = 0x22
	RCExpired      = 0x23
	RCPolicyCC     = 0x24
	RCBinding      = 0x25
	RCCurve        = 0x26
	RCECCPoint     = 0x27
)

var fmt1Msg = map[RCFmt1]string{
	RCAsymmetric:   "asymmetric algorithm not supported or not correct",
	RCAttributes:   "inconsistent attributes",
	RCHash:         "hash algorithm not supported or not appropriate",
	RCValue:        "value is out of range or is not correct for the context",
	RCHierarchy:    "hierarchy is not enabled or is not correct for the use",
	RCKeySize:      "key size is not supported",
	RCMGF:          "mask generation function not supported",
	RCMode:         "mode of operation not supported",
	RCType:         "the type of the value is not appropriate for the use",
	RCHandle:       "the handle is not correct for the use",
	RCKDF:          "unsupported key derivation function or function not appropriate for use",
	RCRange:        "value was out of allowed range",
	RCAuthFail:     "the authorization HMAC check failed and DA counter incremented",
	RCNonce:        "invalid nonce size or nonce value mismatch",
	RCPP:           "authorization requires assertion of PP",
	RCScheme:       "unsupported or incompatible scheme",
	RCSize:         "structure is the wrong size",
	RCSymmetric:    "unsupported symmetric algorithm or key size, or not appropriate for instance",
	RCTag:          "incorrect structure tag",
	RCSelector:     "union selector is incorrect",
	RCInsufficient: "the TPM was unable to unmarshal a value because there were not enough octets in the input buffer",
	RCSignature:    "the signature is not valid",
	RCKey:          "key fields are not compatible with the selected use",
	RCPolicyFail:   "a policy check failed",
	RCIntegrity:    "integrity check failed",
	RCTicket:       "invalid ticket",
	RCReservedBits: "reserved bits not set to zero as required",
	RCBadAuth:      "authorization failure without DA implications",
	RCExpired:      "the policy has expired",
	RCPolicyCC:     "the commandCode in the policy is not the commandCode of the command or the command code in a policy command references a command that is not implemented",
	RCBinding:      "public and sensitive portions of an object are not cryptographically bound",
	RCCurve:        "curve not supported",
	RCECCPoint:     "point is not on the required curve",
}

// Warning codes.
const (
	RCContextGap     RCWarn = 0x01
	RCObjectMemory          = 0x02
	RCSessionMemory         = 0x03
	RCMemory                = 0x04
	RCSessionHandles        = 0x05
	RCObjectHandles         = 0x06
	RCLocality              = 0x07
	RCYielded               = 0x08
	RCCanceled              = 0x09
	RCTesting               = 0x0A
	RCReferenceH0           = 0x10
	RCReferenceH1           = 0x11
	RCReferenceH2           = 0x12
	RCReferenceH3           = 0x13
	RCReferenceH4           = 0x14
	RCReferenceH5           = 0x15
	RCReferenceH6           = 0x16
	RCReferenceS0           = 0x18
	RCReferenceS1           = 0x19
	RCReferenceS2           = 0x1A
	RCReferenceS3           = 0x1B
	RCReferenceS4           = 0x1C
	RCReferenceS5           = 0x1D
	RCReferenceS6           = 0x1E
	RCNVRate                = 0x20
	RCLockout               = 0x21
	RCRetry                 = 0x22
	RCNVUnavailable         = 0x23
)

var warnMsg = map[RCWarn]string{
	RCContextGap:     "gap for context ID is too large",
	RCObjectMemory:   "out of memory for object contexts",
	RCSessionMemory:  "out of memory for session contexts",
	RCMemory:         "out of shared object/session memory or need space for internal operations",
	RCSessionHandles: "out of session handles",
	RCObjectHandles:  "out of object handles",
	RCLocality:       "bad locality",
	RCYielded:        "the TPM has suspended operation on the command; forward progress was made and the command may be retried",
	RCCanceled:       "the command was canceled",
	RCTesting:        "TPM is performing self-tests",
	RCReferenceH0:    "the 1st handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH1:    "the 2nd handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH2:    "the 3rd handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH3:    "the 4th handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH4:    "the 5th handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH5:    "the 6th handle in the handle area references a transient object or session that is not loaded",
	RCReferenceH6:    "the 7th handle in the handle area references a transient object or session that is not loaded",
	RCReferenceS0:    "the 1st authorization session handle references a session that is not loaded",
	RCReferenceS1:    "the 2nd authorization session handle references a session that is not loaded",
	RCReferenceS2:    "the 3rd authorization session handle references a session that is not loaded",
	RCReferenceS3:    "the 4th authorization session handle references a session that is not loaded",
	RCReferenceS4:    "the 5th authorization session handle references a session that is not loaded",
	RCReferenceS5:    "the 6th authorization session handle references a session that is not loaded",
	RCReferenceS6:    "the 7th authorization session handle references a session that is not loaded",
	RCNVRate:         "the TPM is rate-limiting accesses to prevent wearout of NV",
	RCLockout:        "authorizations for objects subject to DA protection are not allowed at this time because the TPM is in DA lockout mode",
	RCRetry:          "the TPM was not able to start the command",
	RCNVUnavailable:  "the command may require writing of NV and NV is not current accessible",
}

// Indexes for arguments, handles and sessions.
const (
	RC1 RcIndex = 0x01
	RC2         = 0x02
	RC3         = 0x03
	RC4         = 0x04
	RC5         = 0x05
	RC6         = 0x06
	RC7         = 0x07
	RC8         = 0x08
	RC9         = 0x09
	RCA         = 0x0A
	RCB         = 0x0B
	RCC         = 0x0C
	RCD         = 0x0D
	RCE         = 0x0E
	RCF         = 0x0F
)

const unknownCode = "unknown error code"

// Error is returned for all Format 0 errors from the TPM. It is used for general
// errors not specific to a parameter, handle or session.
type Error struct {
	Code RCFmt0
}

func (e Error) Error() string {
	msg := fmt0Msg[e.Code]
	if msg == "" {
		msg = unknownCode
	}
	return fmt.Sprintf("error code 0x%x : %s", e.Code, msg)
}

// VendorError represents a vendor-specific error response. These types of responses
// are not decoded and Code contains the complete response code.
type VendorError struct {
	Code uint32
}

func (e VendorError) Error() string {
	return fmt.Sprintf("vendor error code 0x%x", e.Code)
}

// Warning is typically used to report transient errors.
type Warning struct {
	Code RCWarn
}

func (w Warning) Error() string {
	msg := warnMsg[w.Code]
	if msg == "" {
		msg = unknownCode
	}
	return fmt.Sprintf("warning code 0x%x : %s", w.Code, msg)
}

// ParameterError describes an error related to a parameter, and the parameter number.
type ParameterError struct {
	Code      RCFmt1
	Parameter RcIndex
}

func (e ParameterError) Error() string {
	msg := fmt1Msg[e.Code]
	if msg == "" {
		msg = unknownCode
	}
	return fmt.Sprintf("parameter %d, error code 0x%x : %s", e.Parameter, e.Code, msg)
}

// HandleError describes an error related to a handle, and the handle number.
type HandleError struct {
	Code   RCFmt1
	Handle RcIndex
}

func (e HandleError) Error() string {
	msg := fmt1Msg[e.Code]
	if msg == "" {
		msg = unknownCode
	}
	return fmt.Sprintf("handle %d, error code 0x%x : %s", e.Handle, e.Code, msg)
}

// SessionError describes an error related to a session, and the session number.
type SessionError struct {
	Code    RCFmt1
	Session RcIndex
}

func (e SessionError) Error() string {
	msg := fmt1Msg[e.Code]
	if msg == "" {
		msg = unknownCode
	}
	return fmt.Sprintf("session %d, error code 0x%x : %s", e.Session, e.Code, msg)
}

// Decode a TPM2 response code and return the appropriate error. Logic
// according to the "Response Code Evaluation" chart in Part 1 of the TPM 2.0
// spec.
func decodeResponse(code tpmutil.ResponseCode) error {
	if code == tpmutil.RCSuccess {
		return nil
	}
	if code&0x180 == 0 { // Bits 7:8 == 0 is a TPM1 error
		return fmt.Errorf("response status 0x%x", code)
	}
	if code&0x80 == 0 { // Bit 7 unset
		if code&0x400 > 0 { // Bit 10 set, vendor specific code
			return VendorError{uint32(code)}
		}
		if code&0x800 > 0 { // Bit 11 set, warning with code in bit 0:6
			return Warning{RCWarn(code & 0x7f)}
		}
		// error with code in bit 0:6
		return Error{RCFmt0(code & 0x7f)}
	}
	if code&0x40 > 0 { // Bit 6 set, code in 0:5, parameter number in 8:11
		return ParameterError{RCFmt1(code & 0x3f), RcIndex((code & 0xf00) >> 8)}
	}
	if code&0x800 == 0 { // Bit 11 unset, code in 0:5, handle in 8:10
		return HandleError{RCFmt1(code & 0x3f), RcIndex((code & 0x700) >> 8)}
	}
	// Code in 0:5, Session in 8:10
	return SessionError{RCFmt1(code & 0x3f), RcIndex((code & 0x700) >> 8)}
}
