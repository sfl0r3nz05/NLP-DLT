const express = require("express");

const { queryMNO } = require("../controller/queryMNO");

const route = express.Router();

route.get("/", queryMNO);

module.exports = route;