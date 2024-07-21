package yookassa

import (
	"log/slog"
	"net/http"
	"net/http/httputil"

	"github.com/go-faster/errors"
	"github.com/wildwind123/yookassa/ogencl"
)

type LoggingTransport struct {
	Logger *slog.Logger
	Level  slog.Level
}

func (s *LoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	bytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)
	// err is returned after dumping the response

	respBytes, _ := httputil.DumpResponse(resp, true)

	if s.Logger != nil {
		if resp.StatusCode != http.StatusOK {
			s.Logger.Error("wrong status code", slog.Int("code", resp.StatusCode), slog.String("res", string(respBytes)), slog.String("req_out", string(bytes)))
		}
		if s.Level == slog.LevelDebug {
			s.Logger.Info("response", slog.String("yookassa response", string(respBytes)))
			s.Logger.Info("response", slog.String("yookassa DumpRequestOut", string(bytes)))
		}
	}

	return resp, err
}

func YookassaError(err error) *ogencl.YookassaErrorStatusCode {
	if err == nil {
		return nil
	}
	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr == nil {
		return nil
	}
	unwrappedErr1 := errors.Unwrap(unwrappedErr)
	if unwrappedErr1 == nil {
		return nil
	}

	yookassaErr, ok := unwrappedErr1.(*ogencl.YookassaErrorStatusCode)
	if !ok {
		return nil
	}
	return yookassaErr
}
