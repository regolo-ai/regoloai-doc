```sh
curl -i --location $ENDPOINT \
     --header 'Content-Type: application/json' \
     --header 'Accept: application/json' \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     --data '{
               "model": "mistralai/Mistral-7B-Instruct-v0.2",
               "messages": [{
                   "role": "user",
                   "content": "Tell me about Rome in a concise manner"
               }],
               "stream": "True"
             }'
