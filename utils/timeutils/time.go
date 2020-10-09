package timeutils

import "time"

var CstZone *time.Location

func init() {
	CstZone = time.FixedZone("CST", 8*3600)
}
