// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package feed

import (
	"context"
	"douyin/shared/rpc/kitex_gen/base"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type DouyinFeedRequest struct {
	LatestTime int64 `thrift:"latest_time,1" frugal:"1,default,i64" json:"latest_time"`
	UserId     int64 `thrift:"user_id,2" frugal:"2,default,i64" json:"user_id"`
}

func NewDouyinFeedRequest() *DouyinFeedRequest {
	return &DouyinFeedRequest{}
}

func (p *DouyinFeedRequest) InitDefault() {
	*p = DouyinFeedRequest{}
}

func (p *DouyinFeedRequest) GetLatestTime() (v int64) {
	return p.LatestTime
}

func (p *DouyinFeedRequest) GetUserId() (v int64) {
	return p.UserId
}
func (p *DouyinFeedRequest) SetLatestTime(val int64) {
	p.LatestTime = val
}
func (p *DouyinFeedRequest) SetUserId(val int64) {
	p.UserId = val
}

var fieldIDToName_DouyinFeedRequest = map[int16]string{
	1: "latest_time",
	2: "user_id",
}

func (p *DouyinFeedRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinFeedRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinFeedRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.LatestTime = v
	}
	return nil
}

func (p *DouyinFeedRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserId = v
	}
	return nil
}

func (p *DouyinFeedRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("douyin_feed_request"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinFeedRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("latest_time", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.LatestTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinFeedRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_id", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinFeedRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinFeedRequest(%+v)", *p)
}

func (p *DouyinFeedRequest) DeepEqual(ano *DouyinFeedRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.LatestTime) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserId) {
		return false
	}
	return true
}

func (p *DouyinFeedRequest) Field1DeepEqual(src int64) bool {

	if p.LatestTime != src {
		return false
	}
	return true
}
func (p *DouyinFeedRequest) Field2DeepEqual(src int64) bool {

	if p.UserId != src {
		return false
	}
	return true
}

type DouyinFeedResponse struct {
	StatusCode int32         `thrift:"status_code,1" frugal:"1,default,i32" json:"status_code"`
	StatusMsg  string        `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
	VideoList  []*base.Video `thrift:"video_list,3" frugal:"3,default,list<base.Video>" json:"video_list"`
	NextTime   int64         `thrift:"next_time,4" frugal:"4,default,i64" json:"next_time"`
}

func NewDouyinFeedResponse() *DouyinFeedResponse {
	return &DouyinFeedResponse{}
}

func (p *DouyinFeedResponse) InitDefault() {
	*p = DouyinFeedResponse{}
}

func (p *DouyinFeedResponse) GetStatusCode() (v int32) {
	return p.StatusCode
}

func (p *DouyinFeedResponse) GetStatusMsg() (v string) {
	return p.StatusMsg
}

func (p *DouyinFeedResponse) GetVideoList() (v []*base.Video) {
	return p.VideoList
}

func (p *DouyinFeedResponse) GetNextTime() (v int64) {
	return p.NextTime
}
func (p *DouyinFeedResponse) SetStatusCode(val int32) {
	p.StatusCode = val
}
func (p *DouyinFeedResponse) SetStatusMsg(val string) {
	p.StatusMsg = val
}
func (p *DouyinFeedResponse) SetVideoList(val []*base.Video) {
	p.VideoList = val
}
func (p *DouyinFeedResponse) SetNextTime(val int64) {
	p.NextTime = val
}

var fieldIDToName_DouyinFeedResponse = map[int16]string{
	1: "status_code",
	2: "status_msg",
	3: "video_list",
	4: "next_time",
}

func (p *DouyinFeedResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_DouyinFeedResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *DouyinFeedResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.StatusCode = v
	}
	return nil
}

func (p *DouyinFeedResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.StatusMsg = v
	}
	return nil
}

func (p *DouyinFeedResponse) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.VideoList = make([]*base.Video, 0, size)
	for i := 0; i < size; i++ {
		_elem := base.NewVideo()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.VideoList = append(p.VideoList, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *DouyinFeedResponse) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.NextTime = v
	}
	return nil
}

func (p *DouyinFeedResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("douyin_feed_response"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *DouyinFeedResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.StatusCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *DouyinFeedResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status_msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.StatusMsg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *DouyinFeedResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("video_list", thrift.LIST, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.VideoList)); err != nil {
		return err
	}
	for _, v := range p.VideoList {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *DouyinFeedResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("next_time", thrift.I64, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.NextTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *DouyinFeedResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DouyinFeedResponse(%+v)", *p)
}

func (p *DouyinFeedResponse) DeepEqual(ano *DouyinFeedResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.StatusCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.StatusMsg) {
		return false
	}
	if !p.Field3DeepEqual(ano.VideoList) {
		return false
	}
	if !p.Field4DeepEqual(ano.NextTime) {
		return false
	}
	return true
}

func (p *DouyinFeedResponse) Field1DeepEqual(src int32) bool {

	if p.StatusCode != src {
		return false
	}
	return true
}
func (p *DouyinFeedResponse) Field2DeepEqual(src string) bool {

	if strings.Compare(p.StatusMsg, src) != 0 {
		return false
	}
	return true
}
func (p *DouyinFeedResponse) Field3DeepEqual(src []*base.Video) bool {

	if len(p.VideoList) != len(src) {
		return false
	}
	for i, v := range p.VideoList {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
func (p *DouyinFeedResponse) Field4DeepEqual(src int64) bool {

	if p.NextTime != src {
		return false
	}
	return true
}

type FeedService interface {
	Feed(ctx context.Context, request *DouyinFeedRequest) (r *DouyinFeedResponse, err error)
}

type FeedServiceClient struct {
	c thrift.TClient
}

func NewFeedServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FeedServiceClient {
	return &FeedServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewFeedServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FeedServiceClient {
	return &FeedServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewFeedServiceClient(c thrift.TClient) *FeedServiceClient {
	return &FeedServiceClient{
		c: c,
	}
}

func (p *FeedServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *FeedServiceClient) Feed(ctx context.Context, request *DouyinFeedRequest) (r *DouyinFeedResponse, err error) {
	var _args FeedServiceFeedArgs
	_args.Request = request
	var _result FeedServiceFeedResult
	if err = p.Client_().Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type FeedServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      FeedService
}

func (p *FeedServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FeedServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FeedServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFeedServiceProcessor(handler FeedService) *FeedServiceProcessor {
	self := &FeedServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Feed", &feedServiceProcessorFeed{handler: handler})
	return self
}
func (p *FeedServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type feedServiceProcessorFeed struct {
	handler FeedService
}

func (p *feedServiceProcessorFeed) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FeedServiceFeedArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Feed", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FeedServiceFeedResult{}
	var retval *DouyinFeedResponse
	if retval, err2 = p.handler.Feed(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Feed: "+err2.Error())
		oprot.WriteMessageBegin("Feed", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Feed", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type FeedServiceFeedArgs struct {
	Request *DouyinFeedRequest `thrift:"request,1" frugal:"1,default,DouyinFeedRequest" json:"request"`
}

func NewFeedServiceFeedArgs() *FeedServiceFeedArgs {
	return &FeedServiceFeedArgs{}
}

func (p *FeedServiceFeedArgs) InitDefault() {
	*p = FeedServiceFeedArgs{}
}

var FeedServiceFeedArgs_Request_DEFAULT *DouyinFeedRequest

func (p *FeedServiceFeedArgs) GetRequest() (v *DouyinFeedRequest) {
	if !p.IsSetRequest() {
		return FeedServiceFeedArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *FeedServiceFeedArgs) SetRequest(val *DouyinFeedRequest) {
	p.Request = val
}

var fieldIDToName_FeedServiceFeedArgs = map[int16]string{
	1: "request",
}

func (p *FeedServiceFeedArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *FeedServiceFeedArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedServiceFeedArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedServiceFeedArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewDouyinFeedRequest()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FeedServiceFeedArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Feed_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FeedServiceFeedArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FeedServiceFeedArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedServiceFeedArgs(%+v)", *p)
}

func (p *FeedServiceFeedArgs) DeepEqual(ano *FeedServiceFeedArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *FeedServiceFeedArgs) Field1DeepEqual(src *DouyinFeedRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type FeedServiceFeedResult struct {
	Success *DouyinFeedResponse `thrift:"success,0,optional" frugal:"0,optional,DouyinFeedResponse" json:"success,omitempty"`
}

func NewFeedServiceFeedResult() *FeedServiceFeedResult {
	return &FeedServiceFeedResult{}
}

func (p *FeedServiceFeedResult) InitDefault() {
	*p = FeedServiceFeedResult{}
}

var FeedServiceFeedResult_Success_DEFAULT *DouyinFeedResponse

func (p *FeedServiceFeedResult) GetSuccess() (v *DouyinFeedResponse) {
	if !p.IsSetSuccess() {
		return FeedServiceFeedResult_Success_DEFAULT
	}
	return p.Success
}
func (p *FeedServiceFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*DouyinFeedResponse)
}

var fieldIDToName_FeedServiceFeedResult = map[int16]string{
	0: "success",
}

func (p *FeedServiceFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FeedServiceFeedResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FeedServiceFeedResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FeedServiceFeedResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewDouyinFeedResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FeedServiceFeedResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Feed_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FeedServiceFeedResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FeedServiceFeedResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeedServiceFeedResult(%+v)", *p)
}

func (p *FeedServiceFeedResult) DeepEqual(ano *FeedServiceFeedResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *FeedServiceFeedResult) Field0DeepEqual(src *DouyinFeedResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}