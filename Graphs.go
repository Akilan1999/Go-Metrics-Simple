package runstats

import (
	"github.com/Akilan1999/Go-Metrics-Simple/collector"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"reflect"
)

/*================== Mallocs is the cumulative count of heap objects allocated and Frees is the cumulative count of heap objects freed ==================*/

// MallocsAndFreesGraph mallocs graph
func (m *MetricsAllSSingleRun) MallocsAndFreesGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Mallocs is the cumulative count of heap objects allocate and Frees is the cumulative count of heap objects free and live objects",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("Mallocs is the cumulative count of heap objects allocate", m.GenerateMallocs()).
		AddSeries("Frees is the cumulative count of heap objects freed", m.GenerateFrees()).
		AddSeries("Live objects", m.GenerateLiveobjects())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

/*================== Generalized function to generate graphs ==================*/

func (m *MetricsAllSSingleRun) GenerateMallocs() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Mallocs})
	}
	return items
}

// GenerateFrees generate graph for frees
func (m *MetricsAllSSingleRun) GenerateFrees() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Frees})
	}
	return items
}

// GenerateLiveobjects for Live objects (Mallocs - free)
func (m *MetricsAllSSingleRun) GenerateLiveobjects() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Mallocs - m.Metrics[i].Frees})
	}
	return items
}

func (m *MetricsAllSSingleRun) Graph(description string, field string, Type ...string) components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    field,
		Subtitle: description,
	}))

	TypePlot := "int"
	if len(Type) != 0 {
		TypePlot = Type[0]
	}

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries(description, m.GenerateData(field, TypePlot))
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

func (m *MetricsAllSSingleRun) GenerateData(field string, TypePlot string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		var metric collector.Fields
		metric = m.Metrics[i]
		value := reflect.Indirect(reflect.ValueOf(&metric)).FieldByName(field)

		if TypePlot == "int" {
			items = append(items, opts.LineData{Value: value.Int()})
		} else if TypePlot == "float" {
			items = append(items, opts.LineData{Value: value.Float()})
		}
	}
	return items
}
