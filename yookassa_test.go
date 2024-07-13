package yookassa

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"testing"

	"github.com/wildwind123/yookassa/ogencl"
)

type C struct{}

func (c *C) BasicAuth(ctx context.Context, operationName string) (ogencl.BasicAuth, error) {
	return ogencl.BasicAuth{
		// Username: "361728",
		// Password: "test_GSlobMkMJ_amUemttt3epjBLx_lUKrhFpYzkvQs_pE0",
	}, nil
}

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

func TestXxx(t *testing.T) {
	t.Skip()
	client := http.Client{
		Transport: &LoggingTransport{
			Logger: slog.Default(),
			Level:  slog.LevelError,
		},
	}
	ogenCl, err := ogencl.NewClient("https://api.yookassa.ru", &C{}, ogencl.WithClient(&client))
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	r, err := ogenCl.V3PaymentsPost(ctx, &ogencl.Payment{
		Amount: ogencl.PaymentAmount{
			Currency: ogencl.PaymentAmountCurrencyRUB,
			Value:    "100",
		},
		Confirmation: ogencl.NewOptPaymentConfirmation(ogencl.PaymentConfirmation{
			Type: ogencl.PaymentConfirmationEmbeddedPaymentConfirmation,
			PaymentConfirmationEmbedded: ogencl.PaymentConfirmationEmbedded{
				Type: ogencl.PaymentConfirmationEmbeddedTypeEmbedded,
			},
		}),
		Capture: ogencl.NewOptBool(true),
		// Description:       ogencl.OptString{},
		SavePaymentMethod: ogencl.NewOptBool(true),
	}, ogencl.V3PaymentsPostParams{
		IdempotenceKey: "asdfas",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("response", r)

}

func TestPaymentInfo(t *testing.T) {
	t.Skip()
	client := http.Client{
		Transport: &LoggingTransport{
			Logger: slog.Default(),
			Level:  slog.LevelError,
		},
	}
	ogenCl, err := ogencl.NewClient("https://api.yookassa.ru", &C{}, ogencl.WithClient(&client))
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	r, err := ogenCl.V3PaymentsPaymentIDGet(ctx, ogencl.V3PaymentsPaymentIDGetParams{
		PaymentID: "2e24435b-000f-5000-9000-1eae900121b6",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("response", r)
}
