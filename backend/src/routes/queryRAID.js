const express = require("express");

const { queryRAID } = require("../controller/queryRAID");

const route = express.Router();

route.get("/", queryRAID);

module.exports = route;