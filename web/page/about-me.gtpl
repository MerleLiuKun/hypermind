<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>About Me - Hypermind</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="A page of Hypermind.">
    <meta name="author" content="hyper-carrot">

    <link href="../css/bootstrap.css" rel="stylesheet">
    <style>
        body {
        padding-top: 60px;
        }
    </style>
    <link href="../css/bootstrap-responsive.css" rel="stylesheet">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
    <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <link rel="shortcut icon" href="../img/favicon.ico">
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="../img/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="../img/apple-touch-icon-114-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="72x72" href="../img/apple-touch-icon-72-precomposed.png">
    <link rel="apple-touch-icon-precomposed" href="../img/apple-touch-icon-57-precomposed.png">
</head>
<body>
<div class="navbar navbar-fixed-top">
    <div class="navbar-inner">
        <div class="container-fluid">
            <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </a>
            <a class="brand" href="#">Hypermind</a>
            <div class="nav-collapse collapse">
                <ul class="nav">
                    <li><a href="/?page={{.homePage}}">Home</a></li>
                    <li><a href="/?page={{.meetingKanbanPage}}">Meeting Kanban</a></li>
                    <li class="active"><a href="/?page={{.aboutMePage}}">About</a></li>
                </ul>
            </div>
            <ul class="nav pull-right">
                {{if .validLogin}}
                <li><a href="#">Hi, {{.loginName}}</a></li>
                <li class="divider-vertical"></li>
                <a class="btn navbar-form pull-right" href="http://{{.serverAddr}}:{{.serverPort}}/logout">Sign Out</a></p>
                {{else}}
                <li class="divider-vertical"></li>
                <a class="btn navbar-form pull-right" href="http://{{.serverAddr}}:{{.serverPort}}/login">I'm Admin</a></p>
                {{end}}
            </ul>

        </div>
    </div>
</div>

<div class="container-fluid">
    <div class="row-fluid">
        <div class="span2">
            <div class="well sidebar-nav">
                <ul class="nav nav-list">
                    <li class="nav-header">About</li>
                    <li class="active"><a href="/?page={{.aboutMePage}}">About Me</a></li>
                    <li><a href="/?page={{.aboutWebsitePage}}">About Website</a></li>
                </ul>
            </div>
        </div>
        <div class="span10">
            <div class="hero-unit">
                <p>
                    My name is Harry Hao.
                    I live in Beijing.
                    I am in Sohu Inc (NSDQ: SOHU) as the position of Dev Leader.
                </p>
                <p>
                    I'm a broad interests software developer. I'm a open source fan, and pay attention to the agile methods and software process improvement.
                    I focus on Clojure and Go programming language, and contribute strength in order to the popularization of them in Chinese community.
                </p>
                <p>
                    My homepage in GitHub is
                    <a class="btn btn-small" href="https://github.com/hyper-carrot">hyper-carrot</a>
                    .
                </p>
            </div>
        </div>
    </div>
</div>

<script src="../js/jquery.js"></script>
<script src="../js/bootstrap.js"></script>

</body>
</html>

