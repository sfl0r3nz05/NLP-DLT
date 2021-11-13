const express = require("express");

// Controllers
const { rejectProposedChanges } = require("../controller/rejectProposedChanges");

const route = express.Router();

route.post("/", rejectProposedChanges);

module.exports = route;