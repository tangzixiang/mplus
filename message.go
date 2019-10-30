package mplus

import (
	"github.com/tangzixiang/mplus/message"
)

type MSGType = message.MSGType
type Message = message.Message
type Callback = message.Callback
type CallbackMessage = message.CallbackMessage

const (
	MSGLangZH = message.MSGLangZH
	MSGLangEN = message.MSGLangEN
)

var (
	Messages                                   = message.Messages
	NewMessage                                 = message.NewMessage
	NewCallbackMessage                         = message.NewCallbackMessage
	NewErrCodeMessage                          = message.NewErrCodeMessage
	SetDefaultLang                             = message.SetDefaultLang
	MessageStatusOK                            = message.MessageStatusOK
	MessageStatusCreated                       = message.MessageStatusCreated
	MessageStatusAccepted                      = message.MessageStatusAccepted
	MessageStatusNonAuthoritativeInfo          = message.MessageStatusNonAuthoritativeInfo
	MessageStatusNoContent                     = message.MessageStatusNoContent
	MessageStatusResetContent                  = message.MessageStatusResetContent
	MessageStatusPartialContent                = message.MessageStatusPartialContent
	MessageStatusMultiStatus                   = message.MessageStatusMultiStatus
	MessageStatusAlreadyReported               = message.MessageStatusAlreadyReported
	MessageStatusIMUsed                        = message.MessageStatusIMUsed
	MessageStatusMultipleChoices               = message.MessageStatusMultipleChoices
	MessageStatusMovedPermanently              = message.MessageStatusMovedPermanently
	MessageStatusFound                         = message.MessageStatusFound
	MessageStatusSeeOther                      = message.MessageStatusSeeOther
	MessageStatusNotModified                   = message.MessageStatusNotModified
	MessageStatusUseProxy                      = message.MessageStatusUseProxy
	MessageStatusTemporaryRedirect             = message.MessageStatusTemporaryRedirect
	MessageStatusPermanentRedirect             = message.MessageStatusPermanentRedirect
	MessageStatusBadRequest                    = message.MessageStatusBadRequest
	MessageStatusUnauthorized                  = message.MessageStatusUnauthorized
	MessageStatusPaymentRequired               = message.MessageStatusPaymentRequired
	MessageStatusForbidden                     = message.MessageStatusForbidden
	MessageStatusNotFound                      = message.MessageStatusNotFound
	MessageStatusMethodNotAllowed              = message.MessageStatusMethodNotAllowed
	MessageStatusNotAcceptable                 = message.MessageStatusNotAcceptable
	MessageStatusProxyAuthRequired             = message.MessageStatusProxyAuthRequired
	MessageStatusRequestTimeout                = message.MessageStatusRequestTimeout
	MessageStatusConflict                      = message.MessageStatusConflict
	MessageStatusGone                          = message.MessageStatusGone
	MessageStatusLengthRequired                = message.MessageStatusLengthRequired
	MessageStatusPreconditionFailed            = message.MessageStatusPreconditionFailed
	MessageStatusRequestEntityTooLarge         = message.MessageStatusRequestEntityTooLarge
	MessageStatusRequestURITooLong             = message.MessageStatusRequestURITooLong
	MessageStatusUnsupportedMediaType          = message.MessageStatusUnsupportedMediaType
	MessageStatusRequestedRangeNotSatisfiable  = message.MessageStatusRequestedRangeNotSatisfiable
	MessageStatusExpectationFailed             = message.MessageStatusExpectationFailed
	MessageStatusTeapot                        = message.MessageStatusTeapot
	MessageStatusMisdirectedRequest            = message.MessageStatusMisdirectedRequest
	MessageStatusUnprocessableEntity           = message.MessageStatusUnprocessableEntity
	MessageStatusLocked                        = message.MessageStatusLocked
	MessageStatusFailedDependency              = message.MessageStatusFailedDependency
	MessageStatusTooEarly                      = message.MessageStatusTooEarly
	MessageStatusUpgradeRequired               = message.MessageStatusUpgradeRequired
	MessageStatusPreconditionRequired          = message.MessageStatusPreconditionRequired
	MessageStatusTooManyRequests               = message.MessageStatusTooManyRequests
	MessageStatusRequestHeaderFieldsTooLarge   = message.MessageStatusRequestHeaderFieldsTooLarge
	MessageStatusUnavailableForLegalReasons    = message.MessageStatusUnavailableForLegalReasons
	MessageStatusInternalServerError           = message.MessageStatusInternalServerError
	MessageStatusNotImplemented                = message.MessageStatusNotImplemented
	MessageStatusBadGateway                    = message.MessageStatusBadGateway
	MessageStatusServiceUnavailable            = message.MessageStatusServiceUnavailable
	MessageStatusGatewayTimeout                = message.MessageStatusGatewayTimeout
	MessageStatusHTTPVersionNotSupported       = message.MessageStatusHTTPVersionNotSupported
	MessageStatusVariantAlsoNegotiates         = message.MessageStatusVariantAlsoNegotiates
	MessageStatusInsufficientStorage           = message.MessageStatusInsufficientStorage
	MessageStatusLoopDetected                  = message.MessageStatusLoopDetected
	MessageStatusNotExtended                   = message.MessageStatusNotExtended
	MessageStatusNetworkAuthenticationRequired = message.MessageStatusNetworkAuthenticationRequired
)
