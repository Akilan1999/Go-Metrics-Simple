#  Go Metrics Simple
This is a really simple project build to track the performance of
your go program. No InfluxDB BS which I am sick of configuring.
The objective is to run your program and the go metrics gets
saved to the JSON file and after the run the simple program
generates a static HTML. 

#### based on the repo: github.com/tevjef/go-runtime-metrics


## How to use it 
1. It's as easy as
```go 
import (
	metrics "github.com/Akilan1999/Go-Metrics-Simple"
)

func main() {
   // Add this to the starting point of your go program
   err := metrics.RunCollector(metrics.DefaultConfig)
   // <Your Go program here> 
   // Add this to the end point of your go program
	 metrics.ComputeDefaultFile()
}

```
2. After your program is complete just open ```metrics.html``` file (The image below is just sample metrics of the ```/example```code. : 
<br>

![Screenshot 2022-09-02 at 22 37 27](https://user-images.githubusercontent.com/31743758/188217178-2cc9b567-02fe-4534-805c-8fd408e86b46.png)



