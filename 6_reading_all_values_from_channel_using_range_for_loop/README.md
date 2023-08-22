## Reading all values from channel using range for loop

Here we have fixed the crash and the most important bit we are saying the goroutine which is responsible to send data to also tell us when you are done sending data closing the channel. Ranging over a channel is an all inclusive operation which reads if data exists but stops if the channel is closed. We have also decoupled read operation from write. Even if the size of word increases, we do not worry about screwing up the read logic.

```
make example_6
```
