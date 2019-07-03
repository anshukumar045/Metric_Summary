package collector
import (
        "fmt"
        "time"
//      "sync"
        "strings"
//      "path/filepath"
        "strconv"
        "regexp"
        "io/ioutil"
//        "github.com/prometheus/client_golang/prometheus"
//        "github.com/prometheus/procfs"
//      "github.com/prometheus/common/log"
)

// Meminfo summary

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}


func myGetMemInfo(){
        for {
                start := time.Now()
                obj := make([]Exp,0)
                for z := 0 ; z < totlaCount ; z++ {
                        dat, err := ioutil.ReadFile("/proc/meminfo")
                        if err != nil {
                                        fmt.Println(err)
                        }
                        var tmp Exp
                        mystr := string(dat)
                        re := regexp.MustCompile(`kB`)
                        nmystr := re.ReplaceAllString(mystr, "")
                        re1 := regexp.MustCompile(`:`)
                        nnmystr := re1.ReplaceAllString(nmystr, "")

                         re2 := regexp.MustCompile(`\(`)
                        nnmystr1 := re2.ReplaceAllString(nnmystr, "_")

                        re3 := regexp.MustCompile(`\)`)
                        nnmystr2 := re3.ReplaceAllString(nnmystr1, "")

                        valstr := standardizeSpaces(nnmystr2)


                        mystrn := strings.Split(valstr, " ")
                        for i:=0; i<len(mystrn) -1; i=i+2 {
                                tmp.name = mystrn[i]
                                tst , _ := strconv.ParseFloat(mystrn[i+1], 64)
                                tmp.value = tst
                                obj = append(obj, tmp)
                        }
                        if z < totlaCount-1 {
                                time.Sleep(time.Duration(interval) * time.Second)
                        }
                }
                final := Sumstats(obj)
                obj = nil
                summarytocsv(final,"meminfostat")
                timesince := time.Since(start)
//                fmt.Println("meminfostat",timesince)
                 if int(timesince/ time.Second) < 29 {
                        time.Sleep(time.Duration(29.00000000 - float64(timesince/ time.Second)) *  time.Second)
//                        fmt.Println("sleeping")
                }
//                timesince1 := time.Since(start)
//                fmt.Println("final execution time Mydiskstats ", timesince1)
        }
}
