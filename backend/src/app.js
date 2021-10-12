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
const query = require("./routes/query.js");
const queryAll = require("./routes/queryAll.js");
const enrollAdmins = require("./routes/enrollAdmins.js");
const registerUsers = require("./routes/registerUsers.js");
const authentication = require("./routes/authentication.js");
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

app.use("/api/query", query);
app.use("/api/queryAll", queryAll);
app.use("/api/enrollAdmins", enrollAdmins);
app.use("/api/registerUsers", registerUsers);
app.use("/api/authentication", authentication);
app.use("/api/registerParticipant", registerParticipant);


app.listen(
  PORT,
  HOST,
  console.info(`Server runing at port http://${HOST}:${PORT}`)
);