const { Kafka } = require('kafkajs');
const Invoice = require('../models/invoice');

const kafka = new Kafka({
  clientId: 'invoice-service',
  brokers: ['localhost:9092'],
});

const consumer = kafka.consumer({ groupId: 'invoice-group' });

const consumeEvents = async () => {
  await consumer.connect();
  await consumer.subscribe({ topic: 'order-events', fromBeginning: true });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const event = JSON.parse(message.value.toString());
      if (event.EventType === 'OrderCreated') {
        const { ID, productId, quantity, price, status } = JSON.parse(event.Payload);
        Invoice.create({
          orderId: ID,
          productId,
          quantity,
          price,
          status,
        }).then(() => {
          console.log('Invoice created successfully: ' + ID);
        }).catch((err) => {
          if (err.code === 11000) { // MongoDB duplicate key error code
            console.log('Duplicate invoice, skipping: ' + ID);
          } else {
            console.log('Failed to create invoice', err);
          }
        });
      }
    },
  });
};

module.exports = { consumeEvents };
