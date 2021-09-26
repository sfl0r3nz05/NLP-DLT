const express = require("express");

const { registerParticipant } = require("../controller/registerParticipant");

const route = express.Router();

route.post("/", registerParticipant);

module.exports = route;