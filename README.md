# goroutines sample code

It is a sample code to illustrate go routines and channels.

## What does it do
It creates 3 different threads and below are there functionalities:
1. One thread starts server on 8000 port to serve GET request and return 1 to 20 numbers as response in JSON format.
2. Another thread hits GET request defined in the above thread to get number and puts it on channel.
3. Last thread reads number put on channel by the second thread and prints its square.


## How to run it
Run below command to run it

```go run main.go```
