const express = require("express");
const mongoose = require("mongoose");
require('dotenv').config();
const { consumeEvents } = require("./kafka/consumer");

const app = express();

mongoose
  .connect(process.env.MONGO_URI)
  .then(() => {
    console.log("Connected to MongoDB");
  })
  .catch((err) => {
    console.log("Failed to connect to MongoDB", err);
  });

app.use(express.json());

app.get("/health", (req, res) => {
  res.send("Invoice service is up and running");
});

app.listen(3001, () => {
  console.log("Invoice service listening on port 3001");
  consumeEvents();
});
