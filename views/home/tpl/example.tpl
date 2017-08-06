{% extends "../../common/tpl/base.tpl" %}

{% block content %}

Example: {{query}}
<p>Context value - {{context("ctxValue").nested}}</p>
<p>Admin URL - {{url("admin.index")}}</p>


{% endblock %}
