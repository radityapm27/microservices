package resolver

import (
	"testing"

	mocksService "rpm/microservices/api/mocks/service"

	"github.com/stretchr/testify/assert"
)

const (
	snapshotDate = "2020-01-01"
	weeklyDate   = "2020-37"
	monthlyDate  = "2020-01"
	yearlyDate   = "2020"
	startDate    = snapshotDate
	endDate      = "2020-01-07"
	contractor   = "MTX"
)

func TestNew(t *testing.T) {

	t.Parallel()

	var (
		test   = assert.New(t)
		apiSvc = new(mocksService.Service)
	)

	t.Run("success new resolver", func(t *testing.T) {
		resolver := New(apiSvc)
		test.NotNil(resolver)
	})

}

func resetService(r *Resolver) *mocksService.Service {
	svc := new(mocksService.Service)
	r.service = svc
	return svc
}
