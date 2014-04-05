package api

import (
    "launchpad.net/gocheck"
    "time"
)

func (s *S) TestMountDirectoryFromTime(c *gocheck.C){
    time := time.Date(2009,time.November,15,21,40,03,0, time.Local)
    var ret = mountDirectoryPathFromTime(time)
    c.Assert(ret, gocheck.Equals, "2009/11/15/21/40/03/")
}
