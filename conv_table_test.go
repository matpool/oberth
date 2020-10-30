package sqltableconv

import (
	"fmt"
	"testing"
)

func TestConvTable(t *testing.T) {
	f := func(s string) string {
		return s + "_change"
	}
	SQL := `SELECT abc.* FROM abc JOIN (SELECT * FROM xyz) AS zyx ON abc.id = zyx.id`
	SQLExpect := `select abc_change.* from abc_change join (select * from xyz_change) as zyx_change on abc_change.id = zyx_change.id`
	sc, err := ConvTable(SQL, f)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sc)
	if sc != SQLExpect {
		t.Error("conv table failed")
	}
}
