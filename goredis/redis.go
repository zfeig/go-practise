package main
import (
    "fmt"
    "github.com/alphazero/Go-Redis"
)

func main() {
    //DefaultSpec()创建一个连接
    //选择host,若需要auth,则password填写
    //spec        := redis.DefaultSpec().Host("192.168.1.111").Db(0).Password("");
    //若连接的本机redis-server,则host可以省略
    spec        := redis.DefaultSpec().Db(0).Password("");
    client, err := redis.NewSynchClientWithSpec (spec);
    
    if err != nil {
        fmt.Println("Connect redis server fail");
        return
    }

    dbkey := "go:test";
    value :=[]byte("Hello world!");
    client.Set(dbkey, value);
    
    getValue ,err:= client.Get(dbkey);
    if err != nil {
        fmt.Println("Get Key fail");
        return
    } else {
        str := string(getValue);
        fmt.Println(str);
    }
    
}
