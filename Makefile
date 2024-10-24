.PHONY: default

EXE= 

ifeq ($(OS),Windows_NT)
	EXE=.exe
endif

.PHONY: udpsender
udpsender:
	go build -o udpsender$(EXE) main.go

.PHONY: clean
clean:
	rm -f udpsender$(EXE)
