const express = require("express");

const { queryAll } = require("../controller/queryAll");

const route = express.Router();

route.post("/", queryAll);

module.exports = route;