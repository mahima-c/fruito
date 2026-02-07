package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error

type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

type DecodeResponseFunc func(context.Context, *http.Response) (response interface{}, err error)

type RequestFunc func(context.Context, *http.Request) (context.Context, error)

type ServerResponseFunc func(context.Context, http.ResponseWriter) context.Context

type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

type ServerFinalizerFunc func(ctx context.Context, code int, r *http.Request)

type Headerer interface {
	Headers() http.Header
}

type StatusCoder interface {
	StatusCode() int
}

type Middleware func(Endpoint) Endpoint

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func DefaultErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}

	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type Server struct {
	e            Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []RequestFunc
	after        []ServerResponseFunc
	errorEncoder ErrorEncoder
	finalizer    []ServerFinalizerFunc
}

type ServerOption func(*Server)

func NewServer(
	e Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
) *Server {
	s := &Server{
		e:            e,
		dec:          dec,
		enc:          enc,
		errorEncoder: DefaultErrorEncoder,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

type interceptingWriter struct {
	http.ResponseWriter
	code    int
	written int64
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var err error

	ctx := r.Context()

	bodyBytes := []byte{}

	contentType := r.Header.Get("Content-Type")

	if !strings.Contains(contentType, "multipart/form-data") {

		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("ctx", ctx, "error", err, "message", "error reading request body")
		}

		// reset the body to re-read
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	iw := &interceptingWriter{w, http.StatusOK, 0}
	if len(s.finalizer) > 0 {
		w = iw
	}

	r = r.WithContext(ctx)
	defer func() {
		if err != nil {
			s.errorEncoder(ctx, err, w)
		}
	}()

	// recover check to catch all panics
	// to be deffered after err encoder so in case of panic error is returned
	defer func() {
		if r := recover(); r != nil {
			internalErr := fmt.Sprintf("%v", r)
			errMsg := fmt.Sprintf(internalErr, "Something went wrong.")
			fmt.Println(errMsg)
		}
	}()

	for _, f := range s.before {
		ctx, err = f(ctx, r)
		if err != nil {
			return
		}
	}

	request, err := s.dec(ctx, r)
	if err != nil {
		return
	}

	response, err := s.e(ctx, request)
	if err != nil {
		return
	}

	for _, f := range s.after {
		ctx = f(ctx, w)
	}

	if err = s.enc(ctx, w, response); err != nil {
		return
	}

}

func ServerErrorEncoder(ee ErrorEncoder) ServerOption {
	return func(s *Server) { s.errorEncoder = ee }
}
