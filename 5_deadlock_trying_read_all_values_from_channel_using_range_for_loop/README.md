## Deadlock trying read all values from channel using range for loop

Here we had a crash because we are ranging over the channel continously reading from it. As we know, reading from a channel is a blocking operating. Once we have real all data from the channel the main goroutine enters blocking state.

```
make example_5
```
