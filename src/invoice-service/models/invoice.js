const mongoose = require('mongoose');

const invoiceSchema = new mongoose.Schema({
  orderId: {
    type: Number,
    required: true,
    unique: true,
  },
  productId: Number,
  quantity: Number,
  price: Number,
  status: String,
});

module.exports = mongoose.model('Invoice', invoiceSchema);
