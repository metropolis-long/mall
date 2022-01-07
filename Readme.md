### run user service
```
go run ./user/rpc/user.go -f ./user/rpc/etc/user.yaml
```
### run order service
```
go run ./order/api/order.go -f ./order/api/etc/order.yaml
```

``` shell
    function fun(){
         echo "这是一句非常牛逼的代码";
    }
    fun();
```

```flow
st=>start: 开始
op=>operation: My Operation
cond=>condition: Yes or No?
e=>end
st->op->cond
cond(yes)->e
cond(no)->op
&```