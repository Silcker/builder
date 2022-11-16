package reader

import (
    "io/ioutil"
    "log"
    "gopkg.in/yaml.v3"
)
 
type conf struct {
   Hits int64 `yaml:"hits"`
   Time int64 `yaml:"time"`
}

func (c *conf) Config(path string) *conf {
   yamlFile, err := ioutil.ReadFile(path)
   if err != nil  { log.Printf("yamlFile.Get err   #%v ", err) }
   
   err = yaml.Unmarshal(yamlFile, c)
   if err != nil { log.Fatalf("Unmarshal: %v", err) }

   return c
}
