package yookassa

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/go-faster/jx"
	"github.com/joho/godotenv"
	"github.com/quri-dev/yookassa/ogencl"
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
	//t.Skip()
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

	r, err := ogenCl.V3PaymentsPost(ctx, &ogencl.ReqPayment{
		Amount: ogencl.Amount{
			Currency: ogencl.AmountCurrencyRUB,
			Value:    "100",
		},
		Confirmation: ogencl.NewOptReqPaymentConfirmation(ogencl.ReqPaymentConfirmation{
			Type: ogencl.PaymentConfirmationEmbeddedReqPaymentConfirmation,
			PaymentConfirmationEmbedded: ogencl.PaymentConfirmationEmbedded{
				Type: ogencl.PaymentConfirmationEmbeddedTypeEmbedded,
			},
		}),
		Capture:           ogencl.NewOptBool(true),
		SavePaymentMethod: ogencl.NewOptBool(true),
		Metadata: ogencl.NewOptMetadata(ogencl.Metadata{
			"user_id": jx.Raw("123"),
		}),
	}, ogencl.V3PaymentsPostParams{
		IdempotenceKey: "foo_456_asdfas21",
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

func TestInfo(t *testing.T) {
	b, _ := os.ReadFile("fixture/refund.success.json")

	d := ogencl.YookassaHookPostReq{}
	err := d.UnmarshalJSON(b)

	if err != nil {
		t.Error(err)
	}
}

func TestRefund(t *testing.T) {
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

	r, err := ogenCl.V3RefundsPost(ctx, &ogencl.ReqRefundPayment{
		Amount: ogencl.Amount{
			Currency: ogencl.AmountCurrencyRUB,
			Value:    "1",
		},
		PaymentID: "2e2e94dc-000f-5000-8000-19f75c59d8f2d",
	}, ogencl.V3RefundsPostParams{
		IdempotenceKey: "2e2e94dc-000f-5000-8000-19f75c59d8f2",
	})
	if err != nil {
		yookassaErr := YookassaError(err)
		if yookassaErr != nil {
			fmt.Println("yoo kassa err", yookassaErr)
		} else {
			fmt.Println("unknown err", err)
		}
	}

	_ = r
}

func TestPaymentsPost(t *testing.T) {
	//t.Skip()
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

	r, err := ogenCl.V3PaymentsPost(ctx, &ogencl.ReqPayment{
		Amount: ogencl.Amount{
			Currency: ogencl.AmountCurrencyRUB,
			Value:    "1",
		},
		Confirmation: ogencl.NewOptReqPaymentConfirmation(ogencl.ReqPaymentConfirmation{
			Type: ogencl.PaymentConfirmationEmbeddedReqPaymentConfirmation,
			PaymentConfirmationEmbedded: ogencl.PaymentConfirmationEmbedded{
				Type:      ogencl.PaymentConfirmationEmbeddedTypeRedirect,
				ReturnURL: ogencl.NewOptString("https://ya.ru"),
			},
		}),
		Capture:           ogencl.NewOptBool(true),
		SavePaymentMethod: ogencl.NewOptBool(true),
		Metadata: ogencl.NewOptMetadata(ogencl.Metadata{
			"user_id": jx.Raw("123"),
		}),
		Receipt: ogencl.NewOptReceipt(ogencl.Receipt{
			Items: []ogencl.ReceiptItem{
				{
					Description: "Привязка карты (будет возврат)",
					Amount: ogencl.Amount{
						Value:    "1",
						Currency: ogencl.AmountCurrencyRUB,
					},
					VatCode:        1,
					Quantity:       1,
					PaymentSubject: ogencl.NewOptReceiptItemPaymentSubject(ogencl.ReceiptItemPaymentSubjectAnother),
					PaymentMode:    ogencl.NewOptReceiptItemPaymentMode(ogencl.ReceiptItemPaymentModeFullPayment),
				},
			},
			Customer: ogencl.NewOptReceiptCustomer(ogencl.ReceiptCustomer{
				Email: ogencl.NewOptString("test@mail.ru"),
			}),
		}),
	}, ogencl.V3PaymentsPostParams{
		IdempotenceKey: "foo_456_asdfas26",
	})
	if err != nil {
		var e *ogencl.YookassaErrorStatusCode
		if errors.As(err, &e) {
			fmt.Println("yookassa error")
		}
		t.Error(err)
	}
	fmt.Println("response", r)

}
