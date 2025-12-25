# not-a-rule-engine

## What am I building?

I'm not exactly sure yet, but it's definitely not a rule-engine, and I should probably try to work out what I'm actually trying to build before I continue.

Updates:
- Hmm, starting to look like some sort of service proxy? Not sure yet.

## What's the goal?

Just felt like writing some good old fashioned code without any AI or tools, like we all used to do.

That in itself isn't a challenge for people like myself who predominantly don't use AI tools, so I am having fun trying to stricly see how far I can go without using any external dependencies outside of the Go std lib.

## Why now?

Sh\*ts and g\*ggles, why not.

## Current way to call API

### Data to send

Array of the following (services are called concurrently within the ingest handler, and for now responses are aggregated and returned at the same time)

1. type -> service type (e.g cache, log, etc.)
2. payload  -> this is going to be basically function name to call and then key value pairs for function params if any.

curl -X POST http://localhost:9000/api/ingest \
    -d '[{"type": 1, "payload": "- payload data -"}, {"type": 2, "payload": {"- payload key -": "- payload key val pairs -"}}]'

## Current expected response from a service will be something in this format.

### Service response

Array ruturned of each service's response.

Each response should contain data if successful (can be empty of course), and any potential errors from the service.

[{"data":null,"error":null},{"data":null,"error":null}]

## Other responses

Not implemented yet, but non-service related responses will be returned as JSON at some stage in a custom error type once I implement the error checking function in handler.go
