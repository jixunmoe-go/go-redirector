package main

import "html/template"

var tplGoGetPage = template.Must(template.New("go-get-page").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta name="go-import" content="{{ .ImportLine }}">
		<meta name="go-source" content="{{ .SourceLine }}">
	</head>
	<body>
		<pre>go get {{ .PackageName }}</pre>
	</body>
</html>`))

var tplBrowserPage = template.Must(template.New("browser-page").Parse(`<!DOCTYPE html>
<html lang="en">
	<body>
		<style>body{font-family:sans-serif}h2{margin:0}main pre{margin-top:0}</style>
		<pre>go get <a href="{{ .BaseWebURL }}">{{ .PackageName }}</a></pre>
		<main>
			<p>Note: this is a private go module repo. You should this host to <code>GONOPROXY</code>.</p>

			<h2>( windows )</h2>
			<pre>set "GONOPROXY=%GONOPROXY%,{{ .DeployHost }}"</pre>

			<h2>( linux )</h2>
			<pre>export GONOPROXY="${GONOPROXY},{{ .DeployHost }}"</pre>
			<p>hint: add this line to <code>/etc/profile.d/gonoproxy.sh</code></p>
		</main>
	</body>
</html>`))
