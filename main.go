//  2018-02-17 02:49
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/allegro/bigcache"
	"github.com/google/gops/agent"
	"github.com/koding/multiconfig"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/afero"
	"github.com/tsingson/fastweb/utils"
	"github.com/tsingson/fastweb/zaplogger"
	"go.uber.org/zap"
)

const (
	PidFileName = "proxy-bigcache-pid"
)

var (
	// 	output = log.New(os.Stdout, "", 0)
	log               *zap.Logger
	gc                *bigcache.BigCache
	config            *Config
	path, currentPath string
	apkEpgConfig      map[string]ApkEpgConfig
	// apkEpgMap         sync.Map
)

func init() {
	afs := afero.NewOsFs()
	// get run path
	{
		path, _ = utils.GetCurrentExecDir()
		configFile := path + "/vkaaa-config-apk.json"
		// load config
		m := &multiconfig.JSONLoader{Path: configFile}
		config = new(Config)
		var err error
		// Populated the config struct
		err = m.Load(config)
		if err != nil {
			fmt.Println("配置文件:  " + configFile + "  无法读取")
			os.Exit(1)
		}
		apkEpgConfig = make(map[string]ApkEpgConfig, 4)
		if len(config.ApkEpgConfig) > 0 {
			for _, b := range config.ApkEpgConfig {
				apkEpgConfig[b.ApkType] = b
				// 	apkEpgMap.Store(b.ApkType, b)
			}
		}

	}
	{
		logPath := path + "/log"
		check, _ := afero.DirExists(afs, logPath)
		if !check {
			afs.MkdirAll(logPath, 0755)
		}

		// log setup

		log = zaplogger.NewZapLog(logPath, config.AaaConfig.LogFilenamePrefix, config.Debug)
	}
	if err := agent.Listen(agent.Options{ConfigDir: currentPath}); err != nil {
		log.Fatal("google gops Init Fail")
	}
}

func main() {
	//
	runtime.GOMAXPROCS(128)

	cntxt := &daemon.Context{
		PidFileName: PidFileName,
		PidFilePerm: 0644,
		LogFileName: path + "/log-proxy.log",
		LogFilePerm: 0640,
		WorkDir:     path,
		Umask:       027,
		Args:        []string{"proxy-bigcache"},
	}

	d, err1 := cntxt.Reborn()
	if err1 != nil {
		log.Fatal("cat's reborn ", zap.Error(err1))
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Info("- - - - - - - - - - - - - - -")
	log.Info("daemon started")
	{
		/**
			 // number of shards (must be a power of 2)
				Shards:1024,
		    // time after which entry can be evicted
				LifeWindow: 10 * time.Minute,
			// rps * lifeWindow, used only in initial memory allocation
				MaxEntriesInWindow: 1000 * 10 * 60,
			// max entry size in bytes, used only in initial memory allocation
				MaxEntrySize: 500,
			// prints information about additional memory allocation
				Verbose: true,
			// cache will not allocate more memory than this limit, value in MB
			// if value is reached then the oldest entries can be overridden for the new ones
			// 0 value means no size limit
				HardMaxCacheSize: 8192,
			// callback fired when the oldest entry is removed because of its expiration time or no space left
			// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
			// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
				OnRemove: nil,
			// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
			// for the new entry, or because delete was called. A constant representing the reason will be passed through.
			// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
			// Ignored if OnRemove is specified.
				OnRemoveWithReason: nil,
		*/
		//

		var err error
		gc, err = bigcache.NewBigCache(bigcache.Config{
			Shards:             1024,             // number of shards (must be a power of 2)
			LifeWindow:         30 * time.Minute, // time after which entry can be evicted
			MaxEntriesInWindow: 1000 * 30 * 60,   // rps * lifeWindow
			MaxEntrySize:       1024 * 1024 * 3,  // max entry size in bytes, used only in initial memory allocation
			Verbose:            true,             // prints information about additional memory allocation
		})
		if err != nil {
			log.Fatal("cache Init Error", zap.Error(err))
			os.Exit(1)
		}
		//

	}

	FasthttpServ(config.AaaConfig.ServerPort, log)
	// Wait forever.

	select {}

}

/**
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
	wg.Add(1)
	go func(i int) {
		log.Println(i)
		wg.Done()
	}(i)
}
wg.Wait()
*/

/**
func authPost() {

	user_id := "583f38b8-2770-4ed1-87ab-21975123f75a"

	mac_address_1 := "9c:f8:db:05:c5:f9"
	mac_address_2 := "9c:f8:db:05:c5:f0"
	release_sn := "vk-v1.0.0-20180401"

	url := "http://50.7.101.250/rpc/auth"

	var requestLogin aaacrypt.ApkLoginAuth
	requestLogin.UserID = user_id
	requestLogin.MacAddress1 = mac_address_1
	requestLogin.MacAddress2 = mac_address_2
	requestLogin.ReleaseSn = release_sn
	postBodyByte, _ := jsoniter.Marshal(requestLogin)
	// proxy request
	resp, err1 := FastPostJson(url, postBodyByte, 15*time.Second)
	if err1 != nil {
		litter.Dump(err1)
	}

	statusCode := resp.StatusCode()
	if statusCode == 200 {
		body := resp.Body()

		kk := make([]RespItem, 1)
		ee := jsoniter.Unmarshal(body, &kk)
		if ee != nil {
			return
		}
		litter.Dump(kk[0])

	}

}
*/
