# Go CLI Quiz App

-  It reads the question andanswers from a CSV File
-  A timer runs in the background  ,The default time     limit is 60 seconds but you  extend the time as you wish using the --time flag
-  The timer and recieving inputs are go routine which communicate via channel to the rest of the program

 **To attend the quiz with the default Time Limit**
    
   ``make dev``

**To extend the time Limit**

   ``go build . && ./go_quiz --time=<TIME_LIMIT>``

**Concepts Learned**

- Parsing Command Line args
- Communicating via channel
- Parsing CSV
- structs in Go
- Formatting Strings

You can change the problems in the csv file as you wish
