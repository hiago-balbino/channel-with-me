## Deadlock trying to read all values

In this example we are reading from a channel which has no more data to be read. It would read first 4 elements and block forever while trying to read for a 5th element which never appears.

```
make example_3
```
