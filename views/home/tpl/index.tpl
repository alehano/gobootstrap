{% extends "../../common/tpl/base_main.tpl" %}

{% block content %}

<h1>Homepage</h1>
<p>Test value: {{testValue}}</p>
<p>Context value: {{context("ctxValue")}}</p>
<p>Admin URL: {{url("admin.index")}}</p>

{% endblock %}
