include $(GOROOT)/src/Make.inc

TARG=dealmap
GOFILES=\
	dealmap.go\
	http.go

include $(GOROOT)/src/Make.pkg

