package gocheck_matchers

import (
	. "gopkg.in/check.v1"
)

type IsTrueSuite struct {}

var _ = Suite(&IsTrueSuite{})

func (suite *IsTrueSuite) TestCheck(c *C) {
	c.Assert(true, IsTrue)
	c.Assert(false, Not(IsTrue))
}
