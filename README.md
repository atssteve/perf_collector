# Perf Collector

[![CircleCI](https://circleci.com/gh/atssteve/perf_collector.svg?style=svg)](https://circleci.com/gh/atssteve/perf_collector)

Perf Collector is a cross platform performance metric collection tool.

## Hacking

This project uses Go modules. It is recommended you upgrade use Go 1.11 or enable the beta features in Go 1.10.

Each controller.metric and controller.config MUST have something in it! it will panic
Also each variables need to match between the config file and the flag names. Take a look 
at the interval values for and example