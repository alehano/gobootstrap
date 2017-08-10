{% extends "base.tpl" %}

{% block content %}

<h1>Admin Login</h1>
<p>Login: admin, Password: admin</p>
<form action="{{url("admin.login")}}" method="post">
    <input name="login" type="text">
    <input name="password" type="password">
    <button>OK</button>
</form>

{% endblock %}
