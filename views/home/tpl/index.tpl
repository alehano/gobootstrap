{% extends "../../common/tpl/base_main.tpl" %}

{% block content %}

<p>Test value: {{testValue}}</p>
<p>Context value: {{context("ctxValue")}}</p>
<p>Admin URL: {{url("admin.index")}}</p>

{% endblock %}
