# ASCII ART REVERSE

  Ascii Art Reverse is a small program designed to read txt files that has ascii art graphical representation as input and outputs to terminal the equivalent of it in normal string.

## USAGE

To run the application, make sure terminal is at root directory of the project, then run below command:

NOTE: It is important to reference the txt file properly. 

Run this code if txt file is at same directory as main.go:
```
go run . --reverse=<FileName.txt>
```

Run this code if txt file is at child directory and main.go at parent directory:
```
go run . --reverse=<dir/FileName.txt>
```

Run this code if txt file is at parent directory and main.go at child directory:
```
go run . --reverse=<../FileName.txt>
```

## AUTOR

* HUSSAIN
* ABDULRAHMAN
* DANIEL
