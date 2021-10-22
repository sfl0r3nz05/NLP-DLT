var fs = require('fs')

module.exports = async function populateLER(valueToUpdate) {
    try {
        fs.readFile(__dirname + "/../../data/LER.json", function (err, data) {
            var obj = JSON.parse(data)
            obj.LER.push({ "name": valueToUpdate })
            var json = JSON.stringify(obj)
            console.log(json);
            fs.writeFile(__dirname + "/../../data/LER.json", json, function (err) {
                if (err) throw err;
            })
        })
        console.log(ler);
    } catch (error) {

    }
}