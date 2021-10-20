const express = require("express");

const { recoverMNO } = require("../controller/recoverMNO");

const route = express.Router();

route.get("/", recoverMNO);

module.exports = route;