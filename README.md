
# WoSgetter

A program to export large amounts of metadata from Web of Science.

## Dependencies

This program depends on several things that are not in this repository:
1. Go compiler (and Go libraries) *(those are not needed if you download the binary release)*
2. Google Chrome (https://www.google.com/chrome/)
3. Chromedriver (http://chromedriver.chromium.org/downloads)

## Building

    git clone https://github.com/x0wllaar/WoSgetter
    cd WoSgetter
    go get .
    go build -o wosgetter.exe .

These commands will download the code, all the necessary Go libraries and build the executable. You will still have to download Chrome and the Chromedriver binary.

## Command line flags

    > .\wosgetter.exe -h
    Usage of wosgetter.exe:
      -chrome string