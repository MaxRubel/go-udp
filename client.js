const dgram = require('dgram');
const client = dgram.createSocket('udp4');

const message = Buffer.from('Hello, UDP server');

client.send(message, 41234, 'localhost', (err) => {
  if (err) {
    console.error('Failed to send message:', err);
    client.close();
  } else {
    console.log('Hello from node');
  }
});

client.on('message', (msg, rinfo) => {
  console.log(`Node received: ${msg} from ${rinfo.address}:${rinfo.port}`);
  client.close();
});

client.on('error', (err) => {
  console.error(`Client error:\n${err.stack}`);
  client.close();
});

client.on('close', () => {
  console.log('Client closed');
});