# xwash-go
xWash应用的golang后端

# api
url: `/api/:building`  
params:   
  - `building`: 对应的宿舍楼，例如`d19`, `xi1`, `xi2`  

return example:  
```json
{
  "data": {
    "d19_E_1.5": {
            "date": "Wed May 25 14:07:28 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九东边1.5楼",
            "status": "USING",
            "timestamp": "1653487648840"
        },
        "d19_W_1.5": {
            "date": "Wed May 25 14:07:28 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九西边1.5楼",
            "status": "USING",
            "timestamp": "1653487648926"
        },
        "d19_E_2.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 35,
            "message": "",
            "name": "校园机MB65-GF03W",
            "location": "东十九东边2.5楼",
            "status": "AVAILABLE",
            "timestamp": "1653487649197"
        },
        "d19_W_2.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九西边2.5楼",
            "status": "USING",
            "timestamp": "1653487649299"
        },
        "d19_E_3.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 29,
            "message": "",
            "name": "校园机MB65-GF03W",
            "location": "东十九东边3.5楼",
            "status": "USING",
            "timestamp": "1653487649381"
        },
        "d19_W_3.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 15,
            "message": "",
            "name": "校园机MB65-GF03W",
            "location": "东十九西边3.5楼",
            "status": "USING",
            "timestamp": "1653487649457"
        },
        "d19_E_4.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九东边4.5楼",
            "status": "USING",
            "timestamp": "1653487649548"
        },
        "d19_W_4.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九西边4.5楼",
            "status": "USING",
            "timestamp": "1653487649635"
        },
        "d19_E_5.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 0,
            "message": "当前设备未绑定",
            "name": "",
            "location": "东十九东边5.5楼",
            "status": "UNKNOWN",
            "timestamp": "1653487649734"
        },
        "d19_W_5.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九西边5.5楼",
            "status": "USING",
            "timestamp": "1653487649850"
        },
        "d19_E_6.5": {
            "date": "Wed May 25 14:07:29 UTC 2022",
            "remainTime": 35,
            "message": "",
            "name": "校园机MB65-GF03W",
            "location": "东十九东边6.5楼",
            "status": "AVAILABLE",
            "timestamp": "1653487649933"
        },
        "d19_W_6.5": {
            "date": "Wed May 25 14:07:30 UTC 2022",
            "remainTime": 0,
            "message": "此设备正在运行中，请使用其他设备",
            "name": "",
            "location": "东十九西边6.5楼",
            "status": "USING",
            "timestamp": "1653487650043"
        }
  },
  "message": ""

}
```

# 洗衣机api

目前支持的洗衣机品牌/厂商：
- U净网页版(UClean)
- U净App版(UCleanApp)
- Zhuam
- Mplink
- Washpayer
- ~~苏打校园(Sodalife)~~

洗衣机接口详细参数信息[点此查看](https://github.com/xWashTeam/xwash-go/blob/main/Machine-api.md)

> 苏打校园目前(2022-05-25)似乎已经不在SCNU里面了


