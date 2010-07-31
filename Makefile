include $(GOROOT)/src/Make.$(GOARCH)

TARG=GoDate
GOFILES=\
	TimeSpan.go\
	DateTime.go\
	

include $(GOROOT)/src/Make.pkg

