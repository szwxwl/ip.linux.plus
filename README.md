# ip.linux.plus
https://ip.linux.plus/

## 使用方法：

### 获取自己的ip：

**api:**

http[s]://ip.linux.plus/

**example:**

curl "ip.linux.plus"

curl "ip.linux.plus/?type=json"

curl "ip.linux.plus/?type=jsonp&callback=test"

### 获取ip的地理信息：

**api:**

http[s]://ip.linux.plus/search/

**example:**

curl "ip.linux.plus/search/"

curl "ip.linux.plus/search/?type=json"

curl "ip.linux.plus/search/?type=jsonp&callback=test"

curl "ip.linux.plus/search/1.1.1.1"

curl "ip.linux.plus/search/1.1.1.1?type=json"

curl "ip.linux.plus/search/1.1.1.1?type=jsonp&callback=test"
