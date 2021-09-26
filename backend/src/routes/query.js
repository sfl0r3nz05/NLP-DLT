const express = require("express");

const { query } = require("../controller/query");

const route = express.Router();

route.post("/", query);

module.exports = route;