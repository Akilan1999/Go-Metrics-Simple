package runstats

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/components"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Akilan1999/Go-Metrics-Simple/collector"
	"github.com/araddon/dateparse"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

const (
	//defaultHost               = "localhost:8086"
	defaultMeasurement = "go.runtime"
	//defaultDatabase           = "stats"
	defaultCollectionInterval = 1 * time.Second
	defaultBatchInterval      = 60 * time.Second
)

// DefaultConfig A configuration with default values.
var DefaultConfig = &Config{}

type Config struct {
	// Measurement to write points to.
	// Default is "go.runtime.<hostname>".
	Measurement string

	// Interval at which to write batched points to InfluxDB.
	// Default is 60 seconds
	BatchInterval time.Duration

	// Interval at which to collect points.
	// Default is 10 seconds
	CollectionInterval time.Duration

	// Disable collecting CPU Statistics. cpu.*
	// Default is false
	DisableCpu bool

	// Disable collecting Memory Statistics. mem.*
	DisableMem bool

	// Disable collecting GC Statistics (requires Memory be not be disabled). mem.gc.*
	DisableGc bool

	// Default is DefaultLogger which exits when the library encounters a fatal error.
	Logger Logger
}

func (config *Config) init() (*Config, error) {
	if config == nil {
		config = DefaultConfig
	}

	if config.Measurement == "" {
		config.Measurement = defaultMeasurement

		if hn, err := os.Hostname(); err != nil {
			config.Measurement += ".unknown"
		} else {
			config.Measurement += "." + hn
		}
	}

	if config.CollectionInterval == 0 {
		config.CollectionInterval = defaultCollectionInterval
	}

	if config.BatchInterval == 0 {
		config.BatchInterval = defaultBatchInterval
	}

	if config.Logger == nil {
		config.Logger = &DefaultLogger{}
	}

	return config, nil
}

func RunCollector(config *Config) (err error) {
	if config, err = config.init(); err != nil {
		return err
	}

	// Make client
	//clnt, err := client.NewHTTPClient(client.HTTPConfig{
	//	Addr:     "http://" + config.Host,
	//	Username: config.Username,
	//	Password: config.Password,
	//})

	if err != nil {
		return errors.Wrap(err, "failed to create influxdb client")
	}

	// Ping InfluxDB to ensure there is a connection
	//if _, _, err := clnt.Ping(5 * time.Second); err != nil {
	//	return errors.Wrap(err, "failed to ping influxdb client")
	//}
	//
	//// Auto create database
	//_, err = queryDB(clnt, fmt.Sprintf("CREATE DATABASE \"%s\"", config.Database))

	if err != nil {
		config.Logger.Fatalln(err)
	}

	_runStats := &runStats{
		logger: config.Logger,
		//client: clnt,
		config: config,
		pc:     make(chan *client.Point),
	}

	bp, err := _runStats.newBatch()

	if err != nil {
		return err
	}

	_runStats.points = bp

	go _runStats.loop(config.BatchInterval)

	_collector := collector.New(_runStats.onNewPoint)
	_collector.PauseDur = config.CollectionInterval
	_collector.EnableCPU = !config.DisableCpu
	_collector.EnableMem = !config.DisableMem
	_collector.EnableGC = !config.DisableGc

	go _collector.Run()

	return nil
}

type runStats struct {
	logger Logger
	client client.Client
	points client.BatchPoints
	config *Config
	pc     chan *client.Point
}

func (r *runStats) onNewPoint(fields collector.Fields) {
	pt, err := client.NewPoint(r.config.Measurement, fields.Tags(), fields.Values(), time.Now())

	if err != nil {
		r.logger.Fatalln(errors.Wrap(err, "error while creating point"))
	}

	// write to file here
	f, err := os.OpenFile("results.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	marshal, _ := json.Marshal(fields)
	if _, err = f.WriteString(string(marshal) + "\n"); err != nil {
		panic(err)
	}

	r.pc <- pt
}

func (r *runStats) newBatch() (bp client.BatchPoints, err error) {
	//bp, err = client.NewBatchPoints(client.BatchPointsConfig{
	//	Database:        r.config.Database,
	//	Precision:       r.config.Precision,
	//	RetentionPolicy: r.config.RetentionPolicy,
	//})

	if err != nil {
		r.logger.Fatalln(errors.Wrap(err, "could not create BatchPoints"))
	}

	return
}

// Write collected points to influxdb periodically
func (r *runStats) loop(interval time.Duration) {
	ticks := time.Tick(interval)

	for {
		select {
		case <-ticks:
			if r.points == nil || len(r.points.Points()) <= 0 {
				continue
			}

			//if err := r.client.Write(r.points); err != nil {
			//	r.logger.Fatalln(errors.Wrap(err, "could not write points to InfluxDB"))
			//	continue
			//}

			//r.points = nil
			//
			//bp, err := r.newBatch()
			//
			//if err != nil {
			//	r.logger.Fatalln(errors.Wrap(err, "could not create BatchPoints"))
			//	continue
			//}
			//
			//r.points = bp

		case pt := <-r.pc:
			//if r.points != nil {
			r.logger.Println(pt.String())

			//r.points.AddPoint(pt)
			//}
		}
	}
}

type Logger interface {
	Println(v ...interface{})
	Fatalln(v ...interface{})
}

type DefaultLogger struct{}

// Println Overwritten function to save result to the result file
func (*DefaultLogger) Println(v ...interface{}) {}
func (*DefaultLogger) Fatalln(v ...interface{}) { log.Fatalln(v) }

//
//func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
//	q := client.Query{
//		Command: cmd,
//	}
//	if response, err := clnt.Query(q); err == nil {
//		if response.Error() != nil {
//			return res, response.Error()
//		}
//		res = response.Results
//	} else {
//		return res, err
//	}
//	return res, nil
//}

// Basic struct with file information
type MetricsAllSSingleRun struct {
	Metrics  []collector.Fields
	Duration []float64
}

func ComputeDefaultFile() {
	file, err := os.Open("results.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var MetricsArray MetricsAllSSingleRun

	scanner := bufio.NewScanner(file)
	i := 0

	var duration float64
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		var metric collector.Fields
		json.Unmarshal(scanner.Bytes(), &metric)
		// Calculate time difference
		if i == 0 {
			duration = 0
		} else {
			t, _ := dateparse.ParseLocal(strconv.FormatInt(metric.Timestamp, 10))
			prevTimeStamp, _ := dateparse.ParseLocal(strconv.FormatInt(MetricsArray.Metrics[i-1].Timestamp, 10))
			Duration := t.Sub(prevTimeStamp)
			duration += Duration.Seconds()
		}
		fmt.Println(duration)
		MetricsArray.Duration = append(MetricsArray.Duration, duration)
		MetricsArray.Metrics = append(MetricsArray.Metrics, metric)
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	MetricsArray.GenerateGraphs()
}

// GenerateGraphs Generate graphs for run
func (m *MetricsAllSSingleRun) GenerateGraphs() {

	page := components.NewPage()
	page.PageTitle = "Go Program metrics"
	page.Layout = "flex"
	page.AddCharts(
		m.Graph("Alloc is bytes of allocated heap object", "Alloc"),
		m.Graph("TotalAlloc is cumulative bytes allocated for heap objects", "TotalAlloc"),
		m.Graph("Sys is the total bytes of memory obtained from the OS", "Sys"),
		m.Graph("Lookups is the number of pointer lookups performed by the runtime", "Lookups"),
		m.MallocsAndFreesGraph(),
		m.Graph("HeapAlloc is bytes of allocated heap objects", "HeapAlloc"),
		m.Graph("HeapSys is bytes of heap memory obtained from the OS", "HeapSys"),
		m.Graph("HeapIdle is bytes in idle (unused) spans.", "HeapIdle"),
		m.Graph("HeapInuse is bytes in in-use spans.", "HeapInuse"),
		m.Graph("HeapReleased is bytes of physical memory returned to the OS.", "HeapReleased"),
		m.Graph("HeapObjects is the number of allocated heap objects.", "HeapObjects"),
		m.Graph("StackInuse is bytes in stack spans.", "StackSys"),
		m.Graph("StackSys is bytes of stack memory obtained from the OS", "StackSys"),
		m.Graph("MSpanInuse is bytes of allocated mspan structures.", "MSpanInuse"),
		m.Graph("MSpanSys is bytes of memory obtained from the OS for mspan.", "MSpanSys"),
		m.Graph("MCacheInuse is bytes of allocated mcache structures.", "MCacheInuse"),
		m.Graph("MCacheSys is bytes of memory obtained from the OS for mcache structures.", "MCacheSys"),
		m.Graph("GCSys is bytes of memory in garbage collection metadata.", "GCSys"),
		m.Graph("OtherSys is bytes of memory in miscellaneous off-heap runtime allocations.", "OtherSys"),
		m.Graph("NextGC is the target heap size of the next GC cycle.", "NextGC"),
		m.Graph("LastGC is the time the last garbage collection finished, as nanoseconds since 1970 (the UNIX epoch).", "LastGC"),
		m.Graph("PauseTotalNs is the cumulative nanoseconds in GC stop-the-world pauses since the program started.", "PauseTotalNs"),
		m.Graph("PauseNs is a circular buffer of recent GC stop-the-world pause times in nanoseconds.", "PauseNs"),
		m.Graph("NumGC is the number of completed GC cycles.", "NumGC"),
		m.Graph("GCCPUFraction is the fraction of this program's available CPU time used by the GC since the program started.", "GCCPUFraction", "float"),
	)

	f, err := os.Create("metrics.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
