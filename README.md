# requests

## 安装
~~~sh
go get -u github.com/ggdream/requests
~~~

## 使用
~~~go
package main

import "github.com/ggdream/requests"


func main(){
    res := requests.Get("https://www.google.cn", nil)
    defer res.Close()

    println(res.Text())
}
~~~