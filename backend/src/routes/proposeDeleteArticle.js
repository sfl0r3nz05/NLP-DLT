const express = require("express");

// Controllers
const { proposeDeleteArticle } = require("../controller/proposeDeleteArticle");

const route = express.Router();

route.post("/", proposeDeleteArticle);

module.exports = route;