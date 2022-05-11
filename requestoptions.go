// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package opensecrets

//go:generate genopts --prefix=API --outfile=requestoptions.go "extraHeaders:map[string]string" "host:string" "customPayload:interface{}"

type APIOption func(*aPIOptionImpl)

type APIOptions interface {
	ExtraHeaders() map[string]string
	Host() string
	CustomPayload() interface{}
}

func APIExtraHeaders(extraHeaders map[string]string) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.extraHeaders = extraHeaders
	}
}
func APIExtraHeadersFlag(extraHeaders *map[string]string) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.extraHeaders = *extraHeaders
	}
}

func APIHost(host string) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.host = host
	}
}
func APIHostFlag(host *string) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.host = *host
	}
}

func APICustomPayload(customPayload interface{}) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.customPayload = customPayload
	}
}
func APICustomPayloadFlag(customPayload *interface{}) APIOption {
	return func(opts *aPIOptionImpl) {
		opts.customPayload = *customPayload
	}
}

type aPIOptionImpl struct {
	extraHeaders  map[string]string
	host          string
	customPayload interface{}
}

func (a *aPIOptionImpl) ExtraHeaders() map[string]string { return a.extraHeaders }
func (a *aPIOptionImpl) Host() string                    { return a.host }
func (a *aPIOptionImpl) CustomPayload() interface{}      { return a.customPayload }

func makeAPIOptionImpl(opts ...APIOption) *aPIOptionImpl {
	res := &aPIOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAPIOptions(opts ...APIOption) APIOptions {
	return makeAPIOptionImpl(opts...)
}
