require("dotenv").config();
const cors = require("cors");
const express = require("express");
const bodyParser = require("body-parser");
// Prometheus metrics definition-----------------------------
const client = require('prom-client');
const promBundle = require("express-prom-bundle");
const promMetrics = promBundle({ includePath: true });
const collectDefaultMetrics = client.collectDefaultMetrics;
collectDefaultMetrics({ gcDurationBuckets: [0.1, 0.2, 0.3] });
// -----------------------------------------------------------

// Routes file
const list = require("./routes/list.js");
const addOrg = require("./routes/addOrg.js");
const queryMNO = require("./routes/queryMNO.js");
const queryRAID = require("./routes/queryRAID.js");
const enrollAdmins = require("./routes/enrollAdmins.js");
const registerUsers = require("./routes/registerUsers.js");
const authentication = require("./routes/authentication.js");
const proposeAddArticle = require("./routes/proposeAddArticle.js");
const querySingleArticle = require("./routes/querySingleArticle.js");
const proposeUpdateArticle = require("./routes/proposeUpdateArticle.js");
const rejectProposedChanges = require("./routes/rejectProposedChanges.js");
const acceptReachAgreement = require("./routes/acceptReachAgreement.js");
const acceptProposedChanges = require("./routes/acceptProposedChanges.js");
const acceptAgreementInitiation = require("./routes/acceptAgreementInitiation.js");
const proposeAgreementInitiation = require("./routes/proposeAgreementInitiation.js");

// Get PORT from env or set default
const PORT = process.env.PORT;
const HOST = process.env.HOST;

// Init express instance
const app = express();

// enable cors policy
app.use(cors());

// show images
app.use("/public", express.static(__dirname + "/public"));
app.use(promMetrics);
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use("/api/list", list);
app.use("/api/addOrg", addOrg);
app.use("/api/queryMNO", queryMNO);
app.use("/api/queryRAID", queryRAID);
app.use("/api/enrollAdmins", enrollAdmins);
app.use("/api/registerUsers", registerUsers);
app.use("/api/authentication", authentication);
app.use("/api/proposeAddArticle", proposeAddArticle);
app.use("/api/querySingleArticle", querySingleArticle);
app.use("/api/acceptReachAgreement", acceptReachAgreement);
app.use("/api/proposeUpdateArticle", proposeUpdateArticle);
app.use("/api/rejectProposedChanges", rejectProposedChanges);
app.use("/api/acceptProposedChanges", acceptProposedChanges);
app.use("/api/acceptAgreementInitiation", acceptAgreementInitiation);
app.use("/api/proposeAgreementInitiation", proposeAgreementInitiation);


app.listen(
  PORT,
  HOST,
  console.info(`Server runing at port http://${HOST}:${PORT}`)
);