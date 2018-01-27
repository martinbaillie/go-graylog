# API
The Graylog API follows a generated, slightly modified Swagger 1.2 format (read: not Swagger 2.0; not OpenAPI 3.0).

The Swagger doc can be found at `<graylog server>/api/api-browser`.
### API Version
- 2.4.0-beta.3+a6b18a2

### How-to: Codegen
> NOTE: What follows is awkward.

1. Install Node-based `api-spec-converter`.
```bash
npm install api-spec-converter
```

2. Patch the local `swagger-converter` module dependency with this [simple diff](https://github.com/martinbaillie/swagger-converter/commit/61ac25a6c4ab837dc3cf01ecb06f4b4f8393588b).

3. Convert.
> NOTE: Assumes locally running Graylog of chosen version.
```bash
api-spec-converter http://localhost:9000/api/api-docs --from=swagger_1 --to=swagger_2 > swagger.json
```

4. Edit `swagger.json` and delete the following definitions (counts as of 2.4):
	- AuthenticationConfig (3)
	- Metric (1)
	- CollectorInput (2)
	- MetricRegistry (1)
	- JsonSchema (1)
	- LookupResult (2)
	- ResultMessage (2)
	- Multimap (1)
	- CollectorConfiguration (5)
	- CollectorConfigurationSnippet (2)
	- Collection (4)
	- CollectorOutput (2)
	- ChunkedOutput (2)
	- CollectorConfigurationListResponse (1)

5. Rename:
	- RestBoolean --> boolean

6. Add basic authentication underneath `schemes`.
```json
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "security": [
    {
      "basicAuth": []
    }
  ],
 ```
 
7. Grab Golang dependencies.
```bash
go get -u github.com/go-openapi/runtime
go get -u golang.org/x/net/context
go get -u golang.org/x/net/context/ctxhttp
go get -u golang.org/x/net/context/ctxhttp
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

8. Golang Codegen.
```bash
swagger generate client --skip-validation -A go-graylog
```

9. Rename Java-ey reverse domain name notation packages in `pkg/plugins_*`.
