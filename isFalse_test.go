package gocheck_matchers

import (
	. "gopkg.in/check.v1"
)

type IsFalseSuite struct {}

var _ = Suite(&IsFalseSuite{})

func (suite *IsFalseSuite) TestCheck(c *C) {
	c.Assert(true, Not(IsFalse))
	c.Assert(false, IsFalse)
}

func (suite *IsFalseSuite) TestCheckFailure(c *C) {
	_, err := IsFalse.Check([]interface{}{"test"}, []string{"v"})
	c.Assert(err, NotNil)
}
