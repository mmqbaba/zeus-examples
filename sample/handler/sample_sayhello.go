package handler

import (
	"context"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	// "zeus-examples/sample/logic/user"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) SayHello(ctx context.Context, req *gomicro.Request, rsp *gomicro.Reply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SayHello")

	// info, err := user.GetInfo(ctx, "001")
	// if err != nil {
	// 	logger.Error("user.GetInfo err:", err)
	// 	return
	// }
	// rsp.Message = "hello world, " + info + "."

	rdc, err := zeusctx.ExtractRedis(ctx)
	if err != nil {
		return
	}
	r := rdc.Get("001")
	if err = r.Err(); err != nil && err != redis.Nil {
		return
	} else {
		err = nil
	}

	logger.Debug(r.Val())

	mgoc, err := zeusctx.ExtractMongo(ctx)
	if err != nil {
		logger.Error(err)
		return
	}

	coll := mgoc.DB("example").Collection("user")
	result := &struct {
		ID   string `json:"id,omitempty" bson:"id,omitempty"`
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	}{}
	filter := bson.M{"name": "mark"}
	err = coll.FindOne(ctx, filter).Decode(result)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(result)

	rsp.Message = "hello world, " + req.Name + "."
	return
}
