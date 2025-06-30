package yookassa

import (
	"log/slog"
	"net/http"
	"net/http/httputil"

	"github.com/go-faster/errors"
	"github.com/quri-dev/yookassa/ogencl"
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

func CancellationDetailsString(detail ogencl.OptPaymentCancellationDetails) string {

	switch detail.Value.Reason.Value {
	case "fraud_suspected":
		return `Выплата заблокирована из-за подозрения в мошенничестве. Следует обратиться к инициатору отмены выплаты за уточнением подробностей или выбрать другой способ получения выплаты или другое платежное средство (например, другую банковскую карту).`
	case "general_decline":
		return `Ошибка. Причина не детализирована.`
	case "insufficient_funds":
		return `На балансе выплат не хватает денег для проведения выплаты. Пополните баланс и повторите запрос.`
	case "issuer_unavailable":
		return `Превышен лимит на разовое зачисление. Можно уменьшить размер выплаты, разбить сумму и сделать несколько выплат, выбрать другой способ получения выплат или другое платежное средство (например, другую банковскую карту).`
	case "periodic_limit_exceeded":
		return `Превышен лимит выплат за период времени (сутки, месяц). Следует выбрать другой способ получения выплаты или другое платежное средство (например, другую банковскую карту).`
	case "recipient_check_failed":
		return `Только для выплат с проверкой получателя. Получатель выплаты не прошел проверку: имя получателя не совпало с именем владельца счета, на который необходимо перевести деньги.`
	case "recipient_not_found":
		return `Для выплат через СБП: получатель не найден — в выбранном банке или платежном сервисе не найден счет, к которому привязан указанный номер телефона.`
	case "rejected_by_payee":
		return `Эмитент отклонил выплату по неизвестным причинам. Следует обратиться к эмитенту за уточнением подробностей или выбрать другой способ получения выплаты или другое платежное средство (например, другую банковскую карту).`
	case "self_employed_annual_limit_exceeded":
		return `Только для выплат самозанятым. Превышен лимит на годовой доход самозанятого.`

	}

	return string(detail.Value.Reason.Value)
}
