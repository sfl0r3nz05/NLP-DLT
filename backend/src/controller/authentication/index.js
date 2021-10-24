const path = require('path');
const readBuffer = require("../../utils/buffer/readBuffer");
const { FileSystemWallet } = require('fabric-network');

const authentication = async (req, res) => {
  const { username, password } = req.body; console.log(username);
  let selectEnv;
  let companies;
  let company;
  let users;

  selectEnv = 3;
  users = await readBuffer(selectEnv);
  user = users.find((user) => user.username.toLowerCase() === username);

  const walletPath = path.join(process.cwd(), 'wallet');
  const wallet = new FileSystemWallet(walletPath);
  var userExists = await wallet.exists(user.username);

  if (user && userExists) {
    if (user.password === password) {
      selectEnv = 2;
      companies = await readBuffer(selectEnv); console.log(companies);
      company = companies.companies.find((company) => company.id === user.company);
      const payload = {
        username: user.username,
        name: user.name,
        surname: user.surname,
        idParticipant: user.idParticipant,
        img: user.img,
        email: user.email,
        path: user.path,
        company: company,
      };
      //const options = { expiresIn: "10h", issuer: "https://172.31.16.137" };
      //const secret = process.env.JWT_SECRET;
      //const token = jwt.sign(payload, secret, options); console.log(token);
      //return res.send({ token: token, user: payload }); //send OK and the token
      return res.send({ user: payload });
    } else res.status(401).send(); //if the password isn't correct
  } else res.status(404).send();
};

module.exports = { authentication };
