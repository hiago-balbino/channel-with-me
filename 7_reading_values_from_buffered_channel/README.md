# Reading all values from a bufferd channel

In this example, we are using a buffered channel where we can set the size of the channel. For example, if we set a channel buffer to size=4 we will write in this channel 4 times and will block until we read from the channel. The channel will be in a block state until we read from it in some place, after reading the channel change to a non-blocking state.

As we can see in this code example we set the size channel to 4, so when we run the code we will see the channel been written 4 times before reading(we set the timeout the observe the behavior). After reading the first letter we can write more letters.

```
make example_7
```
