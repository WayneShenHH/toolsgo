package metrixsvc

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/rcrowley/go-metrics"
)

const fanout = 10

//MetrixServer start a metrix server
func MetrixServer() {

	r := metrics.NewRegistry()

	c := metrics.NewCounter()
	r.Register("foo", c)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				c.Dec(19)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				c.Inc(47)
				time.Sleep(400e6)
			}
		}()
	}

	g := metrics.NewGauge()
	r.Register("bar", g)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				g.Update(19)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				g.Update(47)
				time.Sleep(400e6)
			}
		}()
	}

	gf := metrics.NewGaugeFloat64()
	r.Register("barfloat64", gf)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				g.Update(19.0)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				g.Update(47.0)
				time.Sleep(400e6)
			}
		}()
	}

	hc := metrics.NewHealthcheck(func(h metrics.Healthcheck) {
		if 0 < rand.Intn(2) {
			h.Healthy()
		} else {
			h.Unhealthy(errors.New("baz"))
		}
	})
	r.Register("baz", hc)

	s := metrics.NewExpDecaySample(1028, 0.015)
	//s := metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	r.Register("bang", h)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				h.Update(19)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				h.Update(47)
				time.Sleep(400e6)
			}
		}()
	}

	m := metrics.NewMeter()
	r.Register("quux", m)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				m.Mark(19)
				time.Sleep(300e6)
			}
		}()
		go func() {
			for {
				m.Mark(47)
				time.Sleep(400e6)
			}
		}()
	}

	t := metrics.NewTimer()
	r.Register("hooah", t)
	for i := 0; i < fanout; i++ {
		go func() {
			for {
				t.Time(func() { time.Sleep(300e6) })
			}
		}()
		go func() {
			for {
				t.Time(func() { time.Sleep(400e6) })
			}
		}()
	}

	metrics.RegisterDebugGCStats(r)
	go metrics.CaptureDebugGCStats(r, 5e9)

	metrics.RegisterRuntimeMemStats(r)
	go metrics.CaptureRuntimeMemStats(r, 5e9)
	metrics.Log(r, 60e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	/*
		w, err := syslog.Dial("unixgram", "/dev/log", syslog.LOG_INFO, "metrics")
		if nil != err { log.Fatalln(err) }
		metrics.Syslog(r, 60e9, w)
	*/

	/*
		addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2003")
		metrics.Graphite(r, 10e9, "metrics", addr)
	*/

	/*
		stathat.Stathat(r, 10e9, "example@example.com")
	*/

}

// PrometheusServer start server
func PrometheusServer() {
	prometheusRegistry := prometheus.NewRegistry()
	metricsRegistry := metrics.NewRegistry()
	pClient := prometheusmetrics.NewPrometheusProvider(metricsRegistry, "test", "subsys", prometheusRegistry, 1*time.Second)
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "test",
		Subsystem: "subsys",
		Name:      "counter",
		Help:      "counter",
	})
	prometheusRegistry.Register(gauge)
	go pClient.UpdatePrometheusMetrics()
	gauge.Set(88)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
	//select {}
}

var (
	// 自訂數值型態的測量數據。
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "CPU 目前的溫度。",
	})
	// 計數型態的測量數據，並帶有自訂標籤。
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "硬碟發生錯誤的次數。",
		},
		[]string{"device"},
	)
)

func init() {
	// 測量數據必須註冊才會暴露給外界知道：
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

// Report report system stat for prometheus
func Report() {
	// 配置測量數據的數值。
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// 我們會用 Prometheus 所提供的預設處理函式在 "/metrics" 路徑監控著。
	// 這會暴露我們的數據內容，所以 Prometheus 就能夠獲取這些數據。
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
