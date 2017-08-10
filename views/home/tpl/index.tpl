{% extends "../../common/tpl/base_main.tpl" %}

{% block content %}

<h1>GOBOOTSTRAP</h1>
<p>Handler value: {{testValue}}</p>
<p>Context value: {{context("ctxValue")}}</p>
<p>Admin URL: <a href="{{url("admin.index")}}">{{url("admin.index")}}</a></p>
<p>JSON example: <a href="{{url("home.json")}}">{{url("home.json")}}</a></p>

{% endblock %}
