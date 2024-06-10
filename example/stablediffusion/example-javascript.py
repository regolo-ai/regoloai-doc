const fetch = require('node-fetch');

const url = process.env.ENDPOINT;
const token = process.env.REGOLO_TOKEN;

const data = { data: ["Cat playing the piano"] };

fetch(url, {
        method: 'POST',
            headers: {
                        'Content-Type': 'application/json',
                                'Authorization': `Bearer ${token}`
                                    },
                body: JSON.stringify(data)
                })
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Error:', error));
