package checker

import("context";"net";"net/http";"net/http/httptest";"testing";"time")
func TestHTTP(t *testing.T){s:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.WriteHeader(200)}));defer s.Close();r:=HTTP(context.Background(),s.URL,time.Second);if !r.Up||r.StatusCode!=200{t.Fatalf("%+v",r)}}
func TestTCP(t *testing.T){l,e:=net.Listen("tcp","127.0.0.1:0");if e!=nil{t.Fatal(e)};defer l.Close();go func(){c,_:=l.Accept();if c!=nil{c.Close()}}();r:=TCP(context.Background(),l.Addr().String(),time.Second);if !r.Up{t.Fatalf("%+v",r)}}
