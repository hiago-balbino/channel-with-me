.PHONY: help example_1 example_2 example_3 example_4 example_5 example_6 example_7 example_8 example_9

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
## example_4: Reading all values from channel using for loop
## See folder: 4_reading_all_values_from_channel_using_for_loop
## Run: make example_4
example_4:
	@go run 4_reading_all_values_from_channel_using_for_loop/main.go

## --------------------------------------------
## example_5: Deadlock trying read all values from channel using range for loop
## See folder: 5_deadlock_trying_read_all_values_from_channel_using_range_for_loop
## Run: make example_5
example_5:
	@go run 5_deadlock_trying_read_all_values_from_channel_using_range_for_loop/main.go

## --------------------------------------------
## example_6: Reading all values from channel using range for loop
## See folder: 6_reading_all_values_from_channel_using_range_for_loop
## Run: make example_6
example_6:
	@go run 6_reading_all_values_from_channel_using_range_for_loop/main.go

## --------------------------------------------
## example_7: Reading values from a buffered channel
## See folder: 7_reading_values_from_buffered_channel
## Run: make example_7
example_7:
	@go run 7_reading_values_from_buffered_channel/main.go

## --------------------------------------------
## example_8: Simple graceful shutdown
## See folder: 8_simple_graceful_shutdown
## Run: make example_8
example_8:
	@go run 8_simple_graceful_shutdown/main.go

## --------------------------------------------
## example_9: Reading values from channel using select
## See folder: 9_reading_values_from_channel_using_select
## Run: make example_9
example_9:
	@go run 9_reading_values_from_channel_using_select/main.go
## --------------------------------------------

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
