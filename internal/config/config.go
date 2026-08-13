package config

import ("os"; "gopkg.in/yaml.v3")

type Config struct {
    Server struct { Address string `yaml:"address"` } `yaml:"server"`
    Database string `yaml:"database"`
    Targets []Target `yaml:"targets"`
}
type Target struct {
    Name string `yaml:"name"`
    Type string `yaml:"type"`
    Address string `yaml:"address"`
    IntervalSeconds int `yaml:"interval_seconds"`
    TimeoutSeconds int `yaml:"timeout_seconds"`
}
func Load(path string)(Config,error) {
    b,err:=os.ReadFile(path); if err!=nil{return Config{},err}
    var c Config
    if err=yaml.Unmarshal(b,&c);err!=nil{return Config{},err}
    if c.Server.Address==""{c.Server.Address=":8080"}
    if c.Database==""{c.Database="data/openping.db"}
    return c,nil
}
