package yookassa

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/go-faster/jx"
	"github.com/joho/godotenv"
	"github.com/wildwind123/yookassa/ogencl"
)

type C struct{}

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("cant load dot env")
	}
}

func (c *C) BasicAuth(ctx context.Context, operationName string) (ogencl.BasicAuth, error) {
	return ogencl.BasicAuth{
		Username: os.Getenv("YOOKASSA_USER"),
		Password: os.Getenv("YOOKASSA_PASSWORD"),
	}, nil
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
		Capture:           ogencl.NewOptBool(true),
		SavePaymentMethod: ogencl.NewOptBool(true),
		Metadata: ogencl.NewOptPaymentMetadata(ogencl.PaymentMetadata{
			"user_id": jx.Raw("123"),
		}),
	}, ogencl.V3PaymentsPostParams{
		IdempotenceKey: "foo_456_2",
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

func TestPayments(t *testing.T) {
	t.Skip()
	client := http.Client{
		Transport: &LoggingTransport{
			Logger: slog.Default(),
			Level:  slog.LevelDebug,
		},
	}
	ogenCl, err := ogencl.NewClient("https://api.yookassa.ru", &C{}, ogencl.WithClient(&client))
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	r, err := ogenCl.V3PaymentsGet(ctx)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	fmt.Println("response", r)
}
