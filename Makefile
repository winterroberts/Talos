.PHONY: build clean

TARGET = gameboy-advance
OUTPUT = bin/game.gba
MAIN = .

build:
	@mkdir -p bin
	tinygo build -o $(OUTPUT) -target=$(TARGET) $(MAIN)

clean:
	rm -f bin/*
