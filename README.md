# jsontosql
json to sql tool

# Demo1

json file, test.json:

``` 
{
    "page_position_code": "121",
    "track_event_type": "test",
    "unique_session_id": "swdwd-74ef-2222-aba6-6f17fe64abb7",
    "experiments": [
        "invitee.b",
        "app_price_test_v1.b",
        "repurchase_v2.a",
        "text_subscribe_us.a",
        "online_service.a",
        "abtest_check.a",
        "shipping_fee_us_v1.a"
    ],
    "operation_system": "other",
    "last_url": "http://www.demo.com/product_detail/1236",
    "user_id": 123,
    "new_adlink_id": 0,
    "event_id": 0,
    "reference_type": "product",
    "source": "android",
    "device_language": "en",
    "cls_created_at": "2011-04-07 08:20:01",
    "page_position_page_name": "product_detail",
    "app_version": "1.0",
    "adlink_id": 123,
    "channel": "Facebook",
    "website": "us-android",
    "page_last_url": "http://www.demo.com/product_detail/391912",
    "ip_address": "127.0.0.1",
    "landing_page": "http://www.demo.com/product_detail/456744",
    "home_page_position_id": "home-19_17827_picturehotspot_6-630130_002",
    "product_id": 12345,
    "url": "http://www.demo.com/product_detail/50302?&site_abb=us",
    "country": "United States",
    "created_at": "2011-04-07 08:20:01",
    "position_content": "",
    "session_id": "",
    "website_abb": "us",
    "page_position_type": "50302",
    "page_position_id": "product_detail-50302-121",
    "page_url": "http://www.demo.com/product_detail/50302?&site_abb=us"
}
```

convert to sql:

``` 
(base) ➜  ~ jsontosql-osx -json_path ~/Desktop/test.json
json-----

**********JSON CONVERT TO CREATE TABLE SQL **********

CREATE TABLE IF NOT EXISTS jsonTableName (
`app_version` String,
`page_position_id` String,
`experiments` Array,
`new_adlink_id` UInt64,
`source` String,
`page_url` String,
`unique_session_id` String,
`last_url` String,
`url` String,
`ip_address` String,
`page_position_type` String,
`page_last_url` String,
`product_id` UInt64,
`session_id` String,
`track_event_type` String,
`device_language` String,
`website` String,
`created_at` String,
`position_content` String,
`operation_system` String,
`event_id` UInt64,
`reference_type` String,
`country` String,
`website_abb` String,
`page_position_code` String,
`cls_created_at` String,
`channel` String,
`home_page_position_id` String,
`page_position_page_name` String,
`adlink_id` UInt64,
`landing_page` String,
`user_id` UInt64
)

******************************

```

# Demo 2

``` 
(base) ➜  ~ jsontosql-osx -json '{"name":"xm","age":23}'
json-----

**********JSON CONVERT TO CREATE TABLE SQL **********

CREATE TABLE IF NOT EXISTS jsonTableName (
`name` String,
`age` UInt64
)

******************************
```

# Build

### 1. Mac

```
go build -ldflags -w -o jsontosql-osx
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags -w -o jsontosql-linux
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags -w -o jsontosql-win.exe
```

### 2. Linux

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 
```

### 3. Win

```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build 
fi  
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build 
```

Optimize packet size:

`go build -ldflags -w`

# Download

[MacOsx](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-osx)

[Linux](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-linux)

[Windows](https://github.com/heyuan110/jsontosql/raw/main/jsontosql-win.exe)