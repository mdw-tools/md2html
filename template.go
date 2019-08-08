package mdopen

const Template = `<!doctype html>
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <link href="https://michaelwhatcott.com/css/custom.css" rel="stylesheet">
    <link href="http://localhost:1111/css/custom.css" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Nunito:400,700|Source+Code+Pro:400,700|Ubuntu:400,700&display=swap" rel="stylesheet">
	<link href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/atom-one-light.min.css" rel="stylesheet">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/languages/go.min.js"></script>
	<script>hljs.initHighlightingOnLoad();</script>    
	<title></title>
</head>
<body>
    <div class="body article">
        {{ .Body }}
    </div>
</body>
</html>
`
