const express = require("express");

const { enrollAdmins } = require("../controller/enrollAdmins");

const route = express.Router();

route.post("/", enrollAdmins);

module.exports = route;