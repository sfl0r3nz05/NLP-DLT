const express = require("express");

// Controllers
const { proposeAgreementInitiation } = require("../controller/proposeAgreementInitiation");

const route = express.Router();

route.post("/", proposeAgreementInitiation);

module.exports = route;
