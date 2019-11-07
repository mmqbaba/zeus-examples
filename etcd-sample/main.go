package main

import (
	"context"
	"log"

	etcd "github.com/coreos/etcd/clientv3"
)

func main() {
	cfg := etcd.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	}
	client, err := etcd.New(cfg)
	if err != nil {
		log.Println(err, client)
		return
    }
    watcher := etcd.NewWatcher(client)
	rch := watcher.Watch(context.Background(), "/zeus/dzqz")
	for wresp := range rch {
	    for _, ev := range wresp.Events {
	        switch ev.Type {
                case etcd.EventTypePut:
                    if ev.IsCreate() {
                        log.Println("create")
                    }
                    if ev.IsModify() {
                        log.Println("update")
                    }
	                log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	            case etcd.EventTypeDelete:
	                log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
	        }
	    }
	}
}
