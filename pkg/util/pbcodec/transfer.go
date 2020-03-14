package pbcodec

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"net"
)

// 声明 Transfer 结构体
type Transfer struct {
	Conn          net.Conn       // 连接
	Buf           [1024 * 2]byte // 传输时，使用的缓冲
}

var (
	G_transfer *Transfer
)

func InitTransfer() {
	var (
		pTCPAddr *net.TCPAddr
		conn net.Conn
		err error
	)
	if pTCPAddr, err = net.ResolveTCPAddr("tcp", "127.0.0.1:9090"); err != nil {
		return
	}
	if conn, err = net.DialTCP("tcp", nil, pTCPAddr); err != nil {
		return
	}

	// 定义 Transfer 指针变量
	G_transfer = &Transfer{
		Conn:  conn,
	}
}


func (t * Transfer) SendMsg(dateType int, pbMessage proto.Message) (err error){
	var (
		sendBytes []byte
		readLen   int
	)

	if sendBytes, err = proto.Marshal(pbMessage); err != nil {
		return
	}

	sendLen := uint32(len(sendBytes))
	sendLen++
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], sendLen)

	buf2 := [1]byte{uint8(dateType)}

	if readLen, err = t.Conn.Write(buf[:4]); readLen != 4 && err != nil {
		if readLen == 0 {
			return errors.New("发送数据长度发生异常，长度为0")
		}
		return
	}
	if readLen, err = t.Conn.Write(buf2[:1]); readLen != 1 && err != nil {
		if readLen == 0 {
			return errors.New("发送数据长度发生异常，长度为0")
		}
		return
	}
	// 发送消息
	if readLen, err = t.Conn.Write(sendBytes); err != nil {
		if readLen == 0 {
			return errors.New("检查到服务器关闭，客户端也关闭")
		}
		return
	}
	return
}


// 获取并解析服务器的消息
func (t *Transfer) ReadMsg(response proto.Message) (err error) {
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		return
	}

	// 根据 buf[:4] 转成一个 uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])
	//根据pkglen 读取消息内容
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	if err = proto.Unmarshal(t.Buf[:pkgLen], response); err != nil {
		return
	}
	return
}
