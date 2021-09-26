require("dotenv").config();
const cors = require("cors");
const express = require("express");
const bodyParser = require("body-parser");
// Prometheus metrics definition-----------------------------
const client = require('prom-client');
const promBundle = require("express-prom-bundle");
const promMetrics = promBundle({ includePath: true });
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics({ gcDurationBuckets: [0.1, 0.2, 0.3] });
// -----------------------------------------------------------

// Routes file
const registerParticipant = require("./routes/registerParticipant.js");

// Get PORT from env or set default
const PORT = process.env.PORT;
const HOST = process.env.HOST;

// Init express instance
const app = express();

// enable cors policy
app.use(cors());

// show images
app.use("/public", express.static(__dirname + "/public"));
app.use(promMetrics);
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use("/api/registerParticipant", registerParticipant);

app.listen(
  PORT,
  HOST,
  console.info(`Server runing at port http://${HOST}:${PORT}`)
);