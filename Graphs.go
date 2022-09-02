package runstats

import (
	"github.com/Akilan1999/Go-Metrics-Simple/collector"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"reflect"
)

/*================== Alloc is bytes of allocated heap objects: Graph ==================*/

// AllocGraph Memory allocation graph
func (m *MetricsAllSSingleRun) AllocGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Alloc is bytes of allocated heap object",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("Alloc is bytes of allocated heap object", m.GenerateAlloc())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateAlloc generate graph for memory allocation
func (m *MetricsAllSSingleRun) GenerateAlloc() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Alloc})
	}
	return items
}

/*==================  TotalAlloc is cumulative bytes allocated for heap objects ==================*/

// TotalAllocGraph Memory allocation graph
func (m *MetricsAllSSingleRun) TotalAllocGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "TotalAlloc is cumulative bytes allocated for heap objects",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("TotalAlloc is cumulative bytes allocated for heap objects", m.GenerateTotalAlloc())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateTotalAlloc generate graph for memory allocation
func (m *MetricsAllSSingleRun) GenerateTotalAlloc() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].TotalAlloc})
	}
	return items
}

/*==================  Sys is the total bytes of memory obtained from the OS. ==================*/

// SysGraph sys graph
func (m *MetricsAllSSingleRun) SysGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Sys is the total bytes of memory obtained from the OS",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("Sys is the total bytes of memory obtained from the OS.", m.GenerateSys())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateSys generate graph for memory allocation
func (m *MetricsAllSSingleRun) GenerateSys() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Sys})
	}
	return items
}

/*==================  Lookups is the number of pointer lookups performed by the runtime. ==================*/

// LookupsGraph lookup graph
func (m *MetricsAllSSingleRun) LookupsGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Lookups is the number of pointer lookups performed by the runtime",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("Lookups is the number of pointer lookups performed by the runtime", m.GenerateLookups())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateLookups generate graph for memory allocation
func (m *MetricsAllSSingleRun) GenerateLookups() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].Lookups})
	}
	return items
}

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

// GenerateMallocs generate graph for memory allocation
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

/*================== HeapAlloc is bytes of allocated heap objects. ==================*/

// HeapAllocGraph lookup graph
func (m *MetricsAllSSingleRun) HeapAllocGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapAlloc is bytes of allocated heap objects",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapAlloc is bytes of allocated heap objects", m.GenerateHeapAlloc())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapAlloc generate graph for memory allocation
func (m *MetricsAllSSingleRun) GenerateHeapAlloc() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapAlloc})
	}
	return items
}

/*================== HeapSys estimates the largest size the heap has had. ==================*/

// HeapSysGraph
func (m *MetricsAllSSingleRun) HeapSysGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapAlloc is bytes of allocated heap objects",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapAlloc is bytes of allocated heap objects", m.GenerateHeapSys())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapSys
func (m *MetricsAllSSingleRun) GenerateHeapSys() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapSys})
	}
	return items
}

/*================== HeapIdle is bytes in idle (unused) spans. ==================*/

// HeapIdleGraph
func (m *MetricsAllSSingleRun) HeapIdleGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapAlloc is bytes of allocated heap objects",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapAlloc is bytes of allocated heap objects", m.GenerateHeapIdle())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapIdle
func (m *MetricsAllSSingleRun) GenerateHeapIdle() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapIdle})
	}
	return items
}

/*================== HeapInuse is bytes in in-use spans ==================*/

// HeapInuseGraph
func (m *MetricsAllSSingleRun) HeapInuseGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapInuse is bytes in in-use spans",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapInuse is bytes in in-use spans", m.GenerateHeapInuse())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapInuse
func (m *MetricsAllSSingleRun) GenerateHeapInuse() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapInuse})
	}
	return items
}

/*================== HeapReleased is bytes of physical memory returned to the OS ==================*/

// HeapReleasedGraph
func (m *MetricsAllSSingleRun) HeapReleasedGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapReleased is bytes of physical memory returned to the OS",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapReleased is bytes of physical memory returned to the OS", m.GenerateHeapReleased())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapReleased
func (m *MetricsAllSSingleRun) GenerateHeapReleased() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapReleased})
	}
	return items
}

/*================== HeapObjects is the number of allocated heap objects. ==================*/

// HeapObjectsGraph
func (m *MetricsAllSSingleRun) HeapObjectsGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "HeapObjects is the number of allocated heap objects.",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("HeapObjects is the number of allocated heap objects.", m.GenerateHeapObjects())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateHeapObjects
func (m *MetricsAllSSingleRun) GenerateHeapObjects() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].HeapObjects})
	}
	return items
}

/*=============================================
          Stack memory statistics.

	 Stacks are not considered part of the
     heap, but the runtime can reuse a span
     of heap memory for stack memory, and
	 vice-versa.
=============================================*/

/*================== StackInuse is bytes in stack spans ==================*/

// StackInuseGraph
func (m *MetricsAllSSingleRun) StackInuseGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "StackInuse is bytes in stack spans.",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("StackInuse is bytes in stack spans.", m.GenerateStackInuse())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

// GenerateStackInuse
func (m *MetricsAllSSingleRun) GenerateStackInuse() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].StackInuse})
	}
	return items
}

/*================== StackSys is bytes of stack memory obtained from the OS. ==================*/

func (m *MetricsAllSSingleRun) StackSysGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "StackSys is bytes of stack memory obtained from the OS.",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("StackSys is bytes of stack memory obtained from the OS.", m.GenerateStackSysGraph())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

func (m *MetricsAllSSingleRun) GenerateStackSysGraph() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].StackSys})
	}
	return items
}

/*=============================================
          off-heap memory statistics.

    The following statistics measure
    runtime-internal structures that are not
    allocated from heap memory (usually
	because they are part of implementing
    the heap). Unlike heap or stack memory,
    any memory allocated to these
	structures is dedicated to these structures.

	These are primarily useful for debugging
    runtime memory overheads.
=============================================*/

/*================== MSpanInuse is bytes of allocated mspan structures. ==================*/

func (m *MetricsAllSSingleRun) MspanInUseGraph() components.Charter {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "MSpanInuse is bytes of allocated mspan structures.",
		Subtitle: "",
	}))

	// Put data into instance
	line.SetXAxis(m.Duration).
		AddSeries("MSpanInuse is bytes of allocated mspan structures.", m.GenerateMspanInUse())
	// Where the magic happens
	//f, _ := os.Create("MemoryAllocation.html")
	//line.Render(f)

	return line
}

func (m *MetricsAllSSingleRun) GenerateMspanInUse() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, _ := range m.Metrics {
		items = append(items, opts.LineData{Value: m.Metrics[i].MSpanInuse})
	}
	return items
}

// Trying interfaces

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
