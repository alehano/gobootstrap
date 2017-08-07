{% extends "base.tpl" %}

{% block content %}

<p>Session value: {{context("session").admin_login}}</p>
<p><a href="{{url("admin.logout")}}">Logout</a></p>

{% endblock %}
