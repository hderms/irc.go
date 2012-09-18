var redis = require("redis"), client1 = redis.createClient();
var app = require('http').createServer(handler)
, io = require('socket.io').listen(app)
, fs = require('fs');
app.listen(6161);

function handler (req, res) {

  console.log(req);
  fs.readFile(__dirname + '/index.html',
    function (err, data) {
      if (err) {
        res.writeHead(500);
        return res.end('Error loading index.html');
      }

      res.writeHead(200);
      res.end(data);
    });
}


client1.on("message", function(channel, message) {
  console.log(channel + " : " + message)
  io.sockets.emit('push', { message: message });
});
client1.subscribe("irc_channel");
