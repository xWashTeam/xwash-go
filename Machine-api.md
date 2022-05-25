# 洗衣机api

目前支持的洗衣机品牌/厂商：
- <a href="#ucleanapp" target="_self">U净App版(UCleanApp)</a>
- <a href="#uclean" target="_self">U净网页版(UClean)</a>
- <a href="#zhuam" target="_self">Zhuam</a>
- <a href="#mplink" target="_self">Mplink</a>
- <a href="#washpayer" target="_self">Washpayer</a>
- ~~苏打校园(Sodalife)~~

> 苏打校园目前(2022-05-25)似乎已经不在SCNU里面了

## UCleanApp
Authorization: `Token`  
Method: `POST`  
BaseUrl: `https://phoenix.ujing.online:443/api/v1/devices/scanWasherCode`  
Header:
```json
{
  "Authorization": Token
}
```
Body:
```json
{
  "qrCode": 二维码链接
}
```
ReturnsValue:
  - data.result.createOrderEnabled: `Bool`
    - 说明：机器是否空闲
    - 取值：`true：空闲`, `false：被占用`
  - data.result.status: `Int`
    - 说明：机器状态
    - 取值：`1：使用中`, `其他：未知`
  - data.result.reason: `String`
    - 说明：机器状态说明

## UClean
Authorization: `None`  
Method: `GET`  
BaseUrl: `https://u.zhinengxiyifang.cn/api/Devices`  
Params: 
  - filter: `JsonString` 注：需要拼接字符串 param1 + qrcodeLink + param2
    - param1: `?filter={\"where\":{\"qrCode\":\"`
    - qrcodeLink: 解析洗衣机上二维码得到
    - param2: `\",\"isRemoved\":false},\"scope\":{\"fields\":[\"virtualId\",\"scanSelfClean\",\"hasAutoLaunchDevice\",\"autoLaunchDeviceOutOfStock\",\"isSlotMachine\",\"deviceTypeId\",\"online\",\"status\",\"boxTypeId\",\"sn\"]},\"include\":[{\"relation\":\"store\",\"scope\":{\"fields\":[\"isRemoved\",\"enable\"]}}]}`

ExampleUrl: `https://u.zhinengxiyifang.cn/api/Devices?filter={"where":{"qrCode":"https://q.ujing.com.cn/ucqrc/index.html?cd=0000000000000A0007555202104170176799","isRemoved":false},"scope":{"fields":["virtualId","scanSelfClean","hasAutoLaunchDevice","autoLaunchDeviceOutOfStock","isSlotMachine","deviceTypeId","online","status","boxTypeId","sn"]},"include":[{"relation":"store","scope":{"fields":["isRemoved","enable"]}}]}`

Returns: `JSON`  
ExampleReturns:
```json
[
{
  "no": "1902",
    "sn": "0000DA11138019701217072101850000",
    "qrCode": "https://q.ujing.com.cn/ucqrc/index.html?cd=0000000000000A0007555202104170176799",
    "deviceModel": "0000.da.31025",
    "alipayOnly": 2,
    "registerType": null,
    "brand": 1,
    "readableName": "校园机MB65-GF03W",
    "status": "0",
    "online": true,
    "statusAll": "{\"isLastFrame\":false,\"switch\":\"0\",\"wash\":\"0\",\"status\":\"9\",\"drain_water\":\"0\",\"wash_mode\":\"0\",\"time_remaining\":\"35\",\"fault_warn\":\"0\",\"device_type\":\"DA\",\"msg_type\":\"0404\",\"online\":null,\"deviceName\":\"\",\"signalStrength\":null,\"linkQuality\":0,\"signalCellID_high\":0,\"signalCellID_low\":0,\"signalECL\":0,\"signalPCI\":0,\"signalSNR\":0,\"totleNumber_high\":0,\"totleNumber_low\":0}",
    "remainTime": 35,
    "statusUpdateTime": "2022-05-25T13:08:08.000Z",
    "virtualId": "190c2563f74db39208e1e587fd780b88",
    "errorCode": "E3",
    "scanSelfClean": null,
    "errorTime": "2022-05-24T14:31:54.000Z",
    "endWorkTime": null,
    "shutDownTime": null,
    "offlineTime": "2022-05-22T07:16:30.000Z",
    "addTime": "2022-01-21T04:14:44.000Z",
    "isRemoved": false,
    "hasAutoLaunchDevice": null,
    "autoLaunchDeviceOutOfStock": null,
    "washTemperatureIsEnable": false,
    "isSlotMachine": null,
    "moduleType": 4,
    "isForceDelete": null,
    "usedCounts": null,
    "isScrapped": false,
    "isBindToAlipay": null,
    "bindToAlipayTime": null,
    "washCounts": 0,
    "washCoins": 0,
    "bleVersion": "8200001c",
    "firstBindTime": "2022-01-21T04:14:44.000Z",
    "nbSignal": "{\"device_type\":\"DA\",\"linkQuality\":240,\"msg_type\":\"7c7c\",\"signalCellID_high\":1877,\"signalCellID_low\":46928,\"signalECL\":0,\"signalPCI\":204,\"signalSNR\":265,\"signalStrength\":167,\"totleNumber_high\":0,\"totleNumber_low\":3822,\"updateTime\":\"2022-05-25T04:32:13.472Z\"}",
    "macAddress": "ECC57F519E2D",
    "deviceId2G": null,
    "bleBindStatus": 2,
    "runTimes": 196,
    "simcardId": 254138,
    "deadline": "2025-01-31T23:59:59.000Z",
    "deviceTypeId": 1,
    "errorCodeId": null,
    "franchiseeId": 8290,
    "applyFranchiseeId": null,
    "orderId": 0,
    "serviceSubjectId": 9931,
    "storeId": "5ba34423bca425dd42000005",
    "boxTypeId": null,
    "store": {
      "id": "5ba34423bca425dd42000005",
      "enable": true,
      "isRemoved": false
    }
}
]
```

ReturnsValue:
 - [0].status: `Int`
  - 说明：是否空闲
  - 取值: `0: 空闲`, `1: 使用中`, `其他: 未知`
 - [0].remainTime: `Int` 
  - 说明：剩余时间
  - 取值说明：只有在状态为使用中时剩余时间才是有效值

## Zhuam
Authorization: `None`  
Method: `POST`  
Url: `http://zhua.myclassphp.com/index.php?m=Home&c=User&a=getIndexData`  
Body:
```json
{
  "uid": "831342"
}
```
ReturnsValue:
  - status: `Int`
    - 说明：机器状态
    - 取值：`1: 空闲`, `0：被占用`

## Mplink
Authorization: `None`  
Method: `GET`  
Url: `http://view.mplink.cn/ajax/GetDeviceStatus.ashx`  
Params:
  deviceId: `Int`
    - 说明：id在洗衣机二维码链接中可以获取到

ExampleUrl: `http://view.mplink.cn/ajax/GetDeviceStatus.ashx?deviceID=ea85348b-6dce-49e0-bd05-8a628e0089f6`

ReturnsValue:
  Value: `Int`
    - 说明：机器状态
    - 取值：
      - 1：空闲
      - 0：被占用
      - 4：套餐缺货
      - 99：设备未就绪
      - 100：设备正在运行


## Washpayer
Washpayer洗衣机查询分为两个步骤：  
1. 查询洗衣机的devno(首次查询后可以存储)
2. 查询洗衣机状态

### 查询devno
Authorization: `cookie`  
Method: `GET`  
Url: 洗衣机二维码链接  
Header: 
```json
  "user-agent": "Mozilla/5.0 (Linux; Android 12; MI 8 Build/SQ1D.211205.016.A1; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3185 MMWEBSDK/20211202 Mobile Safari/537.36 MMWEBID/8395 MicroMessenger/8.0.18.2060(0x28001257) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64"
```
CookieExample: `MyUser_session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOENDY0NjA5MzcsInVzZXJfaWQiOiI2MTJjOWU3YjZmMjkyNTMyNTMxNjdkZWUiLCJleHAiOjE2NDkwODzd9.QPPtwQwIRYv1qfjYjMZjMBjZPnnHdKeLBaamOQ`

**注：devno需要从Response的Set-Cookie中获取！！！**  
Cookie名：`DEV_NO_COOKIE_NAME`  


### 查询洗衣机状态
Authorization: `cookie`  
Method: `POST`  
Url: `https://www.washpayer.com/user/startAction`  
Header: 
```json
  "user-agent": "Mozilla/5.0 (Linux; Android 12; MI 8 Build/SQ1D.211205.016.A1; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3185 MMWEBSDK/20211202 Mobile Safari/537.36 MMWEBID/8395 MicroMessenger/8.0.18.2060(0x28001257) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64"
```
Body:   
```json
{
  "devNo": "${devno}",
  "packageId": "2",
  "attachParas": {
    "startKey": "ojqSxwAC-NiAEMqO3asVzc0CO6OI-1630153349410-11770",
  }
}
```
> startKey是必填固定的参数  
> devno是第一步中获取的devno  
> cookie例子：`jwt_auth_domain=MyUser MyUser_session_id=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDY0NjA5MzcsInVzZXJfaWQiOiI2MTJjOWU3YjZmMjkyNTMyNTMxNjdkZWUiLCJleHAiOjE2NDkwODE3Mzd9.n1r3wQwIbRYv1qfjYjMZjMBjZPnnHdKeLBaamOQ`

ReturnsValue:
  - result
    - 说明：洗衣机状态
    - 取值：`999: 正常`
    - **注：Washpayer洗衣机在使用中时，请求会超时，需要以是否超时此来作为判断依据**




