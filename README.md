gRPC protocol types comparison

How to run Example : 
-  Go run server/main.go
-  Go run client/main.go

1st measurement result :
- Execution time of SimpleRPC : 478.135101ms
- Execution time of BidirectionalRPC : 34.089263ms
- Execution time of ServerStreamingRPC : 44.402185ms
- Execution time of ClientStreamingRPC : 54.84521ms

2nd measurement result :
- Execution time of SimpleRPC : 433.355271ms
- Execution time of BidirectionalRPC : 50.163772ms
- Execution time of ServerStreamingRPC : 49.136169ms
- Execution time of ClientStreamingRPC : 22.490456msv
3rd measurement result :
- Execution time of SimpleRPC : 477.868411ms
- Execution time of BidirectionalRPC : 21.179674ms
- Execution time of ServerStreamingRPC : 13.695817ms
- Execution time of ClientStreamingRPC : 99.638916ms

