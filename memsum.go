package collector

import (
        "fmt"
        "os"
        "strconv"
        "strings"
        "encoding/csv"
        "os/exec"
        "github.com/prometheus/client_golang/prometheus"
//        "github.com/prometheus/common/log"
)

const (
        memInfoSubsystem1 = "memory"
)

type meminfoCollector1 struct{}

func init() {
        registerCollector("meminfo_summary", defaultEnabled, NewMeminfoCollector1)
}

// NewMeminfoCollector returns a new Collector exposing memory stats.
func NewMeminfoCollector1() (Collector, error) {
        return &meminfoCollector1{}, nil
}

// Update calls (*meminfoCollector).getMemInfo to get the platform specific
// memory metrics.
func (c *meminfoCollector1) Update(ch chan<- prometheus.Metric) error {
        tmp := updateExporterMetric("summary", "meminfo_summary")
        if tmp == false {
                return nil
        }
        cmd, _ := exec.Command("whoami").Output()
        path := "/home/"+strings.TrimSpace(string(cmd))+"/csv/meminfostat.csv"
        file, err := os.Open(path)
//      fmt.Println("***********MemInfo************")
        if err != nil {
                fmt.Println(err)
        }
        var summarization = []string{"stats"}
        defer file.Close()
        lines,_ := csv.NewReader(file).ReadAll()
        var p25, median, p75, average, min, max float64
        for i , line := range lines {
                if i > 0 {
                         p25, _ = strconv.ParseFloat(line[1],64)
                         median, _ = strconv.ParseFloat(line[2],64)
                         p75, _ = strconv.ParseFloat(line[3],64)
                         average, _ = strconv.ParseFloat(line[4],64)
                         min, _ = strconv.ParseFloat(line[5],64)
                         max, _ = strconv.ParseFloat(line[6],64)

                        var j = line[0] + "_summary"

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                 nil,
                        ),
                        prometheus.GaugeValue, p25,"p25",
                        )

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                 nil,
                        ),
                        prometheus.GaugeValue, median,"median",
                        )

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                 nil,
                        ),
                        prometheus.GaugeValue, p75,"p75",
                        )

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                nil,
                        ),
                        prometheus.GaugeValue, average,"average",
                        )

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                nil,
                        ),
                        prometheus.GaugeValue, min,"min",
                        )

                        ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem1, j),
                                fmt.Sprintf("Memory information field %s.", j),
                                summarization,
                                nil,
                        ),
                        prometheus.GaugeValue, max,"max",
                        )

                        }
                                }



/*        memInfo, err := c.getMemInfo()
        if err != nil {
                return fmt.Errorf("couldn't get meminfo: %s", err)
        }
        log.Debugf("Set node_mem: %#v", memInfo)
        for k, v := range memInfo {
                ch <- prometheus.MustNewConstMetric(
                        prometheus.NewDesc(
                                prometheus.BuildFQName(namespace, memInfoSubsystem, k),
                                fmt.Sprintf("Memory information field %s.", k),
                                nil, nil,
                        ),
                        prometheus.GaugeValue, v,
                )
        }
*/
      return nil
}
