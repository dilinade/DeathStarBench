// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"media"
)

var _ = media.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void WritePlot(i64 req_id, i64 plot_id, string plot,  carrier)")
  fmt.Fprintln(os.Stderr, "  string ReadPlot(i64 req_id, i64 plot_id,  carrier)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := media.NewPlotServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "WritePlot":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "WritePlot requires 4 args")
      flag.Usage()
    }
    argvalue0, err435 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err435 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err436 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err436 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    arg438 := flag.Arg(4)
    mbTrans439 := thrift.NewTMemoryBufferLen(len(arg438))
    defer mbTrans439.Close()
    _, err440 := mbTrans439.WriteString(arg438)
    if err440 != nil { 
      Usage()
      return
    }
    factory441 := thrift.NewTJSONProtocolFactory()
    jsProt442 := factory441.GetProtocol(mbTrans439)
    containerStruct3 := media.NewPlotServiceWritePlotArgs()
    err443 := containerStruct3.ReadField4(context.Background(), jsProt442)
    if err443 != nil {
      Usage()
      return
    }
    argvalue3 := containerStruct3.Carrier
    value3 := argvalue3
    fmt.Print(client.WritePlot(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "ReadPlot":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "ReadPlot requires 3 args")
      flag.Usage()
    }
    argvalue0, err444 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err444 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1, err445 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err445 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg446 := flag.Arg(3)
    mbTrans447 := thrift.NewTMemoryBufferLen(len(arg446))
    defer mbTrans447.Close()
    _, err448 := mbTrans447.WriteString(arg446)
    if err448 != nil { 
      Usage()
      return
    }
    factory449 := thrift.NewTJSONProtocolFactory()
    jsProt450 := factory449.GetProtocol(mbTrans447)
    containerStruct2 := media.NewPlotServiceReadPlotArgs()
    err451 := containerStruct2.ReadField3(context.Background(), jsProt450)
    if err451 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Carrier
    value2 := argvalue2
    fmt.Print(client.ReadPlot(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
