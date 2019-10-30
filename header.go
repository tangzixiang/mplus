package mplus

import (
	"github.com/tangzixiang/mplus/header"
)

// 请求头常量
const (
	HeaderAccept                          = header.Accept
	HeaderAcceptCharset                   = header.AcceptCharset
	HeaderAccessControlAllowCredentials   = header.AccessControlAllowCredentials
	HeaderAccessControlAllowHeader        = header.AccessControlAllowHeaders
	HeaderAccessControlAllowMethods       = header.AccessControlAllowMethods
	HeaderAccessControlAllowOrigin        = header.AccessControlAllowOrigin
	HeaderAccessControlExposeHeader       = header.AccessControlExposeHeaders
	HeaderAccessControlMaxAge             = header.AccessControlMaxAge
	HeaderAccessControlRequestHeader      = header.AccessControlRequestHeaders
	HeaderAccessControlRequestMethod      = header.AccessControlRequestMethod
	HeaderAcceptEncoding                  = header.AcceptEncoding
	HeaderAcceptLanguage                  = header.AcceptLanguage
	HeaderAcceptRanges                    = header.AcceptRanges
	HeaderAllow                           = header.Allow
	HeaderAge                             = header.Age
	HeaderAltSvc                          = header.AltSvc
	HeaderAuthorization                   = header.Authorization
	HeaderCacheControl                    = header.CacheControl
	HeaderCc                              = header.Cc
	HeaderClearSiteData                   = header.ClearSiteData
	HeaderConnection                      = header.Connection
	HeaderContentType                     = header.ContentType
	HeaderContentLocation                 = header.ContentLocation
	HeaderContentRange                    = header.ContentRange
	HeaderContentID                       = header.ContentID
	HeaderContentDisposition              = header.ContentDisposition
	HeaderContentLanguage                 = header.ContentLanguage
	HeaderContentLength                   = header.ContentLength
	HeaderContentEncoding                 = header.ContentEncoding
	HeaderContentTransferEncoding         = header.ContentTransferEncoding
	HeaderCookie                          = header.Cookie
	HeaderCrossOriginResourcePolicy       = header.CrossOriginResourcePolicy
	HeaderContentSecurityPolicyReportOnly = header.ContentSecurityPolicyReportOnly
	HeaderContentSecurityPolicy           = header.ContentSecurityPolicy
	HeaderDNS                             = header.DNS
	HeaderDate                            = header.Date
	HeaderDNT                             = header.DNT
	HeaderDigest                          = header.Digest
	HeaderDkimSignature                   = header.DkimSignature
	HeaderEtag                            = header.Etag
	HeaderEarlyData                       = header.EarlyData
	HeaderExpect                          = header.Expect
	HeaderExpectCT                        = header.ExpectCT
	HeaderExpires                         = header.Expires
	HeaderFeaturePolicy                   = header.FeaturePolicy
	HeaderFrom                            = header.From
	HeaderForwarded                       = header.Forwarded
	HeaderHost                            = header.Host
	HeaderIfUnmodifiedSince               = header.IfUnmodifiedSince
	HeaderIfModifiedSince                 = header.IfModifiedSince
	HeaderIfMatch                         = header.IfMatch
	HeaderIfRange                         = header.IfRange
	HeaderIfNoneMatch                     = header.IfNoneMatch
	HeaderInReplyTo                       = header.InReplyTo
	HeaderKeepAlive                       = header.KeepAlive
	HeaderLargeAllocation                 = header.LargeAllocation
	HeaderLastModified                    = header.LastModified
	HeaderLocation                        = header.Location
	HeaderMessageID                       = header.MessageID
	HeaderMimeVersion                     = header.MimeVersion
	HeaderOrigin                          = header.Origin
	HeaderPublicKeyPinsReportOnly         = header.PublicKeyPinsReportOnly
	HeaderPublicKeyPins                   = header.PublicKeyPins
	HeaderProxyAuthorization              = header.ProxyAuthorization
	HeaderProxyAuthenticate               = header.ProxyAuthenticate
	HeaderPragma                          = header.Pragma
	HeaderRange                           = header.Range
	HeaderReferer                         = header.Referer
	HeaderRetryAfter                      = header.RetryAfter
	HeaderReferrerPolicy                  = header.ReferrerPolicy
	HeaderReceived                        = header.Received
	HeaderReturnPath                      = header.ReturnPath
	HeaderSaveData                        = header.SaveData
	HeaderServer                          = header.Server
	HeaderSecWebSocketAccept              = header.SecWebSocketAccept
	HeaderServerTiming                    = header.ServerTiming
	HeaderSetCookie                       = header.SetCookie
	HeaderSubject                         = header.Subject
	HeaderStrictTransportSecurity         = header.StrictTransportSecurity
	HeaderSourceMap                       = header.SourceMap
	HeaderTE                              = header.TE
	HeaderTimingAllowOrigin               = header.TimingAllowOrigin
	HeaderTk                              = header.Tk
	HeaderTrailer                         = header.Trailer
	HeaderTransferEncoding                = header.TransferEncoding
	HeaderTo                              = header.To
	HeaderUserAgent                       = header.UserAgent
	HeaderUpgradeInsecureRequests         = header.UpgradeInsecureRequests
	HeaderVia                             = header.Via
	HeaderVary                            = header.Vary
	HeaderWWWAuthenticate                 = header.WWWAuthenticate
	HeaderWantDigest                      = header.WantDigest
	HeaderWarning                         = header.Warning
	HeaderForwardedHost                   = header.ForwardedHost
	HeaderForwardedProto                  = header.ForwardedProto
	HeaderFrameOptions                    = header.FrameOptions
	HeaderXSSProtection                   = header.XSSProtection
	HeaderContentTypeOptions              = header.ContentTypeOptions
	HeaderDNSPrefetchControl              = header.DNSPrefetchControl
	HeaderPoweredBy                       = header.PoweredBy
	HeaderImforwards                      = header.Imforwards
	HeaderRequestID                       = header.RequestID
	HeaderForwardedFor                    = header.ForwardedFor
	HeaderRealIP                          = header.RealIP
	HeaderAppEngineRemoteAddr             = header.AppEngineRemoteAddr
)

const (
	ContentTypeJSON              = header.ContentTypeJSON
	ContentTypeForm              = header.ContentTypeForm
	ContentTypeText              = header.ContentTypeText
	ContentTypeXML               = header.ContentTypeXML
	ContentTypeStream            = header.ContentTypeStream
	ContentTypeHTML              = header.ContentTypeHTML
	ContentTypeXML2              = header.ContentTypeXML2
	ContentTypePlain             = header.ContentTypePlain
	ContentTypeMultipartPOSTForm = header.ContentTypeMultipartPOSTForm
	ContentTypePROTOBUF          = header.ContentTypePROTOBUF
	ContentTypeMSGPACK           = header.ContentTypeMSGPACK
	ContentTypeMSGPACK2          = header.ContentTypeMSGPACK2
)

// 请求头分割字符
const (
	SplitSepBlankSpace = header.SplitSepBlankSpace
	SplitSepComma      = header.SplitSepComma
	SplitSepSemicolon  = header.SplitSepSemicolon
)

var (
	GetHeader                  = header.GetHeader
	GetHeaderValues            = header.GetHeaderValues
	SplitHeader                = header.SplitHeader
	GetHeaderRequestID         = header.GetHeaderRequestID
	SetRequestHeader           = header.SetRequestHeader
	SetRequestHeaderIf         = header.SetRequestHeaderIf
	SetRequestHeaders          = header.SetRequestHeaders
	SetRequestHeadersIf        = header.SetRequestHeadersIf
	AddRequestHeader           = header.AddRequestHeader
	AddRequestHeaderIf         = header.AddRequestHeaderIf
	AddRequestHeaders          = header.AddRequestHeaders
	AddRequestHeadersIf        = header.AddRequestHeadersIf
	GetResponseHeader          = header.GetResponseHeader
	GetResponseHeaderValues    = header.GetResponseHeaderValues
	SetResponseHeader          = header.SetResponseHeader
	SetResponseHeaderIf        = header.SetResponseHeaderIf
	SetResponseHeaders         = header.SetResponseHeaders
	SetResponseHeadersIf       = header.SetResponseHeadersIf
	AddResponseHeader          = header.AddResponseHeader
	AddResponseHeaderIf        = header.AddResponseHeaderIf
	AddResponseHeaders         = header.AddResponseHeaders
	AddResponseHeadersIf       = header.AddResponseHeadersIf
	SetRequestHeaderRequestID  = header.SetRequestHeaderRequestID
	SetResponseHeaderRequestID = header.SetResponseHeaderRequestID
	GetClientIP                = header.GetClientIP
)
