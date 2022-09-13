# tradeEngine

1. build and deploy

docker build . -t exam  
docker run -p 8080:8080 -d exam

2. api test

Post 127.0.0.1:8080/order

body 
```
{
    "type": 0,
    "quantity": 3,
    "price": 35
}
```

3. test script
```
api/order/service_test.go
TestTradingAndUpdateOrder 
TestTradingAndUpdateOrderConcurrent 
```
測試並發與非並發狀況


4. 

model/order 
紀錄order訊息

model/trade
紀錄成功的交易資訊 

DataStore 模擬資料儲存

OrderMap 用 key-value 儲存所有 order 資料

PendingOrderMap 則紀錄目前還未被完成的 order 

PendingOrderMap 第一個 key 使用 type+price 可以快速的 match 同樣價位的 order 

第二個 key 使用 timestamp 可以快速取得最小時間達成 FIFO

將每次資料操作使用 lock 保護 達成 concurrent
