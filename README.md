## Channels in Go

This repository aims to document examples of study code using channels in the Go programming language.
Each folder has a single test example and a brief explanation of the logic or error.

### Examples

* [Deadlock without read value](./1_deadlock_without_read_value)
* [Reading only one value](./2_reading_only_one_value)
* [Deadlock trying read all values](./3_deadlock_trying_read_all_values)
* [Reading all values from channel using for loop](./4_reading_all_values_from_channel_using_for_loop)
* [Deadlock trying read all values from channel using range for loop](./5_deadlock_trying_read_all_values_from_channel_using_range_for_loop)
* [Reading all values from channel using range for loop](./6_reading_all_values_from_channel_using_range_for_loop)
* [Reading values from a buffered channel](./7_reading_values_from_buffered_channel)

### Run to see all examples command

```
make help
```
