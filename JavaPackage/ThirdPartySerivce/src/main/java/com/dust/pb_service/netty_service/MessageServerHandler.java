package com.dust.pb_service.netty_service;

import com.dust.pb_service.proto.Message;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class MessageServerHandler extends ChannelInboundHandlerAdapter {
    private static final Logger LOGGER = LoggerFactory.getLogger(MessageServerHandler.class);

    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        LOGGER.info("开始读取客户端发送过来的数据");
        if (msg instanceof Message.ReadAllRequest) {
            LOGGER.info(msg.toString());
            Message.ReadAllResponse response = packReadAllReponse((Message.ReadAllRequest) msg);
            LOGGER.info(response.toString());
            ctx.channel().writeAndFlush(response.toByteArray().length);
            ctx.channel().writeAndFlush(packReadAllReponse((Message.ReadAllRequest) msg));
        } else if (msg instanceof Message.CreateRequest) {
            System.out.println(msg.toString());
        } else if (msg instanceof Message.ReadRequest) {
            System.out.println(msg.toString());
        } else if (msg instanceof Message.UpdateRequest) {
            System.out.println(msg.toString());
        }
    }

    private Message.ReadAllResponse packReadAllReponse(Message.ReadAllRequest request) {
        String MessageToken = request.getToken();
        Message.User user1 = Message.User.newBuilder()
                .setName("test1")
                .setPassword("test2")
                .setPermission("1")
                .build();
        Message.User user2 = Message.User.newBuilder()
                .setName("test1")
                .setPassword("test2")
                .setPermission("1")
                .build();
        return Message.ReadAllResponse.newBuilder()
                .setToken(MessageToken)
                .addUsers(user1)
                .addUsers(user2)
//                .setUsers(0, user1)
//                .setUsers(1, user2)
                .build();
    }
}
