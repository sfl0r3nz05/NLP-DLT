const registerParticipant = async (req, res) => {
  try {
    let data = req.body; 
    console.log(data);
    if (data) res.sendStatus(200);
  } catch (error) {
    res.sendStatus(400);
  }
};

module.exports = { registerParticipant };