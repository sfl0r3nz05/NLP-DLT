const express = require("express");

// Controllers
const { list } = require("../controller/list");

const route = express.Router();

route.get("/", list);

module.exports = route;
