package date

import (
	"errors"

	errors_ "github.com/kaydxh/golang/go/errors"
	datev1_ "github.com/kaydxh/sea/api/protoapi-spec/date"
	"github.com/kaydxh/sea/pkg/sealet/domain/sealet"
)

func DateError(err error, handled bool) (error, bool) {
	if handled {
		return err, true
	}

	if errors.Is(err, sealet.ErrInterval) {
		return errors_.Errorf(datev1_.SeaDateReasonEnum_INTERNAL, err.Error()), true
	}

	return err, false
}

func APIError(err error) error {
	if err == nil {
		return nil
	}
	return errors_.ErrorChain(DateError)(err, false)
}
