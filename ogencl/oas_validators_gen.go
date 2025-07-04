// Code generated by ogen, DO NOT EDIT.

package ogencl

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *Amount) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "currency",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s AmountCurrency) Validate() error {
	switch s {
	case "RUB":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *Payment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.CancellationDetails.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "cancellation_details",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.PaymentMethod.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "payment_method",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Confirmation.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "confirmation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *PaymentCancellationDetails) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Reason.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reason",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s PaymentCancellationDetailsReason) Validate() error {
	switch s {
	case "general_decline":
		return nil
	case "insufficient_funds":
		return nil
	case "rejected_by_payee":
		return nil
	case "rejected_by_timeout":
		return nil
	case "yoo_money_account_closed":
		return nil
	case "payment_article_number_not_found":
		return nil
	case "payment_basket_id_not_found":
		return nil
	case "payment_tru_code_not_found":
		return nil
	case "some_articles_already_refunded":
		return nil
	case "too_many_refunding_articles":
		return nil
	case "call_issuer":
		return nil
	case "canceled_by_merchant":
		return nil
	case "card_expired":
		return nil
	case "country_forbidden":
		return nil
	case "deal_expired":
		return nil
	case "expired_on_capture":
		return nil
	case "expired_on_confirmation":
		return nil
	case "fraud_suspected":
		return nil
	case "identification_required":
		return nil
	case "internal_timeout":
		return nil
	case "invalid_card_number":
		return nil
	case "invalid_csc":
		return nil
	case "issuer_unavailable":
		return nil
	case "payment_method_limit_exceeded":
		return nil
	case "payment_method_restricted":
		return nil
	case "permission_revoked":
		return nil
	case "unsupported_mobile_operator":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *PaymentConfirmation) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *PaymentConfirmationEmbedded) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s PaymentConfirmationEmbeddedType) Validate() error {
	switch s {
	case "embedded":
		return nil
	case "external":
		return nil
	case "redirect":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PaymentConfirmationType) Validate() error {
	switch s {
	case "embedded":
		return nil
	case "redirect":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *PaymentPaymentMethod) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Status.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s PaymentPaymentMethodStatus) Validate() error {
	switch s {
	case "pending":
		return nil
	case "active":
		return nil
	case "inactive":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s PaymentStatus) Validate() error {
	switch s {
	case "pending":
		return nil
	case "waiting_for_capture":
		return nil
	case "succeeded":
		return nil
	case "canceled":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *Receipt) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ReceiptItem) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        true,
			Max:           10,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.VatCode)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "vat_code",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.PaymentSubject.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "payment_subject",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.PaymentMode.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "payment_mode",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ReceiptItemPaymentMode) Validate() error {
	switch s {
	case "full_prepayment":
		return nil
	case "full_payment":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ReceiptItemPaymentSubject) Validate() error {
	switch s {
	case "commodity":
		return nil
	case "job":
		return nil
	case "service":
		return nil
	case "payment":
		return nil
	case "casino":
		return nil
	case "another":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *RefundPayment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ReqPayment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Confirmation.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "confirmation",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Receipt.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "receipt",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ReqPaymentConfirmation) Validate() error {
	switch s.Type {
	case PaymentConfirmationEmbeddedReqPaymentConfirmation:
		if err := s.PaymentConfirmationEmbedded.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *ReqRefundPayment) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "amount",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *V3PaymentsGetOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Items {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "items",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *YookassaError) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Code.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "code",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s YookassaErrorCode) Validate() error {
	switch s {
	case "invalid_request":
		return nil
	case "invalid_credentials":
		return nil
	case "forbidden":
		return nil
	case "not_found":
		return nil
	case "too_many_requests":
		return nil
	case "internal_server_error":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *YookassaErrorStatusCode) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "Response",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s YookassaErrorType) Validate() error {
	switch s {
	case "error":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *YookassaHookPostReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Type.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Event.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "event",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Object.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "object",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s YookassaHookPostReqEvent) Validate() error {
	switch s {
	case "payment.waiting_for_capture":
		return nil
	case "payment.succeeded":
		return nil
	case "payment.canceled":
		return nil
	case "refund.succeeded":
		return nil
	case "payout.succeeded":
		return nil
	case "payout.canceled":
		return nil
	case "deal.closed":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s YookassaHookPostReqObject) Validate() error {
	switch s.Type {
	case RefundPaymentYookassaHookPostReqObject:
		if err := s.RefundPayment.Validate(); err != nil {
			return err
		}
		return nil
	case PaymentYookassaHookPostReqObject:
		if err := s.Payment.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s YookassaHookPostReqType) Validate() error {
	switch s {
	case "notification":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
