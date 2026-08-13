package api

import ("encoding/csv"; "encoding/json"; "net/http"; "strconv"; "time"; "github.com/nasa2468/openping/internal/storage"; "github.com/prometheus/client_golang/prometheus")

type Server struct{Store *storage.Store}
var checksTotal=prometheus.NewCounterVec(prometheus.CounterOpts{Name:"openping_checks_total",Help:"Total checks recorded by OpenPing."},[]string{"target","status"})
var latency=prometheus.NewGaugeVec(prometheus.GaugeOpts{Name:"openping_latency_ms",Help:"Latest latency in milliseconds."},[]string{"target"})
func init(){prometheus.MustRegister(checksTotal,latency)}
func(s *Server)Register(mux *http.ServeMux){mux.HandleFunc("/healthz",func(w http.ResponseWriter,_ *http.Request){w.WriteHeader(http.StatusOK);w.Write([]byte("ok\n"))});mux.HandleFunc("/api/targets",s.targets);mux.HandleFunc("/api/recent",s.recent);mux.HandleFunc("/api/incidents",s.incidents);mux.HandleFunc("/api/export.csv",s.exportCSV);mux.Handle("/metrics",prometheus.Handler())}
func(s *Server)targets(w http.ResponseWriter,_ *http.Request){x,e:=s.Store.Summaries();if e!=nil{http.Error(w,e.Error(),500);return};writeJSON(w,x)}
func(s *Server)recent(w http.ResponseWriter,r *http.Request){n,_:=strconv.Atoi(r.URL.Query().Get("limit"));if n==0{n=100};x,e:=s.Store.Recent(n);if e!=nil{http.Error(w,e.Error(),500);return};writeJSON(w,x)}
func(s *Server)incidents(w http.ResponseWriter,_ *http.Request){x,e:=s.Store.Incidents(100);if e!=nil{http.Error(w,e.Error(),500);return};writeJSON(w,x)}
func(s *Server)exportCSV(w http.ResponseWriter,_ *http.Request){x,e:=s.Store.Recent(1000);if e!=nil{http.Error(w,e.Error(),500);return};w.Header().Set("Content-Type","text/csv");w.Header().Set("Content-Disposition",`attachment; filename="openping-checks.csv"`);c:=csv.NewWriter(w);c.Write([]string{"id","target","up","latency_ms","status_code","error","checked_at"});for _,v:=range x{c.Write([]string{strconv.FormatInt(v.ID,10),v.Target,strconv.FormatBool(v.Up),strconv.FormatInt(v.LatencyMs,10),strconv.Itoa(v.StatusCode),v.Error,v.CheckedAt.Format(time.RFC3339)})};c.Flush()}
func RecordMetrics(c storage.Check){st:="down";if c.Up{st="up"};checksTotal.WithLabelValues(c.Target,st).Inc();latency.WithLabelValues(c.Target).Set(float64(c.LatencyMs))}
func writeJSON(w http.ResponseWriter,v any){w.Header().Set("Content-Type","application/json");json.NewEncoder(w).Encode(v)}
