var http = require('http');
var url = require('url');

http.createServer(function(req, res) {
    res.writeHead(200, {
        'Content-Type': 'application/json'
    });
    if (req.method === 'POST') {
        req.on('data', chunk => { });
        req.on('end', () => {
            res.end(JSON.stringify({"works":"yes"}));
        });
    } else {
        res.end("only POST allowed");
    }
}).listen(7070);
