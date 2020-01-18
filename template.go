package mdopen

const Template = `<!doctype html>
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <link href="https://michaelwhatcott.com/css/custom.css" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Nunito:400,700|Source+Code+Pro:400,700|Ubuntu:400,700&display=swap" rel="stylesheet">
	<title></title>
</head>
<body>
    <div class="body article">
        {{ .Body }}
    </div>
</body>
</html>
`
