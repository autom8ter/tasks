package functions

import (
	"github.com/autom8ter/tasks/docs"
	"github.com/autom8ter/tasks/pkg/util"
	"github.com/autom8ter/tasks/sdk/go/tasks"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"net/http/pprof"
)

type RouterFunc func(r *mux.Router)

func WithProfiling() RouterFunc {
	return func(r *mux.Router) {
		r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	}
}

func WithDocs() RouterFunc {
	return func(r *mux.Router) {
		r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
			bits, err := docs.Asset("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			io.WriteString(w, string(bits))
		})
	}
}

func WithMetrics() RouterFunc {
	return func(r *mux.Router) {
		var (
			inFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "in_flight_requests",
				Help: "A gauge of requests currently being served by the wrapped handler.",
			})

			counter = prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Name: "api_requests_total",
					Help: "A counter for requests to the wrapped handler.",
				},
				[]string{"code", "method"},
			)

			// duration is partitioned by the HTTP method and handler. It uses custom
			// buckets based on the expected request duration.
			duration = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    "request_duration_seconds",
					Help:    "A histogram of latencies for requests.",
					Buckets: []float64{.025, .05, .1, .25, .5, 1},
				},
				[]string{"handler", "method"},
			)

			// responseSize has no labels, making it a zero-dimensional
			// ObserverVec.
			responseSize = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    "response_size_bytes",
					Help:    "A histogram of response sizes for requests.",
					Buckets: []float64{200, 500, 900, 1500},
				},
				[]string{},
			)
		)

		// Register all of the metrics in the standard registry.
		prometheus.MustRegister(inFlightGauge, counter, duration, responseSize)
		var chain http.Handler
		r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pth, _ := route.GetPathTemplate()
			chain = promhttp.InstrumentHandlerInFlight(inFlightGauge,
				promhttp.InstrumentHandlerDuration(duration.MustCurryWith(prometheus.Labels{"handler": pth}),
					promhttp.InstrumentHandlerCounter(counter,
						promhttp.InstrumentHandlerResponseSize(responseSize, route.GetHandler()),
					),
				),
			)
			route = route.Handler(chain)
			return nil
		})
		r.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
	}
}

func WithTaskList(client tasks.TaskServiceClient) RouterFunc {
	return func(router *mux.Router) {
		router.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				http.Error(w, "expected method: GET", http.StatusMethodNotAllowed)
				return
			}
			resp, err := client.List(r.Context(), &empty.Empty{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for {
				tsk, err := resp.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				str, err := util.PBJSONMarshaler.MarshalToString(tsk)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				_, err = io.WriteString(w, str)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

		})
	}
}

func WithTaskUpdate() RouterFunc {
	return func(router *mux.Router) {

	}
}

func WithTaskCreate(client tasks.TaskServiceClient) RouterFunc {
	return func(router *mux.Router) {
		router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "expected method: POST", http.StatusMethodNotAllowed)
				return
			}
			tsk := &tasks.Task{}
			defer r.Body.Close()
			err := util.PBJSONUnmarshaler.Unmarshal(r.Body, tsk)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			resp, err := client.Create(r.Context(), tsk)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			str, err := util.PBJSONMarshaler.MarshalToString(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.WriteString(w, str)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
}

func WithTaskDelete(client tasks.TaskServiceClient) RouterFunc {
	return func(router *mux.Router) {
		router.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "expected method: POST", http.StatusMethodNotAllowed)
				return
			}
			id := &tasks.IDRequest{}
			defer r.Body.Close()
			err := util.PBJSONUnmarshaler.Unmarshal(r.Body, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			resp, err := client.Delete(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			str, err := util.PBJSONMarshaler.MarshalToString(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.WriteString(w, str)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
}

func WithTaskRead(client tasks.TaskServiceClient) RouterFunc {
	return func(router *mux.Router) {
		router.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				http.Error(w, "expected method: POST", http.StatusMethodNotAllowed)
				return
			}
			id := &tasks.IDRequest{}
			defer r.Body.Close()
			err := util.PBJSONUnmarshaler.Unmarshal(r.Body, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			resp, err := client.Read(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			str, err := util.PBJSONMarshaler.MarshalToString(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			_, err = io.WriteString(w, str)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
}
