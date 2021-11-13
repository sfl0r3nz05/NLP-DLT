const express = require("express");

// Controllers
const { acceptProposedChanges } = require("../controller/acceptProposedChanges");

const route = express.Router();

route.post("/", acceptProposedChanges);

module.exports = route;
