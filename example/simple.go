package main

import (
	"fmt"
	"github.com/py60800/tuya"
	"time"
)

// dual pole electrical light switch or
// tuya multi plug power strip with dps 1-8 (same gwId, same key)
var conf2 = `[
    {"gwId":"1582850884f3eb30128e",
     "key":"XXXXXXXXXXX",
     "type":"Switch",
     "name":"sw1",
     "dps": 1 },
    {"gwId":"1582850884f3eb30128e",
     "key":"XXXXXXXXXXX",
     "type":"Switch",
     "name":"sw2",
     "dps": 2 }
     ]`

//single pole electrical
var conf1 = `[
    {"gwId":"1582850884f3eb30128e",
     "key":"XXXXXXXXXXX",
     "type":"Switch",
     "name":"sw1" }
     ]`

func main() {
	dm := tuya.NewDeviceManager(conf1)
	b1, _ := dm.GetDevice("sw1")
	sw1 := b1.(tuya.Switch)
	b, err := sw1.SetW(true, 10*time.Second)
	if err != nil {
		fmt.Println("Set error:", err)
	} else {
		fmt.Println("Set OK", b)
	}
	time.Sleep(2 * time.Second)
	b, err = sw1.SetW(false, 10*time.Second)
	if err != nil {
		fmt.Println("Set error:", err)
	} else {
		fmt.Println("Set OK", b)
	}
}
