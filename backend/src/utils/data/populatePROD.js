var fs = require('fs')

module.exports = async function populatePROD(valueToUpdate) {
    try {
        console.log(valueToUpdate);
        fs.readFile(__dirname + "/../../data/products.json", function (err, data) {
            var obj = JSON.parse(data)
            obj.push({
                "mno1": "TELEFONICA",
                "country_mno1": "ES",
                "mno2": "ORANGE",
                "country_mno2": "FR",
                "ra_name": "RA001",
                "ra_status": "INIT",
                "timestamp": 1634803280
            })
            var json = JSON.stringify(obj)
            console.log(json);
            fs.writeFile(__dirname + "/../../data/products.json", json, function (err) {
                if (err) throw err;
            })
        })
        console.log(ler);
    } catch (error) {

    }
}