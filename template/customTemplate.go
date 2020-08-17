package template

// CustomTemplate User defined template
var CustomTemplate = `
{
	"stdOutput": {
		"spec":{
		{{- with .Spec -}}
		"numberOfConnections":{{ .NumberOfConnections }}
		
		{{- if .IsTimedTest -}}
		,"testType":"timed","testDurationSeconds":{{ .TestDuration.Seconds }}
		{{- else -}}
		,"testType":"number-of-requests","numberOfRequests":{{ .NumberOfRequests }}
		{{- end -}}
		
		,"method":"{{ .Method }}","url":{{ .URL | printf "%q" }}
		
		{{- with .Headers -}}
		,"headers":[
		{{- range $index, $header :=  . -}}
		{{- if ne $index 0 -}},{{- end -}}
		{"key":{{ .Key | printf "%q" }},"value":{{ .Value | printf "%q" }}}
		{{- end -}}
		]
		{{- end -}}
		
		{{- if .BodyFilePath -}}
		,"bodyFilePath":{{ .BodyFilePath | printf "%q" }}
		{{- else -}}
		,"body":{{ .Body | printf "%q" }}
		{{- end -}}
		
		{{- if .CertPath -}}
		,"certPath":{{ .CertPath | printf "%q" }}
		{{- end -}}
		{{- if .KeyPath -}}
		,"keyPath":{{ .KeyPath | printf "%q" }}
		{{- end -}}
		
		,"stream":{{ .Stream }},"timeoutSeconds":{{ .Timeout.Seconds }}
		
		{{- if .IsFastHTTP -}}
		,"client":"fasthttp"
		{{- end -}}
		{{- if .IsNetHTTPV1 -}}
		,"client":"net/http.v1"
		{{- end -}}
		{{- if .IsNetHTTPV2 -}}
		,"client":"net/http.v2"
		{{- end -}}
		
		{{- with .Rate -}}
		,"rate":{{ . }}
		{{- end -}}
		{{- end -}}
		},
		
		{{- with .Result -}}
		"result":{"bytesRead":{{ .BytesRead -}}
		,"bytesWritten":{{ .BytesWritten -}}
		,"timeTakenSeconds":{{ .TimeTaken.Seconds -}}
		
		,"req1xx":{{ .Req1XX -}}
		,"req2xx":{{ .Req2XX -}}
		,"req3xx":{{ .Req3XX -}}
		,"req4xx":{{ .Req4XX -}}
		,"req5xx":{{ .Req5XX -}}
		,"others":{{ .Others -}}
		
		{{- with .Errors -}}
		,"errors":[
		{{- range $index, $error :=  . -}}
		{{- if ne $index 0 -}},{{- end -}}
		{"description":{{ .Error | printf "%q" }},"count":{{ .Count }}}
		{{- end -}}
		]
		{{- end -}}
		
		{{- with .LatenciesStats (FloatsToArray 0.5 0.75 0.9 0.95 0.99) -}}
		,"latency":{"mean":{{ .Mean -}}
		,"stddev":{{ .Stddev -}}
		,"max":{{ .Max -}}
		
		{{- if WithLatencies -}}
		,"percentiles":{
		{{- range $pc, $lat := .Percentiles }}
		{{- if ne $pc 0.5 -}},{{- end -}}
		{{- printf "\"%2.0f\":%d" (Multiply $pc 100) $lat -}}
		{{- end -}}
		}
		{{- end -}}
		
		}
		{{- end -}}
		
		{{- with .RequestsStats (FloatsToArray 0.5 0.75 0.9 0.95 0.99) -}}
		,"rps":{"mean":{{ .Mean -}}
		,"stddev":{{ .Stddev -}}
		,"max":{{ .Max -}}
		,"percentiles":{
		{{- range $pc, $rps := .Percentiles }}
		{{- if ne $pc 0.5 -}},{{- end -}}
		{{- printf "\"%2.0f\":%f" (Multiply $pc 100) $rps -}}
		{{- end -}}
		}}
		{{- end -}}
		}}
		{{- end -}}
	,"CodeFrequency": {{with .Result.CodeList}}{{ParseMap (.)}}{{end}} }
		`
