
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
            Path to Chrome(ium) binary (default "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe")
      -driver string
            Path to Chromedriver binary (default "./chromedriver")
      -filen int
            File number to start with (default 1)
      -from int
            What record to start from (default 1)
      -o string
            Where to save the reports (default ".")
      -to int
            What record to end with

## How to use

1. Install Google Chrome (or Chromium) and download the Chromedriver binary. 
2. Start the program with the relevant command line flags. 
3. A Chrome window will open.
4. In this window, navigate to the WoS page with the relevant results
5. In the command line window, press Enter to start the download

### Continuing interrupted downloads