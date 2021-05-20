// Code generated by goa v3.4.2, DO NOT EDIT.
//
// updown HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/upload_download/design -o
// $(GOPATH)/src/goa.design/examples/upload_download

package server

import (
	"context"
	"net/http"
	"strconv"

	updown "goa.design/examples/upload_download/gen/updown"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadResponse returns an encoder for responses returned by the updown
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the updown upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			dir         string
			contentType string
			err         error

			params = mux.Vars(r)
		)
		dir = params["dir"]
		contentTypeRaw := r.Header.Get("Content-Type")
		if contentTypeRaw != "" {
			contentType = contentTypeRaw
		} else {
			contentType = "multipart/form-data; boundary=goa"
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("contentType", contentType, "multipart/[^;]+; boundary=.+"))
		if err != nil {
			return nil, err
		}
		payload := NewUploadPayload(dir, contentType)

		return payload, nil
	}
}

// EncodeUploadError returns an encoder for errors returned by the upload
// updown endpoint.
func EncodeUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid_media_type":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadInvalidMediaTypeResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "invalid_multipart_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadInvalidMultipartRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDownloadResponse returns an encoder for responses returned by the
// updown download endpoint.
func EncodeDownloadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*updown.DownloadResult)
		val := res.Length
		lengths := strconv.FormatInt(val, 10)
		w.Header().Set("Content-Length", lengths)
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDownloadRequest returns a decoder for requests sent to the updown
// download endpoint.
func DecodeDownloadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			filename string

			params = mux.Vars(r)
		)
		filename = params["filename"]
		payload := filename

		return payload, nil
	}
}

// EncodeDownloadError returns an encoder for errors returned by the download
// updown endpoint.
func EncodeDownloadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid_file_path":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadInvalidFilePathResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
