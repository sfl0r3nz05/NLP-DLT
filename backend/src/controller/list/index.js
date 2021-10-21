const readBuffer = require("../../utils/buffer/readBuffer");

const list = async (req, res) => {
  try {
    const selectEnv = 1;
    const objs = await readBuffer(selectEnv);

    if (objs) res.status(200).send(objs);
    else res.status(404).send("Products not available");
  } catch (error) {
    console.error(error);
    res.sendStatus(400);
  }
};

module.exports = { list };
