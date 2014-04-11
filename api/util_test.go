package api

import (
    "launchpad.net/gocheck"
    "time"
    "os"
)

func (s *S) TestMountDirectoryFromTime(c *gocheck.C){
    time := time.Date(2009,time.November,15,21,40,03,0, time.Local)
    var ret = mountDirectoryPathFromTime(time)
    c.Assert(ret, gocheck.Equals, "2009/11/15/21/40/")
}

func (s *S) TestCreateDir(c *gocheck.C){
    err := CreateDir("tmp/test")
    c.Assert(err, gocheck.IsNil)

    finfo, err := os.Stat("tmp/test")

    c.Assert(err, gocheck.IsNil)
    c.Assert(finfo.IsDir(), gocheck.Equals, true)
    c.Assert(finfo.Mode().String(), gocheck.Equals, "drwxr-xr-x")

    os.RemoveAll("tmp/test")
}
