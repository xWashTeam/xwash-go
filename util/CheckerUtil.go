package util

import (
  "github.com/tidwall/gjson"
  "github.com/go-resty/resty/v2"
  "time"
  "strconv"
)

type Machine struct {
  Name     string
  Building string
  Belong   string
  Location string
  Link     string
}

type QueryResult struct {
  Name        string
  Location    string
  Status      string
  UpdatedTime string
  Message     string
  RemainTime  int8
}

func Check(machine Machine) QueryResult {
  var qr QueryResult
  qr.Name = machine.Name
  qr.Location = machine.Location
  timeLocation, _ := time.LoadLocation("Local")
  qr.UpdatedTime = time.Now().In(timeLocation).Format("2006-01-02T15:04:05-0700")

  if (machine.Belong == "uClean") {
    qrt := extractUClean(requestUClean(machine))
    qr.Status = qrt.Status
    qr.Message = qrt.Message
    qr.RemainTime = qrt.RemainTime
  }
  return qr
}

func requestUClean(machine Machine) string {
  link := machine.Link
  url := "https://u.zhinengxiyifang.cn/api/Devices?filter={\"where\":{\"qrCode\":\"" + link + "\",\"isRemoved\":false},\"scope\":{\"fields\":[\"virtualId\",\"scanSelfClean\",\"hasAutoLaunchDevice\",\"autoLaunchDeviceOutOfStock\",\"isSlotMachine\",\"deviceTypeId\",\"online\",\"status\",\"boxTypeId\",\"sn\"]},\"include\":[{\"relation\":\"store\",\"scope\":{\"fields\":[\"isRemoved\",\"enable\"]}}]}"
  resp, _ := resty.New().R().Get(url)
  return resp.String()
}

func extractUClean(respString string) QueryResult {
  var qr QueryResult
  if !gjson.Valid(respString) {
    return qr
  }
  resp := gjson.Parse(respString[1 : len(respString)-1])
  const (
    NORMAL = iota
    USING
    UNKNOWN
  )
  status := resp.Get("status").Int()
  if status == 0 {
    qr.Status = "空闲"
  } else if status == 1 {
    qr.Status = "使用中"
  } else {
    qr.Status = "未知"
  }
  qr.RemainTime = int8(resp.Get("remainTime").Int())
  return qr
}

func (qr QueryResult) String() string{
  return "QueryResult: {" + 
    "\n  Name: " + qr.Name + 
    "\n  Location: " + qr.Location + 
    "\n  Status: " + qr.Status + 
    "\n  UpdatedTime: " + qr.UpdatedTime + 
    "\n  Message: " + qr.Message + 
    "\n  RemainTime: " + strconv.Itoa(int(qr.RemainTime)) +
    "\n}"
}
