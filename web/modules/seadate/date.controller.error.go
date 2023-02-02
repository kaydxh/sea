package seadate

import (
	"errors"

	errors_ "github.com/kaydxh/golang/go/errors"
	v1 "github.com/kaydxh/sea/api/protoapi-spec/sea-date/v1"
	"github.com/kaydxh/sea/pkg/seadate/domain/date"
)

func DateError(err error, handled bool) (error, bool) {
	if handled {
		return err, true
	}

	if errors.Is(err, date.ErrInterval) {
		return errors_.Errorf(v1.SeaDateReasonEnum_INTERNAL, err.Error()), true
	}

	return err, false
}

func APIError(err error) error {
	if err == nil {
		return nil
	}
	return errors_.ErrorChain(DateError)(err, false)
}
