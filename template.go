package mdopen

const Template = `<!doctype html>
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <link href="https://michaelwhatcott.com/css/custom.css" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Nunito:400,700|Source+Code+Pro:400,700|Ubuntu:400,700&display=swap" rel="stylesheet">
	<style>
		body, nav a      { font-family: 'Nunito', serif; }
		code, pre, .date { font-family: 'Source Code Pro', monospace; }
		h1, h2           { font-family: 'Ubuntu', sans-serif; }
		
		body, .tldr { font-size: smaller; }
		
		@media screen and (max-width: 800px) {
			body { font-size: 1.2em; }
		}
		
		blockquote p { font-style: italic; }
		
		nav, h1, footer, .tldr { text-align: center; }
		
		nav a { text-decoration: none; }
		
		body { line-height: 1.5; }
		
		/**********************************************/
		
		body       { background-color: #fff; }
		blockquote { background-color: #ddd; }
		code       { background-color: #ddd; }
		pre        { background-color: #ddd; }
		nav a      { background-color: #ccc; }
	
		blockquote { border-left: 10px solid #ccc; }
	
		nav a             { color: #666; }
		body              { color: #666; }
		strong            { color: #444; }
		h2                { color: #444; }
		h1                { color: #444; }
		a, a:visited      { color: #333; }
	
		/**********************************************/
		
		nav ul { list-style: none; }
		
		img.music { box-shadow: 0 0 18px #ccc; }
		
		img.music  { border-radius: 5px; }
		pre        { border-radius: 5px; }
		code       { border-radius: 5px; }
		nav a      { border-radius: 5px; }
		blockquote { border-radius: 5px; }
		
		pre { overflow-x: auto }
		
		/**********************************************/
		
		body              { margin: 0 auto; }
		img.music         { margin: 40px auto; }
		audio             { margin: 0 auto; }
		blockquote        { margin: 0 auto; }
		blockquote:before { margin-right: 0.25em; }
		nav               { margin-bottom: 20px; }
		
		body        { padding: 1em; }
		code        { padding: 3px; }
		blockquote  { padding: .1em 1em; }
		pre         { padding: 1em; }
		td          { padding: 0 10px; }
		nav a       { padding: 5px; }
		nav ul      { padding-inline-start: 0; }
		
		/**********************************************/
		
		body      { max-width: 42em; }
		img.music { max-width: 100%; }
		
		audio     { display: table; }
		nav ul li { display: inline; }
		
		div.clear { clear: both; }
	</style>
	<title></title>
</head>
<body>
    <div class="body article">
        {{ .Body }}
    </div>
</body>
</html>
`
