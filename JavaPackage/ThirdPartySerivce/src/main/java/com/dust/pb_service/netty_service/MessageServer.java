package com.dust.pb_service.netty_service;

import com.dust.pb_service.netty_service.codec.CustomProtobufDecoder;
import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.*;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.codec.LengthFieldBasedFrameDecoder;
import io.netty.handler.codec.LengthFieldPrepender;
import io.netty.handler.codec.protobuf.ProtobufEncoder;
import io.netty.handler.logging.LogLevel;
import io.netty.handler.logging.LoggingHandler;


public class MessageServer {
//    private static final Logger LOGGER = LoggerFactory.getLogger(NettyServerListener.class);
    ServerBootstrap serverBootstrap;
    // boss
    EventLoopGroup bossGroup = new NioEventLoopGroup();
    // work
    EventLoopGroup workerGroup = new NioEventLoopGroup(2);
    /**
     * 启动服务器
     */
    public void startServer() {
        try {
            serverBootstrap = new ServerBootstrap()
                    .group(bossGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                    .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        protected void initChannel(SocketChannel socketChannel) throws Exception {
                            ChannelPipeline pipeline = socketChannel.pipeline();
                            // 负责通过4字节Header指定的Body长度将消息切割
                            pipeline.addLast("frameDecoder",
                                    new LengthFieldBasedFrameDecoder(1048576, 0, 4, 0, 4));
                            // 负责将frameDecoder处理后的完整的一条消息的protobuf字节码转成ProtocolMessage对象
                            pipeline.addLast("protobufDecoder",new CustomProtobufDecoder());
                            // 负责将写入的字节码加上4字节Header前缀来指定Body长度
                            pipeline.addLast("frameEncoder", new LengthFieldPrepender(4));
                            // 负责将ProtocolMessage对象转成protobuf字节码
                            pipeline.addLast("protobufEncoder", new ProtobufEncoder());
                            pipeline.addLast(new MessageServerHandler());
                        }
                    })
                    .option(ChannelOption.SO_KEEPALIVE, true)
                    .option(ChannelOption.SO_BACKLOG, 100)
                    .handler(new LoggingHandler(LogLevel.INFO));;
            Channel channel = serverBootstrap.bind(9090).sync().channel();
            channel.closeFuture().sync();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
