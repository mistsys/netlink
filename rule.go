package netlink

import (
	"fmt"
	"net"
)

// Rule represents a netlink rule.
type Rule struct {
	Priority          int
	Family            int
	Table             int
	Mark              int
	Mask              int
	TunID             uint
	Goto              int
	Src               *net.IPNet
	Dst               *net.IPNet
	Flow              int
	IifName           string
	OifName           string
	SuppressIfgroup   int
	SuppressPrefixlen int
	Invert            bool
}

func (r Rule) String() string {
	var dst, iif, oif string
	if r.Dst != nil {
		dst = " to " + r.Dst.String()
	}
	if r.IifName != "" {
		iif = " iif " + r.IifName
	}
	if r.OifName != "" {
		oif = " oif " + r.OifName
	}
	src := " from all"
	if r.Src != nil {
		src = " from " + r.Src.String()
	}
	return fmt.Sprintf("ip rule %d:%s%s%s%s table %d", r.Priority, src, dst, iif, oif, r.Table)
}

// NewRule return empty rules.
func NewRule() *Rule {
	return &Rule{
		SuppressIfgroup:   -1,
		SuppressPrefixlen: -1,
		Priority:          -1,
		Mark:              -1,
		Mask:              -1,
		Goto:              -1,
		Flow:              -1,
	}
}
