package price

import (
	"math/big"
	"testing"

	"github.com/rbrick/bytes/util"
)

func TestEtherConversion(t *testing.T) {
	price := big.NewInt(1234257003245012000)
	want, _ := big.NewFloat(0).SetString("1.234257003245012")
	converted := util.ConvertHuman(price, 18)

	cf64, _ := converted.Float64()
	wf64, _ := want.Float64()
	if wf64 != cf64 {
		t.Errorf("expected: %f, got: %f\n", wf64, cf64)
	}
}
