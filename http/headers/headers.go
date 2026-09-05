package headers

const (
    // Standard header names
    HeaderAccept            = "Accept"
    HeaderAcceptEncoding    = "Accept-Encoding"
    HeaderAcceptLanguage    = "Accept-Language"
    HeaderAuthorization     = "Authorization"
    HeaderCacheControl      = "Cache-Control"
    HeaderConnection        = "Connection"
    HeaderContentDisposition = "Content-Disposition"
    HeaderContentEncoding   = "Content-Encoding"
    HeaderContentLength     = "Content-Length"
    HeaderContentType       = "Content-Type"
    HeaderCookie            = "Cookie"
    HeaderETag              = "ETag"
    HeaderIfModifiedSince   = "If-Modified-Since"
    HeaderIfNoneMatch       = "If-None-Match"
    HeaderLastModified      = "Last-Modified"
    HeaderLocation          = "Location"
    HeaderOrigin            = "Origin"
    HeaderReferer           = "Referer"
    HeaderRetryAfter        = "Retry-After"
    HeaderSetCookie         = "Set-Cookie"
    HeaderUserAgent         = "User-Agent"
    HeaderWWWAuthenticate   = "WWW-Authenticate"
    HeaderXForwardedFor     = "X-Forwarded-For"
    HeaderXForwardedProto   = "X-Forwarded-Proto"
    HeaderXRequestID        = "X-Request-ID"
    HeaderXCorrelationID    = "X-Correlation-ID"
)

// Common Content-Type values
const (
    ContentTypeJSON          = "application/json"
    ContentTypeXML           = "application/xml"
    ContentTypeYAML          = "application/x-yaml"
    ContentTypeText          = "text/plain"
    ContentTypeHTML          = "text/html"
    ContentTypeForm          = "application/x-www-form-urlencoded"
    ContentTypeMultipart     = "multipart/form-data"
    ContentTypeOctetStream   = "application/octet-stream"
    ContentTypeEventStream   = "text/event-stream"
    ContentTypeProtobuf      = "application/protobuf"
    ContentTypeMsgPack       = "application/msgpack"
)

// Cache-Control directives
const (
    CacheControlNoCache     = "no-cache"
    CacheControlNoStore     = "no-store"
    CacheControlMaxAge      = "max-age"
    CacheControlMustRevalidate = "must-revalidate"
    CacheControlPublic      = "public"
    CacheControlPrivate     = "private"
)
