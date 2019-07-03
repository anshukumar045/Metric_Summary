# Metric_Summary

The logic for summaryMemInfo.go and memsum.go can be repeated for all the metrics for node exporter.

summary.go :-
  It contains the utility to summarize and load the data on to disk under the user running the exporter

summaryMemInfo.go:- 
  Summarise the data and save it to CSV.
  Data is summarized as Percentiles (P25 P75), Median and Average

memsum.go :- 
  Push the data to prometheus as requested
