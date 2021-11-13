const express = require("express");

// Controllers
const { acceptReachAgreement } = require("../controller/acceptReachAgreement");

const route = express.Router();

route.post("/", acceptReachAgreement);

module.exports = route;
