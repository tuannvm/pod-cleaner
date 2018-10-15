{{/* vim: set filetype=mustache: */}}

{{/*
Create a default scoped fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "fullname" -}}
{{- .Release.Name | trunc 63 -}}
{{- end -}}

{{/* Standard labels which follow common convention
https://docs.helm.sh/chart_best_practices/#standard-labels
 */}}
{{- define "labels.standard" -}}
heritage: {{ .Release.Service | quote }}
release: {{ .Release.Name | quote }}
chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
app: {{ template "fullname" $ }}
{{- end -}}
