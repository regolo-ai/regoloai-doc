const fetch = require('node-fetch');

const url = process.env.ENDPOINT;
const token = process.env.REGOLO_TOKEN;

const data = {
	    model: "mistralai/Mistral-7B-Instruct-v0.2",
	    messages: [{ role: "user", content: "Tell me about Rome in a concise manner" }]
};

fetch(url, {
	    method: 'POST',
	    headers: {
		            'Content-Type': 'application/json',
		            'Accept': 'application/json',
		            'Authorization': `Bearer ${token}`
		        },
	    body: JSON.stringify(data)
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
