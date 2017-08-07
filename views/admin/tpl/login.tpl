{% extends "base.tpl" %}

{% block content %}

<h1>Admin Login</h1>
<form action="{{url("admin.login")}}" method="post">
    <input name="login" type="text">
    <input name="password" type="password">
    <button type="submit">OK</button>
</form>

{% endblock %}
