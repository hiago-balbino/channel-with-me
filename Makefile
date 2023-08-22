.PHONY: help example_1 example_2 example_3

## --------------------------------------------
## example_1: Deadlock without read value
## See folder: 1_deadlock_without_read_value
## Run: make example_1
example_1:
	@go run 1_deadlock_without_read_value/main.go

## --------------------------------------------
## example_2: Reading only one value
## See folder: 2_reading_only_one_value 
## Run: make example_2
example_2:
	@go run 2_reading_only_one_value/main.go

## --------------------------------------------
## example_3: Deadlock trying read all values
## See folder: 3_deadlock_trying_read_all_values 
## Run: make example_3
example_3:
	@go run 3_deadlock_trying_read_all_values/main.go

## --------------------------------------------
## > example_4: Given we are reading from the channel as many times as the length of the word variable, we can read all the values, but this code is not elegant enought.
## See: 4_reading_all_values_from_slice_range/main.go
## Run: make example_4
example_4:
	@go run 4_reading_all_values_from_slice_range/main.go

## --------------------------------------------
## > example_5: Here we had a crash because we are ranging over the channel continously reading from it. As we know, reading from a channel is a blocking operating. Once we have real all data from the channel the main goroutine enters blocking state.
## See: 5_deadlock_trying_read_all_values_from_channel_range/main.go
## Run: make example_5
example_5:
	@go run 5_deadlock_trying_read_all_values_from_channel_range/main.go

## --------------------------------------------
## > example_6: Here we have fixed the crash and the most important bit we are saying the goroutine which is responsible to send data to also tell us when you are done sending data closing the channel. Ranging over a channel is an all inclusive operation which reads if data exists but stops if the channel is closed. We have also decoupled read operation from write. Even if the size of word increases, we do not worry about screwing up the read logic.
## See: 6_reading_all_values_from_channel_range/main.go
## Run: make example_6
example_6:
	@go run 6_reading_all_values_from_channel_range/main.go

## --------------------------------------------
## example_7: Reading values from a buffered channel
## See folder: 7_reading_values_from_buffered_channel
## Run: make example_7
example_7:
	@go run 7_reading_values_from_buffered_channel/main.go

## --------------------------------------------

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
