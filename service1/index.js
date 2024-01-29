const express = require("express");
const app = express();
const amqp = require("amqplib/callback_api");

const url = "amqp://localhost";
const queue = "MessageQueue";

let channel = null;
amqp.connect(url, function (err, conn) {
  if (!conn) {
    throw new Error(`AMQP connection not available on ${url}`);
  }
  conn.createChannel(function (err, ch) {
    channel = ch;

    ch.on("error", function (err) {
      console.error("[AMQP] channel error", err.message);
    });

    ch.on("close", function () {
      console.log("[AMQP] channel closed");
    });

    ch.assertQueue(queue);

    ch.consume(queue, (msg) => {
      if (msg !== null) {
        console.log(msg.content.toString());
      } else {
        console.log("Consumer cancelled by server");
      }
    });
  });
});

process.on("exit", (code) => {
  channel.close();
  console.log(`Closing`);
});

app.listen(4000, () => "Server is running...");
