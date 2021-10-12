const express = require("express");

// Controllers
const { authentication } = require("../controller/authentication");

const route = express.Router();

route.post("/", authentication);

module.exports = route;