const express = require("express");

const { registerUsers } = require("../controller/registerUsers");

const route = express.Router();

route.post("/", registerUsers);

module.exports = route;