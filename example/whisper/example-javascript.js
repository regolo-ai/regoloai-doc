const fetch = require('node-fetch');
const fs = require('fs');
const FormData = require('form-data');

const url = "https://api.regolo.ai/v1/models/whisper-large-v3/transcriptions";
const apiKey = "$REGOLOAI_API_KEY";

const form = new FormData();
form.append("model", "whisper-1");
form.append("file", fs.createReadStream("file.mp3"));

fetch(url, {
	  method: 'POST',
	  headers: { "Authorization": `Bearer ${apiKey}`, ...form.getHeaders() },
	  body: form
})
  .then(res => res.json())
  .then(data => console.log(data))
  .catch(err => console.error('Error:', err));
