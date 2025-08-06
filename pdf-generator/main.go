package main

import (
    "encoding/json"
    "log"
    "net/http"
    "bytes"
    "io/ioutil"
)

type CertificateRequest struct {
    Name     string `json:"name"`
    Company  string `json:"company"`
    Position string `json:"position"`
    Duration string `json:"duration"`
}


