## Example `curl` Request for Regolo AI API

```sh
curl -X POST --location $ENDPOINT \
     --header 'Content-Type: application/json' \
     --header "Authorization: Bearer ${REGOLO_TOKEN}" \
     --data '{ "data": ["Cat playing the piano"] }' \
     $ENDPOINT
