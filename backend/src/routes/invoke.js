const express = require("express");

const { invoke } = require("../controller/invoke");

const route = express.Router();

route.post("/", invoke);

module.exports = route;