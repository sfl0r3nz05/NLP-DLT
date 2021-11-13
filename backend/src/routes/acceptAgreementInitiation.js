const express = require("express");

// Controllers
const { acceptAgreementInitiation } = require("../controller/acceptAgreementInitiation");

const route = express.Router();

route.post("/", acceptAgreementInitiation);

module.exports = route;
