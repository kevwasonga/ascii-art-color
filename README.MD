## ASCII Banner Printer
## Overview

The ASCII Banner Printer is a Go program that takes input text as a string and converts it into a banner-like format using predefined ASCII art characters from a specified text file.
Features:

    * Reads banner definitions from a .txt file.
    * Converts input text into ASCII art characters.
    * Supports multiline input through escaped newline characters (\n).

## Prerequisites


    To run this program make sure you have GO 1.16 or above installed on your system. You can download and install Go from golang.org.


## Installation

    Clone the repository  https://learn.zone01kisumu.ke/git/shaokoth/ascii-art.git or download the code: 
## Usage

    Ensure you have a banner definition file (e.g., standard.txt) in the same directory as the program. This file should contain the ASCII art definitions for characters.

    Run the program with the default .txt file:

```bash
go run . <text string>
```
    
## Code Structure

    main.go: The main file containing the code logic.
    asciiArt: Package containing the core functionality.

## Error Handling

The program includes basic error handling for:

    * Invalid file names or paths.
    * Errors during file reading.
    * Unsupported characters not defined in the banner file.

## Source(s)
Ascii manual

## Other Innformation
This project is maintained by 

[ShadrackOkoth]
[KevinWasonga]