package dao

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

type ConfigFile struct {
    // private file file
    DatabaseServer string `json:"database_server"`
    // local addr to listen and serve, default is 127.0.0.1:1315
    DatabaseName string `json:"database_name"`
    // local addr to listen and serve, default is 127.0.0.1:1316
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

// Load file from path
func NewConfigFile(path string) (self *ConfigFile, err error) {

    self = &ConfigFile{DatabaseServer: "127.0.0.1", DatabaseName: "kanal" }

    file_exists, err := exists(path)

    if !file_exists {
        return self, err
    }

    buf_read, err := ioutil.ReadFile(path)
    if err != nil {
        return self, err
    }

    err = json.Unmarshal(buf_read, self)
    if err != nil {
        return
    }

    return
}