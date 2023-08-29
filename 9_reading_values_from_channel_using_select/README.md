## Reading all values from channel using select

This is a simple example in Go that illustrates the use of time-limited channels and context.

The program creates letters and numbers channels and calls the populateChannels function to populate these channels in separate goroutines. While the goroutines are populating the channels, the select loop in the main function waits until a value is available in one of the channels or until the timeout context is canceled. If a value is available on one of the channels, it is printed on the standard output. When the timeout is reached, the program terminates.

```
make example_9
```