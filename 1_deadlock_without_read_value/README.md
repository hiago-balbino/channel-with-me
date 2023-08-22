## Deadlock without read value

In this example, we are trying to read from a channel that will never be written. In Go, if all goroutines are in blocking state forever it will try to crash the progem. A blocking goroutine can only be unblocked by an operation make in another goroutine.

```
make example_1
```
