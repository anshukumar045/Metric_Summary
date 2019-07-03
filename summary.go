/*
  GET THE PERCENTILES MEDIAN AND AVERAGE FOR METRIC FOR 30 SEC AND SAVE IT TO CSV
  PUSH THE VALUES TO PROMETHUES AS REQUESTED
*/
package collector

import (
        "fmt"
//      "time"
      m "math"
        "strings"
//      "reflect"
//      "sync"
        "sort"
        "strconv"
        "os"
        "os/exec"
        "encoding/csv"
)

type Exp struct {
        name string
        value float64
}

type Statsexp struct {
        name    string
        average float64
        median float64
        P25    float64
        P75    float64
        min    float64
        max    float64
}
var interval, duration =  2 ,30

func getInterval(i int, d int)(int,float64,float64,int,int,float64,float64) { // total p25
        totlaCount := int(d/i)
        p25_1 := m.Floor((30 / (float64(i))) * 0.25) -1
        p25_2 := m.Ceil((30 / (float64(i))) * 0.25) -1
        p50_1 := int(m.Floor(30 / (float64(i) * 2))) -1
        p50_2 := int(m.Ceil(30 / (float64(i) * 2))) -1
        p75_1 := m.Floor((30 / (float64(i))) * 0.75) -1
        p75_2 := m.Ceil((30 / (float64(i))) * 0.75) -1
return totlaCount , p25_1 ,p25_2 ,p50_1 ,p50_2 ,p75_1, p75_2
}
var totlaCount, p25_1 ,p25_2 ,p50_1 ,p50_2 ,p75_1, p75_2 = getInterval(interval,duration)

func unique(intSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
func sum(input []float64) float64 {
    sum := 0.00
    for i := range input {
        sum += input[i]
    }
return sum
}
func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func Sumstats(obj []Exp) []Statsexp {
        var keys,ukeys []string
        for i := 0 ; i < len(obj) ; i++ {
                keys = append(keys,obj[i].name)
        }
        ukeys = unique(keys)
        mm := make(map[string][]float64)
        for _,k := range ukeys{
                mm[k] = nil
        }
        for i := 0 ; i < len(obj); i++{
                        for f, _ := range mm {
                                if obj[i].name == f {
                                        mm[f] = append(mm[f],obj[i].value)
                                }
                        }
                }
        final := make([]Statsexp,0)

        for k,v := range mm {
                                var x Statsexp
                                x.name = k
                                x.average = sum(v)/float64(len(v))
                                sort.Float64s(mm[k])
                                x.median = (v[int(p50_1)] + v[int(p50_1)]) / 2     // P25 is 14.5
                                x.P25 = (v[int(p25_1)] + v[int(p25_2)] ) / 2         // P25 is 7.25
                                x.P75 = (v[int(p75_1)] + v[int(p75_1)] ) /2        // P75 is 21.75
                                x.min = v[0]
                                x.max = v[totlaCount-1]
                                final = append(final,x)
        }


/*      for k,v := range mm {
                                fmt.Println("lenth of v", len(v))
                                var x Statsexp
                                x.name = k
                                x.average = sum(v)/len(v)
                                sort.Float64s(mm[k])
                                x.median = (v[14] + v[15]) / 2     // P25 is 14.5
                                x.P25 = (v[7] + v[8] ) / 2         // P25 is 7.25
                                x.P75 = (v[21] + v[22] ) /2        // P75 is 21.75
                                x.min = v[0]
                                x.max = v[29]
                                final = append(final,x)
        }   */
return final
}


func summarytocsv(final []Statsexp, fname string) {
//      filename := "D:\\projects\\src\\test\\lesson\\jasontotext\\objtocsv.csv"
        cmd, _ := exec.Command("whoami").Output()
        path := "/home/"+strings.TrimSpace(string(cmd))+"/csv/"
//      path := "/home/kanshu/csv/"
        if _, err := os.Stat(path); os.IsNotExist(err) {
                os.Mkdir(path, 0755)
        }

        filename := path + fname +".csv"
        err := os.Remove(filename)
        if err != nil {
                fmt.Println(err)
        }
        f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
    }
        writer := csv.NewWriter(f)
        defer f.Close()
        var data1 = [][]string{{"Name","P25","Median","p75","Avg","Min","Max"}}
        for _ , value := range data1 {
        err = writer.Write(value)
        checkError("Cannot write to file", err)
                writer.Flush()
   }
        for i := 0; i<len(final); i++ {
        var data = [][]string{{final[i].name, FloatToString(final[i].P25), FloatToString(final[i].median), FloatToString(final[i].P75), FloatToString(final[i].average), FloatToString(final[i].min), FloatToString(final[i].max)}}
                for _ , value := range data {
                        err = writer.Write(value)
                        checkError("Cannot write to file", err)
                        writer.Flush()
                }
    }
}

func checkError(message string, err error) {
    if err != nil {
        fmt.Println(message, err)
    }
}
