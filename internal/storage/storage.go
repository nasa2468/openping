package storage

import ("database/sql"; "os"; "path/filepath"; "time"; _ "github.com/mattn/go-sqlite3")

type Store struct{DB *sql.DB}
type Check struct{ID int64 `json:"id"`; Target string `json:"target"`; Up bool `json:"up"`; LatencyMs int64 `json:"latency_ms"`; StatusCode int `json:"status_code,omitempty"`; Error string `json:"error,omitempty"`; CheckedAt time.Time `json:"checked_at"`}
type Summary struct{Target string `json:"target"`; Up bool `json:"up"`; Checks int64 `json:"checks"`; UptimePct float64 `json:"uptime_percent"`; AvgLatencyMs float64 `json:"avg_latency_ms"`; LastCheck time.Time `json:"last_check"`}

func Open(path string)(*Store,error){
    if d:=filepath.Dir(path);d!="."{if err:=os.MkdirAll(d,0755);err!=nil{return nil,err}}
    db,err:=sql.Open("sqlite3",path);if err!=nil{return nil,err}
    _,err=db.Exec(`CREATE TABLE IF NOT EXISTS checks(id INTEGER PRIMARY KEY AUTOINCREMENT,target TEXT NOT NULL,up INTEGER NOT NULL,latency_ms INTEGER NOT NULL,status_code INTEGER NOT NULL,error TEXT NOT NULL,checked_at DATETIME NOT NULL); CREATE INDEX IF NOT EXISTS idx_checks_target_time ON checks(target,checked_at);`)
    if err!=nil{db.Close();return nil,err};return &Store{DB:db},nil
}
func(s *Store)Add(c Check)error{u:=0;if c.Up{u=1};_,err:=s.DB.Exec(`INSERT INTO checks(target,up,latency_ms,status_code,error,checked_at)VALUES(?,?,?,?,?,?)`,c.Target,u,c.LatencyMs,c.StatusCode,c.Error,c.CheckedAt.UTC());return err}
func(s *Store)Recent(limit int)([]Check,error){
    if limit<=0||limit>1000{limit=100};rows,err:=s.DB.Query(`SELECT id,target,up,latency_ms,status_code,error,checked_at FROM checks ORDER BY id DESC LIMIT ?`,limit);if err!=nil{return nil,err};defer rows.Close()
    var out []Check;for rows.Next(){var c Check;var u int;if err:=rows.Scan(&c.ID,&c.Target,&u,&c.LatencyMs,&c.StatusCode,&c.Error,&c.CheckedAt);err!=nil{return nil,err};c.Up=u==1;out=append(out,c)};return out,rows.Err()
}
func(s *Store)Summaries()([]Summary,error){
    rows,err:=s.DB.Query(`SELECT target,COUNT(*),MAX(checked_at),CAST(SUM(up) AS REAL)*100.0/COUNT(*),AVG(latency_ms),(SELECT up FROM checks c2 WHERE c2.target=c1.target ORDER BY checked_at DESC LIMIT 1) FROM checks c1 GROUP BY target ORDER BY target`)
    if err!=nil{return nil,err};defer rows.Close();var out []Summary
    for rows.Next(){var x Summary;var u int;if err:=rows.Scan(&x.Target,&x.Checks,&x.LastCheck,&x.UptimePct,&x.AvgLatencyMs,&u);err!=nil{return nil,err};x.Up=u==1;out=append(out,x)};return out,rows.Err()
}
func(s *Store)Incidents(limit int)([]Check,error){
    if limit<=0{limit=100};rows,err:=s.DB.Query(`SELECT id,target,up,latency_ms,status_code,error,checked_at FROM checks WHERE up=0 ORDER BY id DESC LIMIT ?`,limit);if err!=nil{return nil,err};defer rows.Close()
    var out []Check;for rows.Next(){var c Check;var u int;if err:=rows.Scan(&c.ID,&c.Target,&u,&c.LatencyMs,&c.StatusCode,&c.Error,&c.CheckedAt);err!=nil{return nil,err};c.Up=u==1;out=append(out,c)};return out,rows.Err()
}
