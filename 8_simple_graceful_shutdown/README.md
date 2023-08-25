## Simple graceful shutdown

In this example, we created a really simple graceful shutdown using channels and select statements.

We configured the context with a 10-second timeout and this is passed as a parameter to the graceful shutdown function.

In the main function, we call the graceful shutdown function in a separate goroutine and use a WaitGroup just to ensure we really can print some messages or do whatever we want before exiting.

Inside de graceful shutdown function we create a signal channel that will be notified by syscall.SIGINT or syscall.SIGTERM. We have an infinite loop with select statements where we can receive a context timeout, a signal notify, and in the last one we are just printing a message every 2 seconds. While running the application if we interrupt the process with CTRL/Cmd + C for example, we can see the finishing message inside de graceful shutdown functions and then the finishing message from the main goroutine.

```
make example_8
```
