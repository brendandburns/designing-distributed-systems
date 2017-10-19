const http = require('http');
const fs = require('fs');

const port = 8080;
const path = process.env.MEDIA_PATH;

const requestHandler = (request, response) => {  
	console.log(request.url);
	fs.readdir(path + '/*.mp4', (err, items) => {
		var msg = {
			'kind': 'ItemList',
			'apiVersion': 'v1',
			'items': []
		};
		if (!items) {
			return msg;
		}
		for (var i = 0; i < items.length; i++) {
			msg.items.push(items[i]);
		}
		response.end(JSON.stringify(msg));
	});
}

const server = http.createServer(requestHandler);

server.listen(port, (err) => {  
	if (err) {
		return console.log('Error starting server', err);
	}

	console.log(`server is active on ${port}`)
});
