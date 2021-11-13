const recoverMNO = require("../../utils/recoverMNO");
const readBuffer = require("../../utils/buffer/readBuffer");

const list = async (req, res) => {
  try {
    const selectEnv = 1;
    const objs = await readBuffer(selectEnv);
    console.log(objs);
    userInfo = JSON.parse(req.query.ID)
    mno = await recoverMNO(userInfo.username)
    var output = objs.filter(function (obj) { return obj.mno1 == mno || obj.mno2 == mno; })
    if (output) res.status(200).send(output);
    else res.status(404).send("Products not available");
  } catch (error) {
    console.error(error);
    res.sendStatus(400);
  }
};

module.exports = { list };
