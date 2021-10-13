const express = require("express");

// Controllers
const { addOrg } = require("../controller/addOrg");

const route = express.Router();

route.post("/", addOrg);

module.exports = route;
