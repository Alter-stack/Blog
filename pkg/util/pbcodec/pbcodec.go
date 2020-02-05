package pbcodec

import (
	service "github.com/Alter/blog/service/pb_service/protofile"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)


type Codec interface {
	Decode([]byte) (proto.Message, error)
	Encode(proto.Message) ([]byte, error)
}



type MessagePbCodec struct {

}


func (m *MessagePbCodec) Decode(b []byte, message proto.Message) error {
	if err := proto.Unmarshal(b, message); err != nil {
		return err
	}
	return nil
}

func (m *MessagePbCodec) Encode(p proto.Message) ([]byte, error) {
	user, ok :=  p.(*service.User)
	if !ok {
		return nil, errors.New("Not")
	}
	res, err := proto.Marshal(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

