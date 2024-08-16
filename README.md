# scanport

用来扫描对应ip开放的端口情况

## 构建

```
git clone https://github.com/DashingYoungMan/scanport.git

cd scanport

go build
```

## 使用方法

```
scanport -i 要扫描的ip -s 开始端口 -e 结束端口
```

## 参数

```
-i 127.0.0.1  # ip地址
-s  # 扫描开始端口
-e  # 扫描结束端口
```
