## Reading only one value

Here as before(example_1) the main gorountine was waiting on reading from a channel and hoping that some other goroutine would send data. We sent all the letters to channel, but we got only one element printed because we read only once from the channel. 

```
make example_2
```
