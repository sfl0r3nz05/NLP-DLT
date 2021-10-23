const query = require("../query/query");

module.exports = async function recoverMNO(user) {
    try {
        let method = "recoverMNO";
        let noArg = 0
        value = await query(method, noArg, "", user);
        return value
    } catch (error) {
        console.error(error);
        return error
    }
};