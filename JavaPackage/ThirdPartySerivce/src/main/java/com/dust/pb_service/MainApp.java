package com.dust.pb_service;


import com.dust.pb_service.netty_service.MessageServer;
import io.netty.channel.epoll.Epoll;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;


public class MainApp {

    static final Logger logger = LoggerFactory.getLogger(MainApp.class);

    public static void main(String[] args) {
        logger.info("Epoll.isAvailable()={}", Epoll.isAvailable());
        MessageServer messageServer = new MessageServer();
        messageServer.startServer();
//        new Thread(new Runnable() {
//            @Override
//            public void run() {
//                new MessageServer().startServer();
//            }
//        }).start();

    }
}

