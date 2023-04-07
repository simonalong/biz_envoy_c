package service

import (
	biz_config "biz-c-service/biz-config"
	"biz-c-service/pojo/domain"
	"context"
	"gorm.io/gorm/clause"
	netHttp "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/config"
	"github.com/isyscore/isc-gobase/extend/redis"
	"github.com/isyscore/isc-gobase/http"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/isyscore/isc-gobase/server/rsp"
	baseTime "github.com/isyscore/isc-gobase/time"
)

func Cf(c *gin.Context) {
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	logger.Info("接口 FrontCfOkOk %v", headers)

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	_, _, data, err := http.GetOfStandard(config.GetValueString("biz.url.f")+"/api/f/cf/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfOkOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.NewClient()
		ctx := context.Background()
		rdb.Set(ctx, "Cf", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}
	rsp.SuccessOfStandard(c, data)
}

func Bc(c *gin.Context) {
	//logger.Info("接口 FrontCfOkOk ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	_, _, data, err := http.GetOfStandard(config.GetValueString("biz.url.f")+"/api/f/cf/"+haveMysql, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfOkOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb, _ := redis.NewClient()
		ctx := context.Background()
		rdb.Set(ctx, "Cf", baseTime.TimeToStringYmdHms(time.Now()), time.Minute*1)
	}
	rsp.SuccessOfStandard(c, data)
}

func FrontCfOkFile(c *gin.Context) {
	//获取文件头
	file, err := c.FormFile("upload")
	if err != nil {
		rsp.FailedOfStandard(c, 500, "上传失败: "+err.Error())
		return
	}
	//获取文件名
	fileName := file.Filename
	//fmt.Println("文件名：", fileName)
	//保存文件到服务器本地
	//SaveUploadedFile(文件头，保存路径)
	pathPre := config.GetValueString("file.save.path")
	if err := c.SaveUploadedFile(file, pathPre+fileName); err != nil {
		rsp.FailedOfStandard(c, 500, "保存失败: "+err.Error())
		return
	}
	rsp.SuccessOfStandard(c, "ok")
}

func FrontCfOkOk(c *gin.Context) {
	////logger.Info("接口 FrontCfOkOk ")
	//headers := netHttp.Header{}
	//if c.GetHeader("x-request-id") != "" {
	//	headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
	//	headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
	//	headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
	//	headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
	//	headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
	//	headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
	//	headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	//}
	//
	//data, err := http.GetOfStandard(config.GetValueString("biz.url.f") + "/api/f/cf/ok/ok", headers, map[string]string{})
	//if err != nil {
	//	logger.Error("接口 FrontCfOkOk 报错", err)
	//	rsp.FailedOfStandard(c, 500, err.Error())
	//	return
	//}

	valueRspChan := make(chan *string)

	select {
	case <-c.Request.Context().Done():
		//cleanMachineConnect(c, appWatchKey)
		return
	case res := <-valueRspChan:
		{
			rsp.SuccessOfStandard(c, res)
		}
	}
	//rsp.SuccessOfStandard(c, "data")
}

func FrontCfOkStop(c *gin.Context) {
	//logger.Info("接口 FrontCfOkStop ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	_, _, data, err := http.GetOfStandard(config.GetValueString("biz.url.f")+"/api/f/cf/ok/stop", headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfOkStop 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}
	rsp.SuccessOfStandard(c, data)
}

func FrontCfStopOk(c *gin.Context) {
	//logger.Info("接口 FrontCfStopOk ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	_, _, data, err := http.GetOfStandard(config.GetValueString("biz.url.f")+"/api/f/cf/stop/ok", headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfStopOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}
	rsp.SuccessOfStandard(c, data)
}

func FrontCfStopStop(c *gin.Context) {
	//logger.Info("接口 FrontCfStopOk ")
	headers := netHttp.Header{}
	if c.GetHeader("x-request-id") != "" {
		headers["x-request-id"] = []string{c.GetHeader("x-request-id")}
		headers["x-b3-traceid"] = []string{c.GetHeader("x-b3-traceid")}
		headers["x-b3-spanid"] = []string{c.GetHeader("x-b3-spanid")}
		headers["x-b3-parentspanid"] = []string{c.GetHeader("x-b3-parentspanid")}
		headers["x-b3-sampled"] = []string{c.GetHeader("x-b3-sampled")}
		headers["x-b3-flags"] = []string{c.GetHeader("x-b3-flags")}
		headers["x-ot-span-context"] = []string{c.GetHeader("x-ot-span-context")}
	}

	haveRedis := c.Param("haveRedis")
	haveMysql := c.Param("haveMysql")
	haveEtcd := c.Param("haveEtcd")
	_, _, _, err := http.GetOfStandard(config.GetValueString("biz.url.f")+"/api/f/cf/"+haveRedis+"/"+haveMysql+"/"+haveEtcd, headers, map[string]string{})
	if err != nil {
		logger.Error("接口 FrontCfStopOk 报错", err)
		rsp.FailedOfStandard(c, 500, err.Error())
		return
	}

	if haveRedis == "true" {
		rdb := biz_config.RedisDb
		ctx := context.Background()
		rs := rdb.Set(ctx, "FrontCfOkOk", baseTime.TimeToStringYmdHms(time.Now()), time.Hour*1)
		if rs.Err() != nil {
			logger.Error("调用redis设置失败：%v", rs.Err().Error())
		} else {
			logger.Info("调用redis设置成功")
		}
	}

	if haveMysql == "true" {
		biz_config.Db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&domain.BizEnvoyC{ServiceName: "biz-c-service", Times: baseTime.TimeToStringYmdHms(time.Now())})
	}

	if haveEtcd == "true" {
		ctx := context.Background()
		biz_config.EtcdClient.Put(ctx, "biz-c-service.key", baseTime.TimeToStringYmdHms(time.Now()))
	}

	rsp.SuccessOfStandard(c, "ok")
}
