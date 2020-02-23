package com.dust.pb_service.netty_service.codec;

import com.dust.pb_service.proto.Message;
import com.google.protobuf.MessageLite;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufUtil;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.MessageToMessageDecoder;

import java.util.List;

public class CustomProtobufDecoder extends MessageToMessageDecoder<ByteBuf> {


    private Exception decodeException;

    @Override
    protected void decode(ChannelHandlerContext ctx, ByteBuf msg, List<Object> out) throws Exception {
        final byte[] array;
        final int offset;
        final int length = msg.readableBytes();

        if (length < 1) {
            return;
        }

        int messageTypeByte = msg.readByte();
        if (msg.hasArray()) {
            array = msg.array();
            offset = msg.arrayOffset() + msg.readerIndex();
        } else {
            array = ByteBufUtil.getBytes(msg, msg.readerIndex(), length-1, false);
            offset = 0;
        }

        int readableLen= msg.readableBytes();

        //反序列化
        MessageLite result = decodeBody(messageTypeByte, array, offset, readableLen);
        out.add(result);

    }

    private MessageLite decodeBody(int dataType, byte[] array, int offset, int length) throws Exception {

        if (dataType == Message.MyMessage.DataType.ReadAllRequest.getNumber()) {
            return Message.ReadAllRequest
                    .getDefaultInstance()
                    .getParserForType()
                    .parseFrom(array, offset, length);
        } else if (dataType == Message.MyMessage.DataType.ReadRequest.getNumber()) {
            return Message.ReadRequest
                    .getDefaultInstance()
                    .getParserForType()
                    .parseFrom(array, offset, length);
        } else if (dataType == Message.MyMessage.DataType.CreateRequest.getNumber()) {
            return Message.CreateRequest
                    .getDefaultInstance()
                    .getParserForType()
                    .parseFrom(array, offset, length);
        } else if (dataType == Message.MyMessage.DataType.UpdateRequest.getNumber()) {
            return Message.UpdateRequest
                    .getDefaultInstance()
                    .getParserForType()
                    .parseFrom(array, offset, length);
        } else {
            throw decodeException;
        }
        // or throw exception
    }
    /**
     * int转byte[]
     * 该方法将一个int类型的数据转换为byte[]形式，因为int为32bit，而byte为8bit所以在进行类型转换时，知会获取低8位，
     * 丢弃高24位。通过位移的方式，将32bit的数据转换成4个8bit的数据。注意 &0xff，在这当中，&0xff简单理解为一把剪刀，
     * 将想要获取的8位数据截取出来。
     * @param i 一个int数字
     * @return byte[]
     */
    public static byte[] int2ByteArray(int i){
        byte[] result=new byte[4];
        result[0]=(byte)((i >> 24)& 0xFF);
        result[1]=(byte)((i >> 16)& 0xFF);
        result[2]=(byte)((i >> 8)& 0xFF);
        result[3]=(byte)(i & 0xFF);
        return result;
    }
    /**
     * byte[]转int
     * 利用int2ByteArray方法，将一个int转为byte[]，但在解析时，需要将数据还原。同样使用移位的方式，将适当的位数进行还原，
     * 0xFF为16进制的数据，所以在其后每加上一位，就相当于二进制加上4位。同时，使用|=号拼接数据，将其还原成最终的int数据
     * @param bytes byte类型数组
     * @return int数字
     */
    public static int bytes2Int(byte[] bytes){
        int num=bytes[3] & 0xFF;
        num |=((bytes[2] <<8)& 0xFF00);
        num |=((bytes[1] <<16)& 0xFF0000);
        num |=((bytes[0] <<24)& 0xFF0000);
        return num;
    }
}


//        while (msg.readableBytes() > 4) { // 如果可读长度小于包头长度，退出。
//            msg.markReaderIndex();
//
//            // 获取包头中的body长度
//            byte low = msg.readByte();
//            byte high = msg.readByte();
//            short s0 = (short) (low & 0xff);
//            short s1 = (short) (high & 0xff);
//            s1 <<= 8;
//            short length = (short) (s0 | s1);
//
//            // 获取包头中的protobuf类型
//            msg.readByte();
//            byte dataType = msg.readByte();
//
//            // 如果可读长度小于body长度，恢复读指针，退出。
//            if (msg.readableBytes() < length) {
//                msg.resetReaderIndex();
//                return;
//            }
//
//            // 读取body
//            ByteBuf bodyByteBuf = msg.readBytes(length);
//
//            int readableLen= bodyByteBuf.readableBytes();
//            if (bodyByteBuf.hasArray()) {
//                array = bodyByteBuf.array();
//                offset = bodyByteBuf.arrayOffset() + bodyByteBuf.readerIndex();
//            } else {
//                array = new byte[readableLen];
//                bodyByteBuf.getBytes(bodyByteBuf.readerIndex(), array, 0, readableLen);
//                offset = 0;
//            }
//
//            //反序列化
//            MessageLite result = decodeBody(dataType, array, offset, readableLen);
//            out.add(result);
//        }